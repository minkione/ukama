package stubs

import (
	"github.com/google/uuid"
)

type SimCardInfoReq struct {
	Iccid string `path:"iccid" validate:"required"`
}

// TODO: update
type NetworkValidationReq struct {
	network uuid.UUID `path:"network" validate:"required"`
	org     uuid.UUID `path:"org" validate:"required"`
}

type DeleteSimReq struct {
	Imsi string `path:"imsi" validate:"required"`
}
