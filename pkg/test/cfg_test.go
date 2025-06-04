package test

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/testutil"

	aliyun "github.com/starudream/sign-task/pkg/aliyun/config"
	"github.com/starudream/sign-task/pkg/cfg"
	"github.com/starudream/sign-task/pkg/cron"
	douyu "github.com/starudream/sign-task/pkg/douyu/config"
	kuro "github.com/starudream/sign-task/pkg/kuro/config"
	miyoushe "github.com/starudream/sign-task/pkg/miyoushe/config"
	skland "github.com/starudream/sign-task/pkg/skland/config"
	tieba "github.com/starudream/sign-task/pkg/tieba/config"
	volcengine "github.com/starudream/sign-task/pkg/volcengine/config"
)

func Test(t *testing.T) {
	config.Set("aliyun.accounts", []aliyun.Account{{
		Id:     "aliyun_id",
		Secret: "aliyun_secret",
	}})
	config.Set("douyu.accounts", []douyu.Account{{
		Phone: "douyu_phone",
		Did:   "douyu_did",
		Ltp0:  "douyu_ltp0",
		Room:  9999,
		Assigns: []douyu.Assign{
			{
				Count: 1,
			},
			{
				Room: 9999,
				All:  true,
			},
		},
		IgnoreExpiredCheck: false,
	}})
	config.Set("kuro.accounts", []kuro.Account{{
		Phone:   "kuro_phone",
		DevCode: "kuro_dev_code",
		Token:   "kuro_token",
	}})
	config.Set("miyoushe.accounts", []miyoushe.Account{{
		Phone: "miyoushe_phone",
		Device: miyoushe.Device{
			Id:      "device_id",
			Type:    "device_type",
			Name:    "device_name",
			Model:   "device_model",
			Version: "device_version",
			Channel: "device_channel",
		},
		Mid:         "miyoushe_mid",
		SToken:      "miyoushe_stoken",
		Uid:         "miyoushe_uid",
		CToken:      "miyoushe_ctoken",
		SignGameIds: []string{"6"},
	}})
	config.Set("skland.accounts", []skland.Account{{
		Phone: "skland_phone",
		Cred:  "skland_cred",
		Token: "skland_token",
	}})
	config.Set("tieba.accounts", []tieba.Account{{
		Phone: "tieba_phone",
		BDUSS: "tieba_bduss",
	}})
	config.Set("volcengine.accounts", []volcengine.Account{{
		Id:     "volcengine_id",
		Secret: "volcengine_secret",
	}})

	for _, v := range []string{
		"aliyun",
		"douyu",
		"kuro",
		"miyoushe",
		"skland",
		"tieba",
		"volcengine",
		"geetest",
	} {
		config.Set(v+".cron", cron.Config{
			Disable: true,
			Spec:    "0 0 12 * * *",
			Startup: false,
			Jitter:  10,
		})
	}

	testutil.LogNoErr(t, cfg.Save())
}
