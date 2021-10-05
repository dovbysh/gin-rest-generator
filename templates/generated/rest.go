package generated

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	// @Gorg success_cb fmt {"yaml":{"package":"github.com/dovbysh/gin-rest-generator/templates/generated", "cb":"XlsxCb"}, "func":{"cb":"func(c *gin.Context, rows []Row){ c.YAML(http.StatusOK, rows); }"}}
	GetRowsFromSomething(ctx context.Context, dt time.Time, b bool, i int, s string) ([]Row, error)

	// GetRowsFromSomethingNoSuccessCb documentation
	// @Tags the_tag
	// @Param X-Request-Id header string false "Request-Id"
	// @Router /some-report-no-success-cb/rows [GET]
	// @Gorg param dt tags binding:"required"
	// @Gorg param dt comment Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
	GetRowsFromSomethingNoSuccessCb(ctx context.Context, dt time.Time, b bool, i int, s string) ([]Row, error)

	// GetRowsFromSomethingContextOnly documentation
	// @Param authorization header string false "Authorization: Bearer"
	// @Tags the_tag
	// @Param X-Request-Id header string false "Request-Id"
	// @Router /some-report-context-only/rows [GET]
	GetRowsFromSomethingContextOnly(ctx context.Context) ([]Row, error)

	// GetRowsFromSomethingContextOnlyWithSuccessCb documentation
	// @Param authorization header string false "Authorization: Bearer"
	// @Tags the_tag
	// @Param X-Request-Id header string false "Request-Id"
	// @Router /some-report-context-only-with-success-cb/rows [GET]
	// @Gorg success_cb fmt {"func":{"cb":"func(c *gin.Context, rows []Row){ c.YAML(http.StatusOK, rows); }"}}
	GetRowsFromSomethingContextOnlyWithSuccessCb(ctx context.Context) ([]Row, error)
}

func XlsxCb(c *gin.Context, rows []Row) {
	c.YAML(http.StatusOK, rows)
}
