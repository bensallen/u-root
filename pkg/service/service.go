package service

import (
	"github.com/u-root/u-root/pkg/service/onfail"
	"github.com/u-root/u-root/pkg/service/state"
)

type Unit struct {
	Name            string
	Description     string
	EnvironmentFile string
	PIDFile         string
	Type            Type
	OnFail          onfail.Action
	State           state.Value
	Before          []string
	After           []string
	Requires        []string
}

type Type int

const (
	Simple  Type = 0
	Forking Type = 1
)

type Servicer interface {
	Start() error
	Stop() error
	Reload() error
	Restart() error
	Status() error
	Unit() Unit
}

func Start(Servicer) error {
	return nil
}
