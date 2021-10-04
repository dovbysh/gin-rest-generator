package generated

import "time"

type Row struct {
	Id int       `json:"id"`
	F  float64   `json:"f"`
	S  string    `json:"s"`
	Dt time.Time `json:"dt"`
	B  bool      `json:"b"`
}
