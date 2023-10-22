/*
SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing a Asset and Token
type SmartContract struct {
	contractapi.Contract
}

type Asset struct {
	ID       string   `json:"id"`
	DocType  string   `json:"doctype"`
	Desc     string   `json:"desc"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Owner    []string `json:"owner"`
	ISActive bool     `json:"isActive"`
}

type Token struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Symbol         string   `json:"symbol"`
	DocType        string   `json:"doctype"`
	AssetID        string   `json:"assetid"`
	TotalToken     int      `json:"totatCount"`
	AvailableToken int      `json:"avaCount"`
	ReserveToken   int      `json:"resCount"`
	Owner          []string `json:"owner"`
	PricePerToken  float32  `json:"pricePerToken"`
}

func (s *SmartContract) CreateAsset(ctx contractapi.TransactionContextInterface, assetInputString string) error {
	var assetInput Asset
	err := json.Unmarshal([]byte(assetInputString), &assetInput)
	if err != nil {
		return fmt.Errorf("Error while doing unmarshal of input string : %v", err.Error())
	}
	fmt.Println("Input String :", assetInput)

	//Validate input Parameters
	if len(strings.TrimSpace(assetInput.ID)) == 0 {
		return fmt.Errorf("Asset Id should not be empty")
	}

	if assetInput.DocType != "ASSET" {
		return fmt.Errorf(`Doc Type for Asset should be "ASSET"`)
	}

	for index, owner := range assetInput.Owner {
		if len(strings.TrimSpace(owner)) == 0 {
			return fmt.Errorf("Owner %v is null", index+1)
		}
	}

	//Check if asset ID is present or not
	isExist, err := s.IsAssetExist(ctx, assetInput.ID)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("Asset already exist with ID : %v", assetInput.ID)
	}

	assetAsBytes, err := json.Marshal(assetInput)
	if err != nil {
		return fmt.Errorf("Error while doing Marshat of Asset recors : %v", err.Error())
	}

	err = ctx.GetStub().PutState(assetInput.ID, assetAsBytes)
	if err != nil {
		return fmt.Errorf("Error while inserting data to couchDB : %v", err.Error())
	}
	return nil
}

func (s *SmartContract) IsAssetExist(ctx contractapi.TransactionContextInterface, assetId string) (bool, error) {
	assetDetails, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return false, fmt.Errorf("Error while pulling state : %v", err.Error())
	}
	return assetDetails != nil, nil
}

func (s *SmartContract) QueryAsset(ctx contractapi.TransactionContextInterface, assetId string) (*Asset, error) {
	assetAsBytes, err := ctx.GetStub().GetState(assetId)

	if err != nil {
		return nil, fmt.Errorf("Failed to read data from workd state %s", err.Error())
	}
	if assetAsBytes == nil {
		return nil, fmt.Errorf("record not found for %s", assetId)
	}
	assetRecord := new(Asset)
	_ = json.Unmarshal(assetAsBytes, assetRecord)
	return assetRecord, nil
}

/**************************************************/

func (s *SmartContract) CreateToken(ctx contractapi.TransactionContextInterface, tokenInputString string) error {
	var tokenInput Token
	err := json.Unmarshal([]byte(tokenInputString), &tokenInput)
	if err != nil {
		return fmt.Errorf("Error while doing unmarshal of input string : %v", err.Error())
	}
	fmt.Println("Input String :", tokenInput)

	//Validate input Parameters
	if len(strings.TrimSpace(tokenInput.ID)) == 0 {
		return fmt.Errorf("Token Id should not be empty")
	}

	if len(strings.TrimSpace(tokenInput.Name)) == 0 {
		return fmt.Errorf("Token name should not be empty")
	}

	if len(strings.TrimSpace(tokenInput.Symbol)) == 0 {
		return fmt.Errorf("Token Symbol should not be empty")
	}

	if tokenInput.DocType != "TOKEN" {
		return fmt.Errorf(`Doc Type for Asset should be "TOKEN"`)
	}

	if tokenInput.TotalToken <= 0 {
		return fmt.Errorf("Total Token should be +ve")
	}

	if tokenInput.PricePerToken <= 0 {
		return fmt.Errorf("Price per token should be +ve")
	}

	//Check if token ID is present or not
	isExist, err := s.IsTokenExist(ctx, tokenInput.ID)
	if err != nil {
		return err
	}
	if isExist {
		return fmt.Errorf("Token already exist with ID : %v", tokenInput.ID)
	}

	//Check if asset ID is present or not
	isExist, err = s.IsAssetExist(ctx, tokenInput.AssetID)
	if err != nil {
		return err
	}
	if !isExist {
		return fmt.Errorf("Asset does not exist with ID : %v", tokenInput.AssetID)
	}

	//Fetching Owner from asset ID
	assetDetails, err := s.QueryAsset(ctx, tokenInput.AssetID)
	tokenInput.Owner = assetDetails.Owner
	tokenInput.AvailableToken = tokenInput.TotalToken

	tokenAsBytes, err := json.Marshal(tokenInput)
	if err != nil {
		return fmt.Errorf("Error while doing Marshat of Token records : %v", err.Error())
	}

	err = ctx.GetStub().PutState(tokenInput.ID, tokenAsBytes)
	if err != nil {
		return fmt.Errorf("Error while inserting data to couchDB : %v", err.Error())
	}
	return nil
}

func (s *SmartContract) IsTokenExist(ctx contractapi.TransactionContextInterface, tokenId string) (bool, error) {
	tokenDetails, err := ctx.GetStub().GetState(tokenId)
	if err != nil {
		return false, fmt.Errorf("Error while pulling state : %v", err.Error())
	}
	return tokenDetails != nil, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(SmartContract))
	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
