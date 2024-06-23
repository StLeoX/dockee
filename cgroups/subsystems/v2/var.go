package v2

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
