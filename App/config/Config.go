package config

import (
	"errors"
	"log/slog"

	"gopkg.in/ini.v1"
)

type confMain struct {
	MaxFps          int
	Fov             float32
	CameraMovespeed float32
}

type confDev struct {
	DebugMode bool
	Dev       bool
}
type Config struct {
	Main confMain
	Dev  confDev
}

var Cfg Config

func InitConfig(path string) error {
	con, err := ini.Load(path)
	if err != nil {
		slog.Info("Generating default config file...")
		con = ini.Empty()
		_, _ = con.NewSection("Main")
		_, _ = con.Section("Main").NewKey("MaxFps", "144")
		_, _ = con.Section("Main").NewKey("Fov", "45")
		_, _ = con.Section("Main").NewKey("CameraMovespeed", "50")
		_, _ = con.NewSection("Dev")
		_, _ = con.Section("Dev").NewKey("Debug", "false")
		_, _ = con.Section("Dev").NewKey("Dev", "false")

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
