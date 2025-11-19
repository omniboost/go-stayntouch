package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	stayntouch "github.com/omniboost/go-stayntouch"
)

func TestDirectBillsGet(t *testing.T) {
	req := client.NewDirectBillsGet()
	req.QueryParams().HotelID = hotelID
	req.QueryParams().Date = stayntouch.Date{time.Date(2024, 10, 8, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
