package vo

import "time"

type Time struct {
	value time.Time
}

func NewTime() Time {
	return Time{value: time.Now().UTC()}
}

func NewTimeFromTime(t time.Time) Time {
	return Time{value: t.UTC()}
}

func (t Time) Value() time.Time {
	return t.value
}

func (t Time) String() string {
	return t.value.Format(time.RFC3339)
}
