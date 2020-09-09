[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-1f425f.svg)](http://golang.org)
[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/internetarchive/isodos.svg)](https://github.com/internetarchive/isodos)
[![](https://img.shields.io/badge/godoc-reference-blue.svg)](https://godoc.org/github.com/internetarchive/isodos)
[![Go Report Card](https://goreportcard.com/badge/github.com/internetarchive/isodos)](https://goreportcard.com/report/github.com/internetarchive/isodos)
[![GitHub license](https://img.shields.io/github/license/internetarchive/isodos.svg)](https://github.com/internetarchive/isodos/blob/master/LICENSE)

# isodos
Go module and CLI tool to interact with Internet Archive's Isodos API

## Introduction

This project is a client to the **Isodos** (Είσοδος) API built to allow Internet Archive's partners to send us batches of URLs to archive.

## What is this project?

This Isodos client is actually two distinct things.

First, it's a Go module that anyone can integrate in an existing pipeline to use the Isodos API more easily. Second, it's a CLI tool to process lists of URLs.

## Warning

**Isodos is NOT a public API.**

Isodos require a pair of whitelisted archive.org S3-like access key and secret key. You can find yours at: https://archive.org/account/s3.php.

Your archive.org account need to be whitelisted by Internet Archive's staff to start using Isodos with it.

## The CLI tool

### Installation

There are 2 ways to install the client, if you have the Go toolchain installed on your machine you can simply do:

```bash
go get -u github.com/internetarchive/isodos
```

And then you will be able to access it directly by typing `isodos`.

If you don't have the Go toolchain, you can also download a release from the [Releases page](https://github.com/internetarchive/isodos/releases), extract it, make it executable, and you are good to go!

### Usage

If you use `isodos` for the first time, you will be prompted for your S3 access/secret keys, for the default project you want to use Isodos for, and the path where you want to put the config file. (by default, isodos choose `~/.isodos.json`)

After that, sending a list of URLs to Isodos for processing is as easy as:

```bash
isodos send list YOUR-SEEDS-LIST.TXT
```

Isodos will print the JSON response that the API return, containing the digest of the request and its UUID.

## The module

Here is an example to use isodos as a module:
```go
package main

import (
	"encoding/json"
	"log"

	"github.com/internetarchive/isodos"
)

func main() {
	seeds := []string{"https://archive.org", "https://youtube.com", "https://google.com"}

	client := isodos.Init("YOUR-S3-KEY", "YOUR-S3-SECRET", "YOUR-PROJECT")
	response, err := client.Send(seeds, true)
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(b))
}
```
