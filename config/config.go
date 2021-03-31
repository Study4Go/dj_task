package config

import (
	"github.com/BurntSushi/toml"
	"time"
)

// A config is config info struct
type config struct {
	Server     server `toml:"server"`
	Redis      redis  `toml:"redis"`
	LogPath    string `toml:"busi_log_path"`
	AccLogPath string `toml:"access_log_path"`
}

type server struct {
	IP   string
	Port int
}

type redis struct {
	// ip:port
	Address string
	// max idle connect
	MaxIdle int
	// max connect
	MaxActive int
	// after this time(second),idle connect will be recycled
	IdleTimeout time.Duration
}

// Config is base config for project
var Config config

// InitConfig to initial base config
func InitConfig(confPath string, port int) {
	if _, err := toml.DecodeFile(confPath, &Config); err != nil {
		panic("initial config failed")
	}
	if port != -1 {
		Config.Server.Port = port
	}
}
