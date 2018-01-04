// Code generated by protoc-gen-go. DO NOT EDIT.
// source: storage.proto

/*
Package zconfig is a generated protocol buffer package.

It is generated from these files:
	storage.proto
	fw.proto
	devconfig.proto
	devcommon.proto
	netcmn.proto
	appconfig.proto
	vm.proto
	network.proto
	netconfig.proto
	baseosconfig.proto

It has these top-level messages:
	SignatureInfo
	DatastoreConfig
	Image
	Drive
	ACEMatch
	ACEAction
	ACE
	MapServer
	ZedServer
	DeviceLispDetails
	EdgeDevConfig
	UUIDandVersion
	Ipv4Spec
	Ipv6Spec
	NameToEid
	EIDAllocation
	Lispspec
	L2Spec
	AppInstanceConfig
	VmConfig
	NetworkParam
	NetworkServiceReq
	NetworkServiceResp
	NetworkIf
	NetworkConfig
	NetworkAdapter
	OSKeyTags
	OSVerDetails
	BaseOSConfig
*/
package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DsType int32

const (
	DsType_DsUnknown DsType = 0
	DsType_DsHttp    DsType = 1
	DsType_DsHttps   DsType = 2
	DsType_DsS3      DsType = 3
	DsType_DsSFTP    DsType = 4
)

var DsType_name = map[int32]string{
	0: "DsUnknown",
	1: "DsHttp",
	2: "DsHttps",
	3: "DsS3",
	4: "DsSFTP",
}
var DsType_value = map[string]int32{
	"DsUnknown": 0,
	"DsHttp":    1,
	"DsHttps":   2,
	"DsS3":      3,
	"DsSFTP":    4,
}

func (x DsType) String() string {
	return proto.EnumName(DsType_name, int32(x))
}
func (DsType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Format int32

const (
	Format_FmtUnknown Format = 0
	Format_Raw        Format = 1
	Format_QCOW       Format = 2
	Format_QCOW2      Format = 3
	Format_VHD        Format = 4
	Format_VMDK       Format = 5
	Format_OVA        Format = 6
	Format_VHDX       Format = 7
)

var Format_name = map[int32]string{
	0: "FmtUnknown",
	1: "Raw",
	2: "QCOW",
	3: "QCOW2",
	4: "VHD",
	5: "VMDK",
	6: "OVA",
	7: "VHDX",
}
var Format_value = map[string]int32{
	"FmtUnknown": 0,
	"Raw":        1,
	"QCOW":       2,
	"QCOW2":      3,
	"VHD":        4,
	"VMDK":       5,
	"OVA":        6,
	"VHDX":       7,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}
func (Format) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Target int32

const (
	Target_TgtUnknown Target = 0
	Target_Disk       Target = 1
	Target_Kernel     Target = 2
	Target_Initrd     Target = 3
	Target_RamDisk    Target = 4
)

var Target_name = map[int32]string{
	0: "TgtUnknown",
	1: "Disk",
	2: "Kernel",
	3: "Initrd",
	4: "RamDisk",
}
var Target_value = map[string]int32{
	"TgtUnknown": 0,
	"Disk":       1,
	"Kernel":     2,
	"Initrd":     3,
	"RamDisk":    4,
}

func (x Target) String() string {
	return proto.EnumName(Target_name, int32(x))
}
func (Target) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DriveType int32

const (
	DriveType_Unclassified DriveType = 0
	DriveType_CDROM        DriveType = 1
	DriveType_HDD          DriveType = 2
	DriveType_NET          DriveType = 3
)

var DriveType_name = map[int32]string{
	0: "Unclassified",
	1: "CDROM",
	2: "HDD",
	3: "NET",
}
var DriveType_value = map[string]int32{
	"Unclassified": 0,
	"CDROM":        1,
	"HDD":          2,
	"NET":          3,
}

func (x DriveType) String() string {
	return proto.EnumName(DriveType_name, int32(x))
}
func (DriveType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type SignatureInfo struct {
	Intercertsurl string `protobuf:"bytes,1,opt,name=intercertsurl" json:"intercertsurl,omitempty"`
	Signercerturl string `protobuf:"bytes,2,opt,name=signercerturl" json:"signercerturl,omitempty"`
	Signature     []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignatureInfo) Reset()                    { *m = SignatureInfo{} }
func (m *SignatureInfo) String() string            { return proto.CompactTextString(m) }
func (*SignatureInfo) ProtoMessage()               {}
func (*SignatureInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *SignatureInfo) GetIntercertsurl() string {
	if m != nil {
		return m.Intercertsurl
	}
	return ""
}

func (m *SignatureInfo) GetSignercerturl() string {
	if m != nil {
		return m.Signercerturl
	}
	return ""
}

func (m *SignatureInfo) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type DatastoreConfig struct {
	Id       string `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	DType    DsType `protobuf:"varint,1,opt,name=dType,enum=DsType" json:"dType,omitempty"`
	Fqdn     string `protobuf:"bytes,2,opt,name=fqdn" json:"fqdn,omitempty"`
	ApiKey   string `protobuf:"bytes,3,opt,name=apiKey" json:"apiKey,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password" json:"password,omitempty"`
	// depending on datastore types, it could be bucket or path
	Dpath string `protobuf:"bytes,5,opt,name=dpath" json:"dpath,omitempty"`
}

func (m *DatastoreConfig) Reset()                    { *m = DatastoreConfig{} }
func (m *DatastoreConfig) String() string            { return proto.CompactTextString(m) }
func (*DatastoreConfig) ProtoMessage()               {}
func (*DatastoreConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DatastoreConfig) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DatastoreConfig) GetDType() DsType {
	if m != nil {
		return m.DType
	}
	return DsType_DsUnknown
}

func (m *DatastoreConfig) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *DatastoreConfig) GetApiKey() string {
	if m != nil {
		return m.ApiKey
	}
	return ""
}

func (m *DatastoreConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *DatastoreConfig) GetDpath() string {
	if m != nil {
		return m.Dpath
	}
	return ""
}

type Image struct {
	Id             string          `protobuf:"bytes,100,opt,name=id" json:"id,omitempty"`
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion" json:"uuidandversion,omitempty"`
	// it could be relative path/name as well
	Name    string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Sha256  string `protobuf:"bytes,3,opt,name=sha256" json:"sha256,omitempty"`
	Iformat Format `protobuf:"varint,4,opt,name=iformat,enum=Format" json:"iformat,omitempty"`
	// if its signed image
	Siginfo *SignatureInfo `protobuf:"bytes,5,opt,name=siginfo" json:"siginfo,omitempty"`
	DsId    string         `protobuf:"bytes,6,opt,name=dsId" json:"dsId,omitempty"`
}

func (m *Image) Reset()                    { *m = Image{} }
func (m *Image) String() string            { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()               {}
func (*Image) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Image) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Image) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *Image) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Image) GetSha256() string {
	if m != nil {
		return m.Sha256
	}
	return ""
}

