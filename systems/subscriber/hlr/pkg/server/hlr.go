package server

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		Imsi: sub.Imsi,
		Key:  sub.Key,
		Amf:  sub.Amf,
		Op:   sub.Op,
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

func (s *HlrRecordServer) Add(c context.Context, a *pb.AddRecordReq) (*pb.AddRecordResp, error) {
	sub := a.Record

	dbSub, err := grpcHlrRecordToDb(sub, a.Record.Network.name)
	if err != nil {
		return nil, err
	}
	err = s.hlrRepo.Add(a.Org, dbSub)

	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	return &pb.AddRecordResp{}, err
}

func (s *HlrRecordServer) Update(c context.Context, req *pb.UpdateRecordReq) (*pb.UpdateRecordResp, error) {
	imsi, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error getting imsi")
	}

	dbSub, err := grpcHlrRecordToDb(req.Imsi, imsi.Network.Name)
	if err != nil {
		return nil, err
	}

	err = s.hlrRepo.Update(req.Imsi, dbSub)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	return &pb.UpdateRecordResp{}, nil
}

func (s *HlrRecordServer) Delete(c context.Context, req *pb.DeleteRecordReq) (resp *pb.DeleteRecordResp, err error) {
	var delHlrRecord *db.HlrRecord
	switch req.IdOneof.(type) {
	case *pb.DeleteRecordReq_HlrRecord:

		delHlrRecord, err = s.hlrRepo.GetByImsi(req.GetImsi())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "imsi")
		}

	case *pb.DeleteRecordReq_UserId:
		uuid, err := uuid.Parse(req.GetUserId())
		if err != nil {
			logrus.Errorf("Error parsing uuid %s. Error: %s", uuid, err)
			return nil, fmt.Errorf("error parsing uuid")
		}

		imsis, err := s.hlrRepo.GetHlrRecordByUserUuid(uuid)
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "imsi")
		}

		if len(imsis) == 1 {
			delHlrRecord = imsis[0]
		} else if len(imsis) == 0 {
			return nil, status.Error(codes.NotFound, "imsi not found")
		} else {
			return nil, status.Error(codes.Internal, "invalid number of imsis found")
		}
	}

	err = s.hlrRepo.Delete(delHlrRecord.HlrRecord)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	return &pb.DeleteRecordResp{}, nil

}

func (s *HlrRecordServer) UpdateGuti(c context.Context, req *pb.UpdateGutiReq) (*pb.UpdateGutiResp, error) {
	imsi, err := s.hlrRepo.GetByImsi(req.Imsi)
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
	imsi, err := s.hlrRepo.GetByImsi(req.Imsi)
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

func grpcHlrRecordToDb(sub *pb.Record, netName string) (*db.HlrRecord, error) {

	dbSub := &db.HlrRecord{
		Imsi:           sub.Imsi,
		UserUuid:       userId,
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
