---

# Hyperledger Fabric Smart Contract

This repository contains a Hyperledger Fabric smart contract written in GoLang. The smart contract consists of several functions to execute transactions on the blockchain.

## Smart Contract Functions

### 1. CreateAsset

Creates a new asset on the blockchain.

**Input Structure:**

```json
{
  "id": "asset01",
  "doctype": "ASSET",
  "desc": "Asset 01 Desc",
  "name": "XYX property",
  "address": "XYZ property - Street 01",
  "owner": ["Rohan", "Rahul"],
  "isActive": true
}
```

**Invoke Command:**

```sh
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n assettoken --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"CreateAsset","Args":["{\"id\":\"asset01\",\"doctype\":\"ASSET\",\"desc\":\"Asset 01 Desc\",\"name\":\"XYX property\",\"address\":\"XYZ property - Street 01\",\"owner\":[\"Rohan\",\"Rahul\"],\"isActive\":true}"]}'
```

### 2. MintToken

Mints a new token on the blockchain.

**Input Structure:**

```json
{
  "id": "token01",
  "name": "token01",
  "symbol": "TST",
  "doctype": "TOKEN",
  "assetid": "asset01",
  "totalCount": 550,
  "pricePerToken": 10
}
```

**Invoke Command:**

```sh
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n assettoken --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"MintToken","Args":["{\"id\":\"token01\",\"name\":\"token01\",\"symbol\":\"TST\",\"doctype\":\"TOKEN\",\"assetid\":\"asset01\",\"totalCount\":550,\"pricePerToken\":10}"]}'
```

### 3. BalanceOf

Retrieves the balance of a specific token for a given owner.

**Invoke Command:**

```sh
peer chaincode query -C mychannel -n assettoken -c '{"Args":["BalanceOf", "{\"owner\":\"Rahul\",\"tokenid\":\"token01\"}"]}' | jq .
```

### 4. TransferToken

Transfers a specific amount of tokens from one owner to another.

**Input Structure:**

```json
{
  "tokenid": "token01",
  "sender": "Rahul",
  "receiver": "Kiran",
  "amountToTransfer": 10
}
```

**Invoke Command:**

```sh
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n assettoken --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"TransferToken","Args":["{\"tokenid\":\"token01\",\"sender\":\"Rahul\",\"receiver\":\"Kiran\",\"amountToTransfer\":10}"]}'
```

### 5. TransferAsset

Transfers ownership of an asset to one or more recipients.

**Input Structure:**

```json
{
  "assetId": "asset01",
  "to": ["Kiran", "Amit"]
}
```

**Invoke Command:**

```sh
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n assettoken --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"TransferAsset","Args":["{\"assetId\":\"asset01\",\"to\":[\"Kiran\",\"Amit\"]}"]}'
```

### 6. BurnToken

Burns a specific number of tokens for a given owner.

**Input Structure:**

```json
{
  "tokenid": "token01",
  "owner": "Rahul",
  "tokenCountToBurn": 3
}
```

**Invoke Command:**

```sh
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile "${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem" -C mychannel -n assettoken --peerAddresses localhost:7051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt" --peerAddresses localhost:9051 --tlsRootCertFiles "${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt" -c '{"function":"BurnToken","Args":["{\"tokenid\":\"token01\",\"owner\":\"Rahul\",\"tokenCountToBurn\":3}"]}'
```

### 7. GetHistory

Retrieves the transaction history of a specific token.

**Input Structure:**

```json
{
  "id": "token01",
  "doctype": "TOKEN",
  "owner": ""
}
```

**Invoke Command:**

```sh
peer chaincode query -C mychannel -n assettoken -

c '{"Args":["GetHistory", "{\"id\":\"token01\",\"doctype\":\"TOKEN\",\"owner\":\"\"}"]}' | jq .
```
