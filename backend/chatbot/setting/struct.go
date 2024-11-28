package setting

var Setting settingStruct

type settingStruct struct {
	ServerPort string `json:"server_port"`
	Model      string `json:"model"`
	MaxToken   int    `json:"max_token"`
}
