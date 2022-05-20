// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web_test

import (
	"context"
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/web"
	"go.thethings.network/lorawan-stack/v3/pkg/applicationserver/io/web/redis"
	"go.thethings.network/lorawan-stack/v3/pkg/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/component"
	componenttest "go.thethings.network/lorawan-stack/v3/pkg/component/test"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	mockis "go.thethings.network/lorawan-stack/v3/pkg/identityserver/mock"
	"go.thethings.network/lorawan-stack/v3/pkg/rpcmetadata"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
	"google.golang.org/grpc"
)

func TestWebhookRegistryRPC(t *testing.T) {
	a, ctx := test.New(t)

	is, isAddr, closeIS := mockis.New(ctx)
	defer closeIS()
	is.ApplicationRegistry().Add(ctx, registeredApplicationID, registeredApplicationKey,
		ttnpb.Right_RIGHT_APPLICATION_SETTINGS_BASIC,
		ttnpb.Right_RIGHT_APPLICATION_DEVICES_READ,
		ttnpb.Right_RIGHT_APPLICATION_DEVICES_WRITE,
		ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_READ,
		ttnpb.Right_RIGHT_APPLICATION_TRAFFIC_DOWN_WRITE)

	c := componenttest.NewComponent(t, &component.Config{
		ServiceBase: config.ServiceBase{
			GRPC: config.GRPC{
				Listen:                      ":0",
				AllowInsecureForCredentials: true,
			},
			Cluster: cluster.Config{
				IdentityServer: isAddr,
			},
		},
	})
	redisClient, flush := test.NewRedis(ctx, "applicationserver_test")
	defer flush()
	defer redisClient.Close()
	webhookReg := &redis.WebhookRegistry{Redis: redisClient, LockTTL: test.Delay << 10}
	if err := webhookReg.Init(ctx); !a.So(err, should.BeNil) {
		t.FailNow()
	}
	srv := web.NewWebhookRegistryRPC(webhookReg, nil)
	c.RegisterGRPC(&mockRegisterer{ctx, srv})
	componenttest.StartComponent(t, c)
	defer c.Close()

	mustHavePeer(ctx, c, ttnpb.ClusterRole_ENTITY_REGISTRY)

	client := ttnpb.NewApplicationWebhookRegistryClient(c.LoopbackConn())
	creds := grpc.PerRPCCredentials(rpcmetadata.MD{
		AuthType:      "Bearer",
		AuthValue:     registeredApplicationKey,
		AllowInsecure: true,
	})

	// Formats.
	{
		res, err := client.GetFormats(ctx, ttnpb.Empty, creds)
		a.So(err, should.BeNil)
		a.So(res.Formats, should.HaveSameElementsDeep, map[string]string{
			"json":     "JSON",
			"protobuf": "Protocol Buffers",
		})
	}

	// Check empty.
	{
		res, err := client.List(ctx, &ttnpb.ListApplicationWebhooksRequest{
			ApplicationIds: registeredApplicationID,
			FieldMask:      ttnpb.FieldMask("base_url"),
		}, creds)
		a.So(err, should.BeNil)
		a.So(res.Webhooks, should.BeEmpty)
	}

	// Add.
	{
		_, err := client.Set(ctx, &ttnpb.SetApplicationWebhookRequest{
			Webhook: &ttnpb.ApplicationWebhook{
				Ids: &ttnpb.ApplicationWebhookIdentifiers{
					ApplicationIds: registeredApplicationID,
					WebhookId:      registeredWebhookID,
				},
				BaseUrl: "http://localhost/test",
				Format:  "json",
			},
			FieldMask: ttnpb.FieldMask("base_url", "format"),
		}, creds)
		a.So(err, should.BeNil)
	}

	// List; assert one.
	{
		res, err := client.List(ctx, &ttnpb.ListApplicationWebhooksRequest{
			ApplicationIds: registeredApplicationID,
			FieldMask:      ttnpb.FieldMask("base_url"),
		}, creds)
		a.So(err, should.BeNil)
		a.So(res.Webhooks, should.HaveLength, 1)
		a.So(res.Webhooks[0].BaseUrl, should.Equal, "http://localhost/test")
	}

	// Get.
	{
		res, err := client.Get(ctx, &ttnpb.GetApplicationWebhookRequest{
			Ids: &ttnpb.ApplicationWebhookIdentifiers{
				ApplicationIds: registeredApplicationID,
				WebhookId:      registeredWebhookID,
			},
			FieldMask: ttnpb.FieldMask("base_url"),
		}, creds)
		a.So(err, should.BeNil)
		a.So(res.BaseUrl, should.Equal, "http://localhost/test")
	}

	// Delete.
	{
		_, err := client.Delete(ctx, &ttnpb.ApplicationWebhookIdentifiers{
			ApplicationIds: registeredApplicationID,
			WebhookId:      registeredWebhookID,
		}, creds)
		a.So(err, should.BeNil)
	}

	// Check empty.
	{
		res, err := client.List(ctx, &ttnpb.ListApplicationWebhooksRequest{
			ApplicationIds: registeredApplicationID,
			FieldMask:      ttnpb.FieldMask("base_url"),
		}, creds)
		a.So(err, should.BeNil)
		a.So(res.Webhooks, should.BeEmpty)
	}
}

func TestTemplateStoreRPC(t *testing.T) {
	ctx := test.Context()

	for _, tc := range []struct {
		name       string
		contents   map[string][]byte
		assertGet  func(*assertions.Assertion, *ttnpb.ApplicationWebhookTemplate, error)
		assertList func(*assertions.Assertion, *ttnpb.ApplicationWebhookTemplates, error)
	}{
		{
			name: "InvalidStore",
			contents: map[string][]byte{
				"templates.yml": []byte(`invalid-yaml`),
			},
			assertGet: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplate, err error) {
				a.So(err, should.NotBeNil)
				a.So(res, should.BeNil)
			},
			assertList: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplates, err error) {
				a.So(err, should.NotBeNil)
				a.So(res, should.BeNil)
			},
		},
		{
			name: "EmptyStore",
			contents: map[string][]byte{
				"templates.yml": []byte(`--- []`),
			},
			assertGet: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplate, err error) {
				a.So(err, should.NotBeNil)
				a.So(res, should.BeNil)
			},
			assertList: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplates, err error) {
				a.So(err, should.BeNil)
				a.So(res, should.NotBeNil)
				a.So(res.Templates, should.BeEmpty)
			},
		},
		{
			name: "NormalStore",
			contents: map[string][]byte{
				"templates.yml": []byte(`---
- foo`),
				"foo.yml": []byte(
					`---
ids:
  template_id: foo
name: Foo
description: Bar`),
			},
			assertGet: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplate, err error) {
				a.So(err, should.BeNil)
				a.So(res, should.NotBeNil)
				a.So(res.Name, should.NotBeEmpty)
				a.So(res.Description, should.BeEmpty)
			},
			assertList: func(a *assertions.Assertion, res *ttnpb.ApplicationWebhookTemplates, err error) {
				a.So(err, should.BeNil)
				a.So(res, should.NotBeNil)
				a.So(res.Templates, should.HaveLength, 1)
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithCancel(ctx)
			defer cancel()

			a := assertions.New(t)

			c := componenttest.NewComponent(t, &component.Config{})

			config := web.TemplatesConfig{
				Static: tc.contents,
			}
			store, err := config.NewTemplateStore(ctx, c)
			a.So(err, should.BeNil)

			c.RegisterGRPC(&mockRegisterer{ctx, web.NewWebhookRegistryRPC(nil, store)})
			componenttest.StartComponent(t, c)
			defer c.Close()

			client := ttnpb.NewApplicationWebhookRegistryClient(c.LoopbackConn())

			getRes, err := client.GetTemplate(ctx, &ttnpb.GetApplicationWebhookTemplateRequest{
				Ids: &ttnpb.ApplicationWebhookTemplateIdentifiers{
					TemplateId: "foo",
				},
			})
			tc.assertGet(a, getRes, err)

			listRes, err := client.ListTemplates(ctx, &ttnpb.ListApplicationWebhookTemplatesRequest{
				FieldMask: ttnpb.FieldMask("name", "description"),
			})
			tc.assertList(a, listRes, err)
		})
	}
}
