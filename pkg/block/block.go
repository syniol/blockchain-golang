package block

import (
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"blockchain-ws/pkg/block/transaction"
	"blockchain-ws/pkg/cryptography"
)

type IBlock interface {
	BlockNumber() uint
	CreatedDate() time.Time
	BlockHash() []byte
	PreviousBlockHash() []byte
	IsValidChain(prevBlockHash []byte, verbose bool) bool

	calculateBlockHash() []byte
	setBlockHash(parent *Block)
}

type Block struct {
	blockNumber       uint
	creationTimestamp time.Time
	blockHash         []byte
	parent            *Block
	previousBlockHash []byte

	NextBlock    *Block
	Transactions []transaction.Serializer
}

func NewBlock(parent *Block, trx ...transaction.Serializer) (*Block, error) {
	blockNum, err := blockNumber(parent)
	if err != nil {
		return nil, err
	}

	block := &Block{
		blockNumber:       blockNum,
		creationTimestamp: time.Time{}.UTC(),
		Transactions:      trx,
	}

	err = block.setBlockHash(parent)
	if err != nil {
		return nil, err
	}

	return block, nil
}

func blockNumber(parent *Block) (uint, error) {
	if parent == nil {
		return 1, nil
	}

	if parent.blockNumber == 0 {
		return 0, errors.New(
			"parent can not have undefined value for block number",
		)
	}

	return parent.blockNumber + 1, nil
}

func (b *Block) BlockNumber() uint {
	return b.blockNumber
}

func (b *Block) CreatedDate() time.Time {
	return b.creationTimestamp
}

func (b *Block) BlockHash() []byte {
	return b.blockHash
}

func (b *Block) PreviousBlockHash() []byte {
	return b.previousBlockHash
}

func (b *Block) IsValidChain(prevBlockHash []byte, verbose bool) bool {
	return true
}

func (b *Block) calculateBlockHash() ([]byte, error) {
	combined := []byte(
		fmt.Sprintf("%d", b.blockNumber) +
			b.creationTimestamp.UTC().String() +
			string(b.previousBlockHash),
	)

	if len(b.Transactions) > 0 {
		combined = append(combined, []byte(b.Transactions[0].ToString())...)
	}

	output := make([]byte, base64.StdEncoding.EncodedLen(len(combined)))
	base64.StdEncoding.Encode(output, combined)

	blockHash, err := cryptography.NewHash(output)
	if err != nil {
		return nil, err
	}

	return blockHash, nil
}

func (b *Block) setBlockHash(parent *Block) error {
	b.previousBlockHash = nil

	if parent != nil {
		b.previousBlockHash = parent.BlockHash()
		parent.NextBlock = b
	}

	var err error
	b.blockHash, err = b.calculateBlockHash()
	if err != nil {
		return err
	}

	return nil
}
