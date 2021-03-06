package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/go-rod/rod/lib/utils"
)

func main() {
	log.Println("setup project...")

	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))

	nodejsDeps()
	golangDeps()

	genDockerIgnore()
}

func nodejsDeps() {
	utils.Exec("npm", "i", "--no-audit", "--no-fund")
}

func golangDeps() {
	homeDir, err := os.UserHomeDir()
	utils.E(err)

	cmd := exec.Command("go", "get",
		"github.com/ysmood/kit/cmd/godev",
		"golang.org/x/tools/cmd/goimports",
		"github.com/client9/misspell/cmd/misspell",
	)
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	cmd.Dir = homeDir
	utils.SetCmdStdPipe(cmd)
	utils.E(cmd.Run())
}

func genDockerIgnore() {
	s, err := utils.ReadString(".gitignore")
	utils.E(err)
	utils.E(utils.OutputFile(".dockerignore", s))
}
