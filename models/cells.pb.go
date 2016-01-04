// Code generated by protoc-gen-gogo.
// source: cells.proto
// DO NOT EDIT!

package models

import proto "github.com/gogo/protobuf/proto"
import math "math"

// discarding unused import gogoproto "github.com/gogo/protobuf/gogoproto"

import fmt "fmt"
import strings "strings"
import github_com_gogo_protobuf_proto "github.com/gogo/protobuf/proto"
import sort "sort"
import strconv "strconv"
import reflect "reflect"
import github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"

import errors "errors"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type CellCapacity struct {
	MemoryMb   int32 `protobuf:"varint,1,opt,name=memory_mb" json:"memory_mb"`
	DiskMb     int32 `protobuf:"varint,2,opt,name=disk_mb" json:"disk_mb"`
	Containers int32 `protobuf:"varint,3,opt,name=containers" json:"containers"`
}

func (m *CellCapacity) Reset()      { *m = CellCapacity{} }
func (*CellCapacity) ProtoMessage() {}

func (m *CellCapacity) GetMemoryMb() int32 {
	if m != nil {
		return m.MemoryMb
	}
	return 0
}

func (m *CellCapacity) GetDiskMb() int32 {
	if m != nil {
		return m.DiskMb
	}
	return 0
}

func (m *CellCapacity) GetContainers() int32 {
	if m != nil {
		return m.Containers
	}
	return 0
}

type Providers struct {
	Parameters []string `protobuf:"bytes,1,rep,name=parameters" json:"parameters,omitempty"`
}

func (m *Providers) Reset()      { *m = Providers{} }
func (*Providers) ProtoMessage() {}

func (m *Providers) GetParameters() []string {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type ProtoRootfsproviders struct {
	RootfsProviders map[string]*Providers `protobuf:"bytes,1,rep,name=rootfs_providers" json:"rootfs_providers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ProtoRootfsproviders) Reset()      { *m = ProtoRootfsproviders{} }
func (*ProtoRootfsproviders) ProtoMessage() {}

func (m *ProtoRootfsproviders) GetRootfsProviders() map[string]*Providers {
	if m != nil {
		return m.RootfsProviders
	}
	return nil
}

type CellPresence struct {
	CellId          string          `protobuf:"bytes,1,opt,name=cell_id" json:"cell_id"`
	RepAddress      string          `protobuf:"bytes,2,opt,name=rep_address" json:"rep_address"`
	Zone            string          `protobuf:"bytes,3,opt,name=zone" json:"zone"`
	Capacity        *CellCapacity   `protobuf:"bytes,4,opt,name=capacity" json:"capacity,omitempty"`
	RootfsProviders RootFSProviders `protobuf:"bytes,5,opt,name=rootfs_providers,customtype=RootFSProviders" json:"rootfs_providers"`
}

func (m *CellPresence) Reset()      { *m = CellPresence{} }
func (*CellPresence) ProtoMessage() {}

func (m *CellPresence) GetCellId() string {
	if m != nil {
		return m.CellId
	}
	return ""
}

func (m *CellPresence) GetRepAddress() string {
	if m != nil {
		return m.RepAddress
	}
	return ""
}

func (m *CellPresence) GetZone() string {
	if m != nil {
		return m.Zone
	}
	return ""
}

func (m *CellPresence) GetCapacity() *CellCapacity {
	if m != nil {
		return m.Capacity
	}
	return nil
}

type CellsResponse struct {
	Error *Error          `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	Cells []*CellPresence `protobuf:"bytes,2,rep,name=cells" json:"cells,omitempty"`
}

func (m *CellsResponse) Reset()      { *m = CellsResponse{} }
func (*CellsResponse) ProtoMessage() {}

func (m *CellsResponse) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *CellsResponse) GetCells() []*CellPresence {
	if m != nil {
		return m.Cells
	}
	return nil
}

