package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streamingfast/dlauncher/launcher"
	"github.com/streamingfast/firehose-solana/snapshotter/app/snapshotter"
)

func init() {
	launcher.RegisterApp(zlog, &launcher.AppDef{
		ID:          "snapshotter",
		Title:       "snapshotter",
		Description: "Manage solana snapshot",
		RegisterFlags: func(cmd *cobra.Command) error {
			cmd.Flags().String("snapshotter-source-bucket", "mainnet-beta-ledger-us-west1", "bucket where solana snapshot are stored")
			cmd.Flags().String("snapshotter-source-prefix", "", "mainnet-beta-ledger-us-west1")
			cmd.Flags().String("snapshotter-destination-bucket", "", "bucket where solana snapshot will be stored and uncompressed")
			cmd.Flags().String("snapshotter-destination-prefix", "sol-mainnet/snapshots", "")
			cmd.Flags().String("snapshotter-working-dir", "{data-dir}/working", "")
			return nil
		},
		InitFunc: func(runtime *launcher.Runtime) (err error) {
			return nil
		},
		FactoryFunc: func(runtime *launcher.Runtime) (launcher.App, error) {
			dataDir := runtime.AbsDataDir

			return snapshotter.New(
				&snapshotter.Config{
					SourceBucket:               viper.GetString("snapshotter-source-bucket"),
					SourceSnapshotsFolder:      viper.GetString("snapshotter-source-prefix"),
					DestinationBucket:          viper.GetString("snapshotter-destination-bucket"),
					DestinationSnapshotsFolder: viper.GetString("snapshotter-destination-prefix"),
					Workdir:                    MustReplaceDataDir(dataDir, viper.GetString("snapshotter-working-dir")),
				},
			), nil
		},
	})
}