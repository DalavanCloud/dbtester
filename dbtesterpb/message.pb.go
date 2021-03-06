// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dbtesterpb/message.proto

package dbtesterpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Operation int32

const (
	Operation_Start     Operation = 0
	Operation_Stop      Operation = 1
	Operation_Heartbeat Operation = 2
)

var Operation_name = map[int32]string{
	0: "Start",
	1: "Stop",
	2: "Heartbeat",
}
var Operation_value = map[string]int32{
	"Start":     0,
	"Stop":      1,
	"Heartbeat": 2,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}
func (Operation) EnumDescriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Request struct {
	Operation        Operation  `protobuf:"varint,1,opt,name=Operation,proto3,enum=dbtesterpb.Operation" json:"Operation,omitempty"`
	TriggerLogUpload bool       `protobuf:"varint,2,opt,name=TriggerLogUpload,proto3" json:"TriggerLogUpload,omitempty"`
	DatabaseID       DatabaseID `protobuf:"varint,3,opt,name=DatabaseID,proto3,enum=dbtesterpb.DatabaseID" json:"DatabaseID,omitempty"`
	DatabaseTag      string     `protobuf:"bytes,4,opt,name=DatabaseTag,proto3" json:"DatabaseTag,omitempty"`
	// PeerIPsString encodes a list of endpoints in string
	// because Protocol Buffer does not have a list or array datatype
	// which is ordered. 'repeated' does not guarantee the ordering.
	PeerIPsString              string                      `protobuf:"bytes,5,opt,name=PeerIPsString,proto3" json:"PeerIPsString,omitempty"`
	IPIndex                    uint32                      `protobuf:"varint,6,opt,name=IPIndex,proto3" json:"IPIndex,omitempty"`
	CurrentClientNumber        int64                       `protobuf:"varint,7,opt,name=CurrentClientNumber,proto3" json:"CurrentClientNumber,omitempty"`
	ConfigClientMachineInitial *ConfigClientMachineInitial `protobuf:"bytes,8,opt,name=ConfigClientMachineInitial" json:"ConfigClientMachineInitial,omitempty"`
	Flag_Etcd_Other            *Flag_Etcd_Other            `protobuf:"bytes,100,opt,name=flag__etcd__other,json=flagEtcdOther" json:"flag__etcd__other,omitempty"`
	Flag_Etcd_Tip              *Flag_Etcd_Tip              `protobuf:"bytes,101,opt,name=flag__etcd__tip,json=flagEtcdTip" json:"flag__etcd__tip,omitempty"`
	Flag_Etcd_V3_2             *Flag_Etcd_V3_2             `protobuf:"bytes,102,opt,name=flag__etcd__v3_2,json=flagEtcdV32" json:"flag__etcd__v3_2,omitempty"`
	Flag_Etcd_V3_3             *Flag_Etcd_V3_3             `protobuf:"bytes,103,opt,name=flag__etcd__v3_3,json=flagEtcdV33" json:"flag__etcd__v3_3,omitempty"`
	Flag_Zookeeper_R3_5_3Beta  *Flag_Zookeeper_R3_5_3Beta  `protobuf:"bytes,200,opt,name=flag__zookeeper__r3_5_3_beta,json=flagZookeeperR353Beta" json:"flag__zookeeper__r3_5_3_beta,omitempty"`
	Flag_Consul_V1_0_2         *Flag_Consul_V1_0_2         `protobuf:"bytes,300,opt,name=flag__consul__v1_0_2,json=flagConsulV102" json:"flag__consul__v1_0_2,omitempty"`
	Flag_Cetcd_Beta            *Flag_Cetcd_Beta            `protobuf:"bytes,400,opt,name=flag__cetcd__beta,json=flagCetcdBeta" json:"flag__cetcd__beta,omitempty"`
	Flag_Zetcd_Beta            *Flag_Zetcd_Beta            `protobuf:"bytes,500,opt,name=flag__zetcd__beta,json=flagZetcdBeta" json:"flag__zetcd__beta,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{0} }

type Response struct {
	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	// DiskSpaceUsageBytes is the data size of the database on disk in bytes.
	// It measures after database is requested to stop.
	DiskSpaceUsageBytes int64 `protobuf:"varint,2,opt,name=DiskSpaceUsageBytes,proto3" json:"DiskSpaceUsageBytes,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptorMessage, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "dbtesterpb.Request")
	proto.RegisterType((*Response)(nil), "dbtesterpb.Response")
	proto.RegisterEnum("dbtesterpb.Operation", Operation_name, Operation_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Transporter service

type TransporterClient interface {
	Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type transporterClient struct {
	cc *grpc.ClientConn
}

func NewTransporterClient(cc *grpc.ClientConn) TransporterClient {
	return &transporterClient{cc}
}

func (c *transporterClient) Transfer(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/dbtesterpb.Transporter/Transfer", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Transporter service

type TransporterServer interface {
	Transfer(context.Context, *Request) (*Response, error)
}

func RegisterTransporterServer(s *grpc.Server, srv TransporterServer) {
	s.RegisterService(&_Transporter_serviceDesc, srv)
}

func _Transporter_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransporterServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dbtesterpb.Transporter/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransporterServer).Transfer(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transporter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dbtesterpb.Transporter",
	HandlerType: (*TransporterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Transporter_Transfer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dbtesterpb/message.proto",
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Operation != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Operation))
	}
	if m.TriggerLogUpload {
		dAtA[i] = 0x10
		i++
		if m.TriggerLogUpload {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DatabaseID != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.DatabaseID))
	}
	if len(m.DatabaseTag) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.DatabaseTag)))
		i += copy(dAtA[i:], m.DatabaseTag)
	}
	if len(m.PeerIPsString) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintMessage(dAtA, i, uint64(len(m.PeerIPsString)))
		i += copy(dAtA[i:], m.PeerIPsString)
	}
	if m.IPIndex != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.IPIndex))
	}
	if m.CurrentClientNumber != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.CurrentClientNumber))
	}
	if m.ConfigClientMachineInitial != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.ConfigClientMachineInitial.Size()))
		n1, err := m.ConfigClientMachineInitial.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Flag_Etcd_Other != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Etcd_Other.Size()))
		n2, err := m.Flag_Etcd_Other.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.Flag_Etcd_Tip != nil {
		dAtA[i] = 0xaa
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Etcd_Tip.Size()))
		n3, err := m.Flag_Etcd_Tip.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.Flag_Etcd_V3_2 != nil {
		dAtA[i] = 0xb2
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Etcd_V3_2.Size()))
		n4, err := m.Flag_Etcd_V3_2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Flag_Etcd_V3_3 != nil {
		dAtA[i] = 0xba
		i++
		dAtA[i] = 0x6
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Etcd_V3_3.Size()))
		n5, err := m.Flag_Etcd_V3_3.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Flag_Zookeeper_R3_5_3Beta != nil {
		dAtA[i] = 0xc2
		i++
		dAtA[i] = 0xc
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Zookeeper_R3_5_3Beta.Size()))
		n6, err := m.Flag_Zookeeper_R3_5_3Beta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	if m.Flag_Consul_V1_0_2 != nil {
		dAtA[i] = 0xe2
		i++
		dAtA[i] = 0x12
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Consul_V1_0_2.Size()))
		n7, err := m.Flag_Consul_V1_0_2.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n7
	}
	if m.Flag_Cetcd_Beta != nil {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x19
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Cetcd_Beta.Size()))
		n8, err := m.Flag_Cetcd_Beta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n8
	}
	if m.Flag_Zetcd_Beta != nil {
		dAtA[i] = 0xa2
		i++
		dAtA[i] = 0x1f
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.Flag_Zetcd_Beta.Size()))
		n9, err := m.Flag_Zetcd_Beta.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func (m *Response) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Response) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Success {
		dAtA[i] = 0x8
		i++
		if m.Success {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.DiskSpaceUsageBytes != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintMessage(dAtA, i, uint64(m.DiskSpaceUsageBytes))
	}
	return i, nil
}

func encodeVarintMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Request) Size() (n int) {
	var l int
	_ = l
	if m.Operation != 0 {
		n += 1 + sovMessage(uint64(m.Operation))
	}
	if m.TriggerLogUpload {
		n += 2
	}
	if m.DatabaseID != 0 {
		n += 1 + sovMessage(uint64(m.DatabaseID))
	}
	l = len(m.DatabaseTag)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	l = len(m.PeerIPsString)
	if l > 0 {
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.IPIndex != 0 {
		n += 1 + sovMessage(uint64(m.IPIndex))
	}
	if m.CurrentClientNumber != 0 {
		n += 1 + sovMessage(uint64(m.CurrentClientNumber))
	}
	if m.ConfigClientMachineInitial != nil {
		l = m.ConfigClientMachineInitial.Size()
		n += 1 + l + sovMessage(uint64(l))
	}
	if m.Flag_Etcd_Other != nil {
		l = m.Flag_Etcd_Other.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Etcd_Tip != nil {
		l = m.Flag_Etcd_Tip.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Etcd_V3_2 != nil {
		l = m.Flag_Etcd_V3_2.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Etcd_V3_3 != nil {
		l = m.Flag_Etcd_V3_3.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Zookeeper_R3_5_3Beta != nil {
		l = m.Flag_Zookeeper_R3_5_3Beta.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Consul_V1_0_2 != nil {
		l = m.Flag_Consul_V1_0_2.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Cetcd_Beta != nil {
		l = m.Flag_Cetcd_Beta.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	if m.Flag_Zetcd_Beta != nil {
		l = m.Flag_Zetcd_Beta.Size()
		n += 2 + l + sovMessage(uint64(l))
	}
	return n
}

func (m *Response) Size() (n int) {
	var l int
	_ = l
	if m.Success {
		n += 2
	}
	if m.DiskSpaceUsageBytes != 0 {
		n += 1 + sovMessage(uint64(m.DiskSpaceUsageBytes))
	}
	return n
}

func sovMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMessage(x uint64) (n int) {
	return sovMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operation", wireType)
			}
			m.Operation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operation |= (Operation(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TriggerLogUpload", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.TriggerLogUpload = bool(v != 0)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatabaseID", wireType)
			}
			m.DatabaseID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DatabaseID |= (DatabaseID(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatabaseTag", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatabaseTag = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerIPsString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerIPsString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPIndex", wireType)
			}
			m.IPIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IPIndex |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CurrentClientNumber", wireType)
			}
			m.CurrentClientNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CurrentClientNumber |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigClientMachineInitial", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ConfigClientMachineInitial == nil {
				m.ConfigClientMachineInitial = &ConfigClientMachineInitial{}
			}
			if err := m.ConfigClientMachineInitial.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 100:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Etcd_Other", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Etcd_Other == nil {
				m.Flag_Etcd_Other = &Flag_Etcd_Other{}
			}
			if err := m.Flag_Etcd_Other.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 101:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Etcd_Tip", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Etcd_Tip == nil {
				m.Flag_Etcd_Tip = &Flag_Etcd_Tip{}
			}
			if err := m.Flag_Etcd_Tip.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 102:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Etcd_V3_2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Etcd_V3_2 == nil {
				m.Flag_Etcd_V3_2 = &Flag_Etcd_V3_2{}
			}
			if err := m.Flag_Etcd_V3_2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 103:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Etcd_V3_3", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Etcd_V3_3 == nil {
				m.Flag_Etcd_V3_3 = &Flag_Etcd_V3_3{}
			}
			if err := m.Flag_Etcd_V3_3.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 200:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Zookeeper_R3_5_3Beta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Zookeeper_R3_5_3Beta == nil {
				m.Flag_Zookeeper_R3_5_3Beta = &Flag_Zookeeper_R3_5_3Beta{}
			}
			if err := m.Flag_Zookeeper_R3_5_3Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 300:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Consul_V1_0_2", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Consul_V1_0_2 == nil {
				m.Flag_Consul_V1_0_2 = &Flag_Consul_V1_0_2{}
			}
			if err := m.Flag_Consul_V1_0_2.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 400:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Cetcd_Beta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Cetcd_Beta == nil {
				m.Flag_Cetcd_Beta = &Flag_Cetcd_Beta{}
			}
			if err := m.Flag_Cetcd_Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 500:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Flag_Zetcd_Beta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMessage
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Flag_Zetcd_Beta == nil {
				m.Flag_Zetcd_Beta = &Flag_Zetcd_Beta{}
			}
			if err := m.Flag_Zetcd_Beta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Response) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Response: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Response: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Success", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Success = bool(v != 0)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskSpaceUsageBytes", wireType)
			}
			m.DiskSpaceUsageBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DiskSpaceUsageBytes |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("dbtesterpb/message.proto", fileDescriptorMessage) }

