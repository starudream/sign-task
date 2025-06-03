package test

import (
	"testing"

	"github.com/starudream/go-lib/core/v2/config"
	"github.com/starudream/go-lib/core/v2/utils/testutil"

	aliyun "github.com/starudream/sign-task/pkg/aliyun/config"
	"github.com/starudream/sign-task/pkg/cfg"
	douyu "github.com/starudream/sign-task/pkg/douyu/config"
	kuro "github.com/starudream/sign-task/pkg/kuro/config"
	miyoushe "github.com/starudream/sign-task/pkg/miyoushe/config"
	skland "github.com/starudream/sign-task/pkg/skland/config"
	tieba "github.com/starudream/sign-task/pkg/tieba/config"
	volcengine "github.com/starudream/sign-task/pkg/volcengine/config"
)

func Test(t *testing.T) {
	config.Set("aliyun", aliyun.Config{
		Accounts: []aliyun.Account{
			{
				Id:     "",
				Secret: "",
			},
		},
	})
	config.Set("douyu", douyu.Config{
		Accounts: []douyu.Account{
			{
				Phone: "",
				Did:   "",
				Ltp0:  "",
				Room:  0,
				Assigns: []douyu.Assign{
					{
						Count: 0,
						Room:  0,
						All:   false,
					},
				},
				IgnoreExpiredCheck: false,
			},
		},
	})
	config.Set("kuro", kuro.Config{
		Accounts: []kuro.Account{
			{
				Phone:   "",
				DevCode: "",
				Token:   "",
			},
		},
	})
	config.Set("miyoushe", miyoushe.Config{
		Accounts: []miyoushe.Account{
			{
				Phone: "",
				Device: miyoushe.Device{
					Id:      "",
					Type:    "",
					Name:    "",
					Model:   "",
					Version: "",
					Channel: "",
				},
				Mid:         "",
				SToken:      "",
				Uid:         "",
				CToken:      "",
				SignGameIds: []string{},
			},
		},
	})
	config.Set("skland", skland.Config{
		Accounts: []skland.Account{
			{
				Phone: "",
				Cred:  "",
				Token: "",
			},
		},
	})
	config.Set("tieba", tieba.Config{
		Accounts: []tieba.Account{
			{
				Phone: "",
				BDUSS: "",
			},
		},
	})
	config.Set("volcengine", volcengine.Config{
		Accounts: []volcengine.Account{
			{
				Id:     "",
				Secret: "",
			},
		},
	})

	// for _, v := range []string{"aliyun", "douyu", "kuro", "miyoushe", "skland", "tieba", "volcengine"} {
	// 	config.Set(v+".cron", cron.Config{
	// 		Disable: true,
	// 		Spec:    "0 0 4,12,20 * * *",
	// 		Startup: false,
	// 		Jitter:  0,
	// 	})
	// }

	testutil.LogNoErr(t, cfg.Save())
}
