package client

import "github.com/google/uuid"

type SimCardInfo struct {
	Imsi           string
	Iccid          string
	Op             []byte
	Amf            []byte
	Key            []byte
	AlgoType       uint32
	UeDlAmbrBps    uint32
	UeUlAmbrBps    uint32
	Sqn            uint32
	CsgIdPrsent    bool
	CsgId          uint32
	DefaultApnName string
}

type NetworkInfo struct {
	Name      string
	NetworkId uuid.UUID
	OrgID     uuid.UUID
	OrgName   string
}

type PolicyControlSimInfo struct {
	Imsi      string
	Iccid     string
	PackageId uuid.UUID
	NetworkId uuid.UUID
	Visitor   bool
}


type PolicyControlSimPackageUpdate struct {
	Imsi      string
	PackageId uuid.UUID
}
