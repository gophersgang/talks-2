// Code generated by protoc-gen-gogo.
// source: term.proto
// DO NOT EDIT!

/*
	Package main is a generated protocol buffer package.

	It is generated from these files:
		term.proto

	It has these top-level messages:
		Term
		Rev
*/
package main

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto/gogo.pb"

import io "io"
import math1 "math"
import fmt "fmt"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"

import math2 "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Term struct {
	TermStr          *string         `protobuf:"bytes,1,req" json:"TermStr,omitempty"`
	NumDocuments     *uint32         `protobuf:"varint,2,req" json:"NumDocuments,omitempty"`
	NumWords         *int32          `protobuf:"varint,3,req" json:"NumWords,omitempty"`
	InteractionsPos  *int32          `protobuf:"varint,4,opt" json:"InteractionsPos,omitempty"`
	InteractionsNeg  *int32          `protobuf:"varint,5,opt" json:"InteractionsNeg,omitempty"`
	HardcodedScore   *int32          `protobuf:"varint,6,opt" json:"HardcodedScore,omitempty"`
	Infogain         *float32        `protobuf:"fixed32,7,opt" json:"Infogain,omitempty"`
	Shotguns         []*Term_Shotgun `protobuf:"bytes,8,rep,name=shotguns" json:"shotguns,omitempty"`
	Clues            []*Term_Clue    `protobuf:"bytes,9,rep,name=clues" json:"clues,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *Term) Reset()         { *m = Term{} }
func (m *Term) String() string { return proto.CompactTextString(m) }
func (*Term) ProtoMessage()    {}

func (m *Term) GetTermStr() string {
	if m != nil && m.TermStr != nil {
		return *m.TermStr
	}
	return ""
}

func (m *Term) GetNumDocuments() uint32 {
	if m != nil && m.NumDocuments != nil {
		return *m.NumDocuments
	}
	return 0
}

func (m *Term) GetNumWords() int32 {
	if m != nil && m.NumWords != nil {
		return *m.NumWords
	}
	return 0
}

func (m *Term) GetInteractionsPos() int32 {
	if m != nil && m.InteractionsPos != nil {
		return *m.InteractionsPos
	}
	return 0
}

func (m *Term) GetInteractionsNeg() int32 {
	if m != nil && m.InteractionsNeg != nil {
		return *m.InteractionsNeg
	}
	return 0
}

func (m *Term) GetHardcodedScore() int32 {
	if m != nil && m.HardcodedScore != nil {
		return *m.HardcodedScore
	}
	return 0
}

func (m *Term) GetInfogain() float32 {
	if m != nil && m.Infogain != nil {
		return *m.Infogain
	}
	return 0
}

func (m *Term) GetShotguns() []*Term_Shotgun {
	if m != nil {
		return m.Shotguns
	}
	return nil
}

func (m *Term) GetClues() []*Term_Clue {
	if m != nil {
		return m.Clues
	}
	return nil
}

type Term_Shotgun struct {
	TermStr          *string  `protobuf:"bytes,1,req" json:"TermStr,omitempty"`
	Potency          *float32 `protobuf:"fixed32,2,req,def=1" json:"Potency,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Term_Shotgun) Reset()         { *m = Term_Shotgun{} }
func (m *Term_Shotgun) String() string { return proto.CompactTextString(m) }
func (*Term_Shotgun) ProtoMessage()    {}

const Default_Term_Shotgun_Potency float32 = 1

func (m *Term_Shotgun) GetTermStr() string {
	if m != nil && m.TermStr != nil {
		return *m.TermStr
	}
	return ""
}

func (m *Term_Shotgun) GetPotency() float32 {
	if m != nil && m.Potency != nil {
		return *m.Potency
	}
	return Default_Term_Shotgun_Potency
}

type Term_Clue struct {
	TermStr          *string  `protobuf:"bytes,1,req" json:"TermStr,omitempty"`
	ClueStr          *string  `protobuf:"bytes,2,req" json:"ClueStr,omitempty"`
	Potency          *float32 `protobuf:"fixed32,3,req,def=1" json:"Potency,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Term_Clue) Reset()         { *m = Term_Clue{} }
func (m *Term_Clue) String() string { return proto.CompactTextString(m) }
func (*Term_Clue) ProtoMessage()    {}

const Default_Term_Clue_Potency float32 = 1

func (m *Term_Clue) GetTermStr() string {
	if m != nil && m.TermStr != nil {
		return *m.TermStr
	}
	return ""
}

func (m *Term_Clue) GetClueStr() string {
	if m != nil && m.ClueStr != nil {
		return *m.ClueStr
	}
	return ""
}

func (m *Term_Clue) GetPotency() float32 {
	if m != nil && m.Potency != nil {
		return *m.Potency
	}
	return Default_Term_Clue_Potency
}

type Rev struct {
	DocumentId       *int32 `protobuf:"varint,1,req" json:"DocumentId,omitempty"`
	Rank             *int32 `protobuf:"varint,2,req" json:"Rank,omitempty"`
	InMeta           *int32 `protobuf:"varint,3,req" json:"InMeta,omitempty"`
	Next             *int32 `protobuf:"varint,4,req" json:"Next,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Rev) Reset()         { *m = Rev{} }
func (m *Rev) String() string { return proto.CompactTextString(m) }
func (*Rev) ProtoMessage()    {}

func (m *Rev) GetDocumentId() int32 {
	if m != nil && m.DocumentId != nil {
		return *m.DocumentId
	}
	return 0
}

func (m *Rev) GetRank() int32 {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return 0
}

func (m *Rev) GetInMeta() int32 {
	if m != nil && m.InMeta != nil {
		return *m.InMeta
	}
	return 0
}

func (m *Rev) GetNext() int32 {
	if m != nil && m.Next != nil {
		return *m.Next
	}
	return 0
}

func init() {
}
func (m *Term) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TermStr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.TermStr = &s
			index = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDocuments", wireType)
			}
			var v uint32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NumDocuments = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumWords", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.NumWords = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InteractionsPos", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InteractionsPos = &v
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InteractionsNeg", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InteractionsNeg = &v
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HardcodedScore", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HardcodedScore = &v
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Infogain", wireType)
			}
			var v uint32
			i := index + 4
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint32(data[i-4])
			v |= uint32(data[i-3]) << 8
			v |= uint32(data[i-2]) << 16
			v |= uint32(data[i-1]) << 24
			v2 := math1.Float32frombits(v)
			m.Infogain = &v2
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shotguns", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Shotguns = append(m.Shotguns, &Term_Shotgun{})
			m.Shotguns[len(m.Shotguns)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clues", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clues = append(m.Clues, &Term_Clue{})
			m.Clues[len(m.Clues)-1].Unmarshal(data[index:postIndex])
			index = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Term_Shotgun) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TermStr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.TermStr = &s
			index = postIndex
		case 2:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Potency", wireType)
			}
			var v uint32
			i := index + 4
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint32(data[i-4])
			v |= uint32(data[i-3]) << 8
			v |= uint32(data[i-2]) << 16
			v |= uint32(data[i-1]) << 24
			v2 := math1.Float32frombits(v)
			m.Potency = &v2
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Term_Clue) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TermStr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.TermStr = &s
			index = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClueStr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := index + int(stringLen)
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(data[index:postIndex])
			m.ClueStr = &s
			index = postIndex
		case 3:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Potency", wireType)
			}
			var v uint32
			i := index + 4
			if i > l {
				return io.ErrUnexpectedEOF
			}
			index = i
			v = uint32(data[i-4])
			v |= uint32(data[i-3]) << 8
			v |= uint32(data[i-2]) << 16
			v |= uint32(data[i-1]) << 24
			v2 := math1.Float32frombits(v)
			m.Potency = &v2
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Rev) Unmarshal(data []byte) error {
	l := len(data)
	index := 0
	for index < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if index >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[index]
			index++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DocumentId", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DocumentId = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rank", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Rank = &v
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field InMeta", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.InMeta = &v
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Next", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if index >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[index]
				index++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Next = &v
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			index -= sizeOfWire
			skippy, err := github_com_gogo_protobuf_proto.Skip(data[index:])
			if err != nil {
				return err
			}
			if (index + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, data[index:index+skippy]...)
			index += skippy
		}
	}
	return nil
}
func (m *Term) Size() (n int) {
	var l int
	_ = l
	if m.TermStr != nil {
		l = len(*m.TermStr)
		n += 1 + l + sovTerm(uint64(l))
	}
	if m.NumDocuments != nil {
		n += 1 + sovTerm(uint64(*m.NumDocuments))
	}
	if m.NumWords != nil {
		n += 1 + sovTerm(uint64(*m.NumWords))
	}
	if m.InteractionsPos != nil {
		n += 1 + sovTerm(uint64(*m.InteractionsPos))
	}
	if m.InteractionsNeg != nil {
		n += 1 + sovTerm(uint64(*m.InteractionsNeg))
	}
	if m.HardcodedScore != nil {
		n += 1 + sovTerm(uint64(*m.HardcodedScore))
	}
	if m.Infogain != nil {
		n += 5
	}
	if len(m.Shotguns) > 0 {
		for _, e := range m.Shotguns {
			l = e.Size()
			n += 1 + l + sovTerm(uint64(l))
		}
	}
	if len(m.Clues) > 0 {
		for _, e := range m.Clues {
			l = e.Size()
			n += 1 + l + sovTerm(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Term_Shotgun) Size() (n int) {
	var l int
	_ = l
	if m.TermStr != nil {
		l = len(*m.TermStr)
		n += 1 + l + sovTerm(uint64(l))
	}
	if m.Potency != nil {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Term_Clue) Size() (n int) {
	var l int
	_ = l
	if m.TermStr != nil {
		l = len(*m.TermStr)
		n += 1 + l + sovTerm(uint64(l))
	}
	if m.ClueStr != nil {
		l = len(*m.ClueStr)
		n += 1 + l + sovTerm(uint64(l))
	}
	if m.Potency != nil {
		n += 5
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Rev) Size() (n int) {
	var l int
	_ = l
	if m.DocumentId != nil {
		n += 1 + sovTerm(uint64(*m.DocumentId))
	}
	if m.Rank != nil {
		n += 1 + sovTerm(uint64(*m.Rank))
	}
	if m.InMeta != nil {
		n += 1 + sovTerm(uint64(*m.InMeta))
	}
	if m.Next != nil {
		n += 1 + sovTerm(uint64(*m.Next))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTerm(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTerm(x uint64) (n int) {
	return sovTerm(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Term) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Term) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TermStr != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTerm(data, i, uint64(len(*m.TermStr)))
		i += copy(data[i:], *m.TermStr)
	}
	if m.NumDocuments != nil {
		data[i] = 0x10
		i++
		i = encodeVarintTerm(data, i, uint64(*m.NumDocuments))
	}
	if m.NumWords != nil {
		data[i] = 0x18
		i++
		i = encodeVarintTerm(data, i, uint64(*m.NumWords))
	}
	if m.InteractionsPos != nil {
		data[i] = 0x20
		i++
		i = encodeVarintTerm(data, i, uint64(*m.InteractionsPos))
	}
	if m.InteractionsNeg != nil {
		data[i] = 0x28
		i++
		i = encodeVarintTerm(data, i, uint64(*m.InteractionsNeg))
	}
	if m.HardcodedScore != nil {
		data[i] = 0x30
		i++
		i = encodeVarintTerm(data, i, uint64(*m.HardcodedScore))
	}
	if m.Infogain != nil {
		data[i] = 0x3d
		i++
		i = encodeFixed32Term(data, i, uint32(math2.Float32bits(*m.Infogain)))
	}
	if len(m.Shotguns) > 0 {
		for _, msg := range m.Shotguns {
			data[i] = 0x42
			i++
			i = encodeVarintTerm(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Clues) > 0 {
		for _, msg := range m.Clues {
			data[i] = 0x4a
			i++
			i = encodeVarintTerm(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Term_Shotgun) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Term_Shotgun) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TermStr != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTerm(data, i, uint64(len(*m.TermStr)))
		i += copy(data[i:], *m.TermStr)
	}
	if m.Potency != nil {
		data[i] = 0x15
		i++
		i = encodeFixed32Term(data, i, uint32(math2.Float32bits(*m.Potency)))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Term_Clue) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Term_Clue) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.TermStr != nil {
		data[i] = 0xa
		i++
		i = encodeVarintTerm(data, i, uint64(len(*m.TermStr)))
		i += copy(data[i:], *m.TermStr)
	}
	if m.ClueStr != nil {
		data[i] = 0x12
		i++
		i = encodeVarintTerm(data, i, uint64(len(*m.ClueStr)))
		i += copy(data[i:], *m.ClueStr)
	}
	if m.Potency != nil {
		data[i] = 0x1d
		i++
		i = encodeFixed32Term(data, i, uint32(math2.Float32bits(*m.Potency)))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Rev) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Rev) MarshalTo(data []byte) (n int, err error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DocumentId != nil {
		data[i] = 0x8
		i++
		i = encodeVarintTerm(data, i, uint64(*m.DocumentId))
	}
	if m.Rank != nil {
		data[i] = 0x10
		i++
		i = encodeVarintTerm(data, i, uint64(*m.Rank))
	}
	if m.InMeta != nil {
		data[i] = 0x18
		i++
		i = encodeVarintTerm(data, i, uint64(*m.InMeta))
	}
	if m.Next != nil {
		data[i] = 0x20
		i++
		i = encodeVarintTerm(data, i, uint64(*m.Next))
	}
	if m.XXX_unrecognized != nil {
		i += copy(data[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Term(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Term(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintTerm(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
