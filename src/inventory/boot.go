package inventory

import (
	"strings"

	"github.com/filanov/bm-inventory/models"
)

type boot struct {
	dependencies IDependencies
}

func newBoot(dependencies IDependencies) *boot {
	return &boot{dependencies: dependencies}
}

func (b *boot) getPxeInterface() string {
	cmdline, err := b.dependencies.ReadFile("/proc/cmdline")
	if err != nil {
		return ""
	}
	prefix := "BOOTIF="
	for _, part := range strings.Split(strings.TrimSpace(string(cmdline)), " ") {
		if strings.HasPrefix(part, prefix) {
			return part[len(prefix):]
		}
	}
	return ""
}

func (b *boot) getCurrentBootMode() (ret string) {
	ret = "bios"
	mode, err := b.dependencies.Stat("/sys/firmware/efi")
	if err != nil {
		return
	}
	if mode.IsDir() {
		ret = "uefi"
	}
	return
}

func (b *boot) getBoot() *models.Boot {
	ret := models.Boot{
		CurrentBootMode: b.getCurrentBootMode(),
		PxeInterface:    b.getPxeInterface(),
	}
	return &ret
}

func GetBoot(dependencies IDependencies) *models.Boot {
	return newBoot(dependencies).getBoot()
}
