## Development only (ubuntu 16.04 tested)

source zeppelin contracts

```console
cd $GOPATH/src/mit-ra-crowdsale-api
rsync -vax vendor/github.com/OpenZeppelin/zeppelin-solidity/contracts/ solidity/zeppelin/
```

install latest solc

```console
sudo add-apt-repository ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install solc
```

build go bindings

```console
$GOPATH/src/mit-ra-crowdsale-api/solidity/bin/abigen --sol solidity/Mit-raCrowdsale.sol --pkg=mit-ra_crowdsale --out=solidity/bindings/mit-ra_crowdsale/Mit-raCrowdsale.go
```
