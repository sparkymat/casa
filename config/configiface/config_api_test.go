package configiface_test

import (
	"reflect"
	"testing"

	"git.orion.home/oxhead/casa/config"
	"git.orion.home/oxhead/casa/config/configiface"
)

func TestConfig(t *testing.T) {
	structType := reflect.TypeOf(&config.Service{})
	interfaceType := reflect.TypeOf((*configiface.ConfigAPI)(nil)).Elem()

	interfaceMethodMap := map[string]interface{}{}

	for i := 0; i < interfaceType.NumMethod(); i++ {
		interfaceMethodMap[interfaceType.Method(i).Name] = struct{}{}
	}

	for i := 0; i < structType.NumMethod(); i++ {
		if _, found := interfaceMethodMap[structType.Method(i).Name]; !found {
			t.Errorf("'%s' not found on interface '%s'", structType.Method(i).Name, interfaceType.Name())
		}
	}
}
