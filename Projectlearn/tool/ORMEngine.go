package tool

import (
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
}

func OrmEngine(cfg *Config)(*Orm,error){
	database :=cfg.Database
	conn:=database.User +":"+database.Password+"@tcp("+database.Host+":"+database.Port+")/"+database.DbName+"?charset="+database.Charset
	engine,err:=xorm.NewEngine("mysql",conn)
	if err!=nil{
		return nil,err
	}
	engine.ShowSQL(database.ShowSQL)
	orm:=new(Orm)
	orm.Engine=engine
	return orm,nil
}