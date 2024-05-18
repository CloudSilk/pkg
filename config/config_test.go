package config

import (
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		t.Fatal(err)
	}
	err = InitFromString(string(data))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(DefaultConfig.DBConfigs)
}
