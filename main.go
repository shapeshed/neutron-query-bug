package main

import (
	"context"
	sdkmath "cosmossdk.io/math"
	"fmt"
	dextypes "github.com/neutron-org/neutron/v4/x/dex/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient(
		"neutron-grpc.polkachu.com:19190",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}

	dexQueryClient := dextypes.NewQueryClient(conn)

	// Properly create an sdkmath.Int
	amountIn := sdkmath.NewInt(1)

	// Create a request with the custom Int type
	query := dextypes.QueryEstimatePlaceLimitOrderRequest{
		Creator:          "neutron1nz852flh6np9xlg9ju3ka6w5txezsxt0j4lypn",
		Receiver:         "neutron1nz852flh6np9xlg9ju3ka6w5txezsxt0j4lypn",
		TokenIn:          "untrn",
		TokenOut:         "ibc/B559A80D62249C8AA07A380E2A2BEA6E5CA9A6F079C912C3A9E9B494105E4F81",
		TickIndexInToOut: int64(0),
		AmountIn:         amountIn, // Correctly initialized sdkmath.Int
		OrderType:        dextypes.LimitOrderType_FILL_OR_KILL,
	}

	// Make the request
	res, err := dexQueryClient.EstimatePlaceLimitOrder(context.Background(), &query)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Response: %+v\n", res)
	}
}
