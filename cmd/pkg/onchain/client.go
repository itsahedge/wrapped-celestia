package onchain

import (
	"context"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

type OnchainClient struct {
	ctx       context.Context
	sdk       *thirdweb.ThirdwebSDK
	contracts map[string]*thirdweb.SmartContract
}

func NewClient() (*OnchainClient, error) {
	ctx := context.Background()
	rpcUrl := "https://rpc.mevblocker.io/"
	sdk, err := thirdweb.NewThirdwebSDK(rpcUrl, nil)
	if err != nil {
		return nil, err
	}
	client := &OnchainClient{
		sdk: sdk,
		ctx: ctx,
	}
	contracts := client.CreatePairContracts([]string{wTIA_Contract_Address})
	client.contracts = contracts
	return client, nil
}

func (c *OnchainClient) CreatePairContracts(contractAddresses []string) map[string]*thirdweb.SmartContract {
	contracts := make(map[string]*thirdweb.SmartContract)
	for _, pair := range contractAddresses {
		contract, err := c.sdk.GetContractFromAbi(pair, wTIA_ABI)
		if err != nil {
			panic(err)
		}
		contracts[pair] = contract
	}
	return contracts
}
