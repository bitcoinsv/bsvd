### Table of Contents
1. [About](#About)
2. [Getting Started](#GettingStarted)
    1. [Installation](#Installation)
        1. [Windows](#WindowsInstallation)
        2. [Linux/BSD/MacOSX/POSIX](#PosixInstallation)
          1. [Gentoo Linux](#GentooInstallation)
    2. [Configuration](#Configuration)
    3. [Controlling and Querying bsvd via bsvctl](#BsvctlConfig)
    4. [Mining](#Mining)
3. [Help](#Help)
    1. [Startup](#Startup)
        1. [Using bootstrap.dat](#BootstrapDat)
    2. [Network Configuration](#NetworkConfig)
    3. [Wallet](#Wallet)
4. [Contact](#Contact)
    1. [IRC](#ContactIRC)
    2. [Mailing Lists](#MailingLists)
5. [Developer Resources](#DeveloperResources)
    1. [Code Contribution Guidelines](#ContributionGuidelines)
    2. [JSON-RPC Reference](#JSONRPCReference)
    3. [The bitcoinsv Bitcoin Cash-related Go Packages](#GoPackages)

<a name="About" />

### 1. About

bsvd is an alternative full node bitcoin cash implementation written in Go (golang).

This project is a port of the [bsvd](https://github.com/bitcoinsv/bsvd) codebase to Bitcoin Cash. It provides a high powered
and reliable blockchain server which makes it a suitable backend to serve blockchain data to lite clients and block explorers
or to power your local wallet.

bsvd does not include any wallet functionality by design as it makes the codebase more modular and easy to maintain. 
The [bsvwallet](https://github.com/bitcoinsv/bsvwallet) is a separate application that provides a secure Bitcoin Cash wallet 
that communicates with your running bsvd instance via the API.

<a name="GettingStarted" />

### 2. Getting Started

<a name="Installation" />

**2.1 Installation**

The easiest way to run the server is to download a pre-built binary. You can find binaries of our latest release for each operating system at the [releases page](https://github.com/bitcoinsv/bsvd/releases).

<a name="WindowsInstallation" />

**2.1.1 Building From Source**<br />

If you prefer to install from source do the following:

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Run the following commands to obtain bsvd, all dependencies, and install it:

```bash
$ go get github.com/bitcoinsv/bsvd
```

This will download and compile `bsvd` and put it in your path.

**2.2 Configuration**

bsvd has a number of [configuration](http://godoc.org/github.com/bitcoinsv/bsvd)
options, which can be viewed by running: `$ bsvd --help`.

<a name="BsvctlConfig" />

**2.3 Controlling and Querying bsvd via bsvctl**

bsvctl is a command line utility that can be used to both control and query bsvd
via [RPC](http://www.wikipedia.org/wiki/Remote_procedure_call).  bsvd does
**not** enable its RPC server by default;  You must configure at minimum both an
RPC username and password or both an RPC limited username and password:

* bsvd.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
* bsvctl.conf configuration file
```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
```
OR
```
[Application Options]
rpclimituser=mylimituser
rpclimitpass=Limitedp4ssw0rd
```
For a list of available options, run: `$ bsvctl --help`

<a name="Mining" />

**2.4 Mining**

bsvd supports the `getblocktemplate` RPC.
The limited user cannot access this RPC.


**1. Add the payment addresses with the `miningaddr` option.**

```
[Application Options]
rpcuser=myuser
rpcpass=SomeDecentp4ssw0rd
miningaddr=12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX
miningaddr=1M83ju3EChKYyysmM2FXtLNftbacagd8FR
```

**2. Add bsvd's RPC TLS certificate to system Certificate Authority list.**

`cgminer` uses [curl](http://curl.haxx.se/) to fetch data from the RPC server.
Since curl validates the certificate by default, we must install the `bsvd` RPC
certificate into the default system Certificate Authority list.

**Ubuntu**

1. Grant root privileges: `# sudo su -`
2. Copy rpc.cert to /usr/share/ca-certificates: `# cp /home/{USER}/.bsvd/rpc.cert /usr/share/ca-certificates/bsvd.crt`
3. Add "bsvd.crt" to /etc/ca-certificates.conf: `# echo bsvd.crt >> /etc/ca-certificates.conf`
4. Update the CA certificate list: `# update-ca-certificates`

**3. Set your mining software url to use https.**

`$ cgminer -o https://127.0.0.1:8334 -u rpcuser -p rpcpassword`

<a name="Help" />

### 3. Help

<a name="Startup" />

**3.1 Startup**

Typically bsvd will run and start downloading the block chain with no extra
configuration necessary, however, there is an optional method to use a
`bootstrap.dat` file that may speed up the initial block chain download process.

<a name="BootstrapDat" />

**3.1.1 bootstrap.dat**

* [Using bootstrap.dat](https://github.com/bitcoinsv/bsvd/tree/master/docs/using_bootstrap_dat.md)

<a name="NetworkConfig" />

**3.1.2 Network Configuration**

* [What Ports Are Used by Default?](https://github.com/bitcoinsv/bsvd/tree/master/docs/default_ports.md)
* [How To Listen on Specific Interfaces](https://github.com/bitcoinsv/bsvd/tree/master/docs/configure_peer_server_listen_interfaces.md)
* [How To Configure RPC Server to Listen on Specific Interfaces](https://github.com/bitcoinsv/bsvd/tree/master/docs/configure_rpc_server_listen_interfaces.md)
* [Configuring bsvd with Tor](https://github.com/bitcoinsv/bsvd/tree/master/docs/configuring_tor.md)

<a name="Wallet" />

**3.1 Wallet**

bsvd was intentionally developed without an integrated wallet for security
reasons.  Please see [bsvwallet](https://github.com/bitcoinsv/bsvwallet) for more
information.

<a name="DeveloperResources" />

### 4. Developer Resources

<a name="ContributionGuidelines" />

* [Code Contribution Guidelines](https://github.com/bitcoinsv/bsvd/tree/master/docs/code_contribution_guidelines.md)

<a name="JSONRPCReference" />

* [JSON-RPC Reference](https://github.com/bitcoinsv/bsvd/tree/master/docs/json_rpc_api.md)
    * [RPC Examples](https://github.com/bitcoinsv/bsvd/tree/master/docs/json_rpc_api.md#ExampleCode)

<a name="GoPackages" />

* The bitcoinsv Bitcoin Cash-related Go Packages:
    * [pcclient](https://github.com/bitcoinsv/bsvd/tree/master/rpcclient) - Implements a
      robust and easy to use Websocket-enabled Bitcoin JSON-RPC client
    * [btcjson](https://github.com/bitcoinsv/bsvd/tree/master/btcjson) - Provides an extensive API
      for the underlying JSON-RPC command and return values
    * [wire](https://github.com/bitcoinsv/bsvd/tree/master/wire) - Implements the
      Bitcoin wire protocol
    * [peer](https://github.com/bitcoinsv/bsvd/tree/master/peer) -
      Provides a common base for creating and managing Bitcoin network peers.
    * [blockchain](https://github.com/bitcoinsv/bsvd/tree/master/blockchain) -
      Implements Bitcoin block handling and chain selection rules
    * [blockchain/fullblocktests](https://github.com/bitcoinsv/bsvd/tree/master/blockchain/fullblocktests) -
      Provides a set of block tests for testing the consensus validation rules
    * [txscript](https://github.com/bitcoinsv/bsvd/tree/master/txscript) -
      Implements the Bitcoin transaction scripting language
    * [bsvec](https://github.com/bitcoinsv/bsvd/tree/master/bsvec) - Implements
      support for the elliptic curve cryptographic functions needed for the
      Bitcoin scripts
    * [database](https://github.com/bitcoinsv/bsvd/tree/master/database) -
      Provides a database interface for the Bitcoin block chain
    * [mempool](https://github.com/bitcoinsv/bsvd/tree/master/mempool) -
      Package mempool provides a policy-enforced pool of unmined bitcoin
      transactions.
    * [bsvutil](https://github.com/bitcoinsv/bsvutil) - Provides Bitcoin-specific
      convenience functions and types
    * [chainhash](https://github.com/bitcoinsv/bsvd/tree/master/chaincfg/chainhash) -
      Provides a generic hash type and associated functions that allows the
      specific hash algorithm to be abstracted.
    * [connmgr](https://github.com/bitcoinsv/bsvd/tree/master/connmgr) -
      Package connmgr implements a generic Bitcoin network connection manager.
