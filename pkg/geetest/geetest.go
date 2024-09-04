package geetest

type V3Param struct {
	Key        string `json:"key,omitempty"`
	ItemId     string `json:"itemid,omitempty"`
	Referer    string `json:"referer,omitempty"`
	GT         string `json:"gt"`
	Challenge  string `json:"challenge"`
	NewCaptcha int    `json:"new_captcha,omitempty"`
	Success    int    `json:"success,omitempty"`
}

type V3Data struct {
	Challenge string `json:"challenge"`
	Validate  string `json:"validate"`
	Seccode   string `json:"seccode,omitempty"`
}

type V4Data struct {
	CaptchaId     string `json:"captcha_id"`
	LotNumber     string `json:"lot_number,omitempty"`
	PassToken     string `json:"pass_token,omitempty"`
	GenTime       string `json:"gen_time,omitempty"`
	CaptchaOutput string `json:"captcha_output,omitempty"`
}
