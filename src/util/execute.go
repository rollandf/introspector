package util

import (
	bytes2 "bytes"
	"os/exec"
)

func getExitCode(err error) int {
	if err == nil {
		return 0
	}
	switch  value := err.(type) {
	case *exec.ExitError:
		return value.ExitCode()
	default:
		return -1
	}
}

func getErrorStr(err error, stderr *bytes2.Buffer) string {
	b := stderr.Bytes()
	if len(b) > 0 {
		return string(b)
	} else if err != nil {
		return err.Error()
	}
	return ""
}

func Execute(command string, args ...string) (string, string, int) {
	cmd := exec.Command(command, args...)
	var stdout, stderr bytes2.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return string(stdout.Bytes()), getErrorStr(err, &stderr), getExitCode(err)
}

func ExecuteShell(command string) (string, string, int) {
	return Execute("bash", "-c", command)
}
