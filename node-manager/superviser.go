// Copyright 2021 dfuse Platform Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package nodemanager

import (
	"fmt"
	"os"
	"strings"

	"github.com/ShinyTrinkets/overseer"
	nodeManager "github.com/streamingfast/node-manager"
	logplugin "github.com/streamingfast/node-manager/log_plugin"
	"github.com/streamingfast/node-manager/metrics"
	"github.com/streamingfast/node-manager/superviser"
	"github.com/streamingfast/solana-go/rpc"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Superviser struct {
	*superviser.Superviser
	name string

	options          *Options
	client           *rpc.Client
	logger           *zap.Logger
	localSnapshotDir string
}

type Options struct {
	BinaryPath          string
	Arguments           []string
	DataDirPath         string
	RCPPort             string
	LogToZap            bool
	DebugFirehoseLogs   bool
	HeadBlockUpdateFunc nodeManager.HeadBlockUpdater
}

func NewSuperviser(appLogger *zap.Logger, nodelogger *zap.Logger, options *Options) (*Superviser, error) {
	// Ensure process manager line buffer is large enough (50 MiB) for our Firehose instrumentation outputting lot's of text.
	overseer.DEFAULT_LINE_BUFFER_SIZE = 50 * 1024 * 1024

	client := rpc.NewClient(fmt.Sprintf("http://127.0.0.1:%s", options.RCPPort))
	s := &Superviser{
		// The arguments field is actually `nil` because arguments are re-computed upon each start
		Superviser: superviser.New(appLogger, options.BinaryPath, nil),
		options:    options,
		logger:     appLogger,
		client:     client,
	}

	s.RegisterLogPlugin(logplugin.NewKeepLastLinesLogPlugin(25, options.DebugFirehoseLogs))

	if options.LogToZap {
		s.RegisterLogPlugin(logplugin.NewToZapLogPlugin(options.DebugFirehoseLogs, nodelogger))
	} else {
		s.RegisterLogPlugin(logplugin.NewToConsoleLogPlugin(options.DebugFirehoseLogs))
	}

	appLogger.Info("created geth superviser", zap.Object("superviser", s))
	return s, nil
}

func (s *Superviser) Start(options ...nodeManager.StartOption) error {
	s.Logger.Info("updating arguments before starting binary")
	s.Superviser.Arguments = s.options.Arguments

	s.logger.Info("starting process", zap.String("command", s.GetCommand()))
	err := s.Superviser.Start(options...)
	if err != nil {
		return err
	}

	return nil
}

func (s *Superviser) IsRunning() bool {
	isRunning := s.Superviser.IsRunning()
	isRunningMetricsValue := float64(0)
	if isRunning {
		isRunningMetricsValue = float64(1)
	}

	metrics.NodeosCurrentStatus.SetFloat64(isRunningMetricsValue)

	return isRunning
}

func (s *Superviser) GetCommand() string {
	return s.options.BinaryPath + " " + strings.Join(s.options.Arguments, " ")
}

func (s *Superviser) HasData() bool {
	_, err := os.Stat(s.options.DataDirPath)
	return err == nil
}

func (s *Superviser) ServerID() (string, error) {
	return os.Hostname()
}

func (s *Superviser) LastSeenBlockNum() uint64 {
	com := rpc.CommitmentRecent
	r, err := s.client.GetSlot(&com)
	if err != nil {
		s.logger.Error("Failed to get last seen slot from rpc client. returning 0", zap.Error(err))
	}

	return uint64(r)
}

func (s *Superviser) GetName() string {
	return "solana-validator"
}

func (s *Superviser) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("binary", s.options.BinaryPath)
	enc.AddArray("arguments", stringArray(s.options.Arguments))
	enc.AddString("data_dir", s.options.DataDirPath)

	return nil
}
