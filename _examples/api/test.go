package api

import (
	"github.com/gin-gonic/gin"

	"github.com/longhaoteng/wineglass/_examples/errors"
	"github.com/longhaoteng/wineglass/api"
	log "github.com/longhaoteng/wineglass/logger"
	"github.com/longhaoteng/wineglass/server"
)

type Test struct {
	*api.API
}

type TestResp struct {
	F1 string `json:"f1"`
	F2 int64  `json:"f2"`
	F3 bool   `json:"f3"`
}

type TestReqPath struct {
	Name string `json:"name" uri:"name" binding:"required,max=12"`
}

type TestReq struct {
	F1 string `json:"f1" binding:"required,max=32"`
	F2 string `json:"f2" binding:"omitempty,url"`
	F3 int64  `json:"f3" binding:"oneof=0 1 2"`
}

func (t *Test) Router(r *gin.Engine) {
	r.GET("/test", t.Test1)
	r.PUT("/test/:name", t.Test2)
}

func (t *Test) Test1(c *gin.Context) {
	t.Handler(c, func(resp *api.Response) error {
		resp.Data = &TestResp{
			F1: "test",
			F2: 100,
			F3: true,
		}
		// t.Data(c, data)
		return nil
	})
}

func (t *Test) Test2(c *gin.Context) {
	testReqPath := &TestReqPath{}
	testReq := &TestReq{}
	if t.Verifies(
		c,
		testReqPath, c.ShouldBindUri,
		testReq, c.ShouldBind,
	) {
		t.Handler(c, func(resp *api.Response) error {
			if testReqPath.Name != "wineglass" {
				return errors.TestErr
			}

			log.Fields(
				"f1", testReq.F1,
				"f2", testReq.F2,
				"f3", testReq.F3,
			).Log(log.InfoLevel)

			return nil
		})
	}
}

func init() {
	server.AddRouters(&Test{})
}
