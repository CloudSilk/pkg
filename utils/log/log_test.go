package log

import (
	"testing"
)

func TestLogrus(t *testing.T) {
	SetServiceName("test")
	Error(nil, "A group of walrus emerges from the ocean")
	Error(nil, "A group of walrus emerges from the ocean")
	UseJSONFormatter()
	Info(nil, "A group of walrus emerges from the ocean")
}
