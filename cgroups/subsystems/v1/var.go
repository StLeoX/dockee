package v1

import (
	"github.com/stleox/dockee/cgroups/subsystems"
)

var (
	SubsystemIns = []subsystems.Subsystem{
		&CpusetSubSystem{},
		&CpuSubSystem{},
		&MemorySubSystem{},
	}
)
