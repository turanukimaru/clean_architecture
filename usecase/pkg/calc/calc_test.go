package calc

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 今のままでもDB接続してテストはできる。デバッグもできる。このシステムでDB接続とか設定とか同管理するか考えないとな。
func TestCalcUseCase_Add(t *testing.T) {
	service := Adder{1, 2}
	ctx := context.Background()
	res, _ := service.Add(ctx)
	assert.Equal(t, res, 3)
}
