package ducati

import (
	"errors"

	"github.com/cloudfoundry-incubator/guardian/gardener"
	"github.com/pivotal-golang/lager"
)

type Ducati struct{}

func (d *Ducati) Capacity() uint64 { return 0 }

func (d *Ducati) Network(log lager.Logger, handle, spec string) (string, error) {
	return "", errors.New("ducati not implemented")
}

var _ gardener.Networker = &Ducati{}
