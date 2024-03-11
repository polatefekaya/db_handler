package player

import "DatabaseHandler/pkg/interfaces"

var En entity

type entity interfaces.Entity

func (a *entity) ToEntity(p *PlayerRoot) *PlayerEntity {
	return nil
}
