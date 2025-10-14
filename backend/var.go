package backend

const (
	configFolder = "Controller"
	configName   = "config.json"
)

type ConfigStruct struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
	CustomDir string `json:"customdir"`
}