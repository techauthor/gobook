package rand

import (
	"fmt"
	"math/rand"
	"time"
)

func ExampleRandIntn() {
	now := time.Now().UnixNano()
	//rand.Seed(now)
	//rand.Intn(10)
	//如果不自行定义随机数种子的话，每次生成的随机数都是一样的。因为在默认因子为1
	//要解决这个问题，可以以时间作为随机因子
	r := rand.New(rand.NewSource(now))
	n := r.Intn(10)
	fmt.Println(n)

	// Output:
	// 2
}
