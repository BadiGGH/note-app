package main

import (
	"fmt"

	"github.com/BadiGGH/note-app/database"
	"github.com/BadiGGH/note-app/handlers"
)

func main() {
	database.RunDB()
	fmt.Println("____________________________________")
	r := handlers.SetRoutes()

	r.Run(":5623")

}
