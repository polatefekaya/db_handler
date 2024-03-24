package db

import "DatabaseHandler/pkg/interfaces/process"

type PlayerProcessDbHandler struct {
}

type IPlayerDbProcess interface {
	StartDb()
	process.IPlayerProcess
}

func (m *PlayerProcessDbHandler) StartDb() {

}
func (m *PlayerProcessDbHandler) PlayerProcess() {
}

func (m *PlayerProcessDbHandler) LeagueProcess() {

}

func (m *PlayerProcessDbHandler) TeamProcess() {

}
