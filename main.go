package main

import (
	"fmt"
	"os"

	"github.com/BadiGGH/note-app/database"
	"github.com/BadiGGH/note-app/handlers"
)

func main() {
	database.RunDB()
	fmt.Println("____________________________________")
	r := handlers.SetRoutes()

	r.Run(os.Getenv("RUNNINGADDR"))

}
