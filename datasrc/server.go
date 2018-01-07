package datasrc

import (
	"github.com/go-kit/kit/log"
)

type dataSrcServer struct {
	logger log.Logger
}

func (ds *dataSrcServer) Start() {
}

func (ds *dataSrcServer) Stop() {

}

// New creates a new data source service
func New(logger log.Logger) (Service, error) {
	return &dataSrcServer{logger: logger}, nil
}
