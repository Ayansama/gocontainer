package main

import (
	"fmt"
	"os"

	"io"

	"github.com/Ayansama/gocontainer/internal/namespace"
	containerPTY "github.com/Ayansama/gocontainer/internal/pty"
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

	case "pty-test":

		session, err := containerPTY.Start()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
		// bridge master -terminal manually so we can interact
		go io.Copy(session.Master, os.Stdin)
		io.Copy(os.Stdout, session.Master)

	default:
		fmt.Println("usuage: gocontainer <run|child>")
		os.Exit(1)
	}

}
