package queue

import (
	"fmt"
	"github.com/duguying/simpleci/model"
	"github.com/duguying/simpleci/runner"
	"time"
)

func startQueue() {
	for {
		time.Sleep(time.Second)
		pending, err := model.GetBuildPending()
		if err != nil {
			topBuilding, err := model.GetBuildQueueTop()
			if err != nil {
			} else {
				fmt.Printf("start build [%d]\n", topBuilding.Id)
				model.UpdateBuild(topBuilding.Id, model.BUILD_RESULT_PENDING, "", "")
				runner.StartCi(topBuilding)
			}
		} else {
			fmt.Println(pending)
		}
	}
}

func StartQueue() {
	go startQueue()
}
