package push

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/nghiango1/deploy/webhook-handler/pkg/logger"
)

const DEFAULT_WORK_DIR = "/tmp/.work"

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func MainRefHandler(event Event) error {
	// Pull the latest code
	isExist, err := exists(DEFAULT_WORK_DIR)
	if err != nil {
		logger.Get().Error(fmt.Sprintln("Can't access to workdir", DEFAULT_WORK_DIR, ",got: ", err.Error()))
		return err
	}
	if isExist {
		if err := os.RemoveAll(DEFAULT_WORK_DIR); err != nil {
			logger.Get().Error(fmt.Sprintln("Can't cleanup the workdir", DEFAULT_WORK_DIR, ",got: ", err.Error()))
			return err
		}
	}

	cmd := exec.Command("git", "clone", "--depth=1", "https://github.com/nghiango1/deploy.git", DEFAULT_WORK_DIR)
	logger.Get().Debug(fmt.Sprint(cmd.String()))
	err = cmd.Run()
	if err != nil {
		logger.Get().Error(fmt.Sprintln("Can't pull latest code to workdir", DEFAULT_WORK_DIR, ",got: ", err.Error()))
		return err
	}

	logger.Get().Debug(fmt.Sprintln("Done pull latest code"))
	return nil
}

func DevRefHandler(event Event) error {
	return MainRefHandler(event)
}
