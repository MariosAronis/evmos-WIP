package osmosis

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	cmn "github.com/evmos/evmos/v14/precompiles/common"
)

const (
	// EventTypeSwap defines the event type for the Osmosis Swap transaction.
	EventTypeSwap = "Swap"
)

// EmitSwapEvent creates a new Swap event emitted on a Swap transaction.
func (p Precompile) EmitSwapEvent(
	ctx sdk.Context,
	stateDB vm.StateDB,
	senderAddr, receiverAddr common.Address,
	amount *big.Int,
	inputDenom, outputDenom, chainPrefix string,
) error {
	// Prepare the event topics
	event := p.ABI.Events[EventTypeSwap]
	topics := make([]common.Hash, 3)

	// The first topic is always the signature of the event.
	topics[0] = event.ID

	var err error
	// sender and receiver are indexed
	topics[1], err = cmn.MakeTopic(senderAddr)
	if err != nil {
		return err
	}
	topics[2], err = cmn.MakeTopic(receiverAddr)
	if err != nil {
		return err
	}

	// Prepare the event data: denom, amount, memo
	arguments := abi.Arguments{event.Inputs[2], event.Inputs[3], event.Inputs[4], event.Inputs[5]}
	packed, err := arguments.Pack(amount, inputDenom, outputDenom, chainPrefix)
	if err != nil {
		return err
	}

	stateDB.AddLog(&ethtypes.Log{
		Address:     p.Address(),
		Topics:      topics,
		Data:        packed,
		BlockNumber: uint64(ctx.BlockHeight()),
	})

	return nil
}