package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestInvoiceGet(t *testing.T) {
	req := client.NewInvoiceGet()
	req.PathParams().ID = "56706959"
	req.QueryParams().HotelID = hotelID
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}
	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
