package gorg

import (
	"strings"
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
		s    []string
		want Gorg
	}{
		{
			s: []string{"\t// @Gorg param from tags binding:\"required\" example:\"2020-01-01T00:00:00Z\""},
			want: Gorg{
				Params: map[string]Param{
					"from": {Name: "from", Tags: "binding:\"required\" example:\"2020-01-01T00:00:00Z\""},
				},
			},
		},
		{
			s: []string{
				"\t// @Gorg param from tags binding:\"required\" example:\"2020-01-01T00:00:00Z\"",
				"\t// @Gorg param from comment very important documentation",
			},
			want: Gorg{
				Params: map[string]Param{
					"from": {
						Name:    "from",
						Tags:    "binding:\"required\" example:\"2020-01-01T00:00:00Z\"",
						Comment: "very important documentation",
					},
				},
			},
		},
		{
			s: []string{
				"\t// @Gorg param from tags binding:\"required\" example:\"2020-01-01T00:00:00Z\"",
				"\t// @Gorg param from comment very important documentation",
				"// @Gorg pager limit offset 1000",
			},
			want: Gorg{
				Params: map[string]Param{
					"from": {
						Name:    "from",
						Tags:    "binding:\"required\" example:\"2020-01-01T00:00:00Z\"",
						Comment: "very important documentation",
					},
				},
				Pager: Pager{
					Exists:     true,
					LimitName:  "limit",
					OffsetName: "offset",
					MaxLimit:   1000,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(strings.Join(tt.s, "---"), func(t *testing.T) {
			g := Gorg{}
			for _, s := range tt.s {
				g.ParseComment(s)
			}
			assert.Equal(t, tt.want, g)
		})
	}
}
