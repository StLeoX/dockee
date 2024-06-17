package v2

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var Root string

func GetCgroupPath(cgroupPath string, autoCreate bool) (s string, err error) {

	var cgroupRoot string
	if cgroupPath != "" {
		Root, err = FindCgroupMountpoint()
		cgroupRoot = FindAbsoluteCgroupMountpoint()

	} else {
		cgroupRoot = Root
		fmt.Println(cgroupRoot)
	}

	_, err = os.Stat(path.Join(cgroupRoot, cgroupPath))
	if err == nil || (autoCreate && os.IsNotExist(err)) {
		if os.IsNotExist(err) {
			if err := os.Mkdir(path.Join(cgroupRoot, cgroupPath), os.ModePerm); err != nil {
				return "", fmt.Errorf("error create cgroup %v", err)
			}
		}
		return path.Join(cgroupRoot, cgroupPath), nil

	} else {
		return "", fmt.Errorf("cgroup path error %v", err)

	}
}
func FindAbsoluteCgroupMountpoint() string {
	return "/sys/fs/cgroup"
}

func FindCgroupMountpoint() (s string, err error) {
	f, err := os.Open("/proc/self/cgroup")
	if err != nil {
		return "", err
	}
	defer func() {
		err = f.Close()
	}()

	rawPath, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}
	arr := strings.Split(string(rawPath), ":")
	fp := arr[len(arr)-1]

	return "/sys/fs/cgroup" + string(fp[:len(fp)-1]), nil

}
