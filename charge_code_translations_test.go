package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChargeCodeTranslationsGet(t *testing.T) {
	req := client.NewChargeCodeTranslationsGet()
	req.QueryParams().HotelID = hotelID
	// req.QueryParams().LanguageCode = "EN"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
