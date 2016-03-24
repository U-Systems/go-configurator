package go_configurator

import (
	"github.com/u-systems/go-configurator/parser"
	"reflect"
	"errors"
	"fmt"
)

type Config struct {
	configStruct interface{}
	field []parser.Field
}


func New(configSource interface{}) (*Config, error) {
	// checking type of config struct
	if reflect.TypeOf(configSource) != reflect.Ptr {
		nil, errors.New(fmt.Sprintf("config source must be a pointer to %T", configSource))
	}
	// check if it struct
	if reflect.TypeOf(configSource).Elem() != reflect.Struct {
		nil, errors.New("config source must be a struct")
	}


}

func (config *Config) Load(provider ConfigProvider) (interface{}, error) {
	panic("not implemented")
}

func (config *Config) Template(provider ConfigProvider) ([]byte, error) {
	panic("not implemented")
}