func (this *CellCapacity) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CellCapacity)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.MemoryMb != that1.MemoryMb {
		return false
	}
	if this.DiskMb != that1.DiskMb {
		return false
	}
	if this.Containers != that1.Containers {
		return false
	}
	return true
}
func (this *Providers) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Providers)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.Parameters) != len(that1.Parameters) {
		return false
	}
	for i := range this.Parameters {
		if this.Parameters[i] != that1.Parameters[i] {
			return false
		}
	}
	return true
}
func (this *ProtoRootfsproviders) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*ProtoRootfsproviders)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if len(this.RootfsProviders) != len(that1.RootfsProviders) {
		return false
	}
	for i := range this.RootfsProviders {
		if !this.RootfsProviders[i].Equal(that1.RootfsProviders[i]) {
			return false
		}
	}
	return true
}
func (this *CellPresence) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CellPresence)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if this.CellId != that1.CellId {
		return false
	}
	if this.RepAddress != that1.RepAddress {
		return false
	}
	if this.Zone != that1.Zone {
		return false
	}
	if !this.Capacity.Equal(that1.Capacity) {
		return false
	}
	if !this.RootfsProviders.Equal(that1.RootfsProviders) {
		return false
	}
	return true
}
func (this *CellsResponse) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*CellsResponse)
	if !ok {
		return false
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.Error.Equal(that1.Error) {
		return false
	}
	if len(this.Cells) != len(that1.Cells) {
		return false
	}
	for i := range this.Cells {
		if !this.Cells[i].Equal(that1.Cells[i]) {
			return false
		}
	}
	return true
}
func (this *CellCapacity) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.CellCapacity{` +
		`MemoryMb:` + fmt.Sprintf("%#v", this.MemoryMb),
		`DiskMb:` + fmt.Sprintf("%#v", this.DiskMb),
		`Containers:` + fmt.Sprintf("%#v", this.Containers) + `}`}, ", ")
	return s
}
func (this *Providers) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.Providers{` +
		`Parameters:` + fmt.Sprintf("%#v", this.Parameters) + `}`}, ", ")
	return s
}
func (this *ProtoRootfsproviders) GoString() string {
	if this == nil {
		return "nil"
	}
	keysForRootfsProviders := make([]string, 0, len(this.RootfsProviders))
	for k, _ := range this.RootfsProviders {
		keysForRootfsProviders = append(keysForRootfsProviders, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForRootfsProviders)
	mapStringForRootfsProviders := "map[string]*Providers{"
	for _, k := range keysForRootfsProviders {
		mapStringForRootfsProviders += fmt.Sprintf("%#v: %#v,", k, this.RootfsProviders[k])
	}
	mapStringForRootfsProviders += "}"
	s := strings.Join([]string{`&models.ProtoRootfsproviders{` +
		`RootfsProviders:` + mapStringForRootfsProviders + `}`}, ", ")
	return s
}
func (this *CellPresence) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.CellPresence{` +
		`CellId:` + fmt.Sprintf("%#v", this.CellId),
		`RepAddress:` + fmt.Sprintf("%#v", this.RepAddress),
		`Zone:` + fmt.Sprintf("%#v", this.Zone),
		`Capacity:` + fmt.Sprintf("%#v", this.Capacity),
		`RootfsProviders:` + fmt.Sprintf("%#v", this.RootfsProviders) + `}`}, ", ")
	return s
}
func (this *CellsResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&models.CellsResponse{` +
		`Error:` + fmt.Sprintf("%#v", this.Error),
		`Cells:` + fmt.Sprintf("%#v", this.Cells) + `}`}, ", ")
	return s
}
func valueToGoStringCells(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func extensionToGoStringCells(e map[int32]github_com_gogo_protobuf_proto.Extension) string {
	if e == nil {
		return "nil"
	}
	s := "map[int32]proto.Extension{"
	keys := make([]int, 0, len(e))
	for k := range e {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)
	ss := []string{}
	for _, k := range keys {
		ss = append(ss, strconv.Itoa(k)+": "+e[int32(k)].GoString())
	}
	s += strings.Join(ss, ",") + "}"
	return s
}
func (m *CellCapacity) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CellCapacity) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0x8
	i++
	i = encodeVarintCells(data, i, uint64(m.MemoryMb))
	data[i] = 0x10
	i++
	i = encodeVarintCells(data, i, uint64(m.DiskMb))
	data[i] = 0x18
	i++
	i = encodeVarintCells(data, i, uint64(m.Containers))
	return i, nil
}

