package rpc

import (
	"fmt"
	"strconv"
	"time"

	firecoreRPC "github.com/streamingfast/firehose-core/rpc"

	"github.com/gagliardetto/solana-go/rpc"
	"github.com/spf13/cobra"
	"github.com/streamingfast/cli/sflags"
	firecore "github.com/streamingfast/firehose-core"
	"github.com/streamingfast/firehose-core/blockpoller"
	"github.com/streamingfast/firehose-solana/block/fetcher"
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

func NewFetchCmd(logger *zap.Logger, tracer logging.Tracer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rpc <first-streamable-block>",
		Short: "fetch blocks from rpc endpoint",
		Args:  cobra.ExactArgs(1),
		RunE:  fetchRunE(logger, tracer),
	}

	cmd.Flags().StringArray("endpoints", []string{}, "List of endpoints to use to fetch different method calls")
	cmd.Flags().String("state-dir", "/data/poller", "interval between fetch")
	cmd.Flags().Duration("interval-between-fetch", 0, "interval between fetch")
	cmd.Flags().Duration("latest-block-retry-interval", time.Second, "interval between fetch")
	cmd.Flags().Duration("max-block-fetch-duration", 3*time.Second, "interval between fetch")
	cmd.Flags().Int("block-fetch-batch-size", 10, "Number of blocks to fetch in a single batch")
	cmd.Flags().String("network", "mainnet", "network to fetch from (mainnet, devnet, testnet) -- only used to patch a known issue on some slots")

	return cmd
}

func fetchRunE(logger *zap.Logger, tracer logging.Tracer) firecore.CommandExecutor {
	return func(cmd *cobra.Command, args []string) (err error) {
		stateDir := sflags.MustGetString(cmd, "state-dir")

		startBlock, err := strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			return fmt.Errorf("unable to parse first streamable block %d: %w", startBlock, err)
		}

		fetchInterval := sflags.MustGetDuration(cmd, "interval-between-fetch")
		maxBlockFetchDuration := sflags.MustGetDuration(cmd, "max-block-fetch-duration")

		logger.Info(
			"launching firehose-solana poller",
			zap.String("state_dir", stateDir),
			zap.Uint64("first_streamable_block", startBlock),
			zap.Duration("interval_between_fetch", fetchInterval),
			zap.Duration("latest_block_retry_interval", sflags.MustGetDuration(cmd, "latest-block-retry-interval")),
		)

		rpcEndpoints := sflags.MustGetStringArray(cmd, "endpoints")
		rpcClients := firecoreRPC.NewClients[*rpc.Client](maxBlockFetchDuration, firecoreRPC.NewRollingStrategyRoundRobin[*rpc.Client]())
		for _, rpcEndpoint := range rpcEndpoints {
			client := rpc.New(rpcEndpoint)
			rpcClients.Add(client)
		}

		latestBlockRetryInterval := sflags.MustGetDuration(cmd, "latest-block-retry-interval")
		var isMainnet bool
		switch sflags.MustGetString(cmd, "network") {
		case "mainnet", "mainnet-beta":
			isMainnet = true
		}

		poller := blockpoller.New[*rpc.Client](
			fetcher.NewRPC(fetchInterval, latestBlockRetryInterval, isMainnet, logger),
			blockpoller.NewFireBlockHandler("type.googleapis.com/sf.solana.type.v1.Block"),
			rpcClients,
			blockpoller.WithLogger[*rpc.Client](logger),
			blockpoller.WithStoringState[*rpc.Client](stateDir),
		)

		err = poller.Run(startBlock, blockpoller.MaxStopBlock, sflags.MustGetInt(cmd, "block-fetch-batch-size"))
		if err != nil {
			return fmt.Errorf("running poller: %w", err)
		}

		return nil
	}
}
