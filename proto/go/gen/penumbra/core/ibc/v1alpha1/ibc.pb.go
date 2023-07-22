// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: penumbra/core/ibc/v1alpha1/ibc.proto

package ibcv1alpha1

import (
	types "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	v1alpha1 "github.com/penumbra-zone/penumbra/proto/go/gen/penumbra/core/crypto/v1alpha1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IbcAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// oneof action {
	// .ibc.core.connection.v1.MsgConnectionOpenInit connection_open_init = 1;
	// .ibc.core.connection.v1.MsgConnectionOpenTry connection_open_try = 2;
	// .ibc.core.connection.v1.MsgConnectionOpenAck connection_open_ack = 3;
	// .ibc.core.connection.v1.MsgConnectionOpenConfirm connection_open_confirm = 4;
	//
	// .ibc.core.channel.v1.MsgChannelOpenInit channel_open_init = 5;
	// .ibc.core.channel.v1.MsgChannelOpenTry channel_open_try = 6;
	// .ibc.core.channel.v1.MsgChannelOpenAck channel_open_ack = 7;
	// .ibc.core.channel.v1.MsgChannelOpenConfirm channel_open_confirm = 8;
	// .ibc.core.channel.v1.MsgChannelCloseInit channel_close_init = 9;
	// .ibc.core.channel.v1.MsgChannelCloseConfirm channel_close_confirm = 10;
	//
	// .ibc.core.channel.v1.MsgRecvPacket recv_packet = 11;
	// .ibc.core.channel.v1.MsgTimeout timeout = 12;
	// .ibc.core.channel.v1.MsgAcknowledgement acknowledgement = 13;
	//
	// .ibc.core.client.v1.MsgCreateClient create_client = 14;
	// .ibc.core.client.v1.MsgUpdateClient update_client = 15;
	// .ibc.core.client.v1.MsgUpgradeClient upgrade_client = 16;
	// .ibc.core.client.v1.MsgSubmitMisbehaviour submit_misbehaviour = 17;
	// }
	RawAction *anypb.Any `protobuf:"bytes,1,opt,name=raw_action,json=rawAction,proto3" json:"raw_action,omitempty"`
}

func (x *IbcAction) Reset() {
	*x = IbcAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IbcAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IbcAction) ProtoMessage() {}

func (x *IbcAction) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IbcAction.ProtoReflect.Descriptor instead.
func (*IbcAction) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{0}
}

func (x *IbcAction) GetRawAction() *anypb.Any {
	if x != nil {
		return x.RawAction
	}
	return nil
}

