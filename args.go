package main

import (
	"flag"
	"fmt"
	"os"
)

func getArgs() string {

	args := os.Args
	numbers := len(args)
	fmt.Println("number of args:", numbers)
	if numbers == 1 {
		currentDir, _ := os.Getwd()
		return currentDir
	}
	path := args[1]
	var showHelp, showAbout, showVersion bool

	flag.BoolVar(&showHelp, "h", false, "help info")
	flag.BoolVar(&showAbout, "a", false, "about info")
	flag.BoolVar(&showVersion, "v", false, "version info")

	flag.Parse()
	if showHelp {
		printHelp()
		os.Exit(0)
	}
	if showAbout {
		printAbout()
		os.Exit(0)
	}
	if showVersion {
		printVersion()
		os.Exit(0)
	}
	return path

}
func printHelp() {
	fmt.Println("Usage: anjv [OPTIONS]")
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(0)
}
func printVersion() {
	fmt.Println("AnJayView version: 1.0")
	os.Exit(0)
}
func printAbout() {
	os.Exit(0)
}
