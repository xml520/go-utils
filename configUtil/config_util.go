package configUtil

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go-utils/structUtil"
	"path/filepath"
	"reflect"
	"strings"
)

type Config struct {
	*viper.Viper
}

func NewConfig() *Config {
	return &Config{viper.New()}
}

// NewConfigLoad 加载配置，根据文件后缀名识别文件类型
func NewConfigLoad(filename string, defaultConfig map[string]any) (*Config, error) {
	c := NewConfig()
	c.SetConfigFile(filename)
	c.SetConfigType(strings.ReplaceAll(filepath.Ext(filename), ".", ""))
	if err := c.ReadInConfig(); err != nil {
		return nil, err
	}
	c.setDefault(defaultConfig)

	c.WatchConfig()
	return c, nil
}

// NewConfigLoadBind 加载配置并且自动识别结构体并且绑定，根据文件后缀名识别文件类型
func NewConfigLoadBind(filename string, defaultConfig map[string]any, v ...any) (*Config, error) {
	c, err := NewConfigLoad(filename, defaultConfig)
	if err != nil {
		return nil, err
	}
	if err = c.unmarshal(v...); err != nil {
		return nil, err
	}
	c.OnConfigChange(func(in fsnotify.Event) {
		if in.Has(fsnotify.Write) {
			c.unmarshal(v...)
		}
	})
	return c, nil
}
func (c *Config) unmarshal(v ...any) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("struct unmarshal %v", p)
		}
	}()
	for _, a := range v {
		s := structUtil.NewStruct(a)
		name := s.Names().LowerCamel()
		val := reflect.ValueOf(a).Elem()
		for _, field := range s.Fields {
			kv := c.Get(fmt.Sprintf("%s.%s", name, field.Names().LowerCamel()))
			if kv != nil {
				val.FieldByName(field.Names().String()).Set(reflect.ValueOf(kv))
			}
		}
	}
	return err
}
func (c *Config) setDefault(v map[string]any) {
	if v != nil {
		for k, val := range v {
			c.SetDefault(k, val)
		}
	}
}
