package main

import (
	"flag"
	"fmt"
	"os"
)

var version = "dev"

func main() {
	name := flag.String("name", "World", "the name to greet")
	greeting := flag.String("greeting", "Hello", "the greeting message")
	showVersion := flag.Bool("version", false, "show version information")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "my-cli - A simple greeting CLI tool\n\n")
		fmt.Fprintf(os.Stderr, "Usage:\n")
		fmt.Fprintf(os.Stderr, "  my-cli [flags]\n\n")
		fmt.Fprintf(os.Stderr, "Flags:\n")
		fmt.Fprintf(os.Stderr, "  -n, --name string        the name to greet (default \"World\")\n")
		fmt.Fprintf(os.Stderr, "  -g, --greeting string    the greeting message (default \"Hello\")\n")
		fmt.Fprintf(os.Stderr, "  -v, --version            show version information\n")
		fmt.Fprintf(os.Stderr, "  -h, --help               show this help message\n")
		fmt.Fprintf(os.Stderr, "\nExamples:\n")
		fmt.Fprintf(os.Stderr, "  my-cli --name Alice --greeting Hi    # Output: Hi, Alice!\n")
		fmt.Fprintf(os.Stderr, "  my-cli -n Bob -g Hey                 # Output: Hey, Bob!\n")
		fmt.Fprintf(os.Stderr, "  my-cli                               # Output: Hello, World!\n")
		fmt.Fprintf(os.Stderr, "\nVersion: %s\n", version)
	}

	// Register short flags
	flag.StringVar(name, "n", "World", "the name to greet (shorthand)")
	flag.StringVar(greeting, "g", "Hello", "the greeting message (shorthand)")
	flag.BoolVar(showVersion, "v", false, "show version information (shorthand)")

	flag.Parse()

	if *showVersion {
		fmt.Printf("my-cli version %s\n", version)
		os.Exit(0)
	}

	fmt.Printf("%s, %s!\n", *greeting, *name)
}
