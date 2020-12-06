package main

import (
	"github.com/yuseferi/longman-csv-translator/app"
)


func main() {

	config, err := app.NewConfig()
	if err != nil {
		panic(err)
	}
	application, err := app.New(config)
	if err != nil {
		panic(err)
	}
	defer application.Close()
	application.Run()
}
