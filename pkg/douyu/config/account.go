package config

type Account struct {
	Phone string `json:"phone" yaml:"phone"`

	Did  string `json:"did"  yaml:"did"`
	Ltp0 string `json:"ltp0" yaml:"ltp0"`

	Room    int      `json:"room"    yaml:"room"`
	Assigns []Assign `json:"assigns" yaml:"assigns" table:",ignore"`

	IgnoreExpiredCheck bool `json:"ignore_expired_check" yaml:"ignore_expired_check" table:",ignore"`
}

type Assign struct {
	Count int  `json:"count,omitempty" yaml:"count,omitempty"`
	Room  int  `json:"room,omitempty"  yaml:"room,omitempty"`
	All   bool `json:"all,omitempty"   yaml:"all,omitempty"`
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
