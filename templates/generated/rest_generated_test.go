package generated

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

type ItoGenerateImplementation struct {
	t *testing.T
}

func (z *ItoGenerateImplementation) GetRowsFromSomethingNoSuccessCb(ctx context.Context, dt time.Time, b bool, i int, s string) ([]Row, error) {
	panic("implement me")
}

func (z *ItoGenerateImplementation) GetRowsFromSomethingContextOnly(ctx context.Context) ([]Row, error) {
	panic("implement me")
}

func (z *ItoGenerateImplementation) GetRowsFromSomethingContextOnlyWithSuccessCb(ctx context.Context) ([]Row, error) {
	panic("implement me")
}

func (z *ItoGenerateImplementation) GetRowsFromSomething(ctx context.Context, dt time.Time, b bool, i int, s string) ([]Row, error) {
	z.t.Log("GetRowsFromSomething executed", ctx, dt, b, i, s)
	res := []Row{
		{
			Id: i,
			F:  0,
			S:  s,
			Dt: dt,
			B:  b,
		},
	}
	return res, nil
}

func TestItoGenerateRest_handlerGetRowsFromSomething(t *testing.T) {
	g := gin.Default()
	r := g.Group("/")
	NewItoGenerateRest(r, &ItoGenerateImplementation{t: t})
	dt := time.Now()
	v := url.Values{}
	v.Add("dt", dt.Format(time.RFC3339Nano))
	uri := UrlGetRowsFromSomething + "?" + v.Encode()
	req := httptest.NewRequest("GET", uri, nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	var (
		resRows []Row
		err     error
		body    []byte
	)
	body = w.Body.Bytes()
	t.Log(string(body))
	err = json.Unmarshal(body, &resRows)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, 1, len(resRows)) {
		return
	}
	assert.True(t, dt.Equal(resRows[0].Dt))

	var resRows2 []Row
	v.Add("fmt", "func")
	uri = UrlGetRowsFromSomething + "?" + v.Encode()
	req = httptest.NewRequest("GET", uri, nil)
	w = httptest.NewRecorder()
	g.ServeHTTP(w, req)
	body = w.Body.Bytes()
	t.Log(string(body))
	err = yaml.Unmarshal(body, &resRows2)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.Equal(t, 1, len(resRows2)) {
		return
	}
	assert.True(t, dt.Equal(resRows2[0].Dt))
}
