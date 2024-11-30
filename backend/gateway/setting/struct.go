package setting

var Setting settingStruct

type settingStruct struct {
	ServerPort           string `json:"server_port"`
	ImageGeneratorServer string `json:"image_generator_server"`
	ChatbotServer        string `json:"chatbot_server"`
}
