package flonfig

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func assertFlonfigFlagsFile(t *testing.T, testFlonfig *Flonfig) {
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
			assert.True(t, ok, "-d flag must return time.Duration")
			assert.Equal(t, time.Duration(0), val, "the value of -d flag must be the same as the default value of time.Duration")
			break
		case "ds":
			val, ok := item.ParsedValue.(time.Duration)
			assert.True(t, ok, "-ds flag must return time.Duration")
			assert.Equal(t, time.Duration(0), val, "the value of -ds flag must be the same as the default value of time.Duration")
		default:
			assert.Failf(t, "Invalid flag key %s", flagKey)
		}
	}
}

func assertFlonfigFlagsData(t *testing.T, testFlonfig *Flonfig) {
	for flagKey, item := range testFlonfig.Flags {
		switch flagKey {
		case "str":
			val, ok := item.ParsedValue.(string)
			assert.True(t, ok, "-str flag must return string")
			assert.Equal(t, "", val, "the value of -str flag must be the same as the default value of string")
			break
		case "int":
			val, ok := item.ParsedValue.(int)
			assert.True(t, ok, "-int flag must return int")
			assert.Equal(t, int(0), val, "the value of -int flag must be the same as the default value of int")
			break
		case "uint":
			val, ok := item.ParsedValue.(uint)
			assert.True(t, ok, "-uint flag must return uint")
			assert.Equal(t, uint(0), val, "the value of -uint flag must be the same as the default value of uint")
			break
		case "uint64":
			val, ok := item.ParsedValue.(uint64)
			assert.True(t, ok, "-uint64 flag must return uint64")
			assert.Equal(t, uint64(0), val, "the value of -uint64 flag must be the same as the default value of uint64")
			break
		case "bool":
			val, ok := item.ParsedValue.(bool)
			assert.True(t, ok, "-bool flag must return bool")
			assert.Equal(t, false, val, "the value of -bool flag must be the same as the default value of boolean")
			break
		case "float64":
			val, ok := item.ParsedValue.(float64)
			assert.True(t, ok, "-float64 flag must return float64")
			assert.Equal(t, float64(0.0), val, "the value of -float64 flag must be the same as the default value of float64")
			break
		case "duration":
			val, ok := item.ParsedValue.(time.Duration)
			assert.True(t, ok, "-duration flag must return time.Duration")
			assert.Equal(t, time.Duration(0), val, "the value of -duration flag must be the same as the default value of time.Duration")
			break
		case "duration-string":
			val, ok := item.ParsedValue.(time.Duration)
			assert.True(t, ok, "-duration-string flag must return time.Duration")
			assert.Equal(t, time.Duration(0), val, "the value of -duration-string flag must be the same as the default value of time.Duration")
		default:
			assert.Failf(t, "Invalid flag key %s", flagKey)
		}
	}
}

func TestNew(t *testing.T) {
	testFlonfig := New()
	assert.NotNil(t, testFlonfig, "New must not return nil")
}

func TestImplementData(t *testing.T) {
	testFlonfig := New()
	configPath := "./config.data.example.toml"
	assert.FileExists(t, configPath, "Config file must be exist")

	raw, err := ioutil.ReadFile(configPath)
	assert.NoError(t, err, "There must be no error while reading config file")
	if err != nil {
		assert.FailNow(t, "Test failed without config file")
	}

	err = testFlonfig.ImplementData(string(raw))
	assert.NoError(t, err, "Implement must not return an error")

	assertFlonfigFlagsData(t, testFlonfig)
}

func TestImplementFile(t *testing.T) {
	testFlonfig := New()
	configPath := "./config.example.toml"
	assert.FileExists(t, configPath, "Config file must be exist")

	err := testFlonfig.ImplementFile(configPath)
	assert.NoError(t, err, "Implement must not return an error")

	assertFlonfigFlagsFile(t, testFlonfig)
}
