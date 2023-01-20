package main

import (
	"os"

	"github.com/gofrs/uuid"
	"github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"

	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/server"

	pkg "github.com/ukama/ukama/systems/subscriber/hlr/pkg"

	"github.com/ukama/ukama/systems/subscriber/hlr/cmd/version"

	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ccmd "github.com/ukama/ukama/systems/common/cmd"
	"github.com/ukama/ukama/systems/common/config"
	ugrpc "github.com/ukama/ukama/systems/common/grpc"
	mb "github.com/ukama/ukama/systems/common/msgBusServiceClient"
	egen "github.com/ukama/ukama/systems/common/pb/gen/events"

	"github.com/ukama/ukama/systems/common/sql"
	"google.golang.org/grpc"
)

var serviceConfig *pkg.Config

func main() {
	ccmd.ProcessVersionArgument(pkg.ServiceName, os.Args, version.Version)

	initConfig()
	hssDb := initDb()
	runGrpcServer(hssDb)
}

// initConfig reads in config file, ENV variables, and flags if set.
func initConfig() {
	serviceConfig = pkg.NewConfig(pkg.ServiceName)
	config.LoadConfig(pkg.ServiceName, serviceConfig)
	pkg.IsDebugMode = serviceConfig.DebugMode
	logrus.Infof("Config: %+v", serviceConfig)
}

func initDb() sql.Db {
	log.Infof("Initializing Database")
	d := sql.NewDb(serviceConfig.DB, serviceConfig.DebugMode)
	err := d.Init(&db.Hlr{}, &db.Guti{}, &db.Tai{})
	if err != nil {
		log.Fatalf("Database initialization failed. Error: %v", err)
	}
	return d
}

func runGrpcServer(gormdb sql.Db) {

	instanceId := os.Getenv("POD_NAME")
	if instanceId == "" {
		/* used on local machines */
		inst, err := uuid.NewV4()
		if err != nil {
			log.Fatalf("Failed to genrate instanceId. Error %s", err.Error())
		}
		instanceId = inst.String()
	}

	mbClient := mb.NewMsgBusClient(serviceConfig.MsgClient.Timeout, pkg.SystemName,
		pkg.ServiceName, instanceId, serviceConfig.Queue.Uri,
		serviceConfig.Service.Uri, serviceConfig.MsgClient.Host, serviceConfig.MsgClient.Exchange,
		serviceConfig.MsgClient.ListenQueue, serviceConfig.MsgClient.PublishQueue,
		serviceConfig.MsgClient.RetryCount,
		serviceConfig.MsgClient.ListenerRoutes)

	log.Debugf("MessageBus Client is %+v", mbClient)
	hlr := db.NewHlrRecordRepo(gormdb)
	guti := db.NewGutiRepo(gormdb)

	// hlr service
	hlrServer, err := server.NewHlrRecordServer(hlr, guti,
		serviceConfig.FactoryHost, serviceConfig.NetworkHost, serviceConfig.PCRFHost, serviceConfig.Org)

	if err != nil {
		log.Fatalf("hlr server initilization failed. Error: %v", err)
	}
	nSrv := server.NewHlrEventServer(hlr, guti)

	rpcServer := ugrpc.NewGrpcServer(serviceConfig.Grpc, func(s *grpc.Server) {
		gen.RegisterHlrRecordServiceServer(s, hlrServer)
		egen.RegisterEventNotificationServiceServer(s, nSrv)
	})

	go msgBusListener(mbClient)

	rpcServer.StartServer()

}

func msgBusListener(m mb.MsgBusServiceClient) {

	if err := m.Register(); err != nil {
		log.Fatalf("Failed to register to Message Client Service. Error %s", err.Error())
	}

	if err := m.Start(); err != nil {
		log.Fatalf("Failed to start to Message Client Service routine for service %s. Error %s", pkg.ServiceName, err.Error())
	}
}
