package gorg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NoParamsRequired(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "// @gorg NoParamsRequired", want: true},
		{s: "// @gorg noparamsrequired", want: true},
		{s: "  @gorg   noparamsrequired  ", want: true},
		{s: "@gorg   noparamsrequired", want: true},
		{s: "@gorg ", want: false},
		{s: "noparamsrequired ", want: false},
		{s: "gorg   noparamsrequired ", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			g := Gorg{}
			g.ParseComment(tt.s)
			if got := g.NoParamsRequired; got != tt.want {
				t.Errorf("NoParamsRequired = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ParamsTags(t *testing.T) {
	tests := []struct {
		s    string
		want map[string]Param
	}{
		{s: "\t// @Gorg param from tags binding:\"required\" example:\"2020-01-01T00:00:00Z\"", want: map[string]Param{
			"from": {Name: "from", Tags: "binding:\"required\" example:\"2020-01-01T00:00:00Z\""},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			g := Gorg{}
			g.ParseComment(tt.s)
			assert.Equal(t, tt.want, g.Params)
		})
	}
}
