package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	stayntouch "github.com/omniboost/go-stayntouch"
)

func TestInvoicesPerDayGet(t *testing.T) {
	req := client.NewInvoicesPerDayGet()
	req.QueryParams().HotelID = hotelID
	req.QueryParams().Date = stayntouch.Date{time.Date(2025, 9, 9, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
