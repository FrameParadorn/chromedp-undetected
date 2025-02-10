//go:build windows

// Package chromedpundetected provides a chromedp context with an undetected
// Chrome browser.
package chromedpundetected

import (
	"errors"
	"os"
	"os/exec"

	"syscall"

	"github.com/chromedp/chromedp"
)

func headlessOpts() (opts []chromedp.ExecAllocatorOption, cleanup func() error, err error) {
	return nil, nil, errors.New("headless mode not supported in darwin")

	os.Setenv("ANGLE_DEFAULT_PLATFORM", "swiftshader")

	opt := chromedp.ModifyCmdFunc(func(cmd *exec.Cmd) {
		// Do nothing on AWS Lambda
		if _, ok := os.LookupEnv("LAMBDA_TASK_ROOT"); ok {
			return
		}

		// Set process attributes to create a new process group on Windows
		if cmd.SysProcAttr == nil {
			cmd.SysProcAttr = &syscall.SysProcAttr{
				CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
			}
		}

		cmd.Env = append(cmd.Env, "ANGLE_DEFAULT_PLATFORM=swiftshader")
		cmd.Env = append(cmd.Env, "DISPLAY=:99")
	})

	// No cleanup function needed on Windows
	return []chromedp.ExecAllocatorOption{opt}, func() error { return nil }, nil
}