// FungibleTokenPacketData defines a struct for the packet payload
// See FungibleTokenPacketData spec:
// https://github.com/cosmos/ibc/tree/master/spec/app/ics-020-fungible-token-transfer#data-structures
type FungibleTokenPacketData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the token denomination to be transferred
	Denom string `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	// the token amount to be transferred
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	// the sender address
	Sender string `protobuf:"bytes,3,opt,name=sender,proto3" json:"sender,omitempty"`
	// the recipient address on the destination chain
	Receiver string `protobuf:"bytes,4,opt,name=receiver,proto3" json:"receiver,omitempty"`
}

func (x *FungibleTokenPacketData) Reset() {
	*x = FungibleTokenPacketData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FungibleTokenPacketData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FungibleTokenPacketData) ProtoMessage() {}

func (x *FungibleTokenPacketData) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FungibleTokenPacketData.ProtoReflect.Descriptor instead.
func (*FungibleTokenPacketData) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{1}
}

func (x *FungibleTokenPacketData) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *FungibleTokenPacketData) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *FungibleTokenPacketData) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *FungibleTokenPacketData) GetReceiver() string {
	if x != nil {
		return x.Receiver
	}
	return ""
}

type Ics20Withdrawal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount *v1alpha1.Amount `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Denom  *v1alpha1.Denom  `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	// the address on the destination chain to send the transfer to
	DestinationChainAddress string `protobuf:"bytes,3,opt,name=destination_chain_address,json=destinationChainAddress,proto3" json:"destination_chain_address,omitempty"`
	// a "sender" penumbra address to use to return funds from this withdrawal.
	// this should be an ephemeral address
	ReturnAddress *v1alpha1.Address `protobuf:"bytes,4,opt,name=return_address,json=returnAddress,proto3" json:"return_address,omitempty"`
	// the height (on Penumbra) at which this transfer expires (and funds are sent
	// back to the sender address?). NOTE: if funds are sent back to the sender,
	// we MUST verify a nonexistence proof before accepting the timeout, to
	// prevent relayer censorship attacks. The core IBC implementation does this
	// in its handling of validation of timeouts.
	TimeoutHeight uint64 `protobuf:"varint,5,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height,omitempty"`
	// the timestamp at which this transfer expires.
	TimeoutTime uint64 `protobuf:"varint,6,opt,name=timeout_time,json=timeoutTime,proto3" json:"timeout_time,omitempty"`
	// the source port that identifies the channel used for the withdrawal
	SourcePort string `protobuf:"bytes,7,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty"`
	// the source channel used for the withdrawal
	SourceChannel string `protobuf:"bytes,8,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty"`
}

func (x *Ics20Withdrawal) Reset() {
	*x = Ics20Withdrawal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ics20Withdrawal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ics20Withdrawal) ProtoMessage() {}

func (x *Ics20Withdrawal) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ics20Withdrawal.ProtoReflect.Descriptor instead.
func (*Ics20Withdrawal) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{2}
}

func (x *Ics20Withdrawal) GetAmount() *v1alpha1.Amount {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Ics20Withdrawal) GetDenom() *v1alpha1.Denom {
	if x != nil {
		return x.Denom
	}
	return nil
}

func (x *Ics20Withdrawal) GetDestinationChainAddress() string {
	if x != nil {
		return x.DestinationChainAddress
	}
	return ""
}

func (x *Ics20Withdrawal) GetReturnAddress() *v1alpha1.Address {
	if x != nil {
		return x.ReturnAddress
	}
	return nil
}

func (x *Ics20Withdrawal) GetTimeoutHeight() uint64 {
	if x != nil {
		return x.TimeoutHeight
	}
	return 0
}

func (x *Ics20Withdrawal) GetTimeoutTime() uint64 {
	if x != nil {
		return x.TimeoutTime
	}
	return 0
}

func (x *Ics20Withdrawal) GetSourcePort() string {
	if x != nil {
		return x.SourcePort
	}
	return ""
}

func (x *Ics20Withdrawal) GetSourceChannel() string {
	if x != nil {
		return x.SourceChannel
	}
	return ""
}

type ClientData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId        string     `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientState     *anypb.Any `protobuf:"bytes,2,opt,name=client_state,json=clientState,proto3" json:"client_state,omitempty"` // NOTE: left as Any to allow us to add more client types later
	ProcessedTime   string     `protobuf:"bytes,3,opt,name=processed_time,json=processedTime,proto3" json:"processed_time,omitempty"`
	ProcessedHeight uint64     `protobuf:"varint,4,opt,name=processed_height,json=processedHeight,proto3" json:"processed_height,omitempty"`
}

func (x *ClientData) Reset() {
	*x = ClientData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientData) ProtoMessage() {}

func (x *ClientData) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientData.ProtoReflect.Descriptor instead.
func (*ClientData) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{3}
}

func (x *ClientData) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *ClientData) GetClientState() *anypb.Any {
	if x != nil {
		return x.ClientState
	}
	return nil
}

func (x *ClientData) GetProcessedTime() string {
	if x != nil {
		return x.ProcessedTime
	}
	return ""
}

func (x *ClientData) GetProcessedHeight() uint64 {
	if x != nil {
		return x.ProcessedHeight
	}
	return 0
}

type ClientCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter uint64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *ClientCounter) Reset() {
	*x = ClientCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientCounter) ProtoMessage() {}

func (x *ClientCounter) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientCounter.ProtoReflect.Descriptor instead.
func (*ClientCounter) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{4}
}

func (x *ClientCounter) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type ConsensusState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsensusState *anypb.Any `protobuf:"bytes,1,opt,name=consensus_state,json=consensusState,proto3" json:"consensus_state,omitempty"`
}

func (x *ConsensusState) Reset() {
	*x = ConsensusState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConsensusState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsensusState) ProtoMessage() {}

