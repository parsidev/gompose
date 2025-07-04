package crud

import (
	"github.com/Lumicrate/gompose/db"
	"github.com/Lumicrate/gompose/http"
	"reflect"
	"strings"
)

func RegisterCRUDRoutes(engine http.HTTPEngine, dbAdapter db.DBAdapter, entity any) {
	// Use reflection to get entity type name, e.g., "User"
	t := reflect.TypeOf(entity)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	entityName := t.Name()
	basePath := "/" + strings.ToLower(entityName) + "s" // simple pluralizer (add 's')

	// GET /entities (list)
	engine.RegisterRoute("GET", basePath, func(ctx http.Context) {
		handleGetAll(ctx, dbAdapter, entity)
	})

	// GET /entities/:id
	engine.RegisterRoute("GET", basePath+"/:id", func(ctx http.Context) {
		handleGetByID(ctx, dbAdapter, entity)
	})

	// POST /entities
	engine.RegisterRoute("POST", basePath, func(ctx http.Context) {
		handleCreate(ctx, dbAdapter, entity)
	})

	// PUT /entities/:id
	engine.RegisterRoute("PUT", basePath+"/:id", func(ctx http.Context) {
		handleUpdate(ctx, dbAdapter, entity)
	})

	// PATCH /entities/:id
	engine.RegisterRoute("PATCH", basePath+"/:id", func(ctx http.Context) {
		handlePatch(ctx, dbAdapter, entity)
	})

	// DELETE /entities/:id
	engine.RegisterRoute("DELETE", basePath+"/:id", func(ctx http.Context) {
		handleDelete(ctx, dbAdapter, entity)
	})
}
