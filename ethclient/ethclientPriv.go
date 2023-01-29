package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

func (ec *Client) LogsPriv(ctx context.Context, groupId string, q ethereum.FilterQuery) ([]types.Log, error) {
	var result []types.Log
	arg, err := toFilterArg(q)
	if err != nil {
		return nil, err
	}
	err = ec.c.CallContext(ctx, &result, "priv_getLogs", groupId, arg)
	return result, err
}

func (ec *Client) FilterLogsPriv(ctx context.Context, groupId string, filterId string) ([]types.Log, error) {
	var result []types.Log
	err := ec.c.CallContext(ctx, &result, "priv_getFilterLogs", groupId, filterId)
	return result, err
}
