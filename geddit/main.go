package main

import (
	"fmt"
	"github.com/foysal-mamun/reddit"
	"log"
)

func main() {

	items, err := reddit.Get("yii")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}
