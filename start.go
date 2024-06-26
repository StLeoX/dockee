package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/stleox/dockee/container"
)

// start 指令
func startContainer(containerName string) {
	info, err := getContainerInfoByName(containerName)
	if err != nil {
		logrus.Errorf("Get container %s info error %v", containerName, err)
		return
	}

	parent, writePipe := container.ReNewParentProcess(info, "")
	if parent == nil {
		log.Errorf("New parent process error")
		return
	}

	if err := parent.Start(); err != nil {
		log.Errorf("New parent process error: %v", err)
	}

	// 把新的容器信息写回配置文件
	info.Status = container.RUNNING
	info.Pid = fmt.Sprintf("%d", parent.Process.Pid)

	newContentBytes, err := json.Marshal(info)
	if err != nil {
		logrus.Errorf("Json marshal %s error %v", containerName, err)
		return
	}
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerName)
	configFilePath := dirURL + container.ConfigName
	if err := ioutil.WriteFile(configFilePath, newContentBytes, 0622); err != nil {
		logrus.Errorf("Write file %s error: %s", configFilePath, err)
	}
	sendInitCommand([]string{info.Command}, writePipe)

	log.Infof("%s started", containerName)
}
