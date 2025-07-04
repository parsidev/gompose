package hooks

type BeforeCreate interface {
	BeforeCreate() error
}

type AfterCreate interface {
	AfterCreate() error
}

type BeforeDelete interface {
	BeforeDelete() error
}

type AfterDelete interface {
	AfterDelete() error
}

type BeforeUpdate interface {
	BeforeUpdate() error
}

type AfterUpdate interface {
	AfterUpdate() error
}

type BeforePatch interface {
	BeforePatch() error
}

type AfterPatch interface {
	AfterPatch() error
}
