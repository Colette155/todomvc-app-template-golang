package db

import (
	"encoding/json"
	"fmt"
	"todomvc-app-template-golang/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
)

var (
	DB *gorm.DB
)

func ReadConfigFromFile(path string, config interface{}) error {
	dir, _ := os.Getwd()
	fmt.Printf("ReadConfigFromFile:%v\n", dir)

	file, _ := os.Open(path)
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)
	return json.Unmarshal(bytes, &config)
}

func InitMysql() {
	var conf configs.DB
	if err := ReadConfigFromFile("configs/db.json", &conf); err != nil {
		panic(err)
	}
	DB, _ = gorm.Open(mysql.Open(conf.DSN), &gorm.Config{})
	return
}
