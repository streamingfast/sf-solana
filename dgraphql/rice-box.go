package dgraphql

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "get_all_registered_tokens.graphql",
		FileModTime: time.Unix(1608159398, 0),

		Content: string("{\n  registeredTokens {\n    address\n    mintAuthority\n    freezeAuthority\n    supply\n    decimals\n    symbol\n    name\n    logo\n    website\n  }\n}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "get_all_tokens.graphql",
		FileModTime: time.Unix(1608159398, 0),

		Content: string("{\n  tokens {\n    address\n    mintAuthority\n    freezeAuthority\n    supply\n    decimals\n  }\n}\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "get_registered_token.graphql",
		FileModTime: time.Unix(1608159398, 0),

		Content: string("query ($account: String!) {\n  registeredToken(address: $account) {\n    address\n    mintAuthority\n    freezeAuthority\n    supply\n    decimals\n    symbol\n    name\n    logo\n    website\n  }\n}\n"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "get_serum_fill.graphql",
		FileModTime: time.Unix(1608563392, 0),

		Content: string("query ($trader: String!, $market: String) {\n  serumFillHistory(trader: $trader, market: $market) {\n    pageInfo {\n      startCursor\n      endCursor\n    }\n    edges {\n      cursor\n      node {\n        orderId\n        side\n        market {\n          address\n          name\n        }\n        baseToken {\n          address\n          name\n        }\n        quoteToken {\n          address\n          name\n        }\n        lotCount\n        price\n        feeTier\n      }\n    }\n  }\n}\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "get_token.graphql",
		FileModTime: time.Unix(1608159398, 0),

		Content: string("query ($account: String!) {\n  token(address: $account) {\n    address\n    mintAuthority\n    freezeAuthority\n    supply\n    decimals\n  }\n}\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1610381019, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "get_all_registered_tokens.graphql"
			file3, // "get_all_tokens.graphql"
			file4, // "get_registered_token.graphql"
			file5, // "get_serum_fill.graphql"
			file6, // "get_token.graphql"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`examples`, &embedded.EmbeddedBox{
		Name: `examples`,
		Time: time.Unix(1610381019, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"get_all_registered_tokens.graphql": file2,
			"get_all_tokens.graphql":            file3,
			"get_registered_token.graphql":      file4,
			"get_serum_fill.graphql":            file5,
			"get_token.graphql":                 file6,
		},
	})
}
