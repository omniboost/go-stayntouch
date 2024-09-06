package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestReservationsGet(t *testing.T) {
	req := client.NewReservationsGet()
	req.QueryParams().HotelID = client.HotelID()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

