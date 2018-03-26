package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/url"
	"os"
	"time"
)

// Compute the password from given input
func generatePass(key string, input ...string) string {

	data := ""

	for _, v := range input {
		data += v + "/"
	}

	secret := []byte(data + key)

	for i := 0; i < 1<<16; i++ {
		sum := sha256.Sum256(secret)
		secret = sum[:]
	}

	password := base64.StdEncoding.EncodeToString(secret)[:16]
	return password
}

// Returns the host of a url, adding http as default protocol if none is provided
func stripUrl(link string) string {

	u, err := url.Parse(link)

	// Check for missing protocol
	if u.Scheme == "" {
		link = "http://" + link
		u, err = url.Parse(link)
	}

	if err != nil {
		fmt.Println("Failed to extract url. Maybe use --nostrip to use it as is?")
		log.Fatal(err)
	}

	return u.Host
}

func main() {

	var key string
	var mUrl string
	var version string
	var nostrip bool

	app := cli.NewApp()
	app.Name = "go-hashpass"
	app.Compiled = time.Now()
	app.Usage = "A hashpass-compatible stateless password manager/genrator"
	app.UsageText = "-k key -d url [-v version] [-ns]"
	app.Description = "A hashpass-compatible stateless password manager/genrator"
	app.HideVersion = true
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Pablo Ovelleiro",
			Email: "pablo1@mailbox.org",
		},
	}

	// Define avaitible flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "Your secret key",
			Destination: &key,
		},

		cli.StringFlag{
			Name:        "url, u",
			Usage:       "The url to use",
			Destination: &mUrl,
		},

		cli.BoolFlag{
			Name:        "nostrip, ns",
			Usage:       "Don't extract domain if a complete link is provided",
			Destination: &nostrip,
		},
		cli.StringFlag{
			Name:        "vesion, v",
			Usage:       "The password version",
			Value:       "",
			Destination: &version,
		},
	}

	app.Action = func(c *cli.Context) error {

		// Check if necessary parameters are provided
		if key != "" && mUrl != "" {

			// Check if the provided url has to be stripped (i.e. extract the domain from a link
			if !nostrip {
				mUrl = stripUrl(mUrl)
			}

			// Check if a password version was provided
			if version != "" {
				fmt.Println(generatePass(key, mUrl, version))
			} else {
				fmt.Println(generatePass(key, mUrl))
			}
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
