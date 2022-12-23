//go:build integration
// +build integration

package integration

import (
	"github.com/stretchr/testify/assert"
	"github.com/ukama/ukama/systems/common/config"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
)

// Before running test for the first time you have to create a test account in Identity manager and provide email and password for it

type TestConfig struct {
	ServiceHost string `default:"localhost:8080"`
}

var testConf *TestConfig

func init() {
	testConf = &TestConfig{}

	logrus.Info("Expected config ", "integration.yaml", " or env vars for ex: BASEDOMAIN")
	config.LoadConfig("integration", testConf)
	logrus.Infof("Config: %+v", testConf)
}

func Test_HlrClientApi(t *testing.T) {

	client := resty.New()

	t.Run("PostGUTI", func(tt *testing.T) {
		resp, err := client.R().
			EnableTrace().
			Post(getApiUrl() + "/v1/hlr/guti")

		if assert.NoError(t, err) {
			assert.Equal(tt, http.StatusBadRequest, resp.StatusCode())
			assert.Contains(tt, resp.String(), "Error invalid length")
		}
	})

	t.Run("PostTAI", func(tt *testing.T) {
		resp, err := client.R().
			EnableTrace().
			Post(getApiUrl() + "/v1/hlr/tai")

		if assert.NoError(t, err) {
			assert.Equal(tt, http.StatusNotFound, resp.StatusCode())
			assert.Contains(tt, resp.String(), "node record not found")
		}
	})

	t.Run("ReadHLRRecord", func(tt *testing.T) {
		resp, err := client.R().
			EnableTrace().
			Post(getApiUrl() + "/v1/hlr/imsi")

		if assert.NoError(t, err) {
			assert.Equal(tt, http.StatusNotFound, resp.StatusCode())
			assert.Contains(tt, resp.String(), "node record not found")
		}
	})

}

func getApiUrl() string {
	return testConf.ServiceHost
}
