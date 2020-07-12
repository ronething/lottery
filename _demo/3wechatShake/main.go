// author: ashing
// time: 2020/7/11 12:27 下午
// mail: axingfly@gmail.com
// Less is more.

// 微信摇一摇

package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 优惠卷，不相同的编码
	giftTypeCouponFix        // 优惠卷，相同的编码
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

const rateMax = 10000

type gift struct {
	id       int
	name     string
	pic      string
	link     string
	gtype    int
	data     string
	dataList []string
	total    int
	left     int
	inuse    bool
	rate     int
	rateMin  int
	rateMax  int
}

var logger *log.Logger

var giftList []*gift

// initGift 初始化奖品的 rate 参数
func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:      1,
		name:    "iphone4",
		pic:     "",
		link:    "",
		gtype:   giftTypeRealLarge,
		data:    "",
		total:   2,
		left:    2,
		inuse:   true,
		rate:    2,
		rateMin: 0,
		rateMax: 0,
	}

	giftList[0] = &g1
	g2 := gift{
		id:      2,
		name:    "apple",
		pic:     "",
		link:    "",
		gtype:   giftTypeRealSmall,
		data:    "",
		total:   5,
		left:    5,
		inuse:   true,
		rate:    100,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[1] = &g2

	// 3 虚拟券，相同的编码
	g3 := gift{
		id:      3,
		name:    "商城满2000元减50元优惠券",
		pic:     "",
		link:    "",
		gtype:   giftTypeCouponFix,
		data:    "mall-coupon-2018",
		total:   50,
		left:    50,
		rate:    2000,
		inuse:   true,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[2] = &g3
	// 4 虚拟券，不相同的编码
	g4 := gift{
		id:       4,
		name:     "商城无门槛直降50元优惠券",
		pic:      "",
		link:     "",
		gtype:    giftTypeCoupon,
		data:     "",
		dataList: []string{"c01", "c02", "c03", "c04", "c05"},
		total:    5,
		left:     5,
		inuse:    true,
		rate:     2000,
		rateMin:  0,
		rateMax:  0,
	}

	giftList[3] = &g4
	// 5 虚拟币
	g5 := gift{
		id:      5,
		name:    "社区10个金币",
		pic:     "",
		link:    "",
		gtype:   giftTypeCoin,
		data:    "10",
		total:   5,
		left:    5,
		inuse:   true,
		rate:    5000,
		rateMin: 0,
		rateMax: 0,
	}
	giftList[4] = &g5

	rateStart := 0
	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = data.rateMin + data.rate
		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}

	for _, i := range giftList {
		fmt.Printf("%v\n", i)
	}
}

// 初始化日志存取目录
func initLog() {
	f, _ := os.Create("/tmp/log/lottery_wechat.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	initGift()
	return app
}

func (l *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		// data.total == 0 表示不限数量
		if data.inuse && (data.total == 0 || (data.total > 0 && data.left > 0)) {
			count++
			total += data.left
		}
	}

	return fmt.Sprintf("当前有效奖品种类数量: %d,限量奖品总数量: %d\n", count, total)
}

// GET http://localhost:8080/lucky
func (l *lotteryController) GetLucky() map[string]interface{} {
	code := luckyCode()
	result := make(map[string]interface{})
	result["success"] = false
	for _, data := range giftList {
		if !data.inuse || (data.total > 0 && data.left <= 0) {
			continue
		}
		if data.rateMin <= int(code) && int(code) <= data.rateMax {
			sendData, ok := "", false
			switch data.gtype {
			case giftTypeCoin:
				ok, sendData = sendCoin(data)
			case giftTypeCoupon:
				ok, sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok, sendData = sendCouponFix(data)
			case giftTypeRealSmall:
				ok, sendData = sendRealSmall(data)
			case giftTypeRealLarge:
				ok, sendData = sendRealLarge(data)
			}
			if ok {
				saveLuckyData(code, data.id, data.name, data.link, sendData, data.left)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = data.data
				break // NOTICE: 有一个被抽到就直接 break 了，所以不会有抽到多个的情况
			}
		}
	}

	return result
}

// saveLuckyData 记录用户的获奖记录
func saveLuckyData(code int32, id int, name string, link string, sendData string, left int) {
	logger.Printf("lucky, code=%d, gift=%d, name=%s, link=%s, data=%s, left=%d ", code, id, name, link, sendData, left)
}

func sendRealLarge(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left--
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

func sendRealSmall(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

func sendCouponFix(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true, data.data
	} else if data.left > 0 {
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

func sendCoupon(data *gift) (bool, string) {
	if data.left > 0 {
		// 还有剩余的奖品
		left := data.left - 1
		data.left = left
		return true, data.dataList[left] // TODO: 理解一下
	} else {
		return false, "奖品已发完"
	}
}

// 虚拟币
func sendCoin(data *gift) (bool, string) {
	if data.total == 0 {
		return true, data.data
	} else if data.left > 0 {
		data.left = data.left - 1
		return true, data.data
	} else {
		return false, "奖品已发完"
	}
}

func luckyCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}
