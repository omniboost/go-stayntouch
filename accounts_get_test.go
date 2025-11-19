package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountsGet(t *testing.T) {
	req := client.NewAccountsGet()
	req.PathParams().ID = "49469"
	req.QueryParams().HotelID = hotelID
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

// func TestAccountsGetAll(t *testing.T) {
// 	req := client.NewAccountsGet()
// 	req.QueryParams().HotelID = hotelID
// 	resp, err := req.All()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	b, _ := json.MarshalIndent(resp, "", "  ")
// 	fmt.Println(string(b))
// }
