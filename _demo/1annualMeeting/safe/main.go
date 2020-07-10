// author: ashing
// time: 2020/7/11 12:16 上午
// mail: axingfly@gmail.com
// Less is more.

// 加锁实现 userList 添加线程安全

package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var userList []string
var mu sync.Mutex

// 抽奖控制器
type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	userList = make([]string, 0)
	mu = sync.Mutex{}

	app.Run(iris.Addr(":8080"))

}

// GET http://localhost:8080/
func (c *lotteryController) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d\n", count)
}

// POST http://localhost:8080/import
func (c *lotteryController) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	mu.Lock()
	defer mu.Unlock()
	count1 := len(userList)
	for _, u := range users {
		u = strings.TrimSpace(u) // 去除空格
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	count2 := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数: %d，成功导入用户数: %d\n", count2, count2-count1)
}

func (c *lotteryController) GetLucky() string {
	mu.Lock()
	defer mu.Unlock()
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()                                // rand 内部运算的随机数
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count)) // rand 计算得到的随机数
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...) // 移除这个用户
		return fmt.Sprintf("当前中奖用户: %s，剩余用户数: %d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		userList = userList[0:0] // []
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("已经没有参与用户，请先通过 /import 导入用户\n")
	}
}
