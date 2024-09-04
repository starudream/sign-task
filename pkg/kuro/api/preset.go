package api

const (
	Addr          = "https://api.kurobbs.com"
	UserAgent     = "okhttp/3.11.0"
	AndroidName   = "com.kurogame.kjq"
	Version       = "2.2.0"
	SourceAndroid = "android"

	GameIdPNS = 2 // 战双
	GameIdMC  = 3 // 鸣潮

	ForumIdPNS3 = 3  // 伊甸闲庭
	ForumIdMC10 = 10 // 今州茶馆

	GT4Id = "3f7e2d848ce0cb7e7d019d621e556ce2"
)

const (
	CodeHasSigned = Code(1511)
)

var (
	GameNameById = map[int]string{
		GameIdPNS: "战双",
		GameIdMC:  "鸣潮",
	}

	ForumIdByGameId = map[int]int{
		GameIdPNS: ForumIdPNS3,
		GameIdMC:  ForumIdMC10,
	}
)
