package generated

// DO NOT EDIT!
// This code is generated with http://github.com/dovbysh/gin-rest-generator tool
// using ../rest.go.tmpl template

import (
	"fmt"
	"net/http"
	"time"

	wrapErr "github.com/Chekunin/wraperr"
	"github.com/gin-gonic/gin"
)

type ItoGenerateRest struct {
	Usecase ItoGenerate
}

func NewItoGenerateRest(router *gin.RouterGroup, usecase ItoGenerate) *ItoGenerateRest {
	result := &ItoGenerateRest{Usecase: usecase}
	result.routes(router)
	return result
}

const (
	UrlGetRowsFromSomething                         = "/some-report/rows"
	UrlGetRowsFromSomethingContextOnly              = "/some-report-context-only/rows"
	UrlGetRowsFromSomethingContextOnlyWithSuccessCb = "/some-report-context-only-with-success-cb/rows"
	UrlGetRowsFromSomethingNoSuccessCb              = "/some-report-no-success-cb/rows"
)

func (r *ItoGenerateRest) routes(router *gin.RouterGroup) {
	router.GET(UrlGetRowsFromSomething, r.handlerGetRowsFromSomething)
	router.GET(UrlGetRowsFromSomethingContextOnly, r.handlerGetRowsFromSomethingContextOnly)
	router.GET(UrlGetRowsFromSomethingContextOnlyWithSuccessCb, r.handlerGetRowsFromSomethingContextOnlyWithSuccessCb)
	router.GET(UrlGetRowsFromSomethingNoSuccessCb, r.handlerGetRowsFromSomethingNoSuccessCb)

}

type RequestGetRowsFromSomething struct {
	Dt  time.Time `json:"dt" form:"dt" binding:"required"` // Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
	B   bool      `json:"b" form:"b"`
	I   int       `json:"i" form:"i"`
	S   string    `json:"s" form:"s"`
	Fmt string    `json:"fmt" form:"fmt"`
}

// GetRowsFromSomething documentation
// @Param authorization header string false "Authorization: Bearer"
// @Tags the_tag
// @Param X-Request-Id header string false "Request-Id"
// @Router /some-report/rows [GET]
// @Gorg param dt tags binding:"required"
// @Gorg param dt comment Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
// @Gorg success_cb fmt {"yaml":{"package":"github.com/dovbysh/gin-rest-generator/templates/generated", "cb":"XlsxCb"}, "func":{"cb":"func(c *gin.Context, rows []Row){ c.YAML(http.StatusOK, rows); }"}}
// @Summary // GetRowsFromSomething documentation
// @description
// @Accept json
// @Produce json
// @Param req query RequestGetRowsFromSomething false "req"
// @Success 200 {array} []Row
func (r *ItoGenerateRest) handlerGetRowsFromSomething(c *gin.Context) {
	var err error
	var res []Row

	var req RequestGetRowsFromSomething
	if err = c.ShouldBindQuery(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query GetRowsFromSomething"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err = r.Usecase.GetRowsFromSomething(
		c.Request.Context(),
		req.Dt,
		req.B,
		req.I,
		req.S,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(res) == 0 {
		res = []Row{}
	}

	switch req.Fmt {

	case "func":
		func(c *gin.Context, rows []Row) { c.YAML(http.StatusOK, rows) }(c, res)
		return

	case "yaml":
		XlsxCb(c, res)
		return

	}

	c.JSON(http.StatusOK, res)

}

type RequestGetRowsFromSomethingContextOnly struct {
}

// GetRowsFromSomethingContextOnly documentation
// @Param authorization header string false "Authorization: Bearer"
// @Tags the_tag
// @Param X-Request-Id header string false "Request-Id"
// @Router /some-report-context-only/rows [GET]
// @Summary // GetRowsFromSomethingContextOnly documentation
// @description
// @Accept json
// @Produce json
// @Param req query RequestGetRowsFromSomethingContextOnly false "req"
// @Success 200 {array} []Row
func (r *ItoGenerateRest) handlerGetRowsFromSomethingContextOnly(c *gin.Context) {
	var err error
	var res []Row

	var req RequestGetRowsFromSomethingContextOnly
	if err = c.ShouldBindQuery(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query GetRowsFromSomethingContextOnly"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err = r.Usecase.GetRowsFromSomethingContextOnly(
		c.Request.Context(),
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(res) == 0 {
		res = []Row{}
	}

	c.JSON(http.StatusOK, res)

}

type RequestGetRowsFromSomethingContextOnlyWithSuccessCb struct {
	Fmt string `json:"fmt" form:"fmt"`
}

// GetRowsFromSomethingContextOnlyWithSuccessCb documentation
// @Param authorization header string false "Authorization: Bearer"
// @Tags the_tag
// @Param X-Request-Id header string false "Request-Id"
// @Router /some-report-context-only-with-success-cb/rows [GET]
// @Gorg success_cb fmt {"func":{"cb":"func(c *gin.Context, rows []Row){ c.YAML(http.StatusOK, rows); }"}}
// @Summary // GetRowsFromSomethingContextOnlyWithSuccessCb documentation
// @description
// @Accept json
// @Produce json
// @Param req query RequestGetRowsFromSomethingContextOnlyWithSuccessCb false "req"
// @Success 200 {array} []Row
func (r *ItoGenerateRest) handlerGetRowsFromSomethingContextOnlyWithSuccessCb(c *gin.Context) {
	var err error
	var res []Row

	var req RequestGetRowsFromSomethingContextOnlyWithSuccessCb
	if err = c.ShouldBindQuery(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query GetRowsFromSomethingContextOnlyWithSuccessCb"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err = r.Usecase.GetRowsFromSomethingContextOnlyWithSuccessCb(
		c.Request.Context(),
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(res) == 0 {
		res = []Row{}
	}

	switch req.Fmt {

	case "func":
		func(c *gin.Context, rows []Row) { c.YAML(http.StatusOK, rows) }(c, res)
		return

	}

	c.JSON(http.StatusOK, res)

}

type RequestGetRowsFromSomethingNoSuccessCb struct {
	Dt time.Time `json:"dt" form:"dt" binding:"required"` // Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
	B  bool      `json:"b" form:"b"`
	I  int       `json:"i" form:"i"`
	S  string    `json:"s" form:"s"`
}

// GetRowsFromSomethingNoSuccessCb documentation
// @Tags the_tag
// @Param X-Request-Id header string false "Request-Id"
// @Router /some-report-no-success-cb/rows [GET]
// @Gorg param dt tags binding:"required"
// @Gorg param dt comment Datetime (RFC3339|RFC3339Nano), ex. 2021-03-01T00:00:00Z
// @Summary // GetRowsFromSomethingNoSuccessCb documentation
// @description
// @Accept json
// @Produce json
// @Param req query RequestGetRowsFromSomethingNoSuccessCb false "req"
// @Success 200 {array} []Row
func (r *ItoGenerateRest) handlerGetRowsFromSomethingNoSuccessCb(c *gin.Context) {
	var err error
	var res []Row

	var req RequestGetRowsFromSomethingNoSuccessCb
	if err = c.ShouldBindQuery(&req); err != nil {
		err = wrapErr.NewWrapErr(fmt.Errorf("binding data from query GetRowsFromSomethingNoSuccessCb"), err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err = r.Usecase.GetRowsFromSomethingNoSuccessCb(
		c.Request.Context(),
		req.Dt,
		req.B,
		req.I,
		req.S,
	)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	if len(res) == 0 {
		res = []Row{}
	}

	c.JSON(http.StatusOK, res)

}
