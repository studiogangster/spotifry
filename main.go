package main

import (
	"fmt"
	"log"

	"github.com/hpcloud/tail"
)

func main() {
	t, err := tail.TailFile("./fluent/data/log/logs.b5c672640c78bd106082e5e4ad322ea49.log", tail.Config{Follow: true})

	log.Println(err)
	if err != nil {
		for line := range t.Lines {
			fmt.Println(line.Text)
		}
	}
}
