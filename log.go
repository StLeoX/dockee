package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/stleox/dockee/container"
)

func logContainer(containerId string) {
	dirURL := fmt.Sprintf(container.DefaultInfoLocation, containerId)
	logFileLocation := dirURL + container.ContainerLogFile
	file, err := os.Open(logFileLocation)
	defer file.Close()
	if err != nil {
		logrus.Errorf("Log container open file %s error %v", logFileLocation, err)
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		logrus.Errorf("Log container read file %s error %v", logFileLocation, err)
		return
	}
	fmt.Fprint(os.Stdout, string(content))
}
