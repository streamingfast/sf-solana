// cmd/firesol/rpc/check_block.go
package rpc

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/streamingfast/firehose-solana/block/fetcher"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewNextBlockCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "next <block-number>",
		Short: "Check if a block exists and return the next available block",
		Args:  cobra.ExactArgs(1),
		RunE:  nextBlockRunE(logger, tracer),
	}

	cmd.Flags().StringArray("endpoints", []string{}, "List of endpoints to use to fetch different method calls")

	return cmd
}

func nextBlockRunE(logger *zap.Logger, _ logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		ctx := cmd.Context()

		startBlock, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("unable to parse block number %d: %w", startBlock, err)
		}

		rpcEndpoints := sflags.MustGetStringArray(cmd, "endpoints")
		rpcClients := firecoreRPC.NewClients[*rpc.Client]()
		for _, rpcEndpoint := range rpcEndpoints {
			client := rpc.New(rpcEndpoint)
			rpcClients.Add(client)
		}

		rpcFetcher := fetcher.NewRPC(rpcClients, 1*time.Second, 1*time.Second, logger)

		var blockExists bool
		for !blockExists {
			select {
			case <-ctx.Done():
				break
			default:

			}

			blockExists = checkBlockExists(ctx, rpcFetcher, startBlock)
			if blockExists {
				break
			}
			startBlock++
		}

		_, _ = fmt.Fprintf(cmd.OutOrStdout(), fmt.Sprintf("%d", startBlock))
		return nil
	}
}

func checkBlockExists(ctx context.Context, rpcFetcher *fetcher.RPCFetcher, blockNumber uint64) bool {
	_, skip, err := rpcFetcher.Fetch(ctx, blockNumber)
	return !skip && err == nil
}
