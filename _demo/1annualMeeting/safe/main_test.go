// author: ashing
// time: 2020/7/10 11:41 下午
// mail: axingfly@gmail.com
// Less is more.

package main

import (
	"fmt"
	"github.com/kataras/iris/httptest"
	"sync"
	"testing"
)

func TestLottery(t *testing.T) {
	e := httptest.New(t, newApp())

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 0\n")
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			e.POST("/import").
				WithFormField(
					"users",
					fmt.Sprintf("test_u%d", i)).
				Expect().Status(httptest.StatusOK)
		}(i)
	}

	wg.Wait()

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 100\n")

	e.GET("/lucky").Expect().Status(httptest.StatusOK).Body()

	e.GET("/").Expect().Status(httptest.StatusOK).
		Body().Equal("当前总共参与抽奖的用户数: 99\n")

	// 加锁后 userList 操作线程安全

}
