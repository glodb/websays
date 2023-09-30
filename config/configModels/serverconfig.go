package configModels

//Structure for reading server config
type ServerConfig struct {
	Address string `json:"address"`
	Port    string `json:"port"`
}