func (m *Providers) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *Providers) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Parameters) > 0 {
		for _, s := range m.Parameters {
			data[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				data[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			data[i] = uint8(l)
			i++
			i += copy(data[i:], s)
		}
	}
	return i, nil
}

func (m *ProtoRootfsproviders) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *ProtoRootfsproviders) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RootfsProviders) > 0 {
		keysForRootfsProviders := make([]string, 0, len(m.RootfsProviders))
		for k, _ := range m.RootfsProviders {
			keysForRootfsProviders = append(keysForRootfsProviders, k)
		}
		github_com_gogo_protobuf_sortkeys.Strings(keysForRootfsProviders)
		for _, k := range keysForRootfsProviders {
			data[i] = 0xa
			i++
			v := m.RootfsProviders[k]
			if v == nil {
				return 0, errors.New("proto: map has nil element")
			}
			msgSize := v.Size()
			mapSize := 1 + len(k) + sovCells(uint64(len(k))) + 1 + msgSize + sovCells(uint64(msgSize))
			i = encodeVarintCells(data, i, uint64(mapSize))
			data[i] = 0xa
			i++
			i = encodeVarintCells(data, i, uint64(len(k)))
			i += copy(data[i:], k)
			data[i] = 0x12
			i++
			i = encodeVarintCells(data, i, uint64(v.Size()))
			n1, err := v.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n1
		}
	}
	return i, nil
}

