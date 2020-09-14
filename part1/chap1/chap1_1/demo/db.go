package demo

import (
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"xorm.io/core"
)

var (
	DB *xorm.Engine
)

func init() {
	DB = initXorm()
}

func initXorm() *xorm.Engine {
	var err error
	c, err := NewMySqlXormConfig()
	if err != nil {
		panic(err)
	}

	engine, err := xorm.NewEngine(c.DriverName, c.GetDataSourceName())
	if err != nil {
		log.Fatalf("%+v\n", errors.New("init xorm engine error."))
		return nil
	}
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, c.TablePrefix)
	engine.SetTableMapper(tbMapper)
	createTableIfNotExist(engine)
	engine.Sync(new(User))
	engine.ShowSQL(c.ShowSQL)
	//todo from rdsc.LogLevel()
	engine.Logger().SetLevel(core.LOG_DEBUG)
	return engine
}

var beans = []interface{}{
	&User{},
}

func createTableIfNotExist(engine *xorm.Engine) {
	for _, bean := range beans {
		exist, err := engine.IsTableExist(bean)
		if err != nil {
			panic(err)
		}
		if !exist {
			if err := engine.CreateTables(bean); err != nil {
				panic(err)
			}
		}
	}
}
