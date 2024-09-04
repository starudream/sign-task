package config

import (
	"sync"

	"github.com/starudream/go-lib/core/v2/config"

	"github.com/starudream/sign-task/pkg/cfg"
)

type Config struct {
	Accounts cfg.Accounts[Account] `json:"accounts" yaml:"accounts"`
}

var (
	_c   = Config{}
	_cMu = sync.Mutex{}
)

func init() {
	_ = config.Unmarshal("kuro", &_c)
}

func C() Config {
	_cMu.Lock()
	defer _cMu.Unlock()
	return _c
}

func Save() error {
	config.Set("kuro.accounts", _c.Accounts)
	return cfg.Save()
}
