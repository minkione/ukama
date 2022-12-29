package client

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/ukama/ukama/systems/common/rest"
)

type PolicyControl struct {
	R *rest.RestClient
}

func NewPolicyControlClient(url string, debug bool) (*PolicyControl, error) {

	f, err := rest.NewRestClient(url, debug)
	if err != nil {
		logrus.Errorf("Can't conncet to %s url.Error %s", url, err.Error)
		return nil, err
	}

	P := &PolicyControl{
		R: f,
	}

	return P, nil
}

func (P *PolicyControl) AddSim(pcrf PolicyControlSimInfo) error {

	errStatus := &ErrorMessage{}

	resp, err := P.R.C.R().
		SetError(errStatus).
		SetBody(pcrf).
		Put(P.R.Url.String() + "v1/pcrf/sims/" + pcrf.Imsi)

	if err != nil {
		logrus.Errorf("Failed to send api request to PCRF. Error %s", err.Error())
		return err
	}

	if !resp.IsSuccess() {
		logrus.Tracef("Failed to add sim to PCRF. HTTP resp code %d and Error message is %s", resp.StatusCode(), errStatus.Message)
		return fmt.Errorf(" failed to add sim to PCRF: %s", errStatus.Message)
	}

	return nil
}

func (P *PolicyControl) UpdateSim(pcrf PolicyControlSimPackageUpdate) error {

	errStatus := &ErrorMessage{}

	resp, err := P.R.C.R().
		SetError(errStatus).
		SetBody(pcrf).
		Patch(P.R.Url.String() + "v1/pcrf/sims/" + pcrf.Imsi)

	if err != nil {
		logrus.Errorf("Failed to send api request to PCRF. Error %s", err.Error())
		return err
	}

	if !resp.IsSuccess() {
		logrus.Tracef("Failed to update sim info in PCRF. HTTP resp code %d and Error message is %s", resp.StatusCode(), errStatus.Message)
		return fmt.Errorf("update failure in PCRF%s", errStatus.Message)
	}

	return nil
}

func (P *PolicyControl) DeleteSim(imsi string) error {

	errStatus := &ErrorMessage{}

	resp, err := P.R.C.R().
		SetError(errStatus).
		Delete(P.R.Url.String() + "v1/pcrf/sims/" + imsi)

	if err != nil {
		logrus.Errorf("Failed to send api request to PCRF. Error %s", err.Error())
		return err
	}

	if !resp.IsSuccess() {
		logrus.Tracef("Failed to delete sim from PCRF. HTTP resp code %d and Error message is %s", resp.StatusCode(), errStatus.Message)
		return fmt.Errorf(" failed to remove sim from PCRF: %s", errStatus.Message)
	}

	return nil
}
