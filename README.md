bsvd
====
[![Build Status](https://travis-ci.org/bitcoinsv/bsvd.png?branch=master)](https://travis-ci.org/bitcoinsv/bsvd)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitcoinsv/bsvd)](https://goreportcard.com/report/github.com/bitcoinsv/bsvd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/bitcoinsv/bsvd)

bsvd is a full node Bitcoin (BSV) implementation written in Go (golang).

This project is a port of the [bchd](https://github.com/gcash/bchd) codebase to Bitcoin (BSV). It provides a high powered
and reliable blockchain server which makes it a suitable backend to serve blockchain data to lite clients and block explorers
or to power your local wallet.

bsvd does not include any wallet functionality by design as it makes the codebase more modular and easy to maintain. 
The [bsvwallet](https://github.com/bitcoinsv/bsvwallet) is a separate application that provides a secure Bitcoin (BSV) wallet 
that communicates with your running bsvd instance via the API.

## Table of Contents

- [Requirements](#requirements)
- [Install](#install)
  - [Install prebuilt packages](#install-pre-built-packages)
  - [Build from Source](#build-from-source)
- [Getting Started](#getting-started)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [License](#license)

## Requirements

[Go](http://golang.org) 1.9 or newer.

## Install

### Install Pre-built Packages

The easiest way to run the server is to download a pre-built binary. You can find binaries of our latest release for each operating system at the [releases page](https://github.com/bitcoinsv/bsvd/releases).

### Build from Source

If you prefer to install from source do the following:

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Run the following commands to obtain btcd, all dependencies, and install it:

```bash
$ go get github.com/bitcoinsv/bsvd
```

This will download and compile `bsvd` and put it in your path.

If you are a bsvd contributor and would like to change the default config file (`bsvd.conf`), make any changes to `sample-bsvd.conf` and then run the following commands:

```bash
$ go-bindata sample-bsvd.conf  # requires github.com/go-bindata/go-bindata/
$ gofmt -s -w bindata.go
```

## Getting Started

To start bsvd with default options just run:

```bash
$ ./bsvd
```

You'll find a large number of runtime options on the help menu. All of which can also be set in a config file.
See the [sample config file](https://github.com/bitcoinsv/bsvd/blob/master/sample-bsvd.conf) for an example of how to use it.

## Docker

Building and running `bsvd` in docker is quite painless. To build the image:

```
docker build . -t bsvd
```

To run the image:

```
docker run bsvd
```

To run `bsvctl` and connect to your `bsvd` instance:

```
# Find the running bsvd container.
docker ps

# Exec bsvctl.
docker exec <container> bsvctl <command>
```

## Documentation

The documentation is a work-in-progress.  It is located in the [docs](https://github.com/bitcoinsv/bsvd/tree/master/docs) folder.

## Contributing

Contributions are definitely welcome! Please read the contributing [guidelines](https://github.com/bitcoinsv/bsvd/blob/master/docs/code_contribution_guidelines.md) before starting.

## Security Disclosures

To report security issues please contact:

Chris Pacia (ctpacia@gmail.com) - GPG Fingerprint: 0150 2502 DD3A 928D CE52 8CB9 B895 6DBF EE7C 105C

or

Josh Ellithorpe (quest@mac.com) - GPG Fingerprint: B6DE 3514 E07E 30BB 5F40  8D74 E49B 7E00 0022 8DDD 

## License

bsvd is licensed under the [copyfree](http://copyfree.org) ISC License.