var fileDescriptorMessage = []byte{
	// 712 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0xdf, 0x4e, 0x1b, 0x47,
	0x14, 0xc6, 0xbd, 0x98, 0x3f, 0xf6, 0x58, 0xa6, 0xee, 0x00, 0xd5, 0xc8, 0x50, 0x77, 0x85, 0x2a,
	0x64, 0x21, 0xd5, 0x06, 0xaf, 0x68, 0xaf, 0x8b, 0x69, 0x8b, 0xa5, 0x26, 0xa0, 0xb1, 0xe1, 0x82,
	0x9b, 0xd1, 0xec, 0xfa, 0x78, 0x59, 0x61, 0xef, 0x6c, 0x66, 0xc7, 0x28, 0xe1, 0x2e, 0x6f, 0x90,
	0xcb, 0x3c, 0x44, 0x1e, 0x84, 0xcb, 0x3c, 0x42, 0x42, 0x5e, 0x21, 0x0f, 0x10, 0xed, 0xac, 0x17,
	0x0f, 0xd8, 0x4e, 0xee, 0x7c, 0xbe, 0xef, 0x9b, 0xdf, 0xce, 0x1f, 0x9f, 0x83, 0x48, 0xdf, 0x55,
	0x10, 0x2b, 0x90, 0x91, 0xdb, 0x1c, 0x41, 0x1c, 0x73, 0x1f, 0x1a, 0x91, 0x14, 0x4a, 0x60, 0x34,
	0x75, 0xaa, 0x7f, 0xf8, 0x81, 0xba, 0x1e, 0xbb, 0x0d, 0x4f, 0x8c, 0x9a, 0xbe, 0xf0, 0x45, 0x53,
	0x47, 0xdc, 0xf1, 0x40, 0x57, 0xba, 0xd0, 0xbf, 0xd2, 0xa5, 0xd5, 0x1d, 0x03, 0xda, 0xe7, 0x8a,
	0xbb, 0x3c, 0x06, 0x16, 0xf4, 0x27, 0x6e, 0xd5, 0x70, 0x07, 0x43, 0xee, 0x33, 0x50, 0x5e, 0xe6,
	0xfd, 0xf6, 0xdc, 0xbb, 0x13, 0xe2, 0x06, 0x20, 0x02, 0x39, 0x07, 0xad, 0x03, 0x9e, 0x08, 0xe3,
	0xf1, 0x70, 0xe2, 0x6e, 0xcf, 0x2c, 0x37, 0xd8, 0x33, 0xa6, 0x67, 0x98, 0x7b, 0x86, 0xe9, 0x89,
	0x70, 0x10, 0xf8, 0xcc, 0x1b, 0x06, 0x10, 0x2a, 0x36, 0xe2, 0xde, 0x75, 0x10, 0x4e, 0x6e, 0x65,
	0xf7, 0x6d, 0x01, 0xad, 0x51, 0x78, 0x35, 0x86, 0x58, 0x61, 0x07, 0x15, 0xcf, 0x22, 0x90, 0x5c,
	0x05, 0x22, 0x24, 0x96, 0x6d, 0xd5, 0xd7, 0x5b, 0x5b, 0x8d, 0x29, 0xa7, 0xf1, 0x68, 0xd2, 0x69,
	0x0e, 0xef, 0xa3, 0x4a, 0x4f, 0x06, 0xbe, 0x0f, 0xf2, 0x7f, 0xe1, 0x5f, 0x44, 0x43, 0xc1, 0xfb,
	0x64, 0xc9, 0xb6, 0xea, 0x05, 0x3a, 0xa3, 0xe3, 0x3f, 0x11, 0x3a, 0x99, 0x5c, 0x5f, 0xe7, 0x84,
	0xe4, 0xf5, 0x17, 0x7e, 0x31, 0xbf, 0x30, 0x75, 0xa9, 0x91, 0xc4, 0x36, 0x2a, 0x65, 0x55, 0x8f,
	0xfb, 0x64, 0xd9, 0xb6, 0xea, 0x45, 0x6a, 0x4a, 0xf8, 0x77, 0x54, 0x3e, 0x07, 0x90, 0x9d, 0xf3,
	0xb8, 0xab, 0x64, 0x10, 0xfa, 0x64, 0x45, 0x67, 0x9e, 0x8a, 0x98, 0xa0, 0xb5, 0xce, 0x79, 0x27,
	0xec, 0xc3, 0x6b, 0xb2, 0x6a, 0x5b, 0xf5, 0x32, 0xcd, 0x4a, 0x7c, 0x80, 0x36, 0xda, 0x63, 0x29,
	0x21, 0x54, 0x6d, 0x7d, 0x4b, 0x2f, 0xc7, 0x23, 0x17, 0x24, 0x59, 0xb3, 0xad, 0x7a, 0x9e, 0xce,
	0xb3, 0xf0, 0x00, 0x55, 0xdb, 0xfa, 0x5e, 0x53, 0xf5, 0x45, 0x7a, 0xab, 0x9d, 0x30, 0x50, 0x01,
	0x1f, 0x92, 0x82, 0x6d, 0xd5, 0x4b, 0xad, 0x3d, 0xf3, 0x6c, 0x8b, 0xd3, 0xf4, 0x3b, 0x24, 0xfc,
	0x1f, 0xfa, 0x59, 0x3f, 0xae, 0xfe, 0x57, 0x31, 0x26, 0xd4, 0x35, 0x48, 0xd2, 0xd7, 0xf8, 0x5f,
	0x4d, 0xfc, 0x4c, 0x88, 0x96, 0x13, 0xe9, 0x1f, 0xe5, 0xf5, 0xcf, 0x92, 0x12, 0xff, 0x8d, 0x7e,
	0x32, 0x33, 0x2a, 0x88, 0x08, 0x68, 0xcc, 0xf6, 0x22, 0x8c, 0x0a, 0x22, 0x5a, 0xca, 0x20, 0xbd,
	0x20, 0xc2, 0x6d, 0x54, 0x31, 0xfd, 0x5b, 0x87, 0xb5, 0xc8, 0x40, 0x33, 0x76, 0x16, 0x31, 0x92,
	0xcc, 0x14, 0x72, 0xe9, 0xb4, 0xe6, 0x40, 0x1c, 0xe2, 0xff, 0x10, 0xe2, 0x98, 0x10, 0x07, 0x0f,
	0xd0, 0x4e, 0x1a, 0x78, 0xec, 0x27, 0xc6, 0xa4, 0xc3, 0x8e, 0x98, 0xc3, 0x5c, 0x50, 0x9c, 0xdc,
	0x5b, 0x9a, 0x58, 0x9f, 0x25, 0xce, 0x5f, 0x40, 0xb7, 0x12, 0xf7, 0x2a, 0xf3, 0xa8, 0x73, 0xe4,
	0x1c, 0x83, 0xe2, 0xf8, 0x0c, 0x6d, 0xa6, 0xcb, 0xd2, 0xb6, 0x64, 0xec, 0xf6, 0x90, 0x1d, 0xb0,
	0x16, 0xf9, 0xb0, 0xa4, 0xf9, 0xf6, 0x2c, 0xff, 0x69, 0x90, 0xae, 0x27, 0x6a, 0x5b, 0x6b, 0x97,
	0x87, 0x07, 0x2d, 0x7c, 0x9a, 0x3d, 0xa7, 0x97, 0x1e, 0x4d, 0xef, 0xf6, 0x5d, 0x7e, 0xd1, 0x7b,
	0x1a, 0xa9, 0xf4, 0x3d, 0xdb, 0x89, 0xa0, 0xb7, 0xf6, 0x48, 0xba, 0x33, 0x48, 0x5f, 0x17, 0x92,
	0xee, 0x9e, 0x93, 0xae, 0x32, 0xd2, 0xee, 0x25, 0x2a, 0x50, 0x88, 0x23, 0x11, 0xc6, 0x90, 0xb4,
	0x48, 0x77, 0xec, 0x79, 0x10, 0xc7, 0x7a, 0x02, 0x14, 0x68, 0x56, 0x26, 0x2d, 0x72, 0x12, 0xc4,
	0x37, 0xdd, 0x88, 0x7b, 0x70, 0x91, 0xcc, 0xd5, 0xe3, 0x37, 0x0a, 0x62, 0xdd, 0xeb, 0x79, 0x3a,
	0xcf, 0xda, 0x6f, 0x1a, 0xf3, 0x04, 0x17, 0xd1, 0x4a, 0x57, 0x71, 0xa9, 0x2a, 0x39, 0x5c, 0x40,
	0xcb, 0x5d, 0x25, 0xa2, 0x8a, 0x85, 0xcb, 0xa8, 0x78, 0x0a, 0x5c, 0x2a, 0x17, 0xb8, 0xaa, 0x2c,
	0xb5, 0xfe, 0x45, 0xa5, 0x9e, 0xe4, 0x61, 0x1c, 0x09, 0xa9, 0x40, 0xe2, 0xbf, 0x50, 0x41, 0x97,
	0x03, 0x90, 0x78, 0xc3, 0x3c, 0xd1, 0x64, 0x60, 0x55, 0x37, 0x9f, 0x8a, 0xe9, 0x11, 0x76, 0x73,
	0xc7, 0x9b, 0xf7, 0x9f, 0x6b, 0xb9, 0xfb, 0x87, 0x9a, 0xf5, 0xf1, 0xa1, 0x66, 0x7d, 0x7a, 0xa8,
	0x59, 0xef, 0xbf, 0xd4, 0x72, 0xee, 0xaa, 0x9e, 0x78, 0xce, 0xb7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x9b, 0x7a, 0x8c, 0xc2, 0x23, 0x06, 0x00, 0x00,
}
