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
		con.NewSection("Microservices")
		con.Section("Microservices").NewKey("pdfToTxtIp", "pdfToTxt")

		con.Section("Microservices").NewKey("txtToJsonIp", "txtToJson")
		con.Section("Microservices").NewKey("txtToJsonPort", "5000")
		con.Section("Microservices").NewKey("pdfToTxtPort", "5001")
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
