package main

import (
	"context"
	"github.com/thirdweb-dev/go-sdk/v2/thirdweb"
	"log"
	"os"
	"os/signal"
)

type Scope struct {
	ctx       context.Context
	sdk       *thirdweb.ThirdwebSDK
	contracts map[string]*thirdweb.SmartContract
}

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	ctx, cancel := context.WithCancel(context.Background())
	//rpcUrl := os.Getenv("RPC_URL")
	rpcUrl := "https://rpc.mevblocker.io/"

	sdk, err := thirdweb.NewThirdwebSDK(rpcUrl, nil)
	if err != nil {
		panic(err)
	}
	scope := &Scope{
		sdk: sdk,
		ctx: ctx,
	}
	contracts := scope.createPairContracts(contractAddresses)
	scope.contracts = contracts

	go scope.run()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case <-c:
		signal.Stop(c)
		cancel()
		<-ctx.Done()
		log.Println("~ ~ ~ ~ exiting ~ ~ ~ ~")
		os.Exit(0)
	}
}

func (s *Scope) run() {
	err := s.getBridgedTo()
	if err != nil {
		log.Println(err)
	}
}
