package kvloader

import (
	"context"
	"fmt"

	"github.com/golang/protobuf/proto"

	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	"github.com/dfuse-io/dfuse-solana/serumhist/keyer"
	"github.com/streamingfast/kvdb/store"
)

func (kv *KVLoader) GetCheckpoint(ctx context.Context) (*pbserumhist.Checkpoint, error) {
	key := keyer.EncodeCheckpoint()

	ctx, cancel := context.WithTimeout(ctx, DatabaseTimeout)
	defer cancel()

	val, err := kv.kvdb.Get(ctx, key)
	if err == store.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("error while reading checkpoint: %w", err)
	}

	out := &pbserumhist.Checkpoint{}
	if err := proto.Unmarshal(val, out); err != nil {
		return nil, err
	}

	return out, nil
}
