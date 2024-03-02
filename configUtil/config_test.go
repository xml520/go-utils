package configUtil

import (
	"fmt"
	"testing"
)

type Database struct {
	Ok string `yaml:"ok"`
}
type Test struct {
	Name  string `yaml:"name"`
	Name2 string `yaml:"name2"`
}

func TestNewConfigLoad(t *testing.T) {
	//var database Database
	var test Test
	var database Database
	_, err := NewConfigLoadBind("config_test.yaml", nil, &test, &database)
	if err != nil {
		panic(err)
	}
	fmt.Println(test, database)
	select {}
}