func (m *Image) GetIformat() Format {
	if m != nil {
		return m.Iformat
	}
	return Format_FmtUnknown
}

func (m *Image) GetSiginfo() *SignatureInfo {
	if m != nil {
		return m.Siginfo
	}
	return nil
}

func (m *Image) GetDsId() string {
	if m != nil {
		return m.DsId
	}
	return ""
}

type Drive struct {
	Image    *Image    `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Maxsize  int64     `protobuf:"varint,2,opt,name=maxsize" json:"maxsize,omitempty"`
	Readonly bool      `protobuf:"varint,5,opt,name=readonly" json:"readonly,omitempty"`
	Preserve bool      `protobuf:"varint,6,opt,name=preserve" json:"preserve,omitempty"`
	Drvtype  DriveType `protobuf:"varint,8,opt,name=drvtype,enum=DriveType" json:"drvtype,omitempty"`
	Target   Target    `protobuf:"varint,9,opt,name=target,enum=Target" json:"target,omitempty"`
}

func (m *Drive) Reset()                    { *m = Drive{} }
func (m *Drive) String() string            { return proto.CompactTextString(m) }
func (*Drive) ProtoMessage()               {}
func (*Drive) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Drive) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

func (m *Drive) GetMaxsize() int64 {
	if m != nil {
		return m.Maxsize
	}
	return 0
}

func (m *Drive) GetReadonly() bool {
	if m != nil {
		return m.Readonly
	}
	return false
}

func (m *Drive) GetPreserve() bool {
	if m != nil {
		return m.Preserve
	}
	return false
}

func (m *Drive) GetDrvtype() DriveType {
	if m != nil {
		return m.Drvtype
	}
	return DriveType_Unclassified
}

func (m *Drive) GetTarget() Target {
	if m != nil {
		return m.Target
	}
	return Target_TgtUnknown
}

func init() {
	proto.RegisterType((*SignatureInfo)(nil), "SignatureInfo")
	proto.RegisterType((*DatastoreConfig)(nil), "DatastoreConfig")
	proto.RegisterType((*Image)(nil), "Image")
	proto.RegisterType((*Drive)(nil), "Drive")
	proto.RegisterEnum("DsType", DsType_name, DsType_value)
	proto.RegisterEnum("Format", Format_name, Format_value)
	proto.RegisterEnum("Target", Target_name, Target_value)
	proto.RegisterEnum("DriveType", DriveType_name, DriveType_value)
}

func init() { proto.RegisterFile("storage.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x94, 0x51, 0x6f, 0x22, 0x37,
	0x10, 0xc7, 0xb3, 0xc0, 0xee, 0xc2, 0x24, 0x10, 0xcb, 0xaa, 0xaa, 0x55, 0x94, 0x28, 0x29, 0xca,
	0x43, 0xc4, 0xc3, 0x46, 0x22, 0x6a, 0x2b, 0xf5, 0xa9, 0x6d, 0xb6, 0x14, 0x14, 0xa5, 0x69, 0x37,
	0x40, 0xab, 0xaa, 0x2f, 0xce, 0xda, 0x6c, 0xac, 0xb0, 0x36, 0xb5, 0x0d, 0x39, 0xf2, 0x61, 0xee,
	0xa3, 0xdc, 0x37, 0xb9, 0xef, 0x72, 0xb2, 0x0d, 0xe4, 0x72, 0xf7, 0x36, 0xff, 0xff, 0x0c, 0x33,
	0xbf, 0xf1, 0x00, 0xd0, 0xd6, 0x46, 0x2a, 0x52, 0xb2, 0x74, 0xa1, 0xa4, 0x91, 0x47, 0x87, 0x94,
	0xad, 0x0a, 0x59, 0x55, 0x52, 0x78, 0xa3, 0xbb, 0x86, 0xf6, 0x3d, 0x2f, 0x05, 0x31, 0x4b, 0xc5,
	0x46, 0x62, 0x26, 0xf1, 0x39, 0xb4, 0xb9, 0x30, 0x4c, 0x15, 0x4c, 0x19, 0xbd, 0x54, 0xf3, 0x24,
	0x38, 0x0b, 0x2e, 0x5a, 0xf9, 0x5b, 0xd3, 0x56, 0x69, 0x5e, 0x0a, 0xef, 0xd8, 0xaa, 0x9a, 0xaf,
	0x7a, 0x63, 0xe2, 0x63, 0x68, 0xe9, 0x6d, 0xf3, 0xa4, 0x7e, 0x16, 0x5c, 0x1c, 0xe4, 0xaf, 0x46,
	0xf7, 0x7d, 0x00, 0x87, 0x19, 0x31, 0xc4, 0x12, 0xb2, 0x6b, 0x29, 0x66, 0xbc, 0xc4, 0x1d, 0xa8,
	0x71, 0x9a, 0x50, 0xd7, 0xac, 0xc6, 0x29, 0x3e, 0x81, 0x90, 0x8e, 0xd7, 0x0b, 0xe6, 0x28, 0x3a,
	0xfd, 0x38, 0xcd, 0xb4, 0x95, 0xb9, 0x77, 0x31, 0x86, 0xc6, 0xec, 0x7f, 0x2a, 0x36, 0xd3, 0x5d,
	0x8c, 0xbf, 0x85, 0x88, 0x2c, 0xf8, 0x0d, 0x5b, 0xbb, 0x89, 0xad, 0x7c, 0xa3, 0xf0, 0x11, 0x34,
	0x17, 0x44, 0xeb, 0x67, 0xa9, 0x68, 0xd2, 0x70, 0x99, 0x9d, 0xc6, 0xdf, 0x40, 0x48, 0x17, 0xc4,
	0x3c, 0x26, 0xa1, 0x4b, 0x78, 0xd1, 0xfd, 0x18, 0x40, 0x38, 0xaa, 0x48, 0xc9, 0xbe, 0xc2, 0xfa,
	0x11, 0x3a, 0xcb, 0x25, 0xa7, 0x44, 0xd0, 0x15, 0x53, 0x9a, 0x4b, 0xe1, 0xf8, 0xf6, 0xfb, 0x87,
	0xe9, 0x64, 0x32, 0xca, 0x88, 0xa0, 0x53, 0x6f, 0xe7, 0x5f, 0x94, 0x59, 0x60, 0x41, 0x2a, 0xb6,
	0x05, 0xb6, 0xb1, 0x05, 0xd6, 0x8f, 0xa4, 0xff, 0xfd, 0x0f, 0x5b, 0x60, 0xaf, 0xf0, 0x77, 0x10,
	0xf3, 0x99, 0x54, 0x15, 0x31, 0x8e, 0xd7, 0x6e, 0x3f, 0x70, 0x32, 0xdf, 0xfa, 0xf8, 0x02, 0x62,
	0xcd, 0x4b, 0x2e, 0x66, 0xd2, 0x91, 0xef, 0xf7, 0x3b, 0xe9, 0x9b, 0x6b, 0xe6, 0xdb, 0xb4, 0x1d,
	0x4c, 0xf5, 0x88, 0x26, 0x91, 0x1f, 0x6c, 0xe3, 0xee, 0x87, 0x00, 0xc2, 0x4c, 0xf1, 0x15, 0xc3,
	0xc7, 0x10, 0x72, 0xbb, 0xe8, 0x66, 0x8d, 0x28, 0x75, 0x6b, 0xe7, 0xde, 0xc4, 0x09, 0xc4, 0x15,
	0x79, 0xa7, 0xf9, 0x8b, 0xe7, 0xae, 0xe7, 0x5b, 0x69, 0xdf, 0x54, 0x31, 0x42, 0xa5, 0x98, 0xaf,
	0x1d, 0x40, 0x33, 0xdf, 0x69, 0xf7, 0xde, 0x8a, 0x69, 0xa6, 0x56, 0xcc, 0x4d, 0x6d, 0xe6, 0x3b,
	0x8d, 0xcf, 0x21, 0xa6, 0x6a, 0x65, 0xec, 0x61, 0x9b, 0x6e, 0x35, 0x48, 0x1d, 0x88, 0xbb, 0xed,
	0x36, 0x85, 0x4f, 0x21, 0x32, 0x44, 0x95, 0xcc, 0x24, 0xad, 0xcd, 0xfe, 0x63, 0x27, 0xf3, 0x8d,
	0xdd, 0x1b, 0x40, 0xe4, 0xbf, 0x0f, 0xb8, 0x0d, 0xad, 0x4c, 0x4f, 0xc4, 0x93, 0x90, 0xcf, 0x02,
	0xed, 0x61, 0xb0, 0x89, 0xa1, 0x31, 0x0b, 0x14, 0xe0, 0x7d, 0x88, 0x7d, 0xac, 0x51, 0x0d, 0x37,
	0xa1, 0x91, 0xe9, 0xfb, 0x2b, 0x54, 0xf7, 0x25, 0xf7, 0x83, 0xf1, 0x9f, 0xa8, 0xd1, 0xfb, 0x0f,
	0x22, 0xff, 0xb2, 0xb8, 0x03, 0x30, 0xa8, 0xcc, 0x6b, 0xa3, 0x18, 0xea, 0x39, 0x79, 0x46, 0x81,
	0xfd, 0xe0, 0x5f, 0xd7, 0x77, 0x7f, 0xa3, 0x1a, 0x6e, 0x41, 0x68, 0xa3, 0x3e, 0xaa, 0xdb, 0xec,
	0x74, 0x98, 0xa1, 0x86, 0xcd, 0x4e, 0x6f, 0xb3, 0x1b, 0x14, 0x5a, 0xeb, 0x6e, 0xfa, 0x0b, 0x8a,
	0x9c, 0x35, 0xcc, 0xfe, 0x41, 0x71, 0xef, 0x77, 0x88, 0x3c, 0xb7, 0xed, 0x3e, 0x2e, 0x3f, 0xeb,
	0x6e, 0x69, 0xb8, 0x7e, 0x42, 0x81, 0xa5, 0xb9, 0x61, 0x4a, 0xb0, 0x39, 0xaa, 0xd9, 0x78, 0x24,
	0xb8, 0x51, 0x14, 0xd5, 0x2d, 0x7c, 0x4e, 0x2a, 0x57, 0xd4, 0xe8, 0xfd, 0x04, 0xad, 0xdd, 0x2b,
	0x61, 0x04, 0x07, 0x13, 0x51, 0xcc, 0x89, 0xd6, 0x7c, 0xc6, 0x19, 0x45, 0x7b, 0x16, 0xec, 0x3a,
	0xcb, 0xef, 0x6e, 0x51, 0x60, 0x29, 0x86, 0x59, 0x86, 0x6a, 0x36, 0xf8, 0xe3, 0xb7, 0x31, 0xaa,
	0xff, 0xfa, 0x33, 0x9c, 0x16, 0xb2, 0x4a, 0x5f, 0x18, 0x65, 0x94, 0xa4, 0xc5, 0x5c, 0x2e, 0x69,
	0xba, 0xb4, 0xc7, 0xe0, 0xc5, 0xe6, 0xbf, 0xe1, 0xdf, 0x93, 0x92, 0x9b, 0xc7, 0xe5, 0x43, 0x5a,
	0xc8, 0xea, 0xd2, 0xd7, 0x5d, 0x92, 0x05, 0xbf, 0x7c, 0x29, 0xdc, 0x4f, 0xf3, 0x21, 0x72, 0x55,
	0x57, 0x9f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xcd, 0xce, 0xbf, 0x52, 0x04, 0x00, 0x00,
}