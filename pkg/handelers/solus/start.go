package solus

import (
	"github.com/mjarkk/multipkg/pkg/types"
)

func Setup() *types.Handeler {
	return &types.Handeler{
		Install: func() {},
		Update:  func() {},
		Remove:  func() {},
	}
}
