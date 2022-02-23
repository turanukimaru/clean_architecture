package calcapi

import (
	"context"
	"github.com/turanukimaru/ca/goa/gen/calc"
	uc "github.com/turanukimaru/ca/usecase/pkg/calc"
	"log"
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
	////	s.logger.Print("calc.add")
	//dummy := dummies.Dummy{}
	//err = dummy.Hello()
	//err = dummy.Allow()
	//err = dummydb.DbAccess()
	usecase := uc.Adder{A: 1, B: 2}
	return usecase.Add(context.Background())
}
