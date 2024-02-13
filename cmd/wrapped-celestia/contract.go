package main

import (
	"fmt"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
)

var contractAddresses = []string{
	contractWTIA,
}

func (s *Scope) createPairContracts(contractAddresses []string) map[string]*thirdweb.SmartContract {
	c := make(map[string]*thirdweb.SmartContract)
	for _, pair := range contractAddresses {
		contract, err := s.sdk.GetContractFromAbi(pair, wTIA_ABI)
		if err != nil {
			panic(err)
		}
		c[pair] = contract
	}
	fmt.Println("created mapping of contract:", c)
	return c
}
