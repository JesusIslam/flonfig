package flonfig

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
)

// Flag is used to unmarshal the JSON
type Flag struct {
	Key          string      `toml:"key"`
	Env          string      `toml:"env,omitempty"`
	Message      string      `toml:"message"`
	ValueType    string      `toml:"value_type"`
	DefaultValue interface{} `toml:"default_value"`
	ParsedValue  interface{} `toml:"-"`
}

// Flags is only used to parse the toml
type Flags struct {
	Flags []*Flag `toml:"flags"`
}

// Flonfig is the core of this package
type Flonfig struct {
	Flags      map[string]*Flag
	ConfigPath string
}

// New will return an empty Flonfig struct pointer
func New() *Flonfig {
	return &Flonfig{}
}

// Implement will load the config file if the target file exists
// or read it as config string if the target file does not exist
// and implement it as flags
func (f *Flonfig) Implement(configpathOrData string) (err error) {
	f.ConfigPath = configpathOrData
	flags := Flags{}

	if isConfig(configpathOrData) {
		_, err = toml.Decode(configpathOrData, &flags)
		if err != nil {
			return
		}
	} else {
		_, err = toml.DecodeFile(configpathOrData, &flags)
		if err != nil {
			return
		}
	}

	f.Flags = map[string]*Flag{}

	// Define the flags first
	for _, fl := range flags.Flags {
		switch fl.ValueType {
		case "string":
			targetPointer := flag.String(fl.Key, fl.DefaultValue.(string), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "int":
			targetPointer := flag.Int(fl.Key, int(fl.DefaultValue.(int64)), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "int64":
			targetPointer := flag.Int64(fl.Key, fl.DefaultValue.(int64), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "uint":
			targetPointer := flag.Uint(fl.Key, uint(fl.DefaultValue.(int64)), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "uint64":
			targetPointer := flag.Uint64(fl.Key, uint64(fl.DefaultValue.(int64)), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "bool":
			targetPointer := flag.Bool(fl.Key, fl.DefaultValue.(bool), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "float64":
			targetPointer := flag.Float64(fl.Key, fl.DefaultValue.(float64), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "duration":
			targetPointer := flag.Duration(fl.Key, time.Duration(fl.DefaultValue.(int64)), fl.Message)
			fl.ParsedValue = targetPointer
			break
		case "duration_string":
			targetPointer := flag.String(fl.Key, fl.DefaultValue.(string), fl.Message)
			fl.ParsedValue = targetPointer
			break
		default:
			err = fmt.Errorf("Invalid value type %s for flag %s", fl.ValueType, fl.Key)
			return
		}

		f.Flags[fl.Key] = fl
	}

	// Parse the flags
	flag.Parse()

	// collect the flag values
	for key, fl := range f.Flags {
		switch fl.ValueType {
		case "string":
			targetPointer := fl.ParsedValue.(*string)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					targetPointer = &envValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "int":
			targetPointer := fl.ParsedValue.(*int)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue int
					tmpValue, err = strconv.Atoi(envValue)
					if err != nil {
						return
					}
					targetPointer = &tmpValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "int64":
			targetPointer := fl.ParsedValue.(*int64)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue int64
					tmpValue, err = strconv.ParseInt(envValue, 10, 64)
					if err != nil {
						return
					}
					targetPointer = &tmpValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "uint":
			targetPointer := fl.ParsedValue.(*uint)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue uint64
					tmpValue, err = strconv.ParseUint(envValue, 10, 64)
					if err != nil {
						return
					}
					tmpUintValue := uint(tmpValue)
					targetPointer = &tmpUintValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "uint64":
			targetPointer := fl.ParsedValue.(*uint64)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue uint64
					tmpValue, err = strconv.ParseUint(envValue, 10, 64)
					if err != nil {
						return
					}
					targetPointer = &tmpValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "bool":
			targetPointer := fl.ParsedValue.(*bool)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue bool
					tmpValue, err = strconv.ParseBool(envValue)
					if err != nil {
						return
					}
					targetPointer = &tmpValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "float64":
			targetPointer := fl.ParsedValue.(*float64)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue float64
					tmpValue, err = strconv.ParseFloat(envValue, 64)
					if err != nil {
						return
					}
					targetPointer = &tmpValue
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "duration":
			targetPointer := fl.ParsedValue.(*time.Duration)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					var tmpValue int64
					tmpValue, err = strconv.ParseInt(envValue, 10, 64)
					if err != nil {
						return
					}
					tmpValueDuration := time.Duration(tmpValue)
					targetPointer = &tmpValueDuration
				}
			}

			f.Flags[key].ParsedValue = *targetPointer
			break
		case "duration_string":
			targetPointer := fl.ParsedValue.(*string)

			if fl.Env != "" {
				envValue, exist := os.LookupEnv(fl.Env)
				if exist {
					targetPointer = &envValue
				}
			}

			var targetDuration time.Duration
			targetDuration, err = time.ParseDuration(*targetPointer)
			if err != nil {
				return
			}

			f.Flags[key].ParsedValue = targetDuration
			break
		default:
			err = fmt.Errorf("Invalid value type %s for flag %s", fl.ValueType, fl.Key)
			return
		}
	}

	return
}

// Windows max 260 characters path
// macOS max characters path 1016 characters
// Linux max 4096 characters path and 255 characters for file name
func isConfig(data string) (itIs bool) {
	switch runtime.GOOS {
	case "windows":
		if len(data) >= 255 {
			itIs = true
		}
	case "linux":
		if len(data) >= 4096 {
			itIs = true
		}
	case "darwin":
		if len(data) >= 1016 {
			itIs = true
		}
	case "dragonfly":
		if len(data) >= 1016 {
			itIs = true
		}
	case "netbsd":
		if len(data) >= 1016 {
			itIs = true
		}
	case "openbsd":
		if len(data) >= 1016 {
			itIs = true
		}
	default:
	}

	return
}
