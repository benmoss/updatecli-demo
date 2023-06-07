package main

import (
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"
	"text/tabwriter"

	_ "k8s.io/client-go"
)

func main() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		log.Printf("Failed to read build info")
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	fmt.Fprintln(w, "path\tversion")
	for _, dep := range bi.Deps {
		row := strings.Join([]string{dep.Path, dep.Version}, "\t")
		fmt.Fprintln(w, row)
	}
}
