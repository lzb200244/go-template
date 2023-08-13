package config

import (
	"go-template/config/autoload"
	"go-template/config/autoload/db"
)

type Configuration struct {
	Mysql   db.Mysql         `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis   db.Redis         `mapstructure:"redis" json:"redis" yaml:"redis"`
	Log     autoload.Log     `mapstructure:"log" json:"log" yaml:"log"`
	Project autoload.Project `mapstructure:"project" json:"project" yaml:"project"`
}
