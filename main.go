package main

import (
	"fmt"
	"github.com/xlizy/common-go/models"
	"math/rand/v2"
)

var Cnt = 0

type Test struct {
	models.PrimaryKey
	Key string `json:"key" gorm:"column:key_value"`
	models.ControlBy
	models.ModelTime
}

func main() {
	gg := 0
	mm := 0
	t1 := 0
	t2 := 0
	t3 := 0
	for i := 0; i < 1000000; i++ {
		t1 = t1 + 1
		r1 := rand.IntN(2)
		if r1 == 0 {
			gg = gg + 1
		}
		if r1 == 1 {
			t2 = t2 + 1
			mm = mm + 1
			r2 := rand.IntN(2)
			if r2 == 0 {
				gg = gg + 1
			}
			if r2 == 1 {
				t3 = t3 + 1
				mm = mm + 1
				r3 := rand.IntN(2)
				if r3 == 0 {
					gg = gg + 1
				}
				if r3 == 0 {
					mm = mm + 1
				}
			}
		}
	}
	fmt.Println("gg:", gg)
	fmt.Println("mm:", mm)
	fmt.Println("t1:", t1)
	fmt.Println("t2:", t2)
	fmt.Println("t3:", t3)
}
