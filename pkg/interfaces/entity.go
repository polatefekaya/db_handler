package interfaces

type Entity struct {
	m IEntity
}

type IEntity interface {
	ToEntity()
}
