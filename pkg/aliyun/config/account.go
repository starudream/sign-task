package config

type Account struct {
	Key    string `json:"key"    yaml:"key"`
	Secret string `json:"secret" yaml:"secret"`
}

func (account Account) GetKey() string {
	return account.Key
}

func AddAccount(account Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	_c.Accounts = _c.Accounts.Add(account)
}

func UpdateAccount(phone string, cb func(account Account) Account) {
	_cMu.Lock()
	defer _cMu.Unlock()
	_c.Accounts = _c.Accounts.Update(phone, cb)
}

func GetAccount(phone string) (Account, bool) {
	accounts := C().Accounts
	return accounts.Get(phone)
}

func GetAccountOrFirst(phones ...string) Account {
	accounts := C().Accounts
	return accounts.GetOrFirst(phones...)
}
