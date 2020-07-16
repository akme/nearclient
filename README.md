# nearclient
[![PkgGoDev](https://pkg.go.dev/badge/github.com/akme/nearclient)](https://pkg.go.dev/github.com/akme/nearclient)
[![Go Report Card](https://goreportcard.com/badge/github.com/akme/nearclient)](https://goreportcard.com/report/github.com/akme/nearclient) 

nearclient is golang client library to interact with NEAR Protocol via RPC API, inspired by ethclient.

***Warning: not production ready***

## Supported methods
- [x] status
- [ ] send transaction (async)
- [ ] send transaction (wait until done)
- [ ] query
- [x] block (doesn't support finality)
- [x] chunk  
- [x] transaction status  
- [x] validators  
- [x] gasprice  

## Install
```bash
go get -u github.com/akme/nearclient
```

## Docs
For docs go to [pkg.go.dev](https://pkg.go.dev/github.com/akme/nearclient)

## Contribute
Pull Requests are welcome!