package cgroups

import (
	"github.com/impact-eintr/dockee/cgroups/subsystems"
	subV1 "github.com/impact-eintr/dockee/cgroups/subsystems/v1"
	subV2 "github.com/impact-eintr/dockee/cgroups/subsystems/v2"
	log "github.com/sirupsen/logrus"
)

type CgroupManager struct {
	Path     string
	Resource *subsystems.ResourceConfig
}

func NewCgroupManager(path string) *CgroupManager {
	return &CgroupManager{
		Path: path,
	}
}

func (c *CgroupManager) Apply2(pid int) error {
	return subV2.SubsystemIns[0].Apply(c.Path, pid)
}

func (c *CgroupManager) Set2(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subV2.SubsystemIns {
		if err := subSysIns.Set(c.Path, res); err != nil {
			log.Warnf("set cgroup fail %v", err)
		}
	}
	return nil
}

func (c *CgroupManager) Destroy2() error {
	if err := subV2.SubsystemIns[0].Remove(c.Path); err != nil {
		log.Warnf("remove cgroup fail %v", err)
	}
	return nil
}

func (c *CgroupManager) Apply(pid int) error {
	for _, subSysIns := range subV1.SubsystemIns {
		if err := subSysIns.Apply(c.Path, pid); err != nil {
			log.Warnf("apply cgroup fail %v", err)
		}
	}
	return nil
}

func (c *CgroupManager) Set(res *subsystems.ResourceConfig) error {
	for _, subSysIns := range subV1.SubsystemIns {
		if err := subSysIns.Set(c.Path, res); err != nil {
			log.Warnf("set cgroup fail %v", err)
		}
	}
	return nil
}

func (c *CgroupManager) Destroy() error {
	for _, subSysIns := range subV1.SubsystemIns {
		if err := subSysIns.Remove(c.Path); err != nil {
			log.Warnf("remove cgroup fail %v", err)
		}
	}

	return nil
}
