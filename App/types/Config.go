package types

import (
	"errors"
	"log/slog"

	"gopkg.in/ini.v1"
)

type conf struct {
	maxFps int
}
type Config struct {
	conf
}

func NewConfig(path string) (*Config, error) {
	con, err := ini.Load(path)
	if err != nil {
		slog.Info("Generating default config file...")
		con = ini.Empty()
		con.NewSection("Main")
		con.Section("Main").NewKey("shit", "idk mann")

		err = con.SaveTo("configMain.ini")
		if err != nil {
			return nil, errors.New("Failed to create new config " + err.Error())
		}
	}
	var mapped Config
	mapPoint := &mapped
	err = con.MapTo(mapPoint)
	if err != nil {
		return nil, errors.New("Failed to parse config, please try again.")
	}
	return &mapped, nil
}
