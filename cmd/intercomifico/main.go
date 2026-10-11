// Command intercomifico is the terminal client for Intercom customer conversations.
package main

import (
	"fmt"
	"os"

	"github.com/izkreny/intercomifico"
)

func main() {
	if err := intercomifico.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
