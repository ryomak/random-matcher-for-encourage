package main

import (
	"os"
)

func main() {
	app := NewApp("divide", "interview divide", "1.0.0")
	app.Run(os.Args)
}
