// author: ashing
// time: 2020/7/11 12:27 下午
// mail: axingfly@gmail.com
// Less is more.

// 微信摇一摇

package main

import (
	"fmt"
	"log"
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

func main() {

}

func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		id:      1,
		name:    "iphone4",
		pic:     "",
		link:    "",
		gtype:   giftTypeRealLarge,
		data:    "",
		total:   1000,
		left:    1000,
		inuse:   true,
		rate:    10000,
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
		total:   5,
		left:    5,
		rate:    5000,
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
		total:    50,
		left:     50,
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

	fmt.Printf("giftList=%v\n", giftList)
}

func main() {

}
