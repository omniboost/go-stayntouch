package stayntouch

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) MarshalSchema() string {
	return d.Time.Format("2006-01-02")
}

type DateTime struct {
	time.Time
}

func (d DateTime) MarshalSchema() string {
	return d.Time.Format(time.RFC3339)
}

type StringFloat float64

func (f *StringFloat) UnmarshalJSON(text []byte) (err error) {
	var flt float64
	err = json.Unmarshal(text, &flt)
	if err == nil {
		*f = StringFloat(flt)
		return err
	}

	// error, so try string
	var s string
	err = json.Unmarshal(text, &s)
	if err != nil {
		return err
	}

	flt, err = strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*f = StringFloat(flt)
	return nil
}
