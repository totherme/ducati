package ducati

import (
	"errors"

	"github.com/pivotal-golang/lager"
)

type Ducati struct{}

func (d *Ducati) Capacity() uint64 { return 0 }

func (d *Ducati) Network(log lager.Logger, handle, spec string) (string, error) {
	return "", errors.New("ducati not implemented")
}

func (d *Ducati) Destroy(log lager.Logger, handle string) error {
	return errors.New("ducati not implemented")
}

func (d *Ducati) NetIn(handle string, hostPort, containerPort uint32) (uint32, uint32, error) {
	return 0,0,errors.New("ducati not implemented")
}
