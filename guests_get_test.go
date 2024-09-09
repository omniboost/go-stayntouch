package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGuestsGet(t *testing.T) {
	req := client.NewGuestsGet()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}



