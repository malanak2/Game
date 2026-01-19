package config

import (
	"errors"
	"log/slog"

	"gopkg.in/ini.v1"
)

type conf struct {
	MaxFps int
	Fov    float32
}
type Config struct {
	Main conf
}

var Cfg Config

func InitConfig(path string) error {
	con, err := ini.Load(path)
	if err != nil {
		slog.Info("Generating default config file...")
		con = ini.Empty()
		_, _ = con.NewSection("Main")
		_, _ = con.Section("Main").NewKey("MaxFps", "144")
		_, _ = con.Section("Main").NewKey("Fov", "60")

		err = con.SaveTo(path)
		if err != nil {
			return errors.New("Failed to create new config " + err.Error())
		}
	}
	mapPoint := &Cfg
	err = con.MapTo(mapPoint)
	if err != nil {
		return errors.New("failed to parse config, please try again")
	}
	return nil
}
