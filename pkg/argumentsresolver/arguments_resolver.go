package argumentsresolver

import (
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ArgumentsInterface interface {
	GetInt(argName string) (arg int, err error)
	GetString(argName string) (arg string, err error)
	GetBool(argName string) (arg bool, err error)
}

type unsupportedArgTypeError struct{}
type unsupportedArgError struct{}

func (m *unsupportedArgTypeError) Error() string {
	return "unsupported argument type passed"
}

func (m *unsupportedArgError) Error() string {
	return "unsupported argument requested"
}

type argType string

const (
	argTypeString argType = "string"
	argTypeInt    argType = "int"
	argTypeBool   argType = "bool"
)

type arguments struct {
	argType         argType
	argShort        string
	argDescription  string
	argValue        interface{}
	argDefaultValue interface{}
}

type argumentsResolver struct {
	arg map[string]*arguments
}

const ArgumentConfigName = "config-file"
const ArgumentConfigVerbose = "verbose"

func New() ArgumentsInterface {
	ar := &argumentsResolver{}
	ar.arg = make(map[string]*arguments)

	if path, err := os.Getwd(); err != nil {
		panic(err)
	} else {
		ar.arg[ArgumentConfigName] = &arguments{
			argType:         argTypeString,
			argDescription:  "Application config file path (.yaml); Default: /application_folder/config.yaml;",
			argDefaultValue: path + "/config.yaml",
			argValue:        nil,
		}
		ar.arg[ArgumentConfigVerbose] = &arguments{
			argType:         argTypeBool,
			argShort:        "v",
			argDescription:  "Set application into verbose mode (bool); Default: false",
			argDefaultValue: false,
			argValue:        nil,
		}
	}

	if err := ar.resolveArguments(); err != nil {
		panic(err)
	}

	return ar
}

func (a *argumentsResolver) resolveArguments() (err error) {
	for k, v := range a.arg {
		switch {
		case v.argType == argTypeString:
			if v.argShort != "" {
				pflag.StringP(k, v.argShort, v.argDefaultValue.(string), v.argDescription)
			} else {
				pflag.String(k, v.argDefaultValue.(string), v.argDescription)
			}
		case v.argType == argTypeInt:
			if v.argShort != "" {
				pflag.IntP(k, v.argShort, v.argDefaultValue.(int), v.argDescription)
			} else {
				pflag.Int(k, v.argDefaultValue.(int), v.argDescription)
			}
		case v.argType == argTypeBool:
			if v.argShort != "" {
				pflag.BoolP(k, v.argShort, v.argDefaultValue.(bool), v.argDescription)
			} else {
				pflag.Bool(k, v.argDefaultValue.(bool), v.argDescription)
			}
		default:
			return &unsupportedArgTypeError{}
		}
	}

	pflag.Parse()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	if err = viper.BindPFlags(pflag.CommandLine); err != nil {
		return
	}

	for k, v := range a.arg {
		switch {
		case v.argType == argTypeString:
			a.arg[k].argValue = viper.GetString(k)
		case v.argType == argTypeInt:
			a.arg[k].argValue = viper.GetInt(k)
		case v.argType == argTypeBool:
			a.arg[k].argValue = viper.GetBool(k)
		default:
			return &unsupportedArgTypeError{}
		}
	}
	err = viper.BindPFlags(pflag.CommandLine)
	return
}

func (a *argumentsResolver) GetInt(argName string) (arg int, err error) {
	if a.arg[argName].argValue != nil {
		return a.arg[argName].argValue.(int), nil
	}
	if a.arg[argName].argDefaultValue == nil {
		return a.arg[argName].argDefaultValue.(int), nil
	}
	return 0, &unsupportedArgError{}
}

func (a *argumentsResolver) GetString(argName string) (arg string, err error) {
	if a.arg[argName].argValue != nil {
		return a.arg[argName].argValue.(string), nil
	}
	if a.arg[argName].argDefaultValue == nil {
		return a.arg[argName].argDefaultValue.(string), nil
	}
	return "", &unsupportedArgError{}
}

func (a *argumentsResolver) GetBool(argName string) (arg bool, err error) {
	if a.arg[argName].argValue != nil {
		return a.arg[argName].argValue.(bool), nil
	}
	if a.arg[argName].argDefaultValue == nil {
		return a.arg[argName].argDefaultValue.(bool), nil
	}
	return false, &unsupportedArgError{}
}
