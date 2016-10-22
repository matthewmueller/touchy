package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"time"

	"github.com/tj/docopt"
)

var version = "0.0.1"

var usage = `
  Usage:
    touchy <gist> <filepath>...
    touchy -h | --help
    touchy --version
  Options:
    -h, --help      Output help information
    -v, --version   Output program version
`

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	fmt.Println(args["<gist>"])

	urlParts, err := url.Parse(args["<gist>"].(string))
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	urlParts.Host = "gist.githubusercontent.com"
	urlParts.Path += "/raw/"
	gist := urlParts.String()

	netClient := &http.Client{
		Timeout: time.Second * 10,
	}

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// https://gist.github.com/matthewmueller/cb33e2c5f6834511cd45f17b59271052
	//     +
	// index.html
	//     =
	// https://gist.githubusercontent.com/matthewmueller/cb33e2c5f6834511cd45f17b59271052/raw/index.html
	for _, filepath := range args["<filepath>"].([]string) {
		filepath = path.Join(pwd, filepath)

		// if the file exists already, just skip it
		if _, error := os.Stat(filepath); error == nil {
			continue
		}

		urlParts, error := url.Parse(gist)
		if error != nil {
			log.Fatalf("error: %s", err)
		}

		urlParts.Path += path.Base(filepath)
		fmt.Println(urlParts.String())
		resp, error := netClient.Get(urlParts.String())
		if error != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, error := ioutil.ReadAll(resp.Body)
		if error != nil {
			log.Fatal(err)
		}

		error = ioutil.WriteFile(filepath, body, 0644)
		if error != nil {
			log.Fatal(error)
		}
	}
}
