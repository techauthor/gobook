package demo

import "fmt"

type MySqlXormConfig struct {
	UserName    string
	Password    string
	Host        string
	Port        int
	Schema      string
	DriverName  string
	TablePrefix string
	ShowSQL     bool
	LogLevel    string
}

func (c MySqlXormConfig) GetDataSourceName() string {
	ds := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		c.UserName,
		c.Password,
		c.Host,
		c.Port,
		c.Schema)
	return ds
}

func NewMySqlXormConfig() (xc *MySqlXormConfig, err error) {
	xc = &MySqlXormConfig{
		UserName:    "gobook",
		Password:    "123456",
		Host:        "localhost",
		Port:        3306,
		Schema:      "gobookdb",
		DriverName:  "mysql",
		TablePrefix: "t_",
		ShowSQL:     true,
		LogLevel:    "debug",
	}
	return
}
