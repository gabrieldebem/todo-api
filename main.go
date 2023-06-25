package main

import (
	"github.com/gabrieldebem/todo-api/routes"
)

func main() {
    r := routes.Init()

    r.Run(":3000")
}
