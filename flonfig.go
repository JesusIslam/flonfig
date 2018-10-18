package flonfig

import (
	"flag"
	"fmt"
	"time"

	"github.com/BurntSushi/toml"
)

// Flag is used to unmarshal the JSON
type Flag struct {
	Key          string      `toml:"key"`
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

// Implement will load the config file and implement it as flags
func (f *Flonfig) Implement(configpath string) (err error) {
	f.ConfigPath = configpath
	flags := Flags{}

	_, err = toml.DecodeFile(configpath, &flags)
	if err != nil {
		return
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
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "int":
			targetPointer := fl.ParsedValue.(*int)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "int64":
			targetPointer := fl.ParsedValue.(*int64)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "uint":
			targetPointer := fl.ParsedValue.(*uint)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "uint64":
			targetPointer := fl.ParsedValue.(*uint64)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "bool":
			targetPointer := fl.ParsedValue.(*bool)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "float64":
			targetPointer := fl.ParsedValue.(*float64)
			f.Flags[key].ParsedValue = *targetPointer
			break
		case "duration":
			targetPointer := fl.ParsedValue.(*time.Duration)
			f.Flags[key].ParsedValue = *targetPointer
			break
		default:
			err = fmt.Errorf("Invalid value type %s for flag %s", fl.ValueType, fl.Key)
			return
		}
	}

	return
}
