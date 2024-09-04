package api

const (
	AddrHypergryph = "https://as.hypergryph.com"
	AddrZonai      = "https://zonai.skland.com"
	UserAgent      = "Skland/1.5.1 (com.hypergryph.skland; build:100501001; Android 33; ) Okhttp/4.11.0"

	AppCodeSkland = "4ca99fa6b56cc2ba"
	Platform      = "1"
	VName         = "1.5.1"
	DId           = "743a446c83032899"

	GameIdArknights = "1" // 明日方舟

	GameAppCodeArknights = "arknights"
)

const (
	CodeGameHasSigned = 10001
)

var (
	GameIdByCode = map[string]string{
		GameAppCodeArknights: GameIdArknights,
	}
)
