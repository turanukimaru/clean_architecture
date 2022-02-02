package calcapi

import (
	"context"
	"github.com/turanukimaru/ca/goa/dummies"
	"github.com/turanukimaru/ca/usecase/pkg/dummycalc"
	"github.com/turanukimaru/gormstart/pkg/dummydb"
	"log"

	"github.com/turanukimaru/goastart/gen/calc"
)

// calc service example implementation.
// The example methods log the requests and return zero values.
type calcsrvc struct {
	logger *log.Logger
}

// NewCalc returns the calc service implementation.
func NewCalc(logger *log.Logger) calc.Service {
	return &calcsrvc{logger}
}

// Add implements add.
func (s *calcsrvc) Add(ctx context.Context, p *calc.AddPayload) (res int, err error) {
	//	s.logger.Print("calc.add")
	dummy := dummies.Dummy{}
	err = dummy.Hello()
	err = dummy.Allow()
	err = dummydb.DbAccess()
	return dummycalc.Add(p.A, p.B), err
}
