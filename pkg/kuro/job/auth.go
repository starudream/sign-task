package job

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/codec/json"
	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"

	"github.com/starudream/sign-task/pkg/geetest"
	"github.com/starudream/sign-task/pkg/kuro/api"
	"github.com/starudream/sign-task/pkg/kuro/config"
)

func Login(account config.Account) error {
	slog.Info("gtv4 id: %s", api.GT4Id)

	gts := fmtutil.Scan("please enter GeeTest json string: ")
	if gts == "" {
		return nil
	}

	gt, err := json.UnmarshalTo[*geetest.V4Data](gts)
	if err != nil {
		return fmt.Errorf("unmarshal geetest v4 data error: %w", err)
	}

	c := api.NewClient(account)

	data1, err := c.SendPhoneCode(gt)
	if err != nil {
		return fmt.Errorf("send phone code error: %w", err)
	}

	if data1.GeeTest {
		return fmt.Errorf("wrong geetest v4 data")
	}

	code := fmtutil.Scan("please enter the verification code you received: ")
	if code == "" {
		return nil
	}

	data2, err := c.LoginByPhoneCode(code)
	if err != nil {
		return fmt.Errorf("login by phone code error: %w", err)
	}

	account.Token = data2.Token

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return fmt.Errorf("save account error: %w", err)
	}

	slog.Info("login success")
	return nil
}