func (x *ConsensusState) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsensusState.ProtoReflect.Descriptor instead.
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{5}
}

func (x *ConsensusState) GetConsensusState() *anypb.Any {
	if x != nil {
		return x.ConsensusState
	}
	return nil
}

type VerifiedHeights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Heights []*types.Height `protobuf:"bytes,1,rep,name=heights,proto3" json:"heights,omitempty"`
}

func (x *VerifiedHeights) Reset() {
	*x = VerifiedHeights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifiedHeights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifiedHeights) ProtoMessage() {}

func (x *VerifiedHeights) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifiedHeights.ProtoReflect.Descriptor instead.
func (*VerifiedHeights) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{6}
}

func (x *VerifiedHeights) GetHeights() []*types.Height {
	if x != nil {
		return x.Heights
	}
	return nil
}

type ConnectionCounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Counter uint64 `protobuf:"varint,1,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *ConnectionCounter) Reset() {
	*x = ConnectionCounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionCounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionCounter) ProtoMessage() {}

func (x *ConnectionCounter) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionCounter.ProtoReflect.Descriptor instead.
func (*ConnectionCounter) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{7}
}

func (x *ConnectionCounter) GetCounter() uint64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

type ClientConnections struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connections []string `protobuf:"bytes,1,rep,name=connections,proto3" json:"connections,omitempty"`
}

func (x *ClientConnections) Reset() {
	*x = ClientConnections{}
	if protoimpl.UnsafeEnabled {
		mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientConnections) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientConnections) ProtoMessage() {}

func (x *ClientConnections) ProtoReflect() protoreflect.Message {
	mi := &file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientConnections.ProtoReflect.Descriptor instead.
func (*ClientConnections) Descriptor() ([]byte, []int) {
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP(), []int{8}
}

func (x *ClientConnections) GetConnections() []string {
	if x != nil {
		return x.Connections
	}
	return nil
}

var File_penumbra_core_ibc_v1alpha1_ibc_proto protoreflect.FileDescriptor

var file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDesc = []byte{
	0x0a, 0x24, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f,
	0x69, 0x62, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x69, 0x62, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x62, 0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x2a, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x69, 0x62, 0x63, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x09, 0x49, 0x62,
	0x63, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x33, 0x0a, 0x0a, 0x72, 0x61, 0x77, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x09, 0x72, 0x61, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x7b, 0x0a, 0x17,
	0x46, 0x75, 0x6e, 0x67, 0x69, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x63,
	0x6b, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x22, 0xa9, 0x03, 0x0a, 0x0f, 0x49, 0x63,
	0x73, 0x32, 0x30, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12, 0x3d, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3a, 0x0a, 0x05,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x65,
	0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6e, 0x6f,
	0x6d, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x19, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x64, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x4d, 0x0a, 0x0e, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70,
	0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x52, 0x0d, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x22, 0xb4, 0x01, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x37, 0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0b, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72,
	0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x29, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x5f, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x29, 0x0a, 0x0d,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x3d, 0x0a, 0x0f, 0x63, 0x6f, 0x6e,
	0x73, 0x65, 0x6e, 0x73, 0x75, 0x73, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x6e,
	0x73, 0x75, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x47, 0x0a, 0x0f, 0x56, 0x65, 0x72, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x12, 0x34, 0x0a, 0x07, 0x68,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x69,
	0x62, 0x63, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x52, 0x07, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x73, 0x22, 0x2d, 0x0a, 0x11, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x22, 0x35, 0x0a, 0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x8c, 0x02, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x69, 0x62,
	0x63, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x49, 0x62, 0x63, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x55, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2d, 0x7a, 0x6f, 0x6e, 0x65,
	0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x69, 0x62, 0x63, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x69, 0x62, 0x63, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03,
	0x50, 0x43, 0x49, 0xaa, 0x02, 0x1a, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x2e, 0x43,
	0x6f, 0x72, 0x65, 0x2e, 0x49, 0x62, 0x63, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xca, 0x02, 0x1a, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65,
	0x5c, 0x49, 0x62, 0x63, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x26,
	0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72, 0x61, 0x5c, 0x43, 0x6f, 0x72, 0x65, 0x5c, 0x49, 0x62,
	0x63, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1d, 0x50, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x72,
	0x61, 0x3a, 0x3a, 0x43, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x49, 0x62, 0x63, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescOnce sync.Once
	file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescData = file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDesc
)

func file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescGZIP() []byte {
	file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescOnce.Do(func() {
		file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescData = protoimpl.X.CompressGZIP(file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescData)
	})
	return file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDescData
}

var file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_penumbra_core_ibc_v1alpha1_ibc_proto_goTypes = []interface{}{
	(*IbcAction)(nil),               // 0: penumbra.core.ibc.v1alpha1.IbcAction
	(*FungibleTokenPacketData)(nil), // 1: penumbra.core.ibc.v1alpha1.FungibleTokenPacketData
	(*Ics20Withdrawal)(nil),         // 2: penumbra.core.ibc.v1alpha1.Ics20Withdrawal
	(*ClientData)(nil),              // 3: penumbra.core.ibc.v1alpha1.ClientData
	(*ClientCounter)(nil),           // 4: penumbra.core.ibc.v1alpha1.ClientCounter
	(*ConsensusState)(nil),          // 5: penumbra.core.ibc.v1alpha1.ConsensusState
	(*VerifiedHeights)(nil),         // 6: penumbra.core.ibc.v1alpha1.VerifiedHeights
	(*ConnectionCounter)(nil),       // 7: penumbra.core.ibc.v1alpha1.ConnectionCounter
	(*ClientConnections)(nil),       // 8: penumbra.core.ibc.v1alpha1.ClientConnections
	(*anypb.Any)(nil),               // 9: google.protobuf.Any
	(*v1alpha1.Amount)(nil),         // 10: penumbra.core.crypto.v1alpha1.Amount
	(*v1alpha1.Denom)(nil),          // 11: penumbra.core.crypto.v1alpha1.Denom
	(*v1alpha1.Address)(nil),        // 12: penumbra.core.crypto.v1alpha1.Address
	(*types.Height)(nil),            // 13: ibc.core.client.v1.Height
}
var file_penumbra_core_ibc_v1alpha1_ibc_proto_depIdxs = []int32{
	9,  // 0: penumbra.core.ibc.v1alpha1.IbcAction.raw_action:type_name -> google.protobuf.Any
	10, // 1: penumbra.core.ibc.v1alpha1.Ics20Withdrawal.amount:type_name -> penumbra.core.crypto.v1alpha1.Amount
	11, // 2: penumbra.core.ibc.v1alpha1.Ics20Withdrawal.denom:type_name -> penumbra.core.crypto.v1alpha1.Denom
	12, // 3: penumbra.core.ibc.v1alpha1.Ics20Withdrawal.return_address:type_name -> penumbra.core.crypto.v1alpha1.Address
	9,  // 4: penumbra.core.ibc.v1alpha1.ClientData.client_state:type_name -> google.protobuf.Any
	9,  // 5: penumbra.core.ibc.v1alpha1.ConsensusState.consensus_state:type_name -> google.protobuf.Any
	13, // 6: penumbra.core.ibc.v1alpha1.VerifiedHeights.heights:type_name -> ibc.core.client.v1.Height
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_penumbra_core_ibc_v1alpha1_ibc_proto_init() }
func file_penumbra_core_ibc_v1alpha1_ibc_proto_init() {
	if File_penumbra_core_ibc_v1alpha1_ibc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IbcAction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FungibleTokenPacketData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ics20Withdrawal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientCounter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConsensusState); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifiedHeights); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionCounter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientConnections); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_penumbra_core_ibc_v1alpha1_ibc_proto_goTypes,
		DependencyIndexes: file_penumbra_core_ibc_v1alpha1_ibc_proto_depIdxs,
		MessageInfos:      file_penumbra_core_ibc_v1alpha1_ibc_proto_msgTypes,
	}.Build()
	File_penumbra_core_ibc_v1alpha1_ibc_proto = out.File
	file_penumbra_core_ibc_v1alpha1_ibc_proto_rawDesc = nil
	file_penumbra_core_ibc_v1alpha1_ibc_proto_goTypes = nil
	file_penumbra_core_ibc_v1alpha1_ibc_proto_depIdxs = nil
}
