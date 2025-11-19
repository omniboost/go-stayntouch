package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestReservationsGet(t *testing.T) {
	req := client.NewReservationsGet()
	req.QueryParams().HotelID = hotelID
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
