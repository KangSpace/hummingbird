package util

import (
	"fmt"
	"github.com/zedseven/bch"
	"testing"
)

func TestBch(t *testing.T) {
	versionECLBins := [][]uint8{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}, {0, 1, 1}, {1, 0, 0}, {1, 0, 1}, {1, 1, 0}, {1, 1, 1}}
	masks := [][]uint8{{0, 0}, {0, 1}, {1, 0}, {1, 1}}
	if nBchConfig, err := bch.CreateConfig(15, 5); err == nil {
		for i := 0; i < len(versionECLBins); i++ {
			for j := 0; j < len(masks); j++ {
				bys := append(versionECLBins[i], masks[j]...)
				if desc, err := bch.Encode(nBchConfig, &bys); err == nil {
					fmt.Printf("%v\n", desc)
				} else {
					panic(err)
				}
			}

		}
	} else {
		panic(err)
	}
}
