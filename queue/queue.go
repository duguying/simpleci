package queue

import (
	"fmt"
	"github.com/duguying/simpleci/model"
	"time"
)

func startQueue() {
	for {
		time.Sleep(time.Second)
		pending, err := model.GetBuildPending()
		if err != nil {
			fmt.Println("not has pending")
			topBuilding, err := model.GetBuildQueueTop()
			if err != nil {
				fmt.Println("not has top one")
			} else {
				fmt.Println("check!!!!")
				fmt.Println(topBuilding)
			}
		} else {
			fmt.Println("in pending")
			fmt.Println(pending)
		}
	}
}

func StartQueue() {
	go startQueue()
}
