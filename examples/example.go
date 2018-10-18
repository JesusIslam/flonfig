package main

import (
	"fmt"

	"github.com/JesusIslam/flonfig"
)

func main() {
	f := flonfig.New()
	err := f.Implement("./config.toml")
	if err != nil {
		panic(err)
	}

	fmt.Println(f.Flags["print"].ParsedValue.(string))
}
