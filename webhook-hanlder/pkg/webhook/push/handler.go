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
		os.RemoveAll(DEFAULT_WORK_DIR)
	}

	if err := os.Mkdir(DEFAULT_WORK_DIR, os.ModeDir); err != nil {
		logger.Get().Error(fmt.Sprintln("Can't create to workdir", DEFAULT_WORK_DIR, ",got: ", err.Error()))
		return err
	}

	exec.Command("git", "clone", "--depth=1", "github.com/nghiango1/deploy", DEFAULT_WORK_DIR)
	return nil
}

func DevRefHandler(event Event) error {
	return MainRefHandler(event)
}
