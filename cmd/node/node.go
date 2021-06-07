package main

import (
	"flag"
	"fmt"
	"syscall"

	"github.com/enix/dothill-csi/pkg/common"
	"github.com/enix/dothill-csi/pkg/node"
	"k8s.io/klog"
)

var kubeletPath = flag.String("kubeletpath", "/var/lib/kubelet", "Kubelet path")
var bind = flag.String("bind", fmt.Sprintf("unix:///var/lib/kubelet/plugins/%s/csi.sock", common.PluginName), "RPC bind URI (can be a UNIX socket path or any URI)")
var chroot = flag.String("chroot", "", "Chroot into a directory at startup (used when running in a container)")

func main() {
	klog.InitFlags(nil)
	flag.Set("logtostderr", "true")
	flag.Parse()

	if *chroot != "" {
		if err := syscall.Chroot(*chroot); err != nil {
			panic(err)
		}
	}

	klog.Infof("starting dothill storage node plugin %s", common.Version)
	node.New(*kubeletPath).Start(*bind)
}
