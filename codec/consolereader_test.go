// Copyright 2019 dfuse Platform Inc.
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

package codec

import (
	"os"
	"testing"

	pbcodec "github.com/dfuse-io/dfuse-solana/pb/dfuse/solana/codec/v1"

	"github.com/dfuse-io/solana-go/text"
	"github.com/stretchr/testify/require"
)

func Test_readTransaction_Start(t *testing.T) {
	tests := []struct {
		name        string
		line        string
		expectedErr error
	}{
		{
			"TRANSACTION good",
			`TRANSACTION START 2KQspsbTqezudMvUTdjYdFHx37M1LKF1iUU6BW6jpmuvBifwsypFiHECipasaF9fUuorz4thi3fsCcmokh2AXVfN 0100030520492cf7f4b78fcb0993860cea450edd343e36187aed68c366e6420b63a08d1509a4358fa8bde757dc294e4562866094db00ce9f2d8f3cd4f243afb9a0e0798c06a7d517192f0aafc6f265e3fb77cc7ada82c529d0be3b136e2d00552000000006a7d51718c774c928566398691d5eb68b5eb8a39b4b6d5c73555b21000000000761481d357474bb7c4d7624ebd3bdb3d8355e73d11043fc0da3538000000000e91b52701a3abe9afe583beff4e75914573207499b463577958da9e66b5bddc10104040102030045020000000200000000000000d09b3b0300000000d19b3b0300000000e0c76fe89639cb93719e58ce9776c14caf83e4ed3e3c007f72f25638c436252a01c112d15f00000000`,
			nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := newParseCtx()
			err := ctx.readTransactionStart(test.line)

			require.Equal(t, test.expectedErr, err)

			err = text.NewEncoder(os.Stdout).Encode(ctx.slot, nil)
			require.NoError(t, err)
			//fmt.Println("out:", string(buf.Bytes()))

		})
	}
}
func Test_readInstruction_Start(t *testing.T) {
	tests := []struct {
		name        string
		trxID       string
		line        string
		expectedErr error
	}{
		{
			name:        "Instruction good",
			trxID:       "2KQspsbTqezudMvUTdjYdFHx37M1LKF1iUU6BW6jpmuvBifwsypFiHECipasaF9fUuorz4thi3fsCcmokh2AXVfN",
			line:        `INSTRUCTION START 2KQspsbTqezudMvUTdjYdFHx37M1LKF1iUU6BW6jpmuvBifwsypFiHECipasaF9fUuorz4thi3fsCcmokh2AXVfN 1 0 Vote111111111111111111111111111111111111111 020000000200000000000000d09b3b0300000000d19b3b0300000000e0c76fe89639cb93719e58ce9776c14caf83e4ed3e3c007f72f25638c436252a01c112d15f00000000`,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := newParseCtx()
			ctx.trxTraceMap[test.trxID] = &pbcodec.TransactionTrace{}
			err := ctx.readInstructionTraceStart(test.line)

			require.Equal(t, test.expectedErr, err)

			text.NewEncoder(os.Stdout).Encode(ctx, nil)

		})
	}
}

func Test_readAccountChange_Start(t *testing.T) {
	tests := []struct {
		name        string
		trxID       string
		line        string
		expectedErr error
	}{
		{
			name:        "Account change",
			trxID:       "GDRAZ4PbkN3gkHguvoQMRGgz8wXZvoniapcLXLDjg4n2adfmUE9apaEK9fs9vVWLXk1Dv7hGLTnNJdnLVHUNqHp",
			line:        `ACCOUNT_CHANGE GDRAZ4PbkN3gkHguvoQMRGgz8wXZvoniapcLXLDjg4n2adfmUE9apaEK9fs9vVWLXk1Dv7hGLTnNJdnLVHUNqHp 1 0 7gds7PbCzmHbJStjxA5L5K8cu2LVUakmd3MDXFHSfcic 01000000d38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7ed38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7e641f00000000000000b29b3b03000000001f000000b39b3b03000000001e000000b49b3b03000000001d000000b59b3b03000000001c000000b69b3b03000000001b000000b79b3b03000000001a000000b89b3b030000000019000000b99b3b030000000018000000ba9b3b030000000017000000bb9b3b030000000016000000bc9b3b030000000015000000bd9b3b030000000014000000be9b3b030000000013000000bf9b3b030000000012000000c09b3b030000000011000000c19b3b030000000010000000c29b3b03000000000f000000c39b3b03000000000e000000c49b3b03000000000d000000c59b3b03000000000c000000c69b3b03000000000b000000c79b3b03000000000a000000c89b3b030000000009000000c99b3b030000000008000000ca9b3b030000000007000000cb9b3b030000000006000000cc9b3b030000000005000000cd9b3b030000000004000000ce9b3b030000000003000000cf9b3b030000000002000000d09b3b03000000000100000001b19b3b030000000001000000000000007d00000000000000d38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f000000000000000131000000000000004d00000000000000386802000000000000000000000000004e00000000000000aa6208000000000038680200000000004f000000000000007e750e0000000000aa620800000000005000000000000000a3c71400000000007e750e0000000000510000000000000004d21a0000000000a3c71400000000005200000000000000589720000000000004d21a000000000053000000000000004348260000000000589720000000000054000000000000003a832b000000000043482600000000005500000000000000ea343100000000003a832b00000000005600000000000000b531370000000000ea343100000000005700000000000000c91b3d0000000000b53137000000000058000000000000001749430000000000c91b3d00000000005900000000000000905d49000000000017494300000000005a00000000000000ddae4f0000000000905d4900000000005b00000000000000bed6550000000000ddae4f00000000005c000000000000001ceb5b0000000000bed65500000000005d00000000000000d20b6200000000001ceb5b00000000005e00000000000000ea0e680000000000d20b6200000000005f00000000000000ce9c6d0000000000ea0e6800000000006000000000000000cce8720000000000ce9c6d00000000006100000000000000fb27790000000000cce8720000000000620000000000000001817f0000000000fb2779000000000063000000000000004dbc85000000000001817f0000000000640000000000000006978b00000000004dbc850000000000650000000000000040ac91000000000006978b0000000000660000000000000073dc97000000000040ac9100000000006700000000000000cada9d000000000073dc970000000000680000000000000094eba30000000000cada9d000000000069000000000000003b11aa000000000094eba300000000006a00000000000000594db000000000003b11aa00000000006b000000000000007e82b60000000000594db000000000006c000000000000004160bc00000000007e82b600000000006d00000000000000b304c200000000004160bc00000000006e000000000000008fdcc70000000000b304c200000000006f00000000000000d9c5cd00000000008fdcc700000000007000000000000000a3d3d30000000000d9c5cd000000000071000000000000005bb8d90000000000a3d3d300000000007200000000000000507bdf00000000005bb8d900000000007300000000000000f84ae50000000000507bdf00000000007400000000000000779dea0000000000f84ae500000000007500000000000000acabef0000000000779dea0000000000760000000000000048f6f40000000000acabef00000000007700000000000000a107fa000000000048f6f4000000000078000000000000004134ff0000000000a107fa00000000007900000000000000719c0301000000004134ff00000000007a0000000000000042b8080100000000719c0301000000007b0000000000000001f70d010000000042b80801000000007c00000000000000914613010000000001f70d01000000007d000000000000007a0c1501000000009146130100000000d09b3b0300000000c112d15f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 01000000d38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7ed38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7e641f00000000000000b39b3b03000000001f000000b49b3b03000000001e000000b59b3b03000000001d000000b69b3b03000000001c000000b79b3b03000000001b000000b89b3b03000000001a000000b99b3b030000000019000000ba9b3b030000000018000000bb9b3b030000000017000000bc9b3b030000000016000000bd9b3b030000000015000000be9b3b030000000014000000bf9b3b030000000013000000c09b3b030000000012000000c19b3b030000000011000000c29b3b030000000010000000c39b3b03000000000f000000c49b3b03000000000e000000c59b3b03000000000d000000c69b3b03000000000c000000c79b3b03000000000b000000c89b3b03000000000a000000c99b3b030000000009000000ca9b3b030000000008000000cb9b3b030000000007000000cc9b3b030000000006000000cd9b3b030000000005000000ce9b3b030000000004000000cf9b3b030000000003000000d09b3b030000000002000000d19b3b03000000000100000001b29b3b030000000001000000000000007d00000000000000d38ded24c2a932d0722db7c1a33f7582b101d92d52bd489f32bc68801ff92e7e0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001f000000000000000131000000000000004d00000000000000386802000000000000000000000000004e00000000000000aa6208000000000038680200000000004f000000000000007e750e0000000000aa620800000000005000000000000000a3c71400000000007e750e0000000000510000000000000004d21a0000000000a3c71400000000005200000000000000589720000000000004d21a000000000053000000000000004348260000000000589720000000000054000000000000003a832b000000000043482600000000005500000000000000ea343100000000003a832b00000000005600000000000000b531370000000000ea343100000000005700000000000000c91b3d0000000000b53137000000000058000000000000001749430000000000c91b3d00000000005900000000000000905d49000000000017494300000000005a00000000000000ddae4f0000000000905d4900000000005b00000000000000bed6550000000000ddae4f00000000005c000000000000001ceb5b0000000000bed65500000000005d00000000000000d20b6200000000001ceb5b00000000005e00000000000000ea0e680000000000d20b6200000000005f00000000000000ce9c6d0000000000ea0e6800000000006000000000000000cce8720000000000ce9c6d00000000006100000000000000fb27790000000000cce8720000000000620000000000000001817f0000000000fb2779000000000063000000000000004dbc85000000000001817f0000000000640000000000000006978b00000000004dbc850000000000650000000000000040ac91000000000006978b0000000000660000000000000073dc97000000000040ac9100000000006700000000000000cada9d000000000073dc970000000000680000000000000094eba30000000000cada9d000000000069000000000000003b11aa000000000094eba300000000006a00000000000000594db000000000003b11aa00000000006b000000000000007e82b60000000000594db000000000006c000000000000004160bc00000000007e82b600000000006d00000000000000b304c200000000004160bc00000000006e000000000000008fdcc70000000000b304c200000000006f00000000000000d9c5cd00000000008fdcc700000000007000000000000000a3d3d30000000000d9c5cd000000000071000000000000005bb8d90000000000a3d3d300000000007200000000000000507bdf00000000005bb8d900000000007300000000000000f84ae50000000000507bdf00000000007400000000000000779dea0000000000f84ae500000000007500000000000000acabef0000000000779dea0000000000760000000000000048f6f40000000000acabef00000000007700000000000000a107fa000000000048f6f4000000000078000000000000004134ff0000000000a107fa00000000007900000000000000719c0301000000004134ff00000000007a0000000000000042b8080100000000719c0301000000007b0000000000000001f70d010000000042b80801000000007c00000000000000914613010000000001f70d01000000007d000000000000007b0c1501000000009146130100000000d19b3b0300000000c112d15f00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000`,
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctx := newParseCtx()
			ctx.trxTraceMap[test.trxID] = &pbcodec.TransactionTrace{
				InstructionTraces: []*pbcodec.InstructionTrace{
					&pbcodec.InstructionTrace{},
				},
			}
			err := ctx.readAccountChange(test.line)
			require.Equal(t, test.expectedErr, err)
			text.NewEncoder(os.Stdout).Encode(ctx, nil)

		})
	}
}
