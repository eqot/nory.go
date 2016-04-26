package main

import (
	"os"

	"./lib/artifact"
	"github.com/codegangsta/cli"
	"github.com/olekukonko/tablewriter"
)

func main() {
	urls := []string{
		"https://search.maven.org/solrsearch/select?rows=20&wt=json&q=okhttp",
		"https://search.maven.org/solrsearch/select?rows=20&wt=json&q=butterknife",
		"https://search.maven.org/solrsearch/select?rows=20&wt=json&q=rxjava",
	}

	app := cli.NewApp()
	app.Name = "nory"
	app.Usage = "make an explosive entrance"
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	app.Commands = []cli.Command{
		{
			Name:    "sequential",
			Aliases: []string{"s"},
			Usage:   "Run tasks sequentially",
			Action: func(c *cli.Context) {
				artifact.Sequential(urls)
			},
		},
		{
			Name:    "parallel",
			Aliases: []string{"p"},
			Usage:   "Run tasks parallelly",
			Action: func(c *cli.Context) {
				artifact.Parallel(urls)
			},
		},
		{
			Name:    "table",
			Aliases: []string{"t"},
			Usage:   "Show table",
			Action: func(c *cli.Context) {
				data := [][]string{
					[]string{"A", "The Good", "500"},
					[]string{"B", "The Very very Bad Man", "288"},
					[]string{"C", "The Ugly", "120"},
					[]string{"D", "The Gopher", "800"},
				}

				table := tablewriter.NewWriter(os.Stdout)
				table.SetHeader([]string{"groupId", "artifactId", "version"})
				table.SetAutoFormatHeaders(false)

				for _, v := range data {
					table.Append(v)
				}
				table.Render()
			},
		},
	}

	app.Run(os.Args)
}