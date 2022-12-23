package rest

import (
	"time"
)

type GUTI struct {
	PlmnId string `json: "plmn_id" validate:"required"`
	Mmegi  uint64 `json: "mmegi" validate:"required"`
	Mmec   uint64 `json: "mmec" validate:"required"`
	Mtmsi  uint64 `json: "mtmsi" validate:"required"`
}

type UpdateGutiReq struct {
	Imsi     string    `json:"imsi" validate:"required"`
	Guti     GUTI      `json: "guti" validate:"required"`
	UpdateAt time.Time `json: "updated_at" validate:"required"`
}

type GetRecordReq struct {
	Imsi string `json:"imsi" validate:"required"`
}

type APN struct {
	name string `json:"name" validate:"required"`
}

type GetRecordResp struct {
	Imsi        string `json:"imsi" validate:"required"`
	Key         []byte `json:"key" validate:"required" size:16`
	Op          []byte `json:"op" validate:"required" size:16`
	Amf         string `json:"amf" validate:"required"`
	Apn         APN    `json:"apn" validate:"required"`
	AlgoType    uint32 `json:"algo_type" validate:"required"`
	UeDlAmbrBps uint32 `json:"ue_dl_ambr_bps" validate:"required"`
	UeUlAmbrBps uint32 `json:"ue_ul_ambr_bps" validate:"required"`
	Sqn         uint64 `json:"sqn" validate:"required"`
	CsgIs       bool   `json:"csg_is_present" validate:"required"`
	csgId       uint32 `json:"csg_id" validate:"required"`
}

type UpdateTaiReq struct {
	Imsi     string    `json:"imsi" validate:"required"`
	PlmnId   string    `json: "plmn_id" validate:"required"`
	Tac      uint64    `json: "tac" validate:"required"`
	UpdateAt time.Time `json: "updated_at" validate:"required"`
}
