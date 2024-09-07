package job

import (
	"fmt"

	"github.com/starudream/sign-task/pkg/geetest"
	"github.com/starudream/sign-task/pkg/miyoushe/api"
)

func verify(c *api.Client) (data *geetest.V3Data, _ error) {
	param, err := c.CreateVerification()
	if err != nil {
		return nil, fmt.Errorf("create verification error: %w", err)
	}

	data, err = dm(param)
	if err != nil {
		return nil, err
	}

	_, err = c.VerifyVerification(data.Challenge, data.Validate)
	if err != nil {
		return nil, fmt.Errorf("verify verification error: %w", err)
	}

	return
}

func dm(param *geetest.V3Param) (*geetest.V3Data, error) {
	param.Referer = api.RefererAct

	if geetest.RRKey() != "" {
		data, err := geetest.RR(param)
		if err == nil {
			return data, nil
		}
	}

	if geetest.TTKey() != "" {
		data, err := geetest.TT(param)
		if err == nil {
			return data, nil
		}
	}

	return nil, fmt.Errorf("no geetest config")
}
