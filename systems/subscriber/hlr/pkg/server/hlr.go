package server

import (
	"context"
	"time"

	"github.com/ukama/ukama/systems/common/grpc"
	pb "github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"
	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HlrRecordServer struct {
	pb.UnimplementedHlrRecordServiceServer
	hlrRepo  db.HlrRecordRepo
	gutiRepo db.GutiRepo
}

func NewHlrRecordServer(hlrRepo db.HlrRecordRepo, gutiRepo db.GutiRepo) *HlrRecordServer {
	return &HlrRecordServer{
		hlrRepo:  hlrRepo,
		gutiRepo: gutiRepo,
	}
}

func (s *HlrRecordServer) Get(c context.Context, r *pb.GetRecordReq) (*pb.GetRecordResp, error) {
	sub, err := s.hlrRepo.GetByImsi(r.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	resp := &pb.GetRecordResp{Record: &pb.Record{
		Imsi:  sub.Imsi,
		Iccid: sub.Iccid,
		Key:   sub.Key,
		Amf:   sub.Amf,
		Op:    sub.Op,
		Apn: &pb.Apn{
			Name: sub.DefaultApnName,
		},
		AlgoType:    sub.AlgoType,
		CsgId:       sub.CsgId,
		CsgIdPrsent: sub.CsgIdPrsent,
		Sqn:         sub.Sqn,
		UeDlAmbrBps: sub.UeDlAmbrBps,
		UeUlAmbrBps: sub.UeDlAmbrBps,
	}}

	return resp, nil
}

func (s *HlrRecordServer) Activate(c context.Context, req *pb.ActivateReq) (*pb.ActivateResp, error) {

	hlr := &db.Hlr{
		Iccid: req.Iccid,
	}

	/* Send Request to SIM Factory */

	/* Send message to PCRF */

	/* Add to HLR */
	err := s.hlrRepo.Add(req.Network, hlr)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "iccid")
	}

	return &pb.ActivateResp{}, err
}

func (s *HlrRecordServer) Update(c context.Context, req *pb.UpdateRecordReq) (*pb.UpdateRecordResp, error) {
	rec, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error getting imsi")
	}

	err = s.hlrRepo.Update(req.Imsi, rec)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	return &pb.UpdateRecordResp{}, nil
}

func (s *HlrRecordServer) Inactivate(c context.Context, req *pb.InactivateReq) (resp *pb.InactivateResp, err error) {
	var delHlrRecord *db.Hlr
	switch req.Id.(type) {
	case *pb.InactivateReq_Imsi:

		delHlrRecord, err = s.hlrRepo.GetByImsi(req.GetImsi())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "imsi")
		}

	case *pb.InactivateReq_Iccid:
		delHlrRecord, err = s.hlrRepo.GetByIccid(req.GetIccid())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "imsi")
		}
	}

	err = s.hlrRepo.Delete(delHlrRecord.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	return &pb.InactivateResp{}, nil

}

func (s *HlrRecordServer) UpdateGuti(c context.Context, req *pb.UpdateGutiReq) (*pb.UpdateGutiResp, error) {
	_, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	err = s.gutiRepo.Update(&db.Guti{
		Imsi:            req.Imsi,
		Plmn_id:         req.Guti.PlmnId,
		Mmegi:           req.Guti.Mmegi,
		Mmec:            req.Guti.Mmec,
		MTmsi:           req.Guti.Mtmsi,
		DeviceUpdatedAt: time.Unix(int64(req.UpdatedAt), 0),
	})
	if err != nil {
		if err.Error() == db.GutiNotUpdatedErr {
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		}
		return nil, grpc.SqlErrorToGrpc(err, "guti")
	}

	return &pb.UpdateGutiResp{}, nil
}

func (s *HlrRecordServer) UpdateTai(c context.Context, req *pb.UpdateTaiReq) (*pb.UpdateTaiResp, error) {
	_, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	err = s.hlrRepo.UpdateTai(req.Imsi, db.Tai{
		PlmId:           req.PlmnId,
		Tac:             req.Tac,
		DeviceUpdatedAt: time.Unix(int64(req.UpdatedAt), 0),
	})

	if err != nil {
		if err.Error() == db.TaiNotUpdatedErr {
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}
		return nil, grpc.SqlErrorToGrpc(err, "tai")
	}

	return &pb.UpdateTaiResp{}, nil
}

func grpcHlrRecordToDb(sub *pb.Record, netName string) (*db.Hlr, error) {

	dbSub := &db.Hlr{
		Imsi:           sub.Imsi,
		Iccid:          sub.Iccid,
		DefaultApnName: sub.Apn.Name,
		Key:            sub.Key,
		Amf:            sub.Amf,
		Op:             sub.Op,
		Network: &db.Network{
			Name: netName,
		},
	}

	return dbSub, nil
}
