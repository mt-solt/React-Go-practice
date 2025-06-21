package random

import (
	"math/rand"
	"time"
)

// 乱数生成器
type RandomGenerator interface {
	Generate() int64
}

type RandomGeneratorImpl struct {
	rnd *rand.Rand
}

func NewRandomGenerator() *RandomGeneratorImpl {
	return &RandomGeneratorImpl{
		rnd: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *RandomGeneratorImpl) Generate() int64 {
	// 16桁の乱数を生成（10^15 から 10^16-1 の範囲）
	return r.rnd.Int63n(9000000000000000) + 1000000000000000
}
