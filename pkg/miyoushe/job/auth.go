package job

import (
	"encoding/base64"
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"

	"github.com/starudream/sign-task/pkg/miyoushe/api"
	"github.com/starudream/sign-task/pkg/miyoushe/config"
)

func Login(account config.Account) error {
	c := api.NewClient(account)

	_, err := c.SendPhoneCode("")
	if err != nil {
		if !api.IsRetCode(err, api.RetCodeSendPhoneNeedGeetest) {
			return fmt.Errorf("send phone code error: %w", err)
		}

		id, aigis := api.GetAigisData(err)
		if aigis == nil {
			return fmt.Errorf("get aigis data empty")
		}

		slog.Info("aigis gt: %s, challenge: %s", aigis.GT, aigis.Challenge)

		gts := fmtutil.Scan("please enter GeeTest json string: ")
		if gts == "" {
			return nil
		}

		b64 := base64.StdEncoding.EncodeToString([]byte(gts))

		_, err = c.SendPhoneCode(fmt.Sprintf("%s;%s", id, b64))
		if err != nil {
			return fmt.Errorf("send phone code error: %w", err)
		}
	}

	code := fmtutil.Scan("please enter the verification code you received: ")
	if code == "" {
		return nil
	}

	data, err := c.LoginByPhoneCode(code)
	if err != nil {
		return fmt.Errorf("login by phone code error: %w", err)
	}

	account.Uid = data.UserInfo.Aid
	account.Mid = data.UserInfo.Mid
	account.SToken = data.Token.Token

	_, err = Refresh(account)
	if err != nil {
		return err
	}

	slog.Info("login success")
	return nil
}

func Refresh(account config.Account) (config.Account, error) {
	c := api.NewClient(account)

	data, err := c.GetCTokenBySToken()
	if err != nil {
		return account, fmt.Errorf("get ctoken error: %w", err)
	}

	account.CToken = data.CookieToken

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}
