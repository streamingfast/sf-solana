package grpc

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	pbserumhist "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/serumhist/v1"
	"github.com/dfuse-io/solana-go"
)

func (s *Server) GetFills(ctx context.Context, request *pbserumhist.GetFillsRequest) (*pbserumhist.FillsResponse, error) {
	var market *solana.PublicKey
	var trader *solana.PublicKey

	if len(request.Trader) != 0 {
		t, err := solana.PublicKeyFromBase58(request.Trader)
		if err != nil {
			return nil, fmt.Errorf("invalid trader address:%s : %w", request.Trader, err)
		}
		trader = &t
	}

	if len(request.Market) != 0 {
		m, err := solana.PublicKeyFromBase58(request.Market)
		if err != nil {
			return nil, fmt.Errorf("invalid market address:%s : %w", request.Trader, err)
		}
		market = &m
	}

	if trader != nil && market == nil {
		zlog.Debug("get fills by trader", zap.Stringer("trader_address", *trader))

		fills, hasMore, err := s.manager.GetFillsByTrader(ctx, *trader, 100)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve fills by trader: %w", err)
		}
		return &pbserumhist.FillsResponse{
			Fill:    fills,
			HasMore: hasMore,
		}, nil
	}

	if market != nil && trader == nil {
		zlog.Debug("get fills by market", zap.Stringer("market_address", *market))

		fills, hasMore, err := s.manager.GetFillsByMarket(ctx, *market, 100)
		if err != nil {
			return nil, fmt.Errorf("unable to retrieve fills by market: %w", err)
		}
		return &pbserumhist.FillsResponse{
			Fill:    fills,
			HasMore: hasMore,
		}, nil
	}

	if market == nil || trader == nil {
		return nil, fmt.Errorf("both trader and market are empty")
	}

	zlog.Debug("get fills by trader and market",
		zap.Stringer("trader_address", *trader),
		zap.Stringer("market_address", *market),
	)

	f, hasMore, err := s.manager.GetFillsByTraderAndMarket(ctx, *trader, *market, 100)
	if err != nil {
		return nil, fmt.Errorf("unable to retrieve fills by trader and market: %w", err)
	}
	return &pbserumhist.FillsResponse{
		Fill:    f,
		HasMore: hasMore,
	}, nil
}
