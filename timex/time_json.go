package timex

import (
	"fmt"
	"time"
)

func (t *Time) UnmarshalJSON(bytes []byte) error {
	v, err := time.Parse(DateTimeZone, string(bytes))
	if err != nil {
		return err
	}

	t.Time = v
	return nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format(DateTimeZone))), nil
}
