package repository

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/marpme/digibyte-rosetta-node/dependency"
	"github.com/marpme/digibyte-rosetta-node/provider"
)

type BlockProvider struct {
	badgerDb *provider.BadgerDB
}

func initializeDatabase(ctx context.Context) (*BlockProvider, error) {
	badgerdb, _ := dependency.InitBadgerDb()
	b := &BlockProvider{
		badgerDb: badgerdb,
	}
	return b, nil
}

func (b *BlockProvider) StoreBlock(keyHash string, block *types.BlockResponse) {

}
