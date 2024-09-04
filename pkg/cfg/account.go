package cfg

import (
	"fmt"

	"github.com/starudream/go-lib/core/v2/slog"
	"github.com/starudream/go-lib/core/v2/utils/osutil"

	"github.com/starudream/sign-task/util"
)

type Account interface {
	GetKey() string
}

type Accounts[T Account] []T

func (accounts Accounts[Account]) Add(account Account) Accounts[Account] {
	for i := range accounts {
		if accounts[i].GetKey() == account.GetKey() {
			accounts[i] = account
			return accounts
		}
	}
	return append(accounts, account)
}

func (accounts Accounts[Account]) Update(key string, cb func(account Account) Account) Accounts[Account] {
	for i := range accounts {
		if accounts[i].GetKey() == key {
			a := accounts[i]
			b := cb(a)
			slog.Info("update account %s, diff: %s", key, util.Diff(a, b))
			accounts[i] = b
		}
	}
	return accounts
}

func (accounts Accounts[Account]) Get(key string) (account Account, _ bool) {
	for i := range accounts {
		if accounts[i].GetKey() == key {
			return accounts[i], true
		}
	}
	return account, false
}

func (accounts Accounts[Account]) GetOrFirst(keys ...string) Account {
	if len(accounts) == 0 {
		osutil.ExitErr(fmt.Errorf("no account exists"))
	}
	for _, key := range keys {
		if key == "" {
			continue
		}
		for i := range accounts {
			if accounts[i].GetKey() == key {
				return accounts[i]
			}
		}
	}
	if len(keys) > 0 {
		osutil.ExitErr(fmt.Errorf("account not exists"))
	}
	return accounts[0]
}
