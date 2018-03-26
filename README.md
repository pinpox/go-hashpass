# go-hashpass


[![Build Status](https://travis-ci.org/binaryplease/go-hashpass.svg?branch=master)](https://travis-ci.org/binaryplease/go-hashpass)
[![GoDoc](https://godoc.org/github.com/binaryplease/go-hashpass?status.svg)](https://godoc.org/github.com/binaryplease/go-hashpass)
[![Go Report Card](https://goreportcard.com/badge/github.com/binaryplease/go-hashpass)](https://goreportcard.com/report/github.com/binaryplease/go-hashpass)
[![codecov](https://codecov.io/gh/binaryplease/go-hashpass/branch/master/graph/badge.svg)](https://codecov.io/gh/binaryplease/go-hashpass)


A hashpass (https://github.com/stepchowfun/hashpass) compatible password manager

Taken from the Hashpass Readme:



> Hashpass is a Chrome extension designed to make passwords less painful. It generates a unique password for every website you use, and you only have to memorize a single secret key.
>
> Hashpass is deterministic, meaning that it will always generate the same password for any given site and secret key. It uses a well-known formula to generate the passwords, so you could even compute them yourself.
>
> A key feature of Hashpass is that it's stateless. Hashpass never writes to the file system or makes network requests. There is no password database.

This is just a client that implements the used algorithm, yielding the same results. It is implemented as CLI tool and can be used without any browser.


NAME:
   go-hashpass - A hashpass-compatible stateless password manager/genrator

USAGE:
   -k key -d url [-v version] [-ns]

DESCRIPTION:
   A hashpass-compatible stateless password manager/genrator

AUTHOR:
   Pablo Ovelleiro <pablo1@mailbox.org>

COMMANDS:
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --key value, -k value     Your secret key
   --url value, -u value     The url to use
   --nostrip, --ns           Don't extract domain if a complete link is provided
   --vesion value, -v value  The password version
   --help, -h                show help
