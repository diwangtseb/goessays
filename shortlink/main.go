package main

import (
	"sync/atomic"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", target)
	r.GET("/:hash", visit)
	r.Run()
}

var (
	counter int64
)

var tm = make(map[string]string)

func target(ctx *gin.Context) {
	target := ctx.PostForm("target")
	id := atomic.AddInt64(&counter, 1)
	id = id + time.Now().UnixNano()
	hashStr := encode(int(id))
	tm[hashStr] = target
	ctx.JSON(200, gin.H{
		"shortlink": "http://" + ctx.Request.Host + "/" + hashStr,
	})
}

func visit(ctx *gin.Context) {
	hash := ctx.Param("hash")
	if target, ok := tm[hash]; ok {
		ctx.Redirect(301, "http://"+ctx.Request.Host+"/"+target)
	}
	ctx.JSON(404, gin.H{})
}

var decTo62 = map[int]string{
	0:  "0",
	1:  "1",
	2:  "2",
	3:  "3",
	4:  "4",
	5:  "5",
	6:  "6",
	7:  "7",
	8:  "8",
	9:  "9",
	10: "a",
	11: "b",
	12: "c",
	13: "d",
	14: "e",
	15: "f",
	16: "g",
	17: "h",
	18: "i",
	19: "j",
	20: "k",
	21: "l",
	22: "m",
	23: "n",
	24: "o",
	25: "p",
	26: "q",
	27: "r",
	28: "s",
	29: "t",
	30: "u",
	31: "v",
	32: "w",
	33: "x",
	34: "y",
	35: "z",
	36: "A",
	37: "B",
	38: "C",
	39: "D",
	40: "E",
	41: "F",
	42: "G",
	43: "H",
	44: "I",
	45: "J",
	46: "K",
	47: "L",
	48: "M",
	49: "N",
	50: "O",
	51: "P",
	52: "Q",
	53: "R",
	54: "S",
	55: "T",
	56: "U",
	57: "V",
	58: "W",
	59: "X",
	60: "Y",
	61: "Z",
}

func encode(num int) string {
	if num == 0 {
		return "0"
	}
	var result string
	for num > 0 {
		result = decTo62[num%62] + result
		num = num / 62
	}
	return result
}

func decode(str string) int {
	var result int
	for i := 0; i < len(str); i++ {
		result = result*62 + int(str[i]-'0') // ascii code
	}
	return result
}
