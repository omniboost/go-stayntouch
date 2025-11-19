package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestHotelChargeCodesGet(t *testing.T) {
	req := client.NewHotelChargeCodesGet()
	req.PathParams().HotelID = hotelID
	// req.QueryParams().LanguageCode = "EN"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestHotelChargeCodesGetAll(t *testing.T) {
	req := client.NewHotelChargeCodesGet()
	req.PathParams().HotelID = hotelID
	// req.QueryParams().LanguageCode = "EN"
	resp, err := req.All()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
