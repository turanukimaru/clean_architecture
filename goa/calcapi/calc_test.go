package calcapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/turanukimaru/goastart/gen/calc"
	"testing"
)

// 今のままでもDB接続してテストはできる。デバッグもできる。このシステムでDB接続とか設定とか同管理するか考えないとな。
func TestCalcsrvc_Add(t *testing.T) {
	service := calcsrvc{nil}
	ctx := context.Background()
	payload := calc.AddPayload{A: 1, B: 2}
	res, _ := service.Add(ctx, &payload)
	assert.Equal(t, res, 3)
}
