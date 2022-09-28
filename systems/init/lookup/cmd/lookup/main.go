package main

import (
	"os"

	"github.com/num30/config"
	uconf "github.com/ukama/ukama/systems/common/config"
	"github.com/ukama/ukama/systems/common/metrics"
	"github.com/ukama/ukama/systems/init/lookup/cmd/version"
	"github.com/ukama/ukama/systems/init/lookup/internal"
	"github.com/ukama/ukama/systems/init/lookup/internal/db"
	"github.com/ukama/ukama/systems/init/lookup/internal/server"
	"gopkg.in/yaml.v3"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	ccmd "github.com/ukama/ukama/systems/common/cmd"
	ugrpc "github.com/ukama/ukama/systems/common/grpc"
	"github.com/ukama/ukama/systems/common/sql"
	generated "github.com/ukama/ukama/systems/init/lookup/pb/gen"
	mb "github.com/ukama/ukama/systems/init/lookup/pkg/msgBusClient"
	"google.golang.org/grpc"
)

var serviceConfig = internal.NewConfig()

func main() {
	ccmd.ProcessVersionArgument("lookup", os.Args, version.Version)

	/* Log level */
	logrus.SetLevel(logrus.TraceLevel)
	log.Infof("Starting the lookup service")

	initConfig()

	metrics.StartMetricsServer(serviceConfig.Metrics)

	db := initDb()

	runGrpcServer(db)

	logrus.Infof("Exiting service %s", internal.ServiceName)

}

func initDb() sql.Db {
	log.Infof("Initializing Database")
	d := sql.NewDb(serviceConfig.DB, serviceConfig.DebugMode)
	err := d.Init(&db.Org{}, &db.Node{}, &db.System{})
	if err != nil {
		log.Fatalf("Database initialization failed. Error: %v", err)
	}
	return d
}

func initConfig() {
	log.Infof("Initializing config")
	serviceConfig = &internal.Config{
		DB: &uconf.Database{
			DbName: internal.ServiceName,
		},
	}

	err := config.NewConfReader(internal.ServiceName).Read(serviceConfig)
	if err != nil {
		log.Fatal("Error reading config ", err)
	} else if serviceConfig.DebugMode {
		b, err := yaml.Marshal(serviceConfig)
		if err != nil {
			logrus.Infof("Config:\n%s", string(b))
		}
	}

	log.Debugf("\nService: %s DB Config: %+v", internal.ServiceName, serviceConfig.DB)

	internal.IsDebugMode = serviceConfig.DebugMode
}

func runGrpcServer(d sql.Db) {
	instanceId := os.Getenv("POD_NAME")

	//var mbClient *mb.MsgBusClient
	mbClient := mb.NewMsgBusClient(serviceConfig.MsgClient.Timeout, internal.SystemName,
		internal.ServiceName, instanceId, serviceConfig.Queue.Uri,
		serviceConfig.MsgClient.Host, serviceConfig.MsgClient.RetryCount,
		serviceConfig.MsgClient.ListnerRoutes)
	grpcServer := ugrpc.NewGrpcServer(*serviceConfig.Grpc, func(s *grpc.Server) {
		srv := server.NewLookupServer(db.NewNodeRepo(d), db.NewOrgRepo(d), db.NewSystemRepo(d), mbClient)
		generated.RegisterLookupServiceServer(s, srv)
	})

	grpcServer.StartServer()
}
