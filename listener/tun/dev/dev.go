package dev

import (
	"bytes"
	"os/exec"
	"runtime"

	"github.com/Dreamacro/clash/log"
)

// TunDevice is cross-platform tun interface
type TunDevice interface {
	Name() string
	URL() string
	MTU() (int, error)
	IsClose() bool
	Close() error
	Read(buff []byte) (int, error)
	Write(buff []byte) (int, error)
}

func SetLinuxAutoRoute() {
	log.Infoln("Tun adapter auto setting global route")
	addLinuxSystemRoute("0")
	//addLinuxSystemRoute("1")
	//addLinuxSystemRoute("2/7")
	//addLinuxSystemRoute("4/6")
	//addLinuxSystemRoute("8/5")
	//addLinuxSystemRoute("16/4")
	//addLinuxSystemRoute("32/3")
	//addLinuxSystemRoute("64/2")
	//addLinuxSystemRoute("128.0/1")
	//addLinuxSystemRoute("198.18.0/16")
}

func RemoveLinuxAutoRoute() {
	log.Infoln("Tun adapter removing global route")
	delLinuxSystemRoute("0")
	//delLinuxSystemRoute("1")
	//delLinuxSystemRoute("2/7")
	//delLinuxSystemRoute("4/6")
	//delLinuxSystemRoute("8/5")
	//delLinuxSystemRoute("16/4")
	//delLinuxSystemRoute("32/3")
	//delLinuxSystemRoute("64/2")
	//delLinuxSystemRoute("128.0/1")
	//delLinuxSystemRoute("198.18.0/16")
}

func addLinuxSystemRoute(net string) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		return
	}
	cmd := exec.Command("route", "add", "-net", net, "meta")
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		log.Errorln("[auto route] Failed to add system route: %s: %s , cmd: %s", err.Error(), stderr.String(), cmd.String())
	}
}

func delLinuxSystemRoute(net string) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" {
		return
	}
	cmd := exec.Command("route", "delete", "-net", net, "meta")
	_ = cmd.Run()
	//if err := cmd.Run(); err != nil {
	//	log.Errorln("[auto route]Failed to delete system route: %s, cmd: %s", err.Error(), cmd.String())
	//}
}