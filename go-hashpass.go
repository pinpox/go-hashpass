package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

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

func main() {

	// key := "mykey"
	// domain := "play.golang.org"
	// fmt.Println("OIs/gtSGGsuWXLDJ") // expected

	var key string
	var domain string
	var version string
	app := cli.NewApp()
	app.Name = "go-hashpass"
	app.Compiled = time.Now()
	app.Usage = "A hashpass-compatible stateless password manager/genrator"
	app.Description = "A hashpass-compatible stateless password manager/genrator"
	app.HideVersion = true
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Pablo Ovelleiro",
			Email: "pablo1@mailbox.org",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "key, k",
			Usage:       "Your secret key",
			Destination: &key,
		},

		cli.StringFlag{
			Name:        "domain, d",
			Usage:       "The domain to use",
			Destination: &domain,
		},
		cli.StringFlag{
			Name:        "vesion, v",
			Usage:       "The password version",
			Value:       "",
			Destination: &version,
		},
	}

	app.Action = func(c *cli.Context) error {
		if key != "" && domain != "" {
			if version != "" {
				fmt.Println(generatePass(key, domain, version))
			} else {
				fmt.Println(generatePass(key, domain))
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
