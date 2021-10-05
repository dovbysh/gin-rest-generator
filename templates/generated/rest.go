package generated

import (
	"context"
	"time"
)

//go:generate gorg gen -g -p github.com/dovbysh/gin-rest-generator/templates/generated -i ItoGenerate -t ../rest -o rest_generated.go

type ItoGenerate interface {

	// GetRowsFromSomething documentation
	// @Param authorization header string false "Authorization: Bearer"
	// @Tags the_tag
	// @Param X-Request-Id header string false "Request-Id"
	// @Router /some-report/rows [GET]
	// @Gorg param dt tags binding:"required"
	// @Gorg param dt comment Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
	GetRowsFromSomething(ctx context.Context, dt time.Time, b bool, i int, s string) ([]Row, error)
}
