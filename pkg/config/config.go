package config

var ServerConfig Config

type Config struct {
	Token string `json:"token"`
}
