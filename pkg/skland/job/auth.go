package job

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/fmtutil"

	"github.com/starudream/sign-task/pkg/skland/api"
	"github.com/starudream/sign-task/pkg/skland/config"
)

func Login(account config.Account) error {
	c := api.NewClient(account)

	err := c.SendPhoneCode()
	if err != nil {
		return fmt.Errorf("send phone code error: %w", err)
	}

	code := fmtutil.Scan("please enter the verification code you received: ")
	if code == "" {
		return nil
	}

	data1, err := c.LoginByPhoneCode(code)
	if err != nil {
		return fmt.Errorf("login by phone code error: %w", err)
	}

	data2, err := c.GrantApp(data1.Token, api.AppCodeSkland)
	if err != nil {
		return fmt.Errorf("grant app error: %w", err)
	}

	data3, err := c.AuthLoginByCode(data2.Code)
	if err != nil {
		return fmt.Errorf("auth login by code error: %w", err)
	}

	account.Cred = data3.Cred
	account.Token = data3.Token

	_, err = Refresh(account)
	if err != nil {
		return err
	}

	slog.Info("login success")
	return nil
}

func Refresh(account config.Account) (config.Account, error) {
	c := api.NewClient(account)

	data, err := c.AuthRefresh(account.Cred)
	if err != nil {
		return account, fmt.Errorf("auth refresh error: %w", err)
	}

	account.Token = data.Token

	config.UpdateAccount(account.Phone, func(config.Account) config.Account { return account })
	err = config.Save()
	if err != nil {
		return account, fmt.Errorf("save account error: %w", err)
	}

	return account, nil
}
