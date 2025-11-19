package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/omniboost/go-stayntouch"
)

func TestBillsRevenueGet(t *testing.T) {
	req := client.NewBillsRevenueGet()
	req.QueryParams().HotelID = fmt.Sprintf("%d", hotelID)
	req.QueryParams().Date = stayntouch.Date{Time: time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().Ledger = "guest"
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}

func TestBillsRevenueGetAll(t *testing.T) {
	req := client.NewBillsRevenueGet()
	req.QueryParams().HotelID = fmt.Sprintf("%d", hotelID)
	req.QueryParams().Date = stayntouch.Date{Time: time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC)}
	req.QueryParams().Ledger = "guest"
	resp, err := req.All(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
