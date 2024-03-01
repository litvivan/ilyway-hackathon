package main

import "github.com/litvivan/ilyway/cmd/app"


func main() {
	application, err := app.NewApplication()
	if err != nil {
		panic("failed to initialize application: " + err.Error())
	}

	application.Run()
}
