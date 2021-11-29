package ants_test

import (
	"testing"

	"github.com/panjf2000/ants/v2"
)

func TestAntsPool(t *testing.T) {
	defer ants.Release()

	pool, err := ants.NewPool(1000)
	if err != nil {
		t.Error("create pool failed: ", err)
		return
	}
	defer pool.Release()

	pool.Submit(func() {
		t.Log("submited")
	})
}
