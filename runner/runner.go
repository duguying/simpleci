package runner

import (
	"fmt"
	"github.com/duguying/simpleci/global"
	"github.com/duguying/simpleci/model"
	"github.com/gojudge/proc"
	"os/exec"
	"strconv"
	"time"
)

func Run(projectId string) map[string]interface{} {
	id, err := strconv.Atoi(projectId)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}

	project := &model.Project{Id: int64(id)}
	has, err := global.Eg.Get(project)
	if err != nil {
		return map[string]interface{}{
			"error": err.Error(),
		}
	}
	if has {
		return map[string]interface{}{
			"error": "project not exist",
		}
	}

	if global.Lc {
		return map[string]interface{}{
			"error": "CI task locked, please wait",
		}
	}

	// gitUrl := project.GitUrl

	go startCi(project)

	return map[string]interface{}{
		"msg": "ci task executed",
	}
}

func startCi(project *model.Project) {
	global.Lc = true
	cwd := os.Getwd()
	workDir := project.WorkDir
	os.Chdir(workDir)

	log1, err := run("git pull")
	if err != nil {
		model.AddBuild(model.BUILD_RESULT_TIMEOUT, log1, "")
		return
	} else {
		log2, err := run("./simpleci.sh")
		if err != nil {
			model.AddBuild(model.BUILD_RESULT_TIMEOUT, log1, log2)
		} else {
			model.AddBuild(model.BUILD_RESULT_FINE, log1, log2)
		}
	}

	os.Chdir(cwd)
	global.Lc = false
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
