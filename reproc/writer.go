package reproc

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/streamingfast/dstore"
	"github.com/streamingfast/sf-solana/types"

	pbsolana "github.com/streamingfast/sf-solana/types/pb/sol/type/v1"

	"github.com/streamingfast/bstream"
)

const BUNDLE_SIZE = 100

type BundleWriter struct {
	w                     bstream.BlockWriter
	startBlockNum         uint64
	exclusiveStopBlockNum uint64
	buf                   *bytes.Buffer
	store                 dstore.Store
}

func NewBundleWriter(startBlockNum uint64, mergedBlockStore dstore.Store) (*BundleWriter, error) {
	buffer := bytes.NewBuffer(nil)
	blockWriter, err := bstream.GetBlockWriterFactory.New(buffer)
	if err != nil {
		return nil, fmt.Errorf("unable to get block writer: %w", err)
	}
	return &BundleWriter{
		w:                     blockWriter,
		buf:                   buffer,
		startBlockNum:         startBlockNum,
		exclusiveStopBlockNum: startBlockNum + BUNDLE_SIZE,
		store:                 mergedBlockStore,
	}, nil
}

var errBundleComplete = errors.New("bundle complete")

func (w *BundleWriter) Write(blkNum uint64, blk *pbsolana.ConfirmedBlock) error {
	if blkNum >= w.exclusiveStopBlockNum {
		return errBundleComplete
	}

	bstreamBlk, err := types.BlockFromPBSolanaProto(blk)
	if err != nil {
		return fmt.Errorf("failed to convert pbsolana block into bstream block: %w", err)
	}
	if err := w.w.Write(bstreamBlk); err != nil {
		return fmt.Errorf("failed to write bstream block: %w", err)
	}

	return nil
}

func (w *BundleWriter) Flush(ctx context.Context) error {
	filename := fileNameForBlocksBundle(w.startBlockNum)
	zlog.Info("flushing merged block files",
		zap.Uint64("start_block_num", w.startBlockNum),
		zap.Uint64("stop_block_num", w.exclusiveStopBlockNum),
		zap.String("filename", filename),
	)

	err := w.store.WriteObject(ctx, filename, w.buf)
	if err != nil {
		return fmt.Errorf("writing block buffer to store: %w", err)
	}
	return nil
}

func (w *BundleWriter) Next() (err error) {
	w.buf = bytes.NewBuffer(nil)
	w.w, err = bstream.GetBlockWriterFactory.New(w.buf)
	if err != nil {
		return fmt.Errorf("unable to get block writer: %w", err)
	}
	w.startBlockNum = w.exclusiveStopBlockNum
	w.exclusiveStopBlockNum = w.startBlockNum + BUNDLE_SIZE
	return nil
}

func (r *Reproc) saveBlock(ctx context.Context, blkNum uint64, blk *pbsolana.ConfirmedBlock, zlogger *zap.Logger) error {
	if tracer.Enabled() {
		zlogger.Debug("writing block to bundle")
	}
	err := r.bundleWriter.Write(blkNum, blk)
	if err == errBundleComplete {
		if err := r.bundleWriter.Flush(ctx); err != nil {
			return fmt.Errorf("unable to flush bundle: %w", err)
		}
		if err := r.bundleWriter.Next(); err != nil {
			return fmt.Errorf("unable to go to next bundle: %w", err)
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to write block to bundle: %w", err)
	}
	return nil
}
func fileNameForBlocksBundle(blockNum uint64) string {
	return fmt.Sprintf("%010d", blockNum)
}