package core

import (
	"Gompose/crud"
	"Gompose/db"
	"Gompose/http"
	"log"
)

type App struct {
	entities    []any
	dbAdapter   db.DBAdapter
	httpEngine  http.HTTPEngine
	middlewares []http.MiddlewareFunc
}

// NewApp creates a new app instance
func NewApp() *App {
	return &App{
		entities:    []any{},
		middlewares: []http.MiddlewareFunc{},
	}
}

func (a *App) Entities() []any {
	return a.entities
}

func (a *App) AddEntity(entity any) *App {
	a.entities = append(a.entities, entity)
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

func (a *App) Run() {
	if err := a.dbAdapter.Init(); err != nil {
		log.Fatalf("DB Init failed: %v", err)
	}

	if err := a.dbAdapter.Migrate(a.entities); err != nil {
		log.Fatalf("DB Migration failed: %v", err)
	}

	for _, m := range a.middlewares {
		a.httpEngine.Use(m)
	}

	for _, e := range a.entities {
		crud.RegisterCRUDRoutes(a.httpEngine, a.dbAdapter, e)
	}

	if err := a.httpEngine.Start(); err != nil {
		log.Fatalf("HTTP Server failed: %v", err)
	}
}
