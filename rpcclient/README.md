rpcclient
=========

[![Build Status](https://travis-ci.org/bitcoinsv/bsvd.png?branch=master)](https://travis-ci.org/bitcoinsv/bsvd)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/bitcoinsv/bsvd/rpcclient)

rpcclient implements a Websocket-enabled Bitcoin JSON-RPC client package written
in [Go](http://golang.org/).  It provides a robust and easy to use client for
interfacing with a Bitcoin RPC server that uses a bsvd/bitcoin core compatible
Bitcoin JSON-RPC API.

## Status

This package is currently under active development.  It is already stable and
the infrastructure is complete.  However, there are still several RPCs left to
implement and the API is not stable yet.

## Documentation

* [API Reference](http://godoc.org/github.com/bitcoinsv/bsvd/rpcclient)
* [bsvd Websockets Example](https://github.com/bitcoinsv/bsvd/tree/master/rpcclient/examples/bsvdwebsockets)
  Connects to a bsvd RPC server using TLS-secured websockets, registers for
  block connected and block disconnected notifications, and gets the current
  block count
* [bsvwallet Websockets Example](https://github.com/bitcoinsv/bsvd/tree/master/rpcclient/examples/bsvwalletwebsockets)
  Connects to a bsvwallet RPC server using TLS-secured websockets, registers for
  notifications about changes to account balances, and gets a list of unspent
  transaction outputs (utxos) the wallet can sign
* [Bitcoin Core HTTP POST Example](https://github.com/bitcoinsv/bsvd/tree/master/rpcclient/examples/bitcoincorehttp)
  Connects to a bitcoin core RPC server using HTTP POST mode with TLS disabled
  and gets the current block count

## Major Features

* Supports Websockets (bsvd/bsvwallet) and HTTP POST mode (bitcoin core)
* Provides callback and registration functions for bsvd/bsvwallet notifications
* Supports bsvd extensions
* Translates to and from higher-level and easier to use Go types
* Offers a synchronous (blocking) and asynchronous API
* When running in Websockets mode (the default):
  * Automatic reconnect handling (can be disabled)
  * Outstanding commands are automatically reissued
  * Registered notifications are automatically reregistered
  * Back-off support on reconnect attempts

## Installation

```bash
$ go get -u github.com/bitcoinsv/bsvd/rpcclient
```

## License

Package rpcclient is licensed under the [copyfree](http://copyfree.org) ISC
License.
