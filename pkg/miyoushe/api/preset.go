package api

const (
	AddrPassport     = "https://passport-api.mihoyo.com"
	AddrHK4E         = "https://hk4e-sdk.mihoyo.com"
	AddrTakumi       = "https://api-takumi.mihoyo.com"
	AddrTakumiRecord = "https://api-takumi-record.mihoyo.com"
	AddrActNap       = "https://act-nap-api.mihoyo.com"
	AddrBBS          = "https://bbs-api.miyoushe.com"

	RefererApp = "https://app.mihoyo.com"
	RefererAct = "https://act.mihoyo.com"

	AppVersion    = "2.73.1"
	AppIdMiyoushe = "bll8iq97cem8" // 米游社

	UserAgent = "Mozilla/5.0 (Linux; Android 13; 22011211C Build/TP1A.220624.014; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/104.0.5112.97 Mobile Safari/537.36 miHoYoBBS/" + AppVersion

	GameIdBH3 = "1" // 崩坏3
	GameIdYS  = "2" // 原神
	GameIdBH2 = "3" // 崩坏学园2
	GameIdWD  = "4" // 未定事件簿
	GameIdDBY = "5" // 大别野
	GameIdSR  = "6" // 崩坏：星穹铁道
	GameIdZZZ = "8" // 绝区零

	GameNameBH3 = "bh3"
	GameNameYS  = "hk4e"
	GameNameSR  = "hkrpg"
	GameNameZZZ = "nap"

	GameBizBH3CN = GameNameBH3 + "_" + cn
	GameBizYSCN  = GameNameYS + "_" + cn
	GameBizSRCN  = GameNameSR + "_" + cn
	GameBizZZZCN = GameNameZZZ + "_" + cn

	cn = "cn"

	ForumIdSR = "53"

	xRpcAigis    = "x-rpc-aigis"
	xRpcSignGame = "x-rpc-signgame"
)

const (
	RetCodeSendPhoneNeedGeetest = -3101 // 发送验证码需要验证码
	RetCodeGameHasSigned        = -5003 // 已签到
	RetCodeForumHasSigned       = 1008  // 打卡失败或重复打卡
	RetCodeForumNeedGeetest     = 1034  // 需要验证码
)

var (
	GameIdByName = map[string]string{
		GameNameBH3: GameIdBH3,
		GameNameYS:  GameIdYS,
		GameNameSR:  GameIdSR,
		GameNameZZZ: GameIdZZZ,
	}

	GameIdByBiz = map[string]string{
		GameBizBH3CN: GameIdBH3,
		GameBizYSCN:  GameIdYS,
		GameBizSRCN:  GameIdSR,
		GameBizZZZCN: GameIdZZZ,
	}

	GameCNNameById = map[string]string{
		GameIdBH3: "崩坏3",
		GameIdYS:  "原神",
		GameIdSR:  "崩坏：星穹铁道",
		GameIdZZZ: "绝区零",
	}
)
