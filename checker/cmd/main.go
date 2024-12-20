package main

import (
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/application"
)

func main() {
	fmt.Println("Starting application")
	application.Init()
	fmt.Println("application started")
}
