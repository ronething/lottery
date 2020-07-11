// author: ashing
// time: 2020/7/11 12:46 上午
// mail: axingfly@gmail.com
// Less is more.

/**
1、即开既得型
2、双色球自选型
*/
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type lotterController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// GET http://localhost:8080
func (c *lotterController) Get() string {
	c.Ctx.Header("Context-Type", "text/html")
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(10)
	var prize string
	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >= 4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1获得一等奖<br/>"+
			"尾号为2或3获得二等奖<br/>"+
			"尾号为4/5/6获得三等奖<br/>"+
			"code=%d<br/>"+
			"很遗憾，没有获奖", code)
	}
	return fmt.Sprintf("尾号为1获得一等奖<br/>"+
		"尾号为2或3获得二等奖<br/>"+
		"尾号为4/5/6获得三等奖<br/>"+
		"code=%d<br/>"+
		"恭喜你获得: %s", code, prize)

}

// GET http://localhost:8080/prize
func (c *lotterController) GetPrize() string {
	c.Ctx.Header("Context-type", "text/html")
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var prize [7]int32
	// 红色球 1-33
	for i := 0; i < 6; i++ {
		prize[i] = r.Int31n(33) + 1 // + 1
	}
	// 最后一位的蓝色球，1-16
	prize[6] = r.Int31n(16) + 1
	return fmt.Sprintf("今日开奖号码是: %v", prize)

}
