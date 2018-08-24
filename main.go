package main

import (
	"fmt"
	"os"

	"github.com/mjarkk/multipkg/pkg/app"
	"github.com/mjarkk/multipkg/pkg/functions"
)

func printHelp() {
	fmt.Printf(`
  Usage: multipkg [options] [command]

  Options:

    -f, --force        Force command
    -y, --yes          Automaticly input yes for next
    --help             Help menu
    --version          App info
    --debug            Log get debug data 

  Commands:

    install|in|i       <program>  Install a program
    reinstall|rein|ri  <program>  Re-install a program
    remove|re|r        <program>  Remove a program from the system
    update|up|u        *<program> Update a program or the complete system
    search|se|s        <program>  Search for programs
    info|inf           <program>  Get info about a specific programs
`)
}

func printVersion() {
	fmt.Println("App version:", version)
}

func main() {
	if funs.InArr(os.Args[1:], "--help") {
		printHelp()
	} else if funs.InArr(os.Args[1:], "--version") {
		printVersion()
	} else {
		app.Run()
	}
}
