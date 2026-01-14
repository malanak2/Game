package types

import (
	"log/slog"
	"time"

	"github.com/fogleman/gg"
)

type App struct {
	appState AppState
	config   *Config

	lastFrameTime time.Time
}

func InitApp(path *string) (*App, error) {
	config, err := NewConfig(*path)
	if err != nil {
		return nil, err
	}
	app := App{AppState{GraphicalCtx{gg.NewContext(100, 100), []*DrawInfo{}}}, config, time.Now()}

	return &app, nil
}

func (a *App) Run() {
	for {
		a.appState.gCtx.Render()
		oldTime := a.lastFrameTime
		a.lastFrameTime = time.Now()
		slog.Info("FPS: ", "fps", 1/a.lastFrameTime.Sub(oldTime).Seconds())
	}
}
