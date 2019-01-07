// Package main implements inlinesql which generates a Go file containing SQL
// queries.
package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
	"time"
)

func removeComments(s []byte) []byte {
	matchComments := regexp.MustCompile("(?s)--.*?\n|/\\*.*?\\*/")
	return matchComments.ReplaceAll(s, nil)
}

func parseSQLFile(in []byte) []string {
	in = removeComments(in)

	// Split file into seperate queries.
	var tmp []string
	tmp = strings.Split(string(in), ";\n")

	// Remove superflous spaces around each query.
	for i := range tmp {
		tmp[i] = strings.TrimSpace(tmp[i])
		tmp[i] = strings.Join(strings.Fields(tmp[i]), " ")
	}

	// Remove empty rows.
	var out []string
	for _, v := range tmp {
		if v == "" {
			continue
		}

		out = append(out, v)
	}

	return out
}

func main() {
	var (
		in, out, pkg string
	)

	flag.StringVar(&in, "in", "", "SQL file to parse.")
	flag.StringVar(&out, "out", "", "File to output Go code to.")
	flag.StringVar(&pkg, "pkg", "main", "Package name to use.")
	flag.Parse()

	if in == "" {
		log.Fatalln("Please define a SQL file to parse.")
	}

	if out == "" {
		log.Fatalln("Please define a output Go file.")
	}

	sqlFile, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatalln(err)
	}

	queries := parseSQLFile(sqlFile)

	outFile, err := os.Create(out)
	if err != nil {
		log.Fatalln(err)
	}
	defer outFile.Close()

	pkgTmpl.Execute(outFile, struct {
		Timestamp time.Time
		Package   string
		Queries   []string
	}{
		Timestamp: time.Now(),
		Package:   pkg,
		Queries:   queries,
	})
}

var pkgTmpl = template.Must(template.New("").Parse(
	`// Package {{ .Package }} was generated automatically by inlinesql at {{ .Timestamp }}.
package {{ .Package }}

// GetQueries returns a pre-parsed slice of SQL queries.
func GetQueries() []string {
	return []string{
		{{- range .Queries }}
		{{ printf "%q" . }},
		{{- end }}
	}
}
`))
