package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/omni-network/omni/lib/errors"

	"github.com/machinebox/graphql"
	"github.com/spf13/cobra"
)

type QueryConfig struct {
	Addr string
}

func newTestnetCmds() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "testnet",
		Short: "Testnet utility commands",
		Args:  cobra.NoArgs,
	}

	cmd.AddCommand(
		newXmsgsCmd(),
	)

	return cmd
}

func newXmsgsCmd() *cobra.Command {
	var cfg QueryConfig

	cmd := &cobra.Command{
		Use:   "xmsg",
		Short: "Fetch all relevant xmsgs for a given address",
		Long: `Performs a query to the testnet explorer graphQL service,
			matching resulting Xmsgs by the provided address in either source or destination address.`,
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return newFetchXmsgsByAccount(cmd.Context(), cfg)
		},
	}

	bindTestnetNewXmsgsConfig(cmd, &cfg)

	return cmd
}

func newFetchXmsgsByAccount(ctx context.Context, cfg QueryConfig) error {
	client := graphql.NewClient("https://explorer.staging.omni.network/query")
	xmsgcount, err := getXmsgCount(ctx, client)
	if err != nil {
		return errors.Wrap(err, "failed to get xmsgcount")
	}

	var query = `
		query XMsgs($cursor: BigInt!, $limit: BigInt!) {
			xmsgs(cursor: $cursor, limit: $limit) {
				TotalCount
				Edges {
					Cursor
					Node {
						StreamOffset
						SourceMessageSender
						DestAddress
						DestGasLimit
						SourceChainID
						DestChainID
						TxHash
						BlockHeight
						BlockHash
						Status
					}
				}
				PageInfo {
					NextCursor
					HasNextPage
				}
			}
		}
	`
	var cursor = xmsgcount
	variables := map[string]any{
		"cursor": fmt.Sprintf("0x%x", cursor),
		"limit":  fmt.Sprintf("0x%x", 25),
	}
	var filteredResults []any
	for i := 0; i < int(xmsgcount); i += 25 {
		respData, err := executeQuery(ctx, client, query, variables)
		if err != nil {
			return err
		}

		// Filter the results
		xmsgs, ok := respData["xmsgs"].(map[string]any)
		if !ok {
			return errors.New("type map[string]any expected for xmsgs")
		}
		edges, ok := xmsgs["Edges"].([]any)
		if !ok {
			return errors.New("type []any expected for Edges")
		}
		for _, edge := range edges {
			node, ok := edge.(map[string]any)["Node"]
			if !ok {
				return errors.New("type map[string]any expected for Node")
			}
			// Convert the SourceMessageSender and DestAddress to strings
			sourceMessageSender, ok := node.(map[string]any)["SourceMessageSender"].(string)
			if !ok {
				return errors.New("type string expected for SourceMessageSender")
			}
			destAddress, ok := node.(map[string]any)["DestAddress"].(string)
			if !ok {
				return errors.New("type string expected for DestAddress")
			}

			// If either the SourceMessageSender or DestAddress matches the provided address,
			// append the node to filteredResults
			if strings.EqualFold(strings.ToLower(sourceMessageSender), strings.ToLower(cfg.Addr)) || strings.EqualFold(strings.ToLower(destAddress), strings.ToLower(cfg.Addr)) {
				filteredResults = append(filteredResults, node)
			}
		}

		// Update the cursor for the next iteration
		pageInfo, ok := xmsgs["PageInfo"].(map[string]any)
		if !ok {
			return errors.New("type map[string]any expected for PageInfo")
		}
		hasNextPage, ok := pageInfo["HasNextPage"].(bool)
		if !ok {
			return errors.New("type bool expected for HasNextPage")
		}
		if !hasNextPage {
			break
		}
		nextCursor, ok := pageInfo["NextCursor"].(string)
		if !ok {
			return errors.New("type string expected for NextCursor")
		}
		variables["cursor"] = nextCursor
	}

	// Limit the results to the most recent 10 messages
	if len(filteredResults) > 10 {
		filteredResults = filteredResults[len(filteredResults)-10:]
	}

	jsonResults, err := json.MarshalIndent(filteredResults, "", "  ")
	if err != nil {
		return errors.Wrap(err, "failed to marshal filtered results")
	}
	fmt.Println(string(jsonResults))

	return nil
}

func getXmsgCount(ctx context.Context, client *graphql.Client) (int64, error) {
	var xmsgCountQuery = `
		query XMsgCount {
			xmsgcount
		}
	`
	xmsgCountData, err := executeQuery(ctx, client, xmsgCountQuery, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to execute xmsgcount query")
	}

	xmsgcountInterface, ok := xmsgCountData["xmsgcount"].(string)
	if !ok {
		return 0, errors.Wrap(err, "xmsgcount is not of type string")
	}

	xmsgcount, err := strconv.ParseInt(xmsgcountInterface[2:], 16, 64)
	if err != nil {
		return 0, errors.Wrap(err, "failed to convert xmsgcount to integer")
	}

	return xmsgcount, nil
}

func executeQuery(ctx context.Context, client *graphql.Client, query string, variables map[string]any) (map[string]any, error) {
	req := graphql.NewRequest(query)
	// set variables
	for key, value := range variables {
		req.Var(key, value)
	}

	// run it and capture the response
	var respData map[string]any
	if err := client.Run(ctx, req, &respData); err != nil {
		return nil, errors.Wrap(err, "execute query")
	}

	return respData, nil
}
