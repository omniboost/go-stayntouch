package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRateGet(t *testing.T) {
	req := client.NewRateGet()
	req.PathParams().ID = 20450
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}





