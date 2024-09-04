package config

type Account struct {
	Phone   string `json:"phone"    yaml:"phone"`
	DevCode string `json:"dev_code" yaml:"dev_code"`
	Token   string `json:"token"    yaml:"token"   table:",ignore"`
}

func (account Account) GetKey() string {
	return account.Phone
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
