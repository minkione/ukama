package server

import (
	"context"
	"testing"

	mocks "github.com/ukama/ukama/systems/subscriber/hlr/mocks"
	pb "github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"
	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"

	"github.com/stretchr/testify/assert"
)

var Factory = "localhost:9090"
var Network = "localhost:9090"
var Pcrf = "localhost:9090"
var Org = "40987edb-ebb6-4f84-a27c-99db7c136127"

var sub = db.Hlr{
	Iccid:          "0123456789012345678912",
	Imsi:           "012345678912345",
	Op:             []byte("0123456789012345"),
	Key:            []byte("0123456789012345"),
	Amf:            []byte("800"),
	AlgoType:       1,
	UeDlAmbrBps:    2000000,
	UeUlAmbrBps:    2000000,
	Sqn:            1,
	CsgIdPrsent:    false,
	CsgId:          0,
	DefaultApnName: "ukama",
}

func TestHlr_Read(t *testing.T) {
	hlrRepo := &mocks.HlrRecordRepo{}
	gutiRepo := &mocks.GutiRepo{}

	reqPb := pb.ReadReq{
		Id: &pb.ReadReq_Iccid{
			Iccid: "0123456789012345678912",
		},
	}

	hlrRepo.On("GetByIccid", reqPb.GetIccid()).Return(&sub, nil).Once()

	s, err := NewHlrRecordServer(hlrRepo, gutiRepo, Factory, Network, Pcrf, Org)
	assert.NoError(t, err)

	hs, err := s.Read(context.TODO(), &reqPb)
	assert.NoError(t, err)

	assert.NotNil(t, hs)

	hlrRepo.AssertExpectations(t)
	gutiRepo.AssertExpectations(t)
	assert.NoError(t, err)
}
