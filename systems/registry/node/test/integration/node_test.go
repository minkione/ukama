//go:build integration
// +build integration

package integration

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ukama/ukama/systems/common/ukama"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/num30/config"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	uconf "github.com/ukama/ukama/systems/common/config"
	pb "github.com/ukama/ukama/systems/registry/node/pb/gen"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

var tConfig *TestConfig
var orgName string

func init() {
	// set org name
	orgName = fmt.Sprintf("node-integration-self-test-org-%d", time.Now().Unix())

	// load config
	tConfig = &TestConfig{}

	err := config.NewConfReader("integration").Read(tConfig)
	if err != nil {
		logrus.Fatal("Error reading config ", err)
	} else if tConfig.DebugMode {
		b, err := yaml.Marshal(tConfig)
		if err != nil {
			logrus.Infof("Config:\n%s", string(b))
		}
	}

	logrus.Info("Expected config ", "integration.yaml", " or env vars for ex: SERVICEHOST")
	logrus.Infof("Config: %+v\n", tConfig)
}

type TestConfig struct {
	uconf.BaseConfig `mapstructure:",squash"`
	ServiceHost      string `default:"localhost:9090"`
}

func Test_FullFlow(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	logrus.Infoln("Connecting to network ", tConfig.ServiceHost)
	conn, err := grpc.DialContext(ctx, tConfig.ServiceHost, grpc.WithInsecure(), grpc.WithBlock())

	if err != nil {
		assert.NoError(t, err, "did not connect: %v", err)

		return
	}

	c := pb.NewNodeServiceClient(conn)
	// keep all used nodes here so we could delete them after test
	ndToClean := []ukama.NodeID{}

	// Contact the server and print out its response.
	node := ukama.NewVirtualHomeNodeId()
	ndToClean = append(ndToClean, node)

	defer cleanupNodes(t, c, ndToClean)

	var r interface{}

	t.Run("AddAndUpdateNode", func(tt *testing.T) {
		nodeName := "HomeNodeX"
		addResp, err := c.AddNode(ctx, &pb.AddNodeRequest{
			Node: &pb.Node{
				NodeId: node.String(),
				State:  pb.NodeState_UNDEFINED,
				Name:   nodeName,
			},
		})
		handleResponse(tt, err, addResp)
		assert.NotNil(tt, addResp.Node)
		assert.Equal(tt, nodeName, addResp.Node.Name)

		r, err = c.UpdateNodeState(ctx, &pb.UpdateNodeStateRequest{NodeId: node.String(), State: pb.NodeState_ONBOARDED})
		handleResponse(tt, err, r)

		nodeResp, err := c.GetNode(ctx, &pb.GetNodeRequest{NodeId: node.String()})
		handleResponse(tt, err, nodeResp)
		assert.Equal(tt, pb.NodeState_ONBOARDED, nodeResp.Node.State)
		assert.Equal(tt, pb.NodeType_HOME, nodeResp.Node.Type)
	})

	tNodeID := ukama.NewVirtualNodeId(ukama.NODE_ID_TYPE_TOWERNODE)
	aNodeID := ukama.NewVirtualNodeId(ukama.NODE_ID_TYPE_AMPNODE)

	t.Run("AddTowerNodeWithAmplifiers", func(tt *testing.T) {
		ndToClean = append(ndToClean, tNodeID)
		_, err := c.AddNode(ctx, &pb.AddNodeRequest{
			Node: &pb.Node{
				NodeId: tNodeID.String(),
				State:  pb.NodeState_UNDEFINED,
			},
		})
		if err != nil {
			assert.FailNow(tt, "AddNode failed", err.Error())
		}

		ndToClean = append(ndToClean, aNodeID)
		_, err = c.AddNode(ctx, &pb.AddNodeRequest{
			Node: &pb.Node{
				NodeId: aNodeID.String(),
				State:  pb.NodeState_UNDEFINED,
			},
		})
		if err != nil {
			assert.FailNow(tt, "AddNode failed", err.Error())
		}

		_, err = c.AttachNodes(ctx, &pb.AttachNodesRequest{
			ParentNodeId:    tNodeID.String(),
			AttachedNodeIds: []string{aNodeID.String()},
		})
		if err != nil {
			assert.FailNow(tt, "AttachNodes failed", err.Error())
		}

		resp, err := c.GetNode(ctx, &pb.GetNodeRequest{NodeId: tNodeID.String()})
		if assert.NoError(tt, err, "GetNode failed") {
			assert.NotNil(tt, resp.Node.Attached)
			assert.Equal(tt, 1, len(resp.Node.Attached))
			assert.Equal(tt, aNodeID.StringLowercase(), resp.Node.Attached[0].NodeId)
		}
	})

	t.Run("DetachNode", func(tt *testing.T) {
		_, err := c.DetachNode(ctx, &pb.DetachNodeRequest{
			DetachedNodeId: aNodeID.String(),
		})
		if assert.NoError(t, err) {
			resp, err := c.GetNode(ctx, &pb.GetNodeRequest{NodeId: tNodeID.String()})
			if assert.NoError(t, err) {
				assert.Nil(t, resp.Node.Attached)
			}
		}
	})
}

func cleanupNodes(tt *testing.T, c pb.NodeServiceClient, nodes []ukama.NodeID) {
	for _, node := range nodes {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		_, err := c.Delete(ctx, &pb.DeleteRequest{NodeId: node.String()})
		if err != nil {
			assert.FailNow(tt, "DeleteNode failed", err.Error())
		}

		_, err = c.GetNode(ctx, &pb.GetNodeRequest{NodeId: node.String()})
		if assert.Error(tt, err) {
			assert.Equal(tt, codes.NotFound, status.Code(err))
		}
	}
}

func Test_Listener(t *testing.T) {
	// Arrange
	nodeID := "UK-SA2136-HNODE-A1-30DF"

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, c, err := CreateRegistryClient()
	defer conn.Close()

	if err != nil {
		assert.NoErrorf(t, err, "did not connect: %+v\n", err)

		return
	}

	_, err = c.AddNode(ctx, &pb.AddNodeRequest{Node: &pb.Node{
		NodeId: nodeID, State: pb.NodeState_UNDEFINED,
	}})

	e, ok := status.FromError(err)
	if ok && e.Code() == codes.AlreadyExists {
		logrus.Infof("node already exist, err %+v\n", err)
	} else {
		assert.NoError(t, err)

		return
	}

	_, err = c.UpdateNodeState(ctx, &pb.UpdateNodeStateRequest{NodeId: nodeID, State: pb.NodeState_PENDING})
	if err != nil {
		assert.FailNow(t, "Failed to update node. Error: %s", err.Error())
	}

	// Act

	// Assert
	assert.NoError(t, err)

	time.Sleep(3 * time.Second)

	nodeResp, err := c.GetNode(ctx, &pb.GetNodeRequest{NodeId: nodeID})
	assert.NoError(t, err)

	if err != nil {
		assert.Equal(t, pb.NodeState_ONBOARDED, nodeResp.Node.State)
	}
}

func CreateRegistryClient() (*grpc.ClientConn, pb.NodeServiceClient, error) {
	logrus.Infoln("Connecting to network ", tConfig.ServiceHost)

	context, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	conn, err := grpc.DialContext(context, tConfig.ServiceHost, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	c := pb.NewNodeServiceClient(conn)

	return conn, c, nil
}

func handleResponse(t *testing.T, err error, r interface{}) {
	t.Helper()

	fmt.Printf("Response: %v\n", r)

	if err != nil {
		assert.FailNow(t, "Request failed: %v\n", err)
	}
}
