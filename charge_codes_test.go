package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChargeCodesGet(t *testing.T) {
	req := client.NewChargeCodesGet()
	req.QueryParams().HotelID = hotelID
	// req.QueryParams().LanguageCode = "EN"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
