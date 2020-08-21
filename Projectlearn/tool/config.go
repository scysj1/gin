package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName string			`json:"app_name"`
	AppMode string			`json:"app_mode"`
	AppHost string			`json:"app_host"`
	AppPort string			`json:"app_port"`
	Database DatabaseConfig `json:"database"`
}
//数据库连接类
type DatabaseConfig struct{
	Driver string `json:"driver"`
	User string `json:"user"`
	Password string `json:"password"`
	Host string `json:"host"`
	Port string `json:"port"`
	DbName string `json:"db_name"`
	Charset string `json:"Charset"`
	ShowSQL bool `json:"show_sql"`
}




var _cfg *Config = nil
func ParseConfig(path string) (*Config,error){
	file ,err :=os.Open(path)
	if err !=nil{
		panic(err)
	}
	defer file.Close()
	reader :=bufio.NewReader(file)
	decoder :=json.NewDecoder(reader)

	if err =decoder.Decode(&_cfg);err!=nil{
		return nil,err
	}
	return _cfg,nil
}