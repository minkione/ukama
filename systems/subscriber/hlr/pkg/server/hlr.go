package server

import (
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"github.com/ukama/ukama/systems/common/grpc"
	pb "github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"
	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/client"
	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HlrRecordServer struct {
	pb.UnimplementedHlrRecordServiceServer
	hlrRepo  db.HlrRecordRepo
	gutiRepo db.GutiRepo
	pcrf     client.PolicyControl
	network  client.Network
	factory  client.Factory
	Org      string
}

func NewHlrRecordServer(hlrRepo db.HlrRecordRepo, gutiRepo db.GutiRepo, factory client.Factory, network client.Network, pcrf client.PolicyControl, org string) (*HlrRecordServer, error) {

	hlr := HlrRecordServer{
		hlrRepo:  hlrRepo,
		gutiRepo: gutiRepo,
		Org:      org,
		factory:  factory,
		network:  network,
		pcrf:     pcrf,
	}
	return &hlr, nil
}

func (s *HlrRecordServer) Read(c context.Context, req *pb.ReadReq) (*pb.ReadResp, error) {
	var sub *db.Hlr
	var err error

	switch req.Id.(type) {
	case *pb.ReadReq_Imsi:

		sub, err = s.hlrRepo.GetByImsi(req.GetImsi())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "error getting imsi")
		}

	case *pb.ReadReq_Iccid:
		sub, err = s.hlrRepo.GetByIccid(req.GetIccid())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "error getting iccid")
		}
	}

	resp := &pb.ReadResp{Record: &pb.Record{
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
		PackageId:   sub.PackageId,
	}}

	logrus.Infof("Subscriber is having %+v", resp)
	return resp, nil
}

func (s *HlrRecordServer) Activate(c context.Context, req *pb.ActivateReq) (*pb.ActivateResp, error) {

	/* Validate network in Org */
	err := s.network.ValidateNetwork(req.Network, s.Org)
	if err != nil {
		return nil, fmt.Errorf("error validating network")
	}

	/* Send Request to SIM Factory */
	sim, err := s.factory.ReadSimCardInfo(req.Iccid)
	if err != nil {
		return nil, fmt.Errorf("error reading iccid from factory")
	}

	/* Send message to PCRF */
	nId, err := uuid.FromString(req.Network)
	if err != nil {
		logrus.Errorf("NetworkId not valid.")
		return nil, err
	}

	pId, err := uuid.FromString(req.PackageId)
	if err != nil {
		logrus.Errorf("PackageId not valid.")
	}

	pcrfData := client.PolicyControlSimInfo{
		Imsi:      sim.Imsi,
		Iccid:     sim.Iccid,
		PackageId: pId,
		NetworkId: nId,
		Visitor:   false, // We will using this flag on roaming in VLR
	}

	err = s.pcrf.AddSim(pcrfData)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error adding to pcrf")
	}

	/* Add to HLR */
	hlr := &db.Hlr{
		Iccid:          req.Iccid,
		Imsi:           sim.Imsi,
		Op:             sim.Op,
		Key:            sim.Key,
		Amf:            sim.Amf,
		AlgoType:       sim.AlgoType,
		UeDlAmbrBps:    sim.UeDlAmbrBps,
		UeUlAmbrBps:    sim.UeUlAmbrBps,
		Sqn:            uint64(sim.Sqn),
		CsgIdPrsent:    sim.CsgIdPrsent,
		CsgId:          sim.CsgId,
		DefaultApnName: sim.DefaultApnName,
		PackageId:      req.PackageId,
	}

	err = s.hlrRepo.Add(hlr)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating hlr")
	}

	return &pb.ActivateResp{}, err
}

func (s *HlrRecordServer) UpdatePackage(c context.Context, req *pb.UpdatePackageReq) (*pb.UpdatePackageResp, error) {
	hlrRecord, err := s.hlrRepo.GetByIccid(req.GetIccid())
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error getting iccid")
	}

	/* We assum that packageId is validated by subscriber. */
	pId, err := uuid.FromString(req.PackageId)
	if err != nil {
		logrus.Errorf("PackageId not valid.")
		return nil, grpc.SqlErrorToGrpc(err, "error invalid package id")
	}

	pD := client.PolicyControlSimPackageUpdate{
		Imsi:      hlrRecord.Imsi,
		PackageId: pId,
	}

	err = s.pcrf.UpdateSim(pD)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating pcrf")
	}

	err = s.hlrRepo.UpdatePackage(hlrRecord.Imsi, req.PackageId)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating hlr")
	}

	return &pb.UpdatePackageResp{}, nil
}

func (s *HlrRecordServer) Inactivate(c context.Context, req *pb.InactivateReq) (*pb.InactivateResp, error) {
	var delHlrRecord *db.Hlr
	var err error

	switch req.Id.(type) {
	case *pb.InactivateReq_Imsi:

		delHlrRecord, err = s.hlrRepo.GetByImsi(req.GetImsi())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "error getting imsi")
		}

	case *pb.InactivateReq_Iccid:
		delHlrRecord, err = s.hlrRepo.GetByIccid(req.GetIccid())
		if err != nil {
			return nil, grpc.SqlErrorToGrpc(err, "error getting iccid")
		}
	}

	err = s.pcrf.DeleteSim(delHlrRecord.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating pcrf")
	}

	err = s.hlrRepo.Delete(delHlrRecord.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating hlr")
	}

	return &pb.InactivateResp{}, nil

}

/*
func (s *HlrRecordServer) Update(c context.Context, req *pb.UpdateRecordReq) (*pb.UpdateRecordResp, error) {
	rec, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error getting imsi")
	}

	err = s.hlrRepo.Update(req.Imsi, rec)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "error updating hlr")
	}

	return &pb.UpdateRecordResp{}, nil
}
*/

func (s *HlrRecordServer) UpdateGuti(c context.Context, req *pb.UpdateGutiReq) (*pb.UpdateGutiResp, error) {
	_, err := s.hlrRepo.GetByImsi(req.Imsi)
	if err != nil {
		return nil, grpc.SqlErrorToGrpc(err, "imsi")
	}

	err = s.gutiRepo.Update(&db.Guti{
		Imsi:            req.Imsi,
		PlmnId:          req.Guti.PlmnId,
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
		PlmnId:          req.PlmnId,
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
