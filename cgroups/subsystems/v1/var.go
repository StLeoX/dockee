package v1

import (
	"github.com/impact-eintr/dockee/cgroups/subsystems"
)

var (
	SubsystemIns = []subsystems.Subsystem{
		&CpusetSubSystem{},
		&CpuSubSystem{},
		&MemorySubSystem{},
	}
)
