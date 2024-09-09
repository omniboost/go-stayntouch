package stayntouch_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestRoomTypeGet(t *testing.T) {
	req := client.NewRoomTypeGet()
	req.PathParams().ID = 434
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}




