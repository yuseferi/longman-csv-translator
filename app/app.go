package app

import (
	"context"
	_ "github.com/jackc/pgx/v4/stdlib"
	"go.uber.org/zap"
	"math"
	"net/http"
	"sync"
)

type Application struct {
	Config                    *Config
	Ctx                       context.Context
	Error                     chan error
	Http                      http.Handler
	Logger                    *zap.Logger
	WaitGroup                 *sync.WaitGroup
	ctxCancel                 context.CancelFunc

}

func New(config *Config) (app *Application, err error) {
	app = &Application{
		Error:     make(chan error, math.MaxUint8),
		WaitGroup: new(sync.WaitGroup),
		Config:    config,
	}

	app.Ctx, app.ctxCancel = context.WithCancel(context.Background())
	defer func() {
		if err != nil {
			app.Close()
		}
	}()

	app.Logger, err = NewLogger(app.Config.Level)
	if err != nil {
		return nil, err
	}
	app.Logger.Debug("debug mode on")

	return app, nil
}

func (app *Application) Close() {
	app.Logger.Debug("Application stops")
}

func (app *Application) Run() {
	// Run Importer
		if err := app.ScrapeAll(); err != nil {
		app.Logger.Panic("HTTP Server start error", zap.Error(err))
	}
}


