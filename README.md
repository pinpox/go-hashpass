# go-hashpass


[![Build Status](https://travis-ci.org/binaryplease/go-hashpass.svg?branch=master)](https://travis-ci.org/binaryplease/go-hashpass)
[![GoDoc](https://godoc.org/github.com/binaryplease/go-hashpass?status.svg)](https://godoc.org/github.com/binaryplease/go-hashpass)
[![codebeat](https://codebeat.co/badges/0a8f30aa-f975-404b-b878-5fab3ae1cc5f)](https://codebeat.co/projects/github-com-urfave-cli)
[![Go Report Card](https://goreportcard.com/badge/binaryplease/go-hashpass)](https://goreportcard.com/report/binaryplease/go-hashpass)
[![top level coverage](https://gocover.io/_badge/github.com/binaryplease/go-hashpass?0 "top level coverage")](http://gocover.io/github.com/binaryplease/go-hashpass) /
[![altsrc coverage](https://gocover.io/_badge/github.com/binaryplease/go-hashpass/altsrc?0 "altsrc coverage")](http://gocover.io/github.com/binaryplease/go-hashpass/altsrc)
A hashpass (https://github.com/stepchowfun/hashpass) compatible password manager

Taken from the Hashpass Readme:



> Hashpass is a Chrome extension designed to make passwords less painful. It generates a unique password for every website you use, and you only have to memorize a single secret key.
>
> Hashpass is deterministic, meaning that it will always generate the same password for any given site and secret key. It uses a well-known formula to generate the passwords, so you could even compute them yourself.
>
> A key feature of Hashpass is that it's stateless. Hashpass never writes to the file system or makes network requests. There is no password database.

This is just a client that implements the used algorithm, yielding the same results. It is implemented as CLI tool and can be used without any browser.
