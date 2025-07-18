package idgenerator

import (
	"fmt"
	"github.com/sony/sonyflake"
)

type SonyFlakeGenerator struct {
	sf *sonyflake.Sonyflake
}

func NewSonyFlakeGenerator() *SonyFlakeGenerator {
	return &SonyFlakeGenerator{
		sf: sonyflake.NewSonyflake(sonyflake.Settings{}),
	}
}

func (g *SonyFlakeGenerator) NewID() (string, error) {
	id, err := g.sf.NextID()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("ORD-%d", id), nil
}
