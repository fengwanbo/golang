package golang

import (
	"fmt"
	"math/rand"
	"time"
)


const VERSION  = 2.2


func Aaa()  {
	rand.Seed(time.Now().Unix())
	nums :=rand.Intn(100)
	fmt.Println(nums)
}
