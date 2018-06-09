package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

type Feature struct {
	Stats map[string]map[string]string `json:"stats"`
}
type Agent struct {
	Name                string
	Browser             string             `json:"browser"`
	Abbr                string             `json:"abbr"`
	Prefix              string             `json:"prefix"`
	Type                string             `json:"type"`
	UsageGlobal         map[string]float64 `json:"usage_global"`
	Versions            []string           `json:"versions"`
	DataPrefixEceptions map[string]string  `json:"prefix_exceptions"`
	VersionList         []Version          `json:"version_list"`
}

type Version struct {
	Version     string `json:"version"`
	ReleaseData int64  `json:"release_date"`
}

type Data struct {
	Agents   map[string]Agent   `json:"agents"`
	Features map[string]Feature `json:"data"`
}

func main() {
	a := cli.NewApp()
	a.Commands = []cli.Command{
		{
			Name:   "browser",
			Action: BrowserCMD,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "o",
					Value: "ciu/browsers/browsers.go",
				},
			},
		},
	}
	if err := a.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
