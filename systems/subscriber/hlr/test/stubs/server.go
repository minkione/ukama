package stubs

import "github.com/ukama/ukama/systems/subscriber/hlr/pkg/client"

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
		HttpServices:      &HttpEndpoints{
			Timeout:     3 * time.Second,
		},
		serverConf:    &rest.HttpConfig,
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

	r.f = rest.NewFizzRouter(r.config.serverConf, "hlr-stubs", "0,.0.0.", r.config.debugMode)
	v1 := r.f.Group("/v1", " ", " HLR service stubs version v1")

	f := v1.Group("/factory", "Sim Factory", "Sim Factory")
	f.GET("/simcards/:iccid", formatDoc("Read Sim card Information", ""), tonic.Handler(r.getSimCard, http.StatusOK))

	n := v1.Group("/networks", "Registry", "Network Factory")
	n.GET("/:network/orgs/:org", formatDoc("Validate Network", ""), tonic.Handler(r.getValidateNetwork, http.StatusOK))

	p := v1.Group("/pcrf", "PCRF", "Policy control")
	g.PUT("/sims/:imsi", formatDoc("Add Sim", ""), tonic.Handler(r.putSim, http.StatusOK))
	g.Delete("/sims/:imsi", formatDoc("Delete Sim", ""), tonic.Handler(r.deleteSim, http.StatusOK))
	g.PATCH("/sims/:imsi", formatDoc("Update Sim Packege Info", ""), tonic.Handler(r.patchSim, http.StatusOK))
	
}

func formatDoc(summary string, description string) []fizz.OperationOption {
	return []fizz.OperationOption{func(info *openapi.OperationInfo) {
		info.Summary = summary
		info.Description = description
	}}
}

func (r *Router) getSimCard(c *gin.Context, req *SimCardInfoReq) (*client.SimCardInfo, error) {
	s := client.SimCardInfo{
	Iccid:	req.Iccid,
	Imsi: "001010123456789",	
	Op: []byte("0123456789012345"),
	Key: []byte("0123456789012345"),
	Amf: []byte("800"),
	AlgoType: 1,
	UeDlAmbrBps: 2000000,
	UeUlAmbrBps: 2000000,
	sqn: 1,
	CsgIdPrsent: false,
	CsgId: 0,
	DefaultApnName: "ukama",
	}

	return &s,nil
}

func (r *Router) getValidateNetwork(c *gin.Context, req *NetworkValidationReq) error {
	/* No implementaion always return success.*/
	return nil
}

func (r *Router) putSim(c *gin.Context, req *client.PolicyControlSimInfo) error {
	/* No implementaion always return success.*/
	return nil
}

func (r *Router) deleteSim(c *gin.Context, req *DeleteSimReq) error {
	/* No implementaion always return success.*/
	return nil
}

func (r *Router) patchSim(c *gin.Context, req *client.PolicyControlSimPackageUpdate) error {
	/* No implementaion always return success.*/
	return nil
}