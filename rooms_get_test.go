package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRoomsGet(t *testing.T) {
	req := client.NewRoomsGet()
	req.QueryParams().HotelID = hotelID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}



