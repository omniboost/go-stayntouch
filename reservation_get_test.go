package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestReservationGet(t *testing.T) {
	req := client.NewReservationGet()
	req.PathParams().ID = 126789824
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
