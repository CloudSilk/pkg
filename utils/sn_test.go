package utils

import (
	"testing"
	"time"
)

func TestGenSN(t *testing.T) {
	t.Log(GenSN(time.Now().UnixNano()))
	t.Log(GenSN64(time.Now().UnixNano()))
	t.Log(GenMac(time.Now().UnixNano()))
	t.Log(GenShortID(""))
}
