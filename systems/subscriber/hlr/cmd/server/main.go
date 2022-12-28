package main

import (
	"os"

	"github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"

	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/server"

	pkg "github.com/ukama/ukama/systems/subscriber/hlr/pkg"

	"github.com/ukama/ukama/systems/subscriber/hlr/cmd/version"

	"github.com/ukama/ukama/systems/subscriber/hlr/pkg/db"

	log "github.com/sirupsen/logrus"
	ccmd "github.com/ukama/ukama/systems/common/cmd"
	"github.com/ukama/ukama/systems/common/config"
	ugrpc "github.com/ukama/ukama/systems/common/grpc"

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
	serviceConfig = pkg.NewConfig()
	config.LoadConfig(pkg.ServiceName, serviceConfig)
	pkg.IsDebugMode = serviceConfig.DebugMode
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

	// hlr service
	hlrServer := server.NewHlrRecordServer(db.NewHlrRecordRepo(gormdb), db.NewGutiRepo(gormdb))

	rpcServer := ugrpc.NewGrpcServer(serviceConfig.Grpc, func(s *grpc.Server) {
		gen.RegisterHlrRecordServiceServer(s, hlrServer)
	})

	rpcServer.StartServer()

}
