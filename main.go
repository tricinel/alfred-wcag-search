package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/tricinel/alfred-wcag-search/internal/alfred"
	"github.com/tricinel/alfred-wcag-search/internal/search"
)

var Version = "dev"

//go:embed assets/wcag.json
var wcagData []byte

func main() {
	versionFlag := flag.Bool("v", false, "print the version")
	flag.Parse()

	if *versionFlag {
		fmt.Println(Version)
		return
	}

	var query string
	if len(os.Args) > 1 {
		query = strings.TrimSpace(os.Args[1])
	}

	if query == "" {
		alfred.Empty()
		return
	}

	results, err := search.Find(query, wcagData)
	if err != nil {
		alfred.Error(err)
		return
	}

	if len(results) == 0 {
		alfred.NoResults(query)
		return
	}

	resp := alfred.Format(results)

	output, _ := json.Marshal(resp)
	fmt.Println(string(output))
}
