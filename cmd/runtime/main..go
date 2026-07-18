package main

import (
	"fmt"
	"os"

	"github.com/Ayansama/gocontainer/internal/namespace"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gocontainer <run|child>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "run":
		if err := namespace.Run(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case "child":
		if err := namespace.Child(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	default:
		fmt.Println("usuage: gocontainer <run|child>")
		os.Exit(1)
	}

}
