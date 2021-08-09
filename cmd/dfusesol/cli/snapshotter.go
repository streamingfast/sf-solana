package cli

import (
	"github.com/dfuse-io/dfuse-solana/snapshotter/app/snapshotter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/streamingfast/dlauncher/launcher"
)

func init() {
	launcher.RegisterApp(&launcher.AppDef{
		ID:          "snapshotter",
		Title:       "snapshotter",
		Description: "Manage solana snapshot",
		MetricsID:   "snapshotter",
		Logger:      launcher.NewLoggingDef("github.com/dfuse-io/snapshotter.*", nil),
		RegisterFlags: func(cmd *cobra.Command) error {
			cmd.Flags().String("snapshotter-source-bucket", "mainnet-beta-ledger-us-west1", "bucket where solana snapshot are stored")
			cmd.Flags().String("snapshotter-source-prefix", "", "mainnet-beta-ledger-us-west1")
			cmd.Flags().String("snapshotter-destination-bucket", "", "bucket where solana snapshot will be stored and uncompressed")
			cmd.Flags().String("snapshotter-destination-prefix", "sol-mainnet/snapshots", "")
			cmd.Flags().String("snapshotter-working-dir", "{dfuse-data-dir}/working", "")
			return nil
		},
		InitFunc: func(runtime *launcher.Runtime) (err error) {
			return nil
		},
		FactoryFunc: func(runtime *launcher.Runtime) (launcher.App, error) {
			dfuseDataDir := runtime.AbsDataDir

			return snapshotter.New(
				&snapshotter.Config{
					SourceBucket:               viper.GetString("snapshotter-source-bucket"),
					SourceSnapshotsFolder:      viper.GetString("snapshotter-source-prefix"),
					DestinationBucket:          viper.GetString("snapshotter-destination-bucket"),
					DestinationSnapshotsFolder: viper.GetString("snapshotter-destination-prefix"),
					Workdir:                    mustReplaceDataDir(dfuseDataDir, viper.GetString("snapshotter-working-dir")),
				},
			), nil
		},
	})
}
