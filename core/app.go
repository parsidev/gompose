package core

import (
	"log"

	"github.com/Lumicrate/gompose/auth"
	"github.com/Lumicrate/gompose/crud"
	"github.com/Lumicrate/gompose/db"
	"github.com/Lumicrate/gompose/docs/swagger"
	"github.com/Lumicrate/gompose/http"
	"github.com/Lumicrate/gompose/i18n"
)

type App struct {
	entities        []registeredEntity
	dbAdapter       db.DBAdapter
	httpEngine      http.HTTPEngine
	middlewares     []http.MiddlewareFunc
	authProvider    auth.AuthProvider
	swaggerProvider *swagger.SwaggerProvider
	localization    *i18n.Translator
}

type registeredEntity struct {
	entity any
	config *crud.Config
}

func NewApp() *App {
	return &App{
		entities:    []registeredEntity{},
		middlewares: []http.MiddlewareFunc{},
	}
}

func (a *App) Entities() []any {
	raw := make([]any, len(a.entities))
	for i, e := range a.entities {
		raw[i] = e.entity
	}
	return raw
}

func (a *App) AddEntity(entity any, opts ...crud.Option) *App {
	cfg := crud.DefaultConfig()
	for _, opt := range opts {
		opt(cfg)
	}
	a.entities = append(a.entities, registeredEntity{entity: entity, config: cfg})
	return a
}

func (a *App) UseDB(adapter db.DBAdapter) *App {
	a.dbAdapter = adapter
	return a
}

func (a *App) UseHTTP(engine http.HTTPEngine) *App {
	a.httpEngine = engine
	return a
}

func (a *App) RegisterMiddleware(m http.MiddlewareFunc) *App {
	a.middlewares = append(a.middlewares, m)
	return a
}

func (a *App) UseAuth(provider auth.AuthProvider) *App {
	a.authProvider = provider
	return a
}

func (a *App) UseI18n(directory, defaultLocale string) *App {
	var err error

	if a.localization, err = i18n.NewI18n(directory, defaultLocale); err != nil {
		log.Fatalf("i18n Init failed: %v", err)
	}

	a.middlewares = append(a.middlewares, a.localization.GetMiddleware())

	return a
}

func (a *App) SetLocale(locale string) *App {
	a.localization = a.localization.SetLocale(locale)

	return a
}

func (a *App) T(messageID string, args ...any) string {
	return a.localization.T(messageID, args...)
}

func (a *App) UseSwagger() *App {
	a.swaggerProvider = swagger.NewSwaggerProvider()
	return a
}

func (a *App) Run() {
	if a.dbAdapter != nil {
		if err := a.dbAdapter.Init(); err != nil {
			log.Fatalf("DB Init failed: %v", err)
		}

		if err := a.dbAdapter.Migrate(a.Entities()); err != nil {
			log.Fatalf("DB Migration failed: %v", err)
		}
	}

	if a.authProvider != nil {
		err := a.authProvider.Init()
		if err != nil {
			return
		}
		a.authProvider.RegisterRoutes(a.httpEngine)
	}

	for _, m := range a.middlewares {
		a.httpEngine.Use(m)
	}

	for _, e := range a.entities {
		crud.RegisterCRUDRoutes(a.httpEngine, a.dbAdapter, e.entity, e.config, a.authProvider)
	}

	if a.swaggerProvider != nil {
		a.swaggerProvider.RegisterRoutes(a.httpEngine)
	}

	if err := a.httpEngine.Start(); err != nil {
		log.Fatalf("HTTP Server failed: %v", err)
	}
}
