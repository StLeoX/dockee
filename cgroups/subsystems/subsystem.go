package subsystems

type ResourceConfig struct {
	MemoryLimit string
	Cpu         string
	Cpuset      string
}

type Subsystem interface {
	Name() string
	Set(cgroupPath string, res *ResourceConfig) error
	Apply(cgroupPath string, pid int) error
	Remove(cgroupPath string) error
}
