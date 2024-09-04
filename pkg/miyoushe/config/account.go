package config

type Account struct {
	Phone  string `json:"phone"  yaml:"phone"`
	Device Device `json:"device" yaml:"device"`

	Mid    string `json:"mid"    yaml:"mid"`
	SToken string `json:"stoken" yaml:"stoken" table:",ignore"`

	Uid    string `json:"uid"    yaml:"uid"`
	CToken string `json:"ctoken" yaml:"ctoken" table:",ignore"`

	SignGameIds []string `json:"sign_game_ids" yaml:"sign_game_ids" table:",ignore"`
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
