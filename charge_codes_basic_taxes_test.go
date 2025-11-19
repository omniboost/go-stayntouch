package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestChargeCodesBasicTaxesGet(t *testing.T) {
	req := client.NewChargeCodesBasicTaxesGet()
	req.PathParams().ID = 11918
	req.QueryParams().ID = "11918"
	req.QueryParams().HotelID = hotelID
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
