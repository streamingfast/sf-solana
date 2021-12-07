package serumhist

import (
	"time"

	pbserumhist "github.com/streamingfast/sf-solana/pb/sf/solana/serumhist/v1"
	"github.com/streamingfast/solana-go"
)

type Ref struct {
	Market      solana.PublicKey
	OrderSeqNum uint64
	BlockNumber uint64
	TrxID       []byte
	TrxIdx      uint32
	InstIdx     uint32
	BlockID     []byte
	Timestamp   time.Time
}

type NewOrder struct {
	Ref
	Trader solana.PublicKey
	Order  *pbserumhist.Order
}

type FillEvent struct {
	Ref
	TradingAccount solana.PublicKey
	Trader         solana.PublicKey
	Fill           *pbserumhist.Fill
}

type OrderExecuted struct {
	Ref
}

type OrderClosed struct {
	Ref
	InstrRef *pbserumhist.InstructionRef
}

type OrderCancelled struct {
	Ref
	InstrRef *pbserumhist.InstructionRef
}

type TradingAccount struct {
	Trader     solana.PublicKey
	Account    solana.PublicKey
	SlotNumber uint64
}
