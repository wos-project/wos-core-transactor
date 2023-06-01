# wos-core-transactor README #

## How-to Build ##

### build for max ###
go build -o wos-core-transactor ./app

### build for linux ###
env GOOS=linux GOARCH=amd64 go build -o wos-core-transactor ./app

### prerequisites for NEAR

Here are the manual steps to mint an NFT.  The wasm file in the repo is build using up to step 4.  To actually mint you have to have near credentials loaded and the contract must be deployed first.

1. install Rust (do not use brew) https://www.rust-lang.org/tools/install
2. install npm cli 
```Console
npm install near-cli -g
```
3. checkout contract repo 
```Console
git clone https://github.com/near-examples/NFT
```
4. edit NFT/nft/src/lib.rs and customize fields so doesn't say Example
5. build wasm file
```Console
./scripts/build.sh
cp non_fungible_token.wasm chains/near
```
6. login or copy near credentials (this creates ~/.near-credentials/testnet/questori.testnet.json)
```Console
near login
```
7. deploy the contract
```Console
near deploy --wasmFile non_fungible_token.wasm --accountId questori.testnet
```
9. initialize contract
```Console
near call questori.testnet new_default_meta '{"owner_id": "'questori.testnet'"}' --accountId questori.testnet
```
9.  view contract details
```Console
near view questori.testnet nft_tokens_for_owner '{"account_id": "'questori.testnet'"}'
```
10.  mint to someone, token_id must be unique
```Console
near call questori.testnet nft_mint '{"token_id": "1", "receiver_id": "'questori.testnet'", "token_metadata": { "title": "Questori Stori", "description": "NFT containing media and metadata for Questori Stori", "media": "https://bafkreiabag3ztnhe5pg7js4bj6sxuvkz3sdf76cjvcuqjoidvnfjz7vwrq.ipfs.dweb.link/", "copies": 1}}' --accountId questori.testnet --deposit 0.01
```
1.   optionally transfer
```Console
near call questori.testnet nft_transfer '{"owner_id": "questori.testnet" , "token_id": "1"}' --accountId questori.testnet --depositYocto 1
```
