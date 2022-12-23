package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/sirupsen/logrus"
	"github.com/ukama/ukama/systems/common/config"
	"github.com/ukama/ukama/systems/common/rest"
	"github.com/ukama/ukama/systems/subscriber/node-gateway/cmd/version"
	"github.com/ukama/ukama/systems/subscriber/node-gateway/pkg"
	"github.com/ukama/ukama/systems/subscriber/node-gateway/pkg/client"

	pb "github.com/ukama/ukama/systems/subscriber/hlr/pb/gen"
	"github.com/wI2L/fizz"
	"github.com/wI2L/fizz/openapi"
)

const NODE_URL_PARAMETER = "node"

type Router struct {
	f       *fizz.Fizz
	clients *Clients
	config  *RouterConfig
}

type RouterConfig struct {
	metricsConfig config.Metrics
	httpEndpoints *pkg.HttpEndpoints
	debugMode     bool
	serverConf    *rest.HttpConfig
}

type Clients struct {
	h hlr
}

type hlr interface {
	UpdateGuti(req *pb.UpdateGutiReq) (*pb.UpdateGutiResp, error)
	UpdateTai(req *pb.UpdateTaiReq) (*pb.UpdateTaiResp, error)
	ReadRecord(req *pb.GetRecordReq) (*pb.GetRecordResp, error)
}

func NewClientsSet(endpoints *pkg.GrpcEndpoints) *Clients {
	c := &Clients{}
	c.h = client.NewHlr(endpoints.Hlr, endpoints.Timeout)

	return c
}

func NewRouter(clients *Clients, config *RouterConfig) *Router {
	r := &Router{
		clients: clients,
		config:  config,
	}

	if !config.debugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r.init()
	return r
}

func NewRouterConfig(svcConf *pkg.Config) *RouterConfig {
	return &RouterConfig{
		metricsConfig: svcConf.Metrics,
		httpEndpoints: &svcConf.HttpServices,
		serverConf:    &svcConf.Server,
		debugMode:     svcConf.DebugMode,
	}
}

func (rt *Router) Run() {
	logrus.Info("Listening on port ", rt.config.serverConf.Port)
	err := rt.f.Engine().Run(fmt.Sprint(":", rt.config.serverConf.Port))
	if err != nil {
		panic(err)
	}
}

func (r *Router) init() {

	r.f = rest.NewFizzRouter(r.config.serverConf, pkg.SystemName, version.Version, r.config.debugMode)
	v1 := r.f.Group("/v1", "Node Gateway for Subscriber system ", "Node Gateway Subscriber system version v1")

	hss := v1.Group("/hlr", "HLR", "Home Location Register")
	hss.POST("/guti", formatDoc("Update GUTI", ""), tonic.Handler(r.postGuti, http.StatusOK))
	hss.POST("/imsi", formatDoc("Get Subscriber", ""), tonic.Handler(r.postReadImsiRecord, http.StatusOK))
	hss.POST("/tai", formatDoc("Update TAI", ""), tonic.Handler(r.postTai, http.StatusOK))
}

func formatDoc(summary string, description string) []fizz.OperationOption {
	return []fizz.OperationOption{func(info *openapi.OperationInfo) {
		info.Summary = summary
		info.Description = description
	}}
}

func (r *Router) postGuti(c *gin.Context, req *UpdateGutiReq) (*pb.UpdateGutiResp, error) {

	return r.clients.h.UpdateGuti(&pb.UpdateGutiReq{})
}

func (r *Router) postReadImsiRecord(c *gin.Context, req *GetRecordReq) (*pb.GetRecordResp, error) {

	return r.clients.h.ReadRecord(&pb.GetRecordReq{})
}

func (r *Router) postTai(c *gin.Context, req *UpdateTaiReq) (*pb.UpdateTaiResp, error) {

	return r.clients.h.UpdateTai(&pb.UpdateTaiReq{})
}