func (m *CellPresence) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CellPresence) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintCells(data, i, uint64(len(m.CellId)))
	i += copy(data[i:], m.CellId)
	data[i] = 0x12
	i++
	i = encodeVarintCells(data, i, uint64(len(m.RepAddress)))
	i += copy(data[i:], m.RepAddress)
	data[i] = 0x1a
	i++
	i = encodeVarintCells(data, i, uint64(len(m.Zone)))
	i += copy(data[i:], m.Zone)
	if m.Capacity != nil {
		data[i] = 0x22
		i++
		i = encodeVarintCells(data, i, uint64(m.Capacity.Size()))
		n2, err := m.Capacity.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	data[i] = 0x2a
	i++
	i = encodeVarintCells(data, i, uint64(m.RootfsProviders.Size()))
	n3, err := m.RootfsProviders.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *CellsResponse) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *CellsResponse) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Error != nil {
		data[i] = 0xa
		i++
		i = encodeVarintCells(data, i, uint64(m.Error.Size()))
		n4, err := m.Error.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if len(m.Cells) > 0 {
		for _, msg := range m.Cells {
			data[i] = 0x12
			i++
			i = encodeVarintCells(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeFixed64Cells(data []byte, offset int, v uint64) int {
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
func encodeFixed32Cells(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintCells(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *CellCapacity) Size() (n int) {
	var l int
	_ = l
	n += 1 + sovCells(uint64(m.MemoryMb))
	n += 1 + sovCells(uint64(m.DiskMb))
	n += 1 + sovCells(uint64(m.Containers))
	return n
}

func (m *Providers) Size() (n int) {
	var l int
	_ = l
	if len(m.Parameters) > 0 {
		for _, s := range m.Parameters {
			l = len(s)
			n += 1 + l + sovCells(uint64(l))
		}
	}
	return n
}

func (m *ProtoRootfsproviders) Size() (n int) {
	var l int
	_ = l
	if len(m.RootfsProviders) > 0 {
		for k, v := range m.RootfsProviders {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
			}
			mapEntrySize := 1 + len(k) + sovCells(uint64(len(k))) + 1 + l + sovCells(uint64(l))
			n += mapEntrySize + 1 + sovCells(uint64(mapEntrySize))
		}
	}
	return n
}

func (m *CellPresence) Size() (n int) {
	var l int
	_ = l
	l = len(m.CellId)
	n += 1 + l + sovCells(uint64(l))
	l = len(m.RepAddress)
	n += 1 + l + sovCells(uint64(l))
	l = len(m.Zone)
	n += 1 + l + sovCells(uint64(l))
	if m.Capacity != nil {
		l = m.Capacity.Size()
		n += 1 + l + sovCells(uint64(l))
	}
	l = m.RootfsProviders.Size()
	n += 1 + l + sovCells(uint64(l))
	return n
}

func (m *CellsResponse) Size() (n int) {
	var l int
	_ = l
	if m.Error != nil {
		l = m.Error.Size()
		n += 1 + l + sovCells(uint64(l))
	}
	if len(m.Cells) > 0 {
		for _, e := range m.Cells {
			l = e.Size()
			n += 1 + l + sovCells(uint64(l))
		}
	}
	return n
}

func sovCells(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCells(x uint64) (n int) {
	return sovCells(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *CellCapacity) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CellCapacity{`,
		`MemoryMb:` + fmt.Sprintf("%v", this.MemoryMb) + `,`,
		`DiskMb:` + fmt.Sprintf("%v", this.DiskMb) + `,`,
		`Containers:` + fmt.Sprintf("%v", this.Containers) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Providers) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Providers{`,
		`Parameters:` + fmt.Sprintf("%v", this.Parameters) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ProtoRootfsproviders) String() string {
	if this == nil {
		return "nil"
	}
	keysForRootfsProviders := make([]string, 0, len(this.RootfsProviders))
	for k, _ := range this.RootfsProviders {
		keysForRootfsProviders = append(keysForRootfsProviders, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForRootfsProviders)
	mapStringForRootfsProviders := "map[string]*Providers{"
	for _, k := range keysForRootfsProviders {
		mapStringForRootfsProviders += fmt.Sprintf("%v: %v,", k, this.RootfsProviders[k])
	}
	mapStringForRootfsProviders += "}"
	s := strings.Join([]string{`&ProtoRootfsproviders{`,
		`RootfsProviders:` + mapStringForRootfsProviders + `,`,
		`}`,
	}, "")
	return s
}
func (this *CellPresence) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CellPresence{`,
		`CellId:` + fmt.Sprintf("%v", this.CellId) + `,`,
		`RepAddress:` + fmt.Sprintf("%v", this.RepAddress) + `,`,
		`Zone:` + fmt.Sprintf("%v", this.Zone) + `,`,
		`Capacity:` + strings.Replace(fmt.Sprintf("%v", this.Capacity), "CellCapacity", "CellCapacity", 1) + `,`,
		`RootfsProviders:` + fmt.Sprintf("%v", this.RootfsProviders) + `,`,
		`}`,
	}, "")
	return s
}
func (this *CellsResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&CellsResponse{`,
		`Error:` + strings.Replace(fmt.Sprintf("%v", this.Error), "Error", "Error", 1) + `,`,
		`Cells:` + strings.Replace(fmt.Sprintf("%v", this.Cells), "CellPresence", "CellPresence", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringCells(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *CellCapacity) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
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
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryMb", wireType)
			}
			m.MemoryMb = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.MemoryMb |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskMb", wireType)
			}
			m.DiskMb = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.DiskMb |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Containers", wireType)
			}
			m.Containers = 0
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Containers |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCells(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCells
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *Providers) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
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
				return fmt.Errorf("proto: wrong wireType = %d for field Parameters", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Parameters = append(m.Parameters, string(data[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCells(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCells
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *ProtoRootfsproviders) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
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
				return fmt.Errorf("proto: wrong wireType = %d for field RootfsProviders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var keykey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				keykey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var stringLenmapkey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLenmapkey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if stringLenmapkey < 0 {
				return ErrInvalidLengthCells
			}
			postStringIndexmapkey := iNdEx + int(stringLenmapkey)
			if postStringIndexmapkey > l {
				return io.ErrUnexpectedEOF
			}
			mapkey := string(data[iNdEx:postStringIndexmapkey])
			iNdEx = postStringIndexmapkey
			var valuekey uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				valuekey |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			var mapmsglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				mapmsglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postmsgIndex := iNdEx + mapmsglen
			if mapmsglen < 0 {
				return ErrInvalidLengthCells
			}
			if postmsgIndex > l {
				return io.ErrUnexpectedEOF
			}
			mapvalue := &Providers{}
			if err := mapvalue.Unmarshal(data[iNdEx:postmsgIndex]); err != nil {
				return err
			}
			iNdEx = postmsgIndex
			if m.RootfsProviders == nil {
				m.RootfsProviders = make(map[string]*Providers)
			}
			m.RootfsProviders[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCells(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCells
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *CellPresence) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
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
				return fmt.Errorf("proto: wrong wireType = %d for field CellId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CellId = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepAddress = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Zone", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + int(stringLen)
			if stringLen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Zone = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Capacity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Capacity == nil {
				m.Capacity = &CellCapacity{}
			}
			if err := m.Capacity.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootfsProviders", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthCells
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RootfsProviders.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCells(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCells
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func (m *CellsResponse) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
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
				return fmt.Errorf("proto: wrong wireType = %d for field Error", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Error == nil {
				m.Error = &Error{}
			}
			if err := m.Error.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cells", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			postIndex := iNdEx + msglen
			if msglen < 0 {
				return ErrInvalidLengthCells
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cells = append(m.Cells, &CellPresence{})
			if err := m.Cells[len(m.Cells)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			var sizeOfWire int
			for {
				sizeOfWire++
				wire >>= 7
				if wire == 0 {
					break
				}
			}
			iNdEx -= sizeOfWire
			skippy, err := skipCells(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCells
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	return nil
}
func skipCells(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for {
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCells
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipCells(data[start:])
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
	ErrInvalidLengthCells = fmt.Errorf("proto: negative length found during unmarshaling")
)