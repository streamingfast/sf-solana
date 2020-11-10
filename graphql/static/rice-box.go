package static

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "favorites.json",
		FileModTime: time.Unix(1588011909, 0),

		Content: string("{\n  \"eos\": [\n    {\n      \"label\": \"EOS - Search Stream (Forward)\",\n      \"query\": \"subscription ($query: String!, $cursor: String, $limit: Int64) {\\n  searchTransactionsForward(query: $query, limit: $limit, cursor: $cursor) {\\n    undo\\n    cursor\\n    trace {\\n    \\tblock {\\n      \\tnum\\n        id\\n        confirmed\\n        timestamp\\n        previous\\n      }\\n      id\\n      matchingActions {\\n      \\taccount\\n        name\\n        json\\n        seq\\n        receiver\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": {\n        \"mainnet\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\",\n        \"jungle\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\",\n        \"kylin\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\",\n        \"local\": \"{\\n  \\\"query\\\": \\\"action:onblock\\\",\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\"\n      }\n    },\n    {\n      \"label\": \"EOS - Search Query (Forward)\",\n      \"query\": \"query ($query: String!, $cursor: String, $limit: Int64, $low: Int64, $high: Int64) {\\n  searchTransactionsForward(query: $query, lowBlockNum: $low, highBlockNum: $high, limit: $limit, cursor: $cursor) {\\n    results {\\n      undo\\n      cursor\\n      trace {\\n        block {\\n          num\\n          id\\n          confirmed\\n          timestamp\\n          previous\\n        }\\n        id\\n        matchingActions {\\n          account\\n          name\\n          json\\n          seq\\n          receiver\\n        }\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"low\\\": -500,\\n  \\\"high\\\": -1,\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\"\n    },\n    {\n      \"label\": \"EOS - Search Stream (Backward)\",\n      \"query\": \"subscription ($query: String!, $cursor: String, $limit: Int64, $low: Int64) {\\n  searchTransactionsBackward(query: $query, lowBlockNum: $low, limit: $limit, cursor: $cursor) {\\n    cursor\\n    trace {\\n    \\tblock {\\n      \\tnum\\n        id\\n        confirmed\\n        timestamp\\n        previous\\n      }\\n      id\\n      matchingActions {\\n      \\taccount\\n        name\\n        json\\n        seq\\n        receiver\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"low\\\": 1,\\n  \\\"limit\\\": 10\\n}\"\n    },\n    {\n      \"label\": \"EOS - Search Query (Backward)\",\n      \"query\": \"query ($query: String!, $cursor: String, $limit: Int64, $low: Int64, $high: Int64) {\\n  searchTransactionsBackward(query: $query, lowBlockNum: $low, highBlockNum: $high, limit: $limit, cursor: $cursor) {\\n    results {\\n      cursor\\n      trace {\\n        block {\\n          num\\n          id\\n          confirmed\\n          timestamp\\n          previous\\n        }\\n        id\\n        matchingActions {\\n          account\\n          name\\n          json\\n          seq\\n          receiver\\n        }\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": \"{\\n  \\\"query\\\": \\\"receiver:eosio.token action:transfer -data.quantity:'0.0001 EOS'\\\",\\n  \\\"low\\\": -500,\\n  \\\"high\\\": -1,\\n  \\\"cursor\\\": \\\"\\\",\\n  \\\"limit\\\": 10\\n}\"\n    },\n    {\n      \"label\": \"EOS - Time Ranges\",\n      \"query\": \"query ($start: Time!, $end: Time!) {\\n  low: blockIDByTime(time: $start) {\\n    num\\n    id\\n  }\\n  high: blockIDByTime(time: $end) {\\n    num\\n    id\\n  }\\n}\\n\",\n      \"variables\": \"{\\\"start\\\": \\\"2019-11-14T12:00:00.000Z\\\", \\\"end\\\": \\\"2019-11-14T17:00:00.000Z\\\"}\\n\\n\"\n    },\n    {\n      \"label\": \"EOS - (Alpha) Get Block By Id\",\n      \"query\": \"query ($blockId: String!) {\\n  block(id: $blockId) {\\n    id\\n    num\\n    dposLIBNum\\n    executedTransactionCount\\n    irreversible\\n    header {\\n      id\\n      num\\n      timestamp\\n      producer\\n      previous\\n    }\\n    transactionTraces(first: 5) {\\n      pageInfo {\\n        startCursor\\n        endCursor\\n      }\\n      edges {\\n        cursor\\n        node {\\n          id\\n          status\\n          topLevelActions {\\n            account\\n            name\\n            receiver\\n            json\\n          }\\n        }\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": {\n        \"mainnet\": \"{\\\"blockId\\\": \\\"063a7e525142f64d7465bbebc690afbb228bff7d7e0ffda31d9a06106fbc1982\\\"}\\n\",\n        \"jungle\": \"{\\\"blockId\\\": \\\"047c7822f396e64b9cbb28cc2b199b8e5a4c33c894b8742eab646e670486bb0d\\\"}\\n\",\n        \"kylin\": \"{\\\"blockId\\\": \\\"05609e94b57cdea5ce4ff8afa89070d37a85923855f0d41efdbd956dbaddb5f7\\\"}\\n\"\n      }\n    },\n    {\n      \"label\": \"EOS - (Alpha) Get Block By Num\",\n      \"query\": \"query ($blockNum: Uint32) {\\n  block(num: $blockNum) {\\n    id\\n    num\\n    dposLIBNum\\n    executedTransactionCount\\n    irreversible\\n    header {\\n      id\\n      num\\n      timestamp\\n      producer\\n      previous\\n    }\\n    transactionTraces(first: 5) {\\n      pageInfo {\\n        startCursor\\n        endCursor\\n      }\\n      edges {\\n        cursor\\n        node {\\n          id\\n          status\\n          topLevelActions {\\n            account\\n            name\\n            receiver\\n            json\\n          }\\n        }\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": {\n        \"mainnet\": \"{\\\"blockNum\\\": 104782163}\\n\",\n        \"jungle\": \"{\\\"blockNum\\\": 75430834}\\n\",\n        \"kylin\": \"{\\\"blockNum\\\": 90340699}\\n\"\n      }\n    },\n    {\n      \"label\": \"EOS - (Alpha) Get Tokens\",\n      \"query\": \"query {\\n  tokens {\\n    blockRef {\\n      id\\n      number\\n    }\\n    pageInfo {\\n      startCursor\\n      endCursor\\n    }\\n    edges {\\n      cursor\\n      node {\\n        symbol\\n        contract\\n        holders\\n        totalSupply\\n        \\n      }\\n    }\\n  }\\n}\",\n      \"variables\": \"\"\n    },\n    {\n      \"label\": \"EOS - (Alpha) Get Token Balances\",\n      \"query\": \"query($contract: String!,$symbol:String!,$limit:\\tUint32, $opts: [ACCOUNT_BALANCE_OPTION!]) {\\n  tokenBalances(contract: $contract, symbol: $symbol,limit: $limit, options: $opts) {\\n    blockRef {\\n      id\\n      number\\n    }\\n    pageInfo {\\n      startCursor\\n      endCursor\\n    }\\n    edges {\\n      node {\\n        account\\n        contract\\n        symbol\\n        precision\\n        amount\\n        balance\\n      }\\n    }\\n  }\\n}\",\n      \"variables\": \"{\\n  \\\"contract\\\": \\\"eosio.token\\\",\\n  \\\"symbol\\\": \\\"EOS\\\",\\n  \\\"opts\\\": [\\\"EOS_INCLUDE_STAKED\\\"],\\n  \\\"limit\\\": 10\\n}\"\n    },\n    {\n      \"label\": \"EOS - (Alpha) Get Account Balances\",\n      \"query\": \"query($account: String!,$limit:\\tUint32, $opts: [ACCOUNT_BALANCE_OPTION!]) {\\n  accountBalances(account: $account,limit: $limit, options: $opts) {\\n    blockRef {\\n      id\\n      number\\n    }\\n    pageInfo {\\n      startCursor\\n      endCursor\\n    }\\n    edges {\\n      node {\\n        account\\n        contract\\n        symbol\\n        precision\\n        amount\\n        balance\\n      }\\n    }\\n  }\\n}\",\n      \"variables\": \"{\\n  \\\"account\\\": \\\"eosio.token\\\",\\n  \\\"opts\\\": [\\\"EOS_INCLUDE_STAKED\\\"],\\n  \\\"limit\\\": 0\\n}\"\n    }\n  ],\n  \"eth\": [\n    {\n      \"label\": \"ETH - Stream Search\",\n      \"query\": \"subscription ($query: String!, $sort: SORT, $low: Int64, $high: Int64, $limit: Int64) {\\n  searchTransactions(indexName: CALLS, query: $query, sort: $sort, lowBlockNum: $low, highBlockNum: $high, limit: $limit) {\\n    undo\\n    cursor\\n    block {\\n      hash\\n      number\\n    }\\n    node {\\n      hash\\n      from\\n      to\\n      value(encoding: ETHER)\\n      gasLimit\\n      gasPrice(encoding: ETHER)\\n      matchingCalls {\\n        from\\n        to\\n        value\\n        inputData\\n        returnData\\n        balanceChanges {\\n          address\\n          newValue(encoding: ETHER)\\n          reason\\n        }\\n        logs {\\n          address\\n          topics\\n          data\\n        }\\n        storageChanges {\\n          key\\n          oldValue\\n          newValue\\n        }\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": \"{\\\"query\\\": \\\"-value:0\\\",\\\"sort\\\":\\\"ASC\\\",\\\"low\\\":-1,\\\"limit\\\":10}\\n\"\n    },\n    {\n      \"label\": \"ETH - Get Transaction\",\n      \"query\": \"query ($hash: String!) {\\n  transaction(hash: $hash) {\\n    hash\\n    from\\n    to\\n    value(encoding: ETHER)\\n    gasLimit\\n    gasPrice(encoding: ETHER)\\n    flatCalls {\\n      from\\n      to\\n      value\\n      inputData\\n      returnData\\n      balanceChanges {\\n        address\\n        newValue(encoding: ETHER)\\n        reason\\n      }\\n      logs {\\n        address\\n        topics\\n        data\\n      }\\n      storageChanges {\\n        key\\n        oldValue\\n        newValue\\n      }\\n    }\\n  }\\n}\\n\",\n      \"variables\": {\n        \"mainnet\": \"{\\\"hash\\\": \\\"df98f1f0c3e962ac829a165e0f54e35a25148d1b1322808b258e49ab66dce697\\\"}\\n\",\n        \"ropsten\": \"{\\\"hash\\\": \\\"9c3a9d3a4634c4953b8b3b02ac22777d82d6132019c8a3ad5c5f03e7da7f8be5\\\"}\\n\"\n      }\n    },\n    {\n      \"label\":\"ETH - Stream Blocks\",\n      \"query\": \"subscription {\\n  blocks {\\n    undo\\n    node {\\n      hash\\n      number\\n      transactionTraces {\\n        edges {\\n          node {\\n            from\\n            to\\n            flatCalls {\\n              callType\\n              depth\\n            }\\n          }\\n        }\\n      }\\n    }\\n  }\\n}\\n\"\n    },\n    {\n      \"label\":\"ETH - (Alpha) Execute\",\n      \"query\": \"query ($from: String!, $to: String!, $value: String, $input: String!, $gasLimit: Uint64!, $gasPrice: String!, $blockRef: BlockRef) {\\n  _alphaExecute(from: $from, to: $to, input: $input, gasLimit: $gasLimit, gasPrice: $gasPrice, value: $value, atBlock: $blockRef) {\\n    hash\\n    from\\n    to\\n    value(encoding: ETHER)\\n    gasLimit\\n    gasPrice(encoding: ETHER)\\n    flatCalls {\\n      from\\n      to\\n      value\\n      inputData\\n      returnData\\n      balanceChanges {\\n        address\\n        newValue(encoding: ETHER)\\n        reason\\n      }\\n      logs {\\n        address\\n        topics\\n        data\\n      }\\n      storageChanges {\\n        key\\n        oldValue\\n        newValue\\n      }\\n    }\\n  }\\n}\\n\",\n      \"alpha\": true,\n      \"variables\": {\n        \"mainnet\": \"{\\n  \\\"from\\\": \\\"965e553fa090e747ff6b85150ca4d8378b3c1711\\\",\\n  \\\"to\\\": \\\"d7b9a9b2f665849c4071ad5af77d8c76aa30fb32\\\",\\n  \\\"input\\\": \\\"\\\",\\n  \\\"value\\\": \\\"b1904d40e790c1\\\",\\n  \\\"gasLimit\\\": \\\"21000\\\",\\n  \\\"gasPrice\\\": \\\"0x8f0d1800\\\",\\n  \\\"blockRef\\\": {\\n    \\\"number\\\": 945754\\n  }\\n}\\n\",\n        \"ropsten\": \"{\\n  \\\"from\\\": \\\"d820c9EAC4E44bDc645C81a8E554A65068E20BE8\\\",\\n  \\\"to\\\": \\\"1056A76b5ffF557d0C70Cef64D45BDB172dBcAea\\\",\\n  \\\"input\\\": \\\"a6f2ae3a\\\",\\n  \\\"value\\\": \\\"4139c1192c560000\\\",\\n  \\\"gasLimit\\\": \\\"237644\\\",\\n  \\\"gasPrice\\\": \\\"77359400\\\",\\n  \\\"blockRef\\\": {\\n    \\\"number\\\": 7267470\\n  }\\n}\\n\"\n      }\n    },\n    {\n      \"label\":\"ETH - (Alpha) Replay\",\n      \"query\": \"query ($hash: String!, $blockRef: BlockRef) {\\n  _alphaReplay(hash: $hash, atBlock: $blockRef) {\\n    hash\\n    from\\n    to\\n    value(encoding: ETHER)\\n    gasLimit\\n    gasPrice(encoding: ETHER)\\n    block {\\n      number\\n    }\\n    flatCalls {\\n      from\\n      to\\n      value\\n      inputData\\n      returnData\\n      balanceChanges {\\n        address\\n        newValue(encoding: ETHER)\\n        reason\\n      }\\n      logs {\\n        address\\n        topics\\n        data\\n      }\\n      storageChanges {\\n        key\\n        oldValue\\n        newValue\\n      }\\n    }\\n  }\\n}\\n      \",\n      \"alpha\": true,\n      \"variables\": {\n        \"mainnet\": \"{\\n  \\\"hash\\\": \\\"5913720c8414ccb9d97dd1aae426146ae9fdf8794dfb9f0321f6b2a0a34eea1a\\\",\\n  \\\"blockRef\\\": {\\n    \\\"number\\\": 75042\\n  }\\n}\\n\",\n        \"ropsten\": \"{\\n  \\\"hash\\\": \\\"9c3a9d3a4634c4953b8b3b02ac22777d82d6132019c8a3ad5c5f03e7da7f8be5\\\",\\n  \\\"blockRef\\\": {\\n    \\\"number\\\": 7267470\\n  }\\n}\\n\"\n      }\n    }\n  ]\n}\n"),
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    "graphiql.html",
		FileModTime: time.Unix(1605015856, 0),

		Content: string("<!DOCTYPE html>\n<!--\n Copyright 2019 dfuse Platform Inc.\n\n Licensed under the Apache License, Version 2.0 (the \"License\");\n you may not use this file except in compliance with the License.\n You may obtain a copy of the License at\n\n      http://www.apache.org/licenses/LICENSE-2.0\n\n Unless required by applicable law or agreed to in writing, software\n distributed under the License is distributed on an \"AS IS\" BASIS,\n WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n See the License for the specific language governing permissions and\n limitations under the License.\n-->\n\n<html>\n<head>\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/graphiql@0.16.0/graphiql.css\"/>\n    <link rel=\"stylesheet\" href=\"graphiql_dfuse_override.css\"/>\n    <link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Lato&display=swap\">\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react.min.js\"></script>\n    <script src=\"https://cdnjs.cloudflare.com/ajax/libs/react/15.5.4/react-dom.min.js\"></script>\n    <script src=\"//unpkg.com/graphiql@0.16.0/graphiql.js\"></script>\n    <script src=\"//unpkg.com/subscriptions-transport-ws@0.8.3/browser/client.js\"></script>\n    <script src=\"//unpkg.com/graphiql-subscriptions-fetcher@0.0.2/browser/client.js\"></script>\n    <script src=\"helper.js\"></script>\n    <script src=\"https://unpkg.com/@dfuse/client@0.3.9/dist/dfuse-client.umd.js\"></script>\n</head>\n<body style=\"width: 100%; height: 100%; margin: 0; overflow: hidden;\">\n\n<div id=\"graphiql\" style=\"height: 100vh;\">Loading...</div>\n\n<script>\n    window.DfuseClientConfig = --== json . ==--;\n\n</script>\n\n<script>\n    const url = new URL(window.location.href)\n    const urlPathSegments = url.toString().split(\"/\");\n    const queryParams = url.searchParams\n\n    const server = urlPathSegments[2];\n    const proto = urlPathSegments[0];\n    const alphaSchema = isAlphaSchemaQueryParamFound(queryParams)\n\n    async function initialize() {\n\n        loadGraphiql()\n    }\n\n    function loadGraphiql() {\n                const subscriptionsClient = new window.SubscriptionsTransportWs.SubscriptionClient((proto == 'https:' ? 'wss://' : 'ws://') + server + '/graphql',\n                    {\n                        reconnect: true,\n                        connectionCallback: (error) => {\n                            if (error != null) {\n                                alert(error.message)\n                            }\n                        }\n                    });\n\n                let activeQuery = fetchQueryProp(queryParams)\n                let activeVariables = fetchVariablesProp(queryParams)\n\n                const graphqlFetcher = graphQLFetcherFactory(queryParams)\n                const subscriptionsFetcher = window.GraphiQLSubscriptionsFetcher.graphQLFetcher(subscriptionsClient, graphqlFetcher);\n\n                ReactDOM.render(\n                    React.createElement(GraphiQL, {\n                        fetcher: subscriptionsFetcher,\n                        query: activeQuery,\n                        variables: activeVariables,\n                        onEditQuery: function(query) {\n                            activeQuery = query || undefined\n                            pushState(url, activeQuery, activeVariables)\n                        },\n                        onEditVariables: function(variables) {\n                            activeVariables = variables || undefined\n                            pushState(url, activeQuery, activeVariables)\n                        },\n                    }),\n                    document.getElementById(\"graphiql\")\n                );\n    }\n    initialize();\n\n</script>\n</body>\n</html>\n\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "graphiql_dfuse_override.css",
		FileModTime: time.Unix(1588011909, 0),

		Content: string("/**\n * Copyright 2019 dfuse Platform Inc.\n *\n * Licensed under the Apache License, Version 2.0 (the \"License\");\n * you may not use this file except in compliance with the License.\n * You may obtain a copy of the License at\n *\n *      http://www.apache.org/licenses/LICENSE-2.0\n *\n * Unless required by applicable law or agreed to in writing, software\n * distributed under the License is distributed on an \"AS IS\" BASIS,\n * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n * See the License for the specific language governing permissions and\n * limitations under the License.\n */\n\n.graphiql-container .topBar {\n    background-image: url('https://www.dfuse.io/hubfs/dfuse-graphiQL-logo-white-03-01.svg');\n    background-size: auto 31px;\n    background-repeat: no-repeat;\n    background-position: 10px 12px;\n    background-position-x: right;\n    height:44px;\n    border-bottom: 1px solid #c9cacf;\n}\n\n.graphiql-container .docExplorerShow, .graphiql-container .historyShow {\n    border-bottom: 1px solid #c9cacf!important;\n}\n\n.CodeMirror-linenumber {\n    color:#9a9ba3;\n}\n\n.graphiql-container .execute-options > li.selected, .graphiql-container .toolbar-menu-items > li.hover, .graphiql-container .toolbar-menu-items > li:active, .graphiql-container .toolbar-menu-items > li:hover, .graphiql-container .toolbar-select-options > li.hover, .graphiql-container .toolbar-select-options > li:active, .graphiql-container .toolbar-select-options > li:hover, .graphiql-container .history-contents > p:hover, .graphiql-container .history-contents > p:active {\n    background: #ff4660;\n    color: #fff;\n}\n\n.graphiql-container .topBarWrap {\n    background-image: linear-gradient(to right, rgb(255, 70, 96) 8%, rgb(65, 17, 160) 93%);\n    background-repeat: no-repeat;\n    background-position: 0px 0px;\n    background-size: cover;\n}\n\n.graphiql-container .topBar .title {\n    display:none;\n}\n\n.graphiql-container .execute-button {\n    background: rgba(255,255,255,.3);\n    border-radius: 17px;\n    border: 0px solid rgba(0, 0, 0, 0.25);\n    -webkit-box-shadow: none !important;\n    box-shadow: none !important;\n    cursor: pointer;\n    fill: #fff;\n    height: 34px;\n    margin: 0;\n    padding: 0;\n    width: 34px;\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n}\n\n.graphiql-container .execute-button svg {\n    position:relative;\n    top:1px;\n    left:1px;\n}\n\n.graphiql-container .execute-button-wrap {\n    margin: 0px 14px 0 7px;\n\n}\n\n\n.graphiql-container .toolbar-button {\n    background: rgba(255,255,255,.3);\n    -webkit-box-shadow: none !important;\n    box-shadow: none !important;\n    border-radius: 5px;\n    color: #fff;\n    cursor: pointer;\n    display: inline-block;\n    margin: 0 5px;\n    padding: 8px 20px 8px;\n    text-decoration: none;\n    text-overflow: ellipsis;\n    white-space: nowrap;\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n}\n\n.docExplorerShow,\n.history-title,\n.doc-explorer-title {\n    font-family: 'Lato', sans sherif;\n    text-transform: uppercase;\n    font-size: 13px;\n    letter-spacing: 1px;\n}\n\n.graphiql-container .toolbar-button:hover,\n.graphiql-container .execute-button:hover {\n    background: rgba(255,255,255,.4);\n}\n\n.graphiql-container .docExplorerShow, .graphiql-container .historyShow {\n    background: #fff;\n\n    border-bottom: 1px solid #d0d0d0;\n    border-right: none;\n    border-top: none;\n    color: #4111A0;\n    cursor: pointer;\n    font-size: 14px;\n    margin: 0;\n    outline: 0;\n    padding: 2px 20px 0 18px;\n}\n\n.graphiql-container .docExplorerShow:before {\n    border-left: 2px solid #4111A0;\n    border-top: 2px solid #4111A0;\n}\n\n.graphiql-container .doc-explorer-title-bar, .graphiql-container .history-title-bar {\n    height:44px;\n    place-items: center;\n}\n\n.graphiql-container .doc-explorer-contents, .graphiql-container .history-contents {\n    top: 57px;\n\n}\n\n.graphiql-container, .graphiql-container button, .graphiql-container input {\n    color: #22244b;\n}\n\n\n.graphiql-container .search-box {\n    border-bottom: none;\n    display: block;\n    font-size: 14px;\n    margin: 10px 0px 32px 0px;\n    position: relative;\n}\n\n.graphiql-container .search-box:before {\n    top: 4px;\n    left: 13px;\n}\n.graphiql-container .search-box > input {\n    -webkit-appearance: none;\n    -moz-appearance: none;\n    box-shadow: none !important;\n    outline: none;\n    background: #f8f8fa;\n    border: 1px solid #e1e1e4 !important;\n    padding: 13px 20px 13px 40px;\n    box-sizing: border-box;\n    margin: 0px 0px;\n    width: 100%;\n    outline: none;\n    color: #586090 !important;\n    border-radius: 8px;\n}\n\n.CodeMirror-scroll {\n    background: #f8f8fa;\n}\n\n.graphiql-container .variable-editor-title {\n    background: #e7e7ec;\n    border-top:  1px solid #dbdbe8;\n    border-bottom:  1px solid #dbdbe8;\n}\n\n\n.graphiql-container .result-window .CodeMirror-gutters, .CodeMirror-gutters {\n    background: #f0f0f5;\n    border-color: #dbdbe8;\n}\n\n.cm-keyword {\n    color: #030b8c;\n}\n\n.cm-punctuation {\n    color: #a2a2a2;\n}\n\n.cm-variable {\n    color: #4bb310;\n}\n\n.cm-atom {\n    color: #b9f;\n}\n\n.cm-property {\n    color: #4f6bc1;\n}\n\n.cm-number {\n    color: #5bcfff;\n}\n\n.cm-attribute {\n    color: #8B2BB9;\n}\n\n.cm-def {\n    color: #e60a0a;\n}\n\n.cm-string {\n    color: #E48F32;\n}\n\n.graphiql-container .type-name {\n    color: #0db4de;\n}"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "helper.js",
		FileModTime: time.Unix(1605016016, 0),

		Content: string("/**\n * Copyright 2019 dfuse Platform Inc.\n *\n * Licensed under the Apache License, Version 2.0 (the \"License\");\n * you may not use this file except in compliance with the License.\n * You may obtain a copy of the License at\n *\n *      http://www.apache.org/licenses/LICENSE-2.0\n *\n * Unless required by applicable law or agreed to in writing, software\n * distributed under the License is distributed on an \"AS IS\" BASIS,\n * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.\n * See the License for the specific language governing permissions and\n * limitations under the License.\n */\n\nfunction refreshToken(server, token, apiKey, onCompletion, onError) {\n    if(token === \"\") {\n        return getToken(apiKey, onCompletion, onError)\n    }\n\n    const jwt = parseJwt(token);\n    const exp = jwt[\"exp\"];\n    const now = Date.now() / 1000;\n\n    console.log(\"exp  : \" + exp);\n    console.log(\"now  : \" + now);\n    const remainingTime = exp - now;\n\n    console.log(\"Time remaining in second: \" + remainingTime);\n    if (remainingTime < 60 * 60) {\n        return getToken(apiKey, onCompletion, onError)\n    }\n\n    onCompletion(token);\n}\n\nfunction getToken(apiKey, onCompletion, onError) {\n    const url = \"https://auth.dfuse.io/v1/auth/issue\";\n    const r = new XMLHttpRequest();\n    r.open(\"POST\", url, false);\n    r.setRequestHeader(\"Content-type\", \"application/json\");\n    r.onreadystatechange = function() {//Call a function when the state changes.\n        if(r.readyState === 4) {\n            if (r.status === 200) {\n                console.log(\"Got new token: \" + r.response);\n                const responseToken =  JSON.parse(r.response)\n                onCompletion(responseToken[\"token\"]);\n            } else {\n                alert(\"Error: \" + r.status + \" - \" + r.response);\n                onError(r.status, r.response);\n            }\n        }\n    };\n\n    r.send('{\"api_key\":\"' + apiKey + '\"}');\n}\n\nfunction graphQLFetcherFactory(urlQueryParams) {\n    let graphqlUrl = \"/graphql\"\n    if (isAlphaSchemaQueryParamFound(urlQueryParams)) {\n        console.log(\"Requesting alpha endpoints\")\n        graphqlUrl += \"?alpha-schema=true\"\n    }\n\n    return (graphQLParams) => {\n        return fetch(graphqlUrl, {\n            method: \"post\",\n//            headers: {\n//                Authorization: \"Bearer \" + token,\n//            },\n            body: JSON.stringify(graphQLParams),\n            credentials: \"include\",\n        }).then(function (response) {\n            return response.text();\n        }).then(function (responseBody) {\n            try {\n                return JSON.parse(responseBody);\n            } catch (error) {\n                return error.message;\n            }\n        }).catch(function (error) {\n            console.log(\"Error:\", error);\n            alert(error)\n        });\n    }\n}\n\nfunction isAlphaSchemaQueryParamFound(urlQueryParams) {\n    if (urlQueryParams.has(\"alpha-schema\")) {\n        return urlQueryParams.get(\"alpha-schema\") === \"true\"\n    }\n\n    if (urlQueryParams.has(\"alphaSchema\")) {\n        return urlQueryParams.get(\"alphaSchema\") === \"true\"\n    }\n\n    return false\n}\n\nfunction fetchQueryProp(queryParams) {\n    if (!queryParams.has(\"query\")) {\n        return undefined\n    }\n\n    try {\n        return window.atob(queryParams.get(\"query\"))\n    } catch (error) {\n        console.error(\"query params 'query' is not a valid base64 object\")\n        return undefined\n    }\n}\n\nfunction fetchVariablesProp(queryParams) {\n    if (!queryParams.has(\"variables\")) {\n        return undefined\n    }\n\n    try {\n        return window.atob(queryParams.get(\"variables\"))\n    } catch (error) {\n        console.error(\"query params 'variables' is not a valid base64 object\")\n        return undefined\n    }\n}\n\nfunction pushState(url, query, variables) {\n    const queryParams = []\n    if (query !== undefined) {\n        queryParams.push(`query=${window.btoa(query)}`)\n    }\n\n    if (variables !== undefined) {\n        queryParams.push(`variables=${window.btoa(variables)}`)\n    }\n\n    if (queryParams.length <= 0) {\n        return\n    }\n\n    window.history.pushState(\"\", \"New Query\", `${url.pathname}?${queryParams.join(\"&\")}`);\n}\n\nasync function getConfig() {\n    if (window.location.hostname === \"localhost\") {\n        return await fetchConfig()\n    }\n\n    const parts = window.location.host.split(\".\");\n    return { network: parts[0], protocol: parts[1] }\n}\n\nfunction fetchFavorites() {\n    console.info(\"Fetching favorites JSON data\")\n    return fetch(\"/graphiql/favorites.json\")\n            .then((response) => response.json())\n            .then((body) => {\n                console.log(\"Got favorites JSON data\")\n                return body\n            })\n            .catch((error) => {\n                console.log(\"Fetch favorites JSON data error\", error);\n                return {}\n            })\n}\n\nfunction fetchConfig() {\n    console.info(\"Fetching config JSON data\")\n    return fetch(\"/graphiql/config.json\")\n            .then((response) => response.json())\n            .then((body) => {\n                console.log(\"Got config JSON data\")\n                return body\n            })\n            .catch((error) => {\n                console.log(\"Fetch config JSON data error\", error);\n                return {}\n            })\n}\n\nfunction getFavoriteFromStorage() {\n    const storageItem = localStorage.getItem(\"graphiql:favorites\");\n    console.log(\"Retrieved client favorites from browser storage\")\n\n    store = { favorites: [] };\n    if (storageItem !== null) {\n        store = JSON.parse(storageItem);\n    }\n\n    return store\n}\n\nfunction setFavoriteFromStorage(store) {\n    console.log(\"Saving client favorites to browser storage\")\n    localStorage.setItem(\"graphiql:favorites\", toJSON(store));\n}\n\nasync function reconfigureGraphiQLStorage(protocol, network, alphaSchema) {\n    const isFirstTime = localStorage.getItem(\"dfuse:graphiql:is_first_time\");\n    if (isFirstTime == null) {\n        // Let's open the history pane the first time the user opens this page\n        localStorage.setItem(\"graphiql:historyPaneOpen\", toJSON(true));\n    }\n\n    const serverFavorites = await setFavorites(protocol, network, alphaSchema)\n\n    localStorage.setItem(\"dfuse:graphiql:is_first_time\", toJSON(false))\n\n    if (isFirstTime == null && serverFavorites.length > 0) {\n        localStorage.setItem(\"graphiql:query\", serverFavorites[0].query);\n\n        if (serverFavorites[0].variables) {\n            localStorage.setItem(\"graphiql:variables\", serverFavorites[0].variables);\n        }\n    }\n}\n\nasync function setFavorites(protocol, network, alphaSchema) {\n    const favoritesByProtocolMap = await fetchFavorites()\n\n    console.log(`Looking for favorites for given ${protocol}/${network} values`)\n    const serverFavorites = favoritesByProtocolMap[protocol]\n    if (serverFavorites == null) {\n        console.log(\"Favorites not found for this protocol/network values.\")\n        return\n    }\n\n    const store = getFavoriteFromStorage()\n\n    // Clear all dfuse managed favorites, we will add them back\n    store[\"favorites\"] = store[\"favorites\"].filter((value) =>\n        // We keep only favorites that don't have the `procotol`\n        value.protocol == null\n    )\n\n    console.log(\"Favorites store prior update\")\n    serverFavorites.reverse().forEach((favorite) => {\n        if (!alphaSchema && favorite.alpha) {\n            return\n        }\n\n        favorite.favorite = true\n        favorite.protocol = protocol\n        if (favorite.variables && typeof favorite.variables === \"object\") {\n            const networkVariables = favorite.variables[network]\n            if (networkVariables != null) {\n                favorite.variables = networkVariables\n            }\n        }\n\n        store[\"favorites\"] = updateFavorite(store[\"favorites\"], favorite);\n    })\n    console.log(\"Favorites store after update\")\n\n    setFavoriteFromStorage(store)\n\n    // We reverse it again because the `reverse` operation is \"in-place\"\n    return serverFavorites.reverse()\n}\n\nfunction updateFavorite(favorites, fav) {\n    const index = favorites.findIndex(f => (f.label === fav.label));\n    if (index >= 0) {\n        console.log(`Updating favorite ${fav.label}`)\n        favorites[index] = fav\n    } else {\n        console.log(`Adding favorite ${fav.label}`)\n        favorites.push(fav)\n    }\n\n    return favorites\n}\n\nfunction toJSON(input) {\n    return JSON.stringify(input)\n}\n"),
	}
	file6 := &embedded.EmbeddedFile{
		Filename:    "schema.graphql",
		FileModTime: time.Unix(1605020183, 0),

		Content: string("schema {\n  query: Queries\n  subscription: Subscription\n  mutation: Mutations\n}\n\ntype Queries {\n}\n\ntype Subscription {\n  orderBook(\n    dexAddress: String!\n  ): OrderBook\n}\n\ntype Mutations {\n}\n\ntype OrderBook {\n  type: OrderBookType!\n  orders: [Order!]!\n}\n\nenum OrderBookType {\n  ASK\n  BID\n}\n\ntype Order {\n}\n\nscalar Int32\nscalar Int64\nscalar JSON\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1605020183, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "favorites.json"
			file3, // "graphiql.html"
			file4, // "graphiql_dfuse_override.css"
			file5, // "helper.js"
			file6, // "schema.graphql"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`build`, &embedded.EmbeddedBox{
		Name: `build`,
		Time: time.Unix(1605020183, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"favorites.json":              file2,
			"graphiql.html":               file3,
			"graphiql_dfuse_override.css": file4,
			"helper.js":                   file5,
			"schema.graphql":              file6,
		},
	})
}
