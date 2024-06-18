import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	liquidity "cloud.google.com/go/liquidity/apiv1"
	"cloud.google.com/go/liquidity/apiv1/liquiditypb"
)

// addLiquidity adds the given amount of liquidity to the given commitment.
func addLiquidity(w io.Writer, projectID, region, commitmentID string, amount int64) error {
	// projectID := "my-project-id"
	// region := "us-central1"
	// commitmentID := "my-commitment"
	// amount := 1000000000
	ctx := context.Background()
	client, err := liquidity.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("liquidity.NewClient: %v", err)
	}
	defer client.Close()

	req := &liquiditypb.AddLiquidityRequest{
		Name:    fmt.Sprintf("projects/%s/locations/%s/liquidityCommitments/%s", projectID, region, commitmentID),
		Amount:  amount,
		Request: nil,
	}

	resp, err := client.AddLiquidity(ctx, req)
	if err != nil {
		return fmt.Errorf("client.AddLiquidity: %v", err)
	}

	b, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return fmt.Errorf("json.MarshalIndent: %v", err)
	}

	fmt.Fprintf(w, "Added liquidity: %s", string(b))
	return nil
}
  
