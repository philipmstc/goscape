package feature

import (
	"fmt"
	"time"
	"philipmstc/goscape/model"
)

func IsAnimated() bool {
	return false 
}

func Load() {
	// eventually, loads feature configurations from CLI or DB 
}

func IsNewGame() bool {
	return true
}

func ShowAnimation(action model.Action) {
	fmt.Print("[     ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#    ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[##   ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[###  ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#### ]")
	time.Sleep(1 * time.Second)
	fmt.Print("\r[#####]")
	fmt.Println()
}