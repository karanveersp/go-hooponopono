package main

import (
	"fmt"
	"time"

	"github.com/MasterDimmy/go-cls"
)

func main() {
	phrases := [4]string{"I love you.", "I'm sorry.", "Please forgive me.", "Thank you."}

	i := 0
	for {
		fmt.Println(phrases[i])
		i++
		time.Sleep(time.Second * 2)
		if i == 4 {
			// reset
			i = 0
			cls.CLS()
		}
	}

}
