package setting

var Setting settingStruct

type settingStruct struct {
	ServerPort string `json:"server_port"`
	MaxToken   int    `json:"max_token"`
}
