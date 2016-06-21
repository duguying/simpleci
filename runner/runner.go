package runner

import (
	"fmt"
	"github.com/duguying/simpleci/global"
	"github.com/duguying/simpleci/model"
	"github.com/gojudge/proc"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func AddRunTask(projectId string, info map[string]interface{}) map[string]interface{} {
	id, err := strconv.Atoi(projectId)
	if err != nil {
		return map[string]interface{}{
			"result": false,
			"error":  err.Error(),
		}
	}

	project := &model.Project{Id: int64(id)}
	has, err := global.Eg.Get(project)
	if err != nil {
		return map[string]interface{}{
			"result": false,
			"error":  err.Error(),
		}
	}
	if has {
		return map[string]interface{}{
			"result": false,
			"error":  "project not exist",
		}
	}

	bid, err := model.AddBuild(model.BUILD_RESULT_PENDING, info["commit_id"].(string), "", "")

	if err != nil {
		return map[string]interface{}{
			"result": false,
			"error":  err.Error(),
		}
	}

	return map[string]interface{}{
		"result":   true,
		"build_id": bid,
	}
}

func StartCi(build *model.Build) {
	cwd, err := os.Getwd()
	if err != nil {
		return
	}

	projectId := build.ProjectId
	project, err := model.GetProject(projectId)

	workDir := project.WorkDir
	os.Chdir(workDir)

	log1, err := run("git pull")
	if err != nil {
		// update
		model.UpdateBuild(build.Id, model.BUILD_RESULT_TIMEOUT, log1, "")
		return
	} else {
		log2, err := run("./simpleci.sh")
		if err != nil {
			model.UpdateBuild(build.Id, model.BUILD_RESULT_TIMEOUT, log1, log2)
		} else {
			model.UpdateBuild(build.Id, model.BUILD_RESULT_FINE, log1, log2)
		}
	}

	os.Chdir(cwd)
}

func run(command string) (string, error) {
	cmd := exec.Command("/bin/sh", "-c", command)
	log := []byte("")
	var err error
	start := time.Now()
	over := false

	go func() {
		log, err = cmd.Output()
		over = true
	}()

	for {
		if time.Now().Nanosecond()-start.Nanosecond() > 60000 { // 1min
			pid := cmd.Process.Pid
			p := proc.GetProc(int64(pid))
			p.KillProcChainReverse()
			return string(log), fmt.Errorf("%s", "out of time")
		} else if over {
			return string(log), nil
		}
	}
}
