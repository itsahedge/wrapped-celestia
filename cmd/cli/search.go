package main

import (
	"fmt"
)

type GetBridged struct{}

func (c *GetBridged) Run() error {
	client, err := TiaClient()
	if err != nil {
		return err
	}
	total, err := client.GetBridgedTo()
	if err != nil {
		return err
	}
	fmt.Print("Total BridgedTo: ", total)
	return nil
}
