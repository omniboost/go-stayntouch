package stayntouch_test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestRoomTypeGet(t *testing.T) {
	req := client.NewRoomTypeGet()
	req.PathParams().ID = 434
	resp, err := req.Do(context.Background())
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
