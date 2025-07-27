package main

import (
	"fmt"

	"github.com/rahulvarma07/note_backend/internal/config"
)

func main() {
	var cnf config.Config = *config.MustLoad()
	fmt.Println(cnf)
}
