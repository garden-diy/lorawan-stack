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

package mqtt

import (
	"context"

	"github.com/gogo/protobuf/proto"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io/mqtt/topics"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

type protobuf struct {
	topics.Layout
}

func (protobuf) FromDownlink(down *ttnpb.DownlinkMessage, _ *ttnpb.GatewayIdentifiers) ([]byte, error) {
	gwDown := &ttnpb.GatewayDown{
		DownlinkMessage: down,
	}
	return proto.Marshal(gwDown)
}

func (protobuf) ToUplink(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.UplinkMessage, error) {
	uplink := &ttnpb.UplinkMessage{}
	if err := proto.Unmarshal(message, uplink); err != nil {
		return nil, err
	}
	return uplink, nil
}

func (protobuf) ToStatus(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.GatewayStatus, error) {
	status := &ttnpb.GatewayStatus{}
	if err := proto.Unmarshal(message, status); err != nil {
		return nil, err
	}
	return status, nil
}

func (protobuf) ToTxAck(message []byte, _ *ttnpb.GatewayIdentifiers) (*ttnpb.TxAcknowledgment, error) {
	ack := &ttnpb.TxAcknowledgment{}
	if err := proto.Unmarshal(message, ack); err != nil {
		return nil, err
	}
	return ack, nil
}

// NewProtobuf returns a format that uses Protocol Buffers marshaling and unmarshaling.
func NewProtobuf(ctx context.Context) Format {
	return &protobuf{
		Layout: topics.New(ctx),
	}
}
