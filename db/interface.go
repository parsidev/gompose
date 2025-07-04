package db

type Pagination struct {
	Limit  int
	Offset int
}

type Sort struct {
	Field     string
	Direction string // "asc" or "desc"
}

type DBAdapter interface {
	Init() error
	Migrate(entities []any) error

	Create(entity any) error
	Update(entity any) error
	Delete(id string, entity any) error

	FindAll(entity any, filters map[string]any, pagination Pagination, sort []Sort) (any, error)
	FindByID(id string, entity any) (any, error)
}
