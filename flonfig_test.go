package flonfig

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	testFlonfig := New()
	assert.NotNil(t, testFlonfig, "New must not return nil")
}

func TestImplement(t *testing.T) {
	testFlonfig := New()
	configPath := "./config.example.toml"
	err := testFlonfig.Implement(configPath)

	assert.Nil(t, err, "Implement must not return an error")

	for flagKey, item := range testFlonfig.Flags {
		switch flagKey {
		case "s":
			val, ok := item.ParsedValue.(string)
			assert.True(t, ok, "-s flag must return string")
			assert.Equal(t, "", val, "the value of -s flag must be the same as the default value of string")
			break
		case "i":
			val, ok := item.ParsedValue.(int)
			assert.True(t, ok, "-i flag must return int")
			assert.Equal(t, int(0), val, "the value of -i flag must be the same as the default value of int")
			break
		case "u":
			val, ok := item.ParsedValue.(uint)
			assert.True(t, ok, "-u flag must return uint")
			assert.Equal(t, uint(0), val, "the value of -u flag must be the same as the default value of uint")
			break
		case "u64":
			val, ok := item.ParsedValue.(uint64)
			assert.True(t, ok, "-u64 flag must return uint64")
			assert.Equal(t, uint64(0), val, "the value of -u64 flag must be the same as the default value of uint64")
			break
		case "b":
			val, ok := item.ParsedValue.(bool)
			assert.True(t, ok, "-b flag must return bool")
			assert.Equal(t, false, val, "the value of -b flag must be the same as the default value of boolean")
			break
		case "f":
			val, ok := item.ParsedValue.(float64)
			assert.True(t, ok, "-f flag must return float64")
			assert.Equal(t, float64(0.0), val, "the value of -f flag must be the same as the default value of float64")
			break
		case "d":
			val, ok := item.ParsedValue.(time.Duration)
			assert.True(t, ok, "-d flag must return string")
			assert.Equal(t, time.Duration(0), val, "the value of -d flag must be the same as the default value of time.Duration")
			break
		default:
			assert.Failf(t, "Invalid flag key %s", flagKey)
		}
	}
}
