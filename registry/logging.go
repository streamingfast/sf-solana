package registry

import (
	"github.com/streamingfast/logging"
	"go.uber.org/zap"
)

var zlog = zap.NewNop()
var traceEnabled = logging.IsTraceEnabled("token", "github.com/dfuse-io/dfuse-solana/token")

func init() {
	logging.Register("github.com/dfuse-io/dfuse-solana/token", &zlog)
}
