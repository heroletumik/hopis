# Note this code is beta.
It is not ready yet for production.

# Horizon
[![Build Status](https://travis-ci.org/stellar/horizon.svg?branch=master)](https://travis-ci.org/stellar/horizon)
[![docs examples](https://sourcegraph.com/api/repos/github.com/stellar/horizon/.badges/docs-examples.svg)](https://sourcegraph.com/github.com/stellar/horizon)

Horizon is the [client facing API](/docs) server
for the Stellar ecosystem. 


## Building

[gb](getgb.io) is used for building horizon.

Given you have a running golang installation, you can install this with:

```bash
go get -u github.com/constabulary/gb/...
```

From within the project directory, simply run `gb build`.  After successful
completion, you should find `bin/horizon` is present in the project directory.

## Regenerating generated code

Horizon uses two go tools you'll need to install:
1. [go-bindata](https://github.com/jteeuwen/go-bindata) is used to bundle test data
1. [go-codegen](https://github.com/nullstyle/go-codegen) is used to generate some boilerplate code

After the above are installed, simply run `gb generate`.

## Running Tests

first, create two local Postgres databases, and start a redis server on port
`6379`

```bash
psql -c 'create database "horizon_test";'
psql -c 'create database "stellar-core_test";'
redis-server
```

then, run the tests like so:

```bash
bash script/run_tests.bash
```
