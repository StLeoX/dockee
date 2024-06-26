package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/stleox/dockee/container"
	_ "github.com/stleox/dockee/nsenter"
)

const (
	ENV_EXEC_PID = "docker_pid"
	ENV_EXEC_CMD = "docker_cmd"
)

func ExecContainer(containerName string, comArray []string) {
	pid, err := GetContainerPidByName(containerName)
	if err != nil {
		logrus.Errorf("Exec container getContainerPidByName %s error %v", containerName, err)
		return
	}
	cmdStr := strings.Join(comArray, " ")
	logrus.Infof("container pid %s", pid)
	logrus.Infof("command %s", cmdStr)

	cmd := exec.Command("/proc/self/exe", "exec")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = os.Setenv(ENV_EXEC_PID, pid)
	if err != nil {
		logrus.Warn("can't set environment variable")
	}
	err = os.Setenv(ENV_EXEC_CMD, cmdStr)
	if err != nil {
		logrus.Warn("can't set environment variable")
	}
	containerEnvs := getEnvsByPid(pid)
	cmd.Env = append(os.Environ(), containerEnvs...)

	if err := cmd.Run(); err != nil {
		logrus.Errorf("Exec container %s error %v", containerName, err)
	}
}

func GetContainerPidByName(containerName string) (string, error) {
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	configFilePath := dirURL + container.ConfigName
	contentBytes, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return "", err
	}
	var containerInfo container.ContainerInfo
	if err := json.Unmarshal(contentBytes, &containerInfo); err != nil {
		return "", err
	}
	pid := containerInfo.Pid
	if strings.Contains(pid, " ") {
		return "", fmt.Errorf("check the status of container %s", containerName)
	}

	return pid, nil
}

func getEnvsByPid(pid string) []string {
	path := fmt.Sprintf("/proc/%s/environ", pid)
	contentBytes, err := ioutil.ReadFile(path)
	if err != nil {
		logrus.Errorf("Read file %s error %v", path, err)
		return nil
	}
	//env split by \u0000
	envs := strings.Split(string(contentBytes), "\u0000")
	return envs
}
