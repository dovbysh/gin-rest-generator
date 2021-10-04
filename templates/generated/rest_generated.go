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
	UrlGetRowsFromSomething = "/some-report/rows"
)

func (r *ItoGenerateRest) routes(router *gin.RouterGroup) {
	router.GET(UrlGetRowsFromSomething, r.handlerGetRowsFromSomething)

}

type RequestGetRowsFromSomething struct {
	Dt time.Time `json:"dt" form:"dt" binding:"required"` // Datetime (RFC3339), ex. 2021-03-01T00:00:00Z
	B  bool      `json:"b" form:"b"`
	I  int       `json:"i" form:"i"`
	S  string    `json:"s" form:"s"`
}

// GetRowsFromSomething documentation
// @Param authorization header string false "Authorization: Bearer"
// @Tags the_tag
// @Param X-Request-Id header string false "Request-Id"
// @Router /some-report/rows [GET]
// @Gorg param dt tags binding:"required"
// @Gorg param dt comment Datetime (RFC3339), ex. 2021-03-01T00:00:00Z
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
	c.JSON(http.StatusOK, res)

}
