module github.com/streamingfast/sf-solana

go 1.15

require (
	cloud.google.com/go/bigtable v1.13.0
	cloud.google.com/go/storage v1.21.0
	github.com/GeertJohan/go.rice v1.0.0
	github.com/ShinyTrinkets/overseer v0.3.0
	github.com/abourget/llerrgroup v0.2.0
	github.com/dustin/go-humanize v1.0.0
	github.com/golang/protobuf v1.5.2
	github.com/graph-gophers/graphql-go v0.0.0-20201027172035-4c772c181653
	github.com/klauspost/compress v1.15.2
	github.com/lithammer/dedent v1.1.0
	github.com/lorenzosaino/go-sysctl v0.1.1
	github.com/manifoldco/promptui v0.8.0
	github.com/mholt/archiver/v3 v3.5.0
	github.com/mr-tron/base58 v1.2.0
	github.com/spf13/cobra v1.3.0
	github.com/spf13/viper v1.10.1
	github.com/streamingfast/bstream v0.0.2-0.20220505155906-a0834b9c5258
	github.com/streamingfast/cli v0.0.4-0.20220113202443-f7bcefa38f7e
	github.com/streamingfast/dauth v0.0.0-20220404140613-a40f4cd81626
	github.com/streamingfast/derr v0.0.0-20220301163149-de09cb18fc70
	github.com/streamingfast/dgraphql v0.0.2-0.20220307143518-466192441cfe
	github.com/streamingfast/dgrpc v0.0.0-20220301153539-536adf71b594
	github.com/streamingfast/dlauncher v0.0.0-20211210162313-cf4aa5fc4878
	github.com/streamingfast/dmetering v0.0.0-20220307162406-37261b4b3de9
	github.com/streamingfast/dmetrics v0.0.0-20210811180524-8494aeb34447
	github.com/streamingfast/dstore v0.1.1-0.20220419183635-aad7bcb15b8e
	github.com/streamingfast/firehose v0.1.1-0.20220427051727-9b108461d3a6
	github.com/streamingfast/jsonpb v0.0.0-20210811021341-3670f0aa02d0
	github.com/streamingfast/kvdb v0.0.2-0.20210811194032-09bf862bd2e3
	github.com/streamingfast/logging v0.0.0-20220405224725-2755dab2ce75
	github.com/streamingfast/merger v0.0.3-0.20220427133815-fb04cd671db1
	github.com/streamingfast/node-manager v0.0.2-0.20220422154052-6a6439016eaf
	github.com/streamingfast/pbgo v0.0.6-0.20220428192744-f80aee7d4688
	github.com/streamingfast/relayer v0.0.2-0.20220307182103-5f4178c54fde
	github.com/streamingfast/sf-solana/types v0.0.0-20220505183419-daee6f2a30b9
	github.com/streamingfast/sf-tools v0.0.0-20220504154702-4b10c7c5fdf1
	github.com/streamingfast/shutter v1.5.0
	github.com/streamingfast/solana-go v0.5.1-0.20220502224452-432fbe84aee8
	github.com/streamingfast/substreams v0.0.5-beta
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/teris-io/shortid v0.0.0-20201117134242-e59966efd125 // indirect
	github.com/test-go/testify v1.1.4
	go.uber.org/zap v1.21.0
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	google.golang.org/api v0.70.0
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)

replace github.com/graph-gophers/graphql-go => github.com/streamingfast/graphql-go v0.0.0-20210204202750-0e485a040a3c
replace github.com/ShinyTrinkets/overseer => github.com/dfuse-io/overseer v0.2.1-0.20191024193921-39856397cf3f
