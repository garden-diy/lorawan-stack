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

package fetch_test

import (
	"testing"

	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/v3/pkg/config"
	. "go.thethings.network/lorawan-stack/v3/pkg/fetch"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test"
	"go.thethings.network/lorawan-stack/v3/pkg/util/test/assertions/should"
)

func TestBucket(t *testing.T) {
	a := assertions.New(t)
	ctx := test.Context()

	conf := config.BlobConfig{Provider: "local"}
	conf.Local.Directory = t.TempDir()

	filename := "file"
	content := []byte("Hello world")

	bucket, err := conf.Bucket(ctx, "bucket", test.HTTPClientProvider)
	a.So(err, should.BeNil)

	err = bucket.WriteAll(ctx, filename, content, nil)
	a.So(err, should.BeNil)

	fetcher := FromBucket(ctx, bucket, "")

	// Reading working file
	{
		fileContent, err := fetcher.File(filename)
		a.So(err, should.BeNil)
		a.So(string(fileContent), should.Equal, string(content))
	}

	// Reading non-working file
	{
		_, err = fetcher.File("non-existing file")
		a.So(err, should.NotBeNil)
	}
}
