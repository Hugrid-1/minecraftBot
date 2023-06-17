package config

import "os"

type AppConfig struct {
	MinecraftServer MinecraftServer
	ServerSettings  ServerSettings
	NotifySettings  NotifySettings
	WorkMode        string
}
type ServerSettings struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}

type MinecraftServer struct {
	IP   string `json:"ip"`
	Port string `json:"port"`
}
type NotifySettings struct {
	Email string `json:"email"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		MinecraftServer: MinecraftServer{
			IP:   os.Getenv("MINECRAFT_SERVER_IP"),
			Port: os.Getenv("MINECRAFT_SERVER_PORT"),
		},
		NotifySettings: NotifySettings{
			Email: os.Getenv("EMAIL_TO_NOTIFY"),
		},
		ServerSettings: ServerSettings{
			IP:   os.Getenv("SERVER_IP"),
			Port: os.Getenv("SERVER_PORT"),
		},
		WorkMode: os.Getenv("WORK_MODE"),
	}
}
