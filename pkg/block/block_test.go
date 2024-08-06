package block_test

import (
	"testing"

	"blockchain-ws/pkg/block"
)

func TestNewBlock(t *testing.T) {
	sut := block.NewBlock

	t.Run("GIVEN block is Genesis", func(t *testing.T) {
		actual, _ := sut(nil)

		if actual.BlockNumber() != 1 {
			t.Errorf("it was expecing block number 1 but git %d", actual.BlockNumber())
		}

		expectedBlockHash := "a716ab20d7a70a4cdb07660e289631c98eaefe8c94d2b84898cf76b0fcf089f8fe235405edf9b97c5a5f291b57bc024f522456fc55e10a07570c78caa6467ce0"
		if string(actual.BlockHash()) != expectedBlockHash {
			t.Errorf("it was expecting %s but got %s", expectedBlockHash, string(actual.BlockHash()))
		}

		if actual.PreviousBlockHash() != nil {
			t.Errorf("It was expecting previous block hash to be nil but got %s", string(actual.PreviousBlockHash()))
		}
	})

	t.Run("GIVEN block is not Genesis", func(t *testing.T) {
		genesisMockBlock, _ := sut(nil)
		actual, _ := sut(genesisMockBlock)

		if actual.BlockNumber() != 2 {
			t.Errorf("it was expecing block number 1 but got %d", actual.BlockNumber())
		}

		expectedBlockHash := "93b6cb43bd7f0ce90fa12c4c4ba5f99ea44497f6bbe903a12f1b4a2c7caba463e469aed4238404db6610158203c920acfef71bea143aa30f7cadb7d82186738a"
		if string(actual.BlockHash()) != expectedBlockHash {
			t.Errorf("it was expecting %s but got %s", expectedBlockHash, string(actual.BlockHash()))
		}

		expectedPreviousBlockHash := "a716ab20d7a70a4cdb07660e289631c98eaefe8c94d2b84898cf76b0fcf089f8fe235405edf9b97c5a5f291b57bc024f522456fc55e10a07570c78caa6467ce0"
		if string(actual.PreviousBlockHash()) != expectedPreviousBlockHash {
			t.Errorf("It was expecting previous block hash to be %s but got %s", expectedPreviousBlockHash, string(actual.PreviousBlockHash()))
		}

		genesisNextHashBlock := string(genesisMockBlock.NextBlock.BlockHash())
		currentBlockHash := string(actual.BlockHash())

		if genesisNextHashBlock != currentBlockHash {
			t.Errorf("it was expecting %s but got %s", currentBlockHash, genesisNextHashBlock)
		}
	})
}
