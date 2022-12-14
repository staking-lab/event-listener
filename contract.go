package listener

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Contract struct {
	Name    string
	ABI     abi.ABI
	Address common.Address

	// event ID <-> event Name mapping
	events map[common.Hash]string
}

func NewContract(name string, abiJSON string, address common.Address) (*Contract, error) {
	c := &Contract{
		Name:    name,
		Address: address,
		events:  make(map[common.Hash]string),
	}

	abi, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		logger.Error("Failed to parse ABI interface", "name", name, "abi", abiJSON, "address", address.Hex(), "err", err)
		return nil, err
	}
	c.ABI = abi
	for _, event := range abi.Events {
		c.events[event.ID] = event.Name
	}
	return c, nil
}
