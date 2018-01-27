package main

import (
	"encoding/json"
	"github.com/urfave/cli"
	//	"net/http"
	"net/url"
	"os"
)

type Format struct {
}

type Result struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Formats []Format `json:"formats",omitempty`
	Url string `json:"url",omitempty`
}

type Extractor interface {
	Match(*url.URL) bool
	Extract(*url.URL) (*Result, error)
}

var Extractors []Extractor

func main() {
	app := cli.NewApp()
	app.Action = func (c *cli.Context) error {
		js := json.NewEncoder(os.Stdout)
		for _, item := range c.Args() {
			u, err := url.Parse(item)
			if err != nil {
				return err
			}
			for _, ext := range Extractors {
				if ext.Match(u) {
					r, err := ext.Extract(u)
					if err != nil {
						return err
					}
					js.Encode(r)
				}
			}
		}
		return nil
	}
	app.Run(os.Args)
}
