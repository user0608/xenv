package main

import (
	_ "embed"
	"fmt"
	"log"
	"log/slog"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
)

//go:embed version
var version string

func main() {
	if len(os.Args) == 2 {
		if os.Args[1] == "version" || os.Args[1] == "--version" {
			fmt.Println(path.Base(os.Args[0]), version)
			os.Exit(0)
		}
	}
	if len(os.Args) > 2 {
		slog.Error("to many args found", "args", strings.Join(os.Args[1:], " "))
	}
	var filename string
	if len(os.Args) == 1 {
		filename = ".env"
	} else {
		filename = os.Args[1]
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	defer f.Close()
	envs, err := godotenv.Parse(f)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	var builder strings.Builder
	for k, v := range envs {
		builder.WriteString(fmt.Sprintf(`%s="%s" `, k, v))
	}
	fmt.Print(builder.String())
}
