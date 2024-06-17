package network

import (
	"fmt"
	"net"
	"testing"
)

func TestAllocate(t *testing.T) {
	// 在192.168.0.0/24网段内分配IP
	_, ipnet, _ := net.ParseCIDR("192.168.0.0/24")
	ip, _ := ipAllocator.Allocate(ipnet)
	t.Logf("alloc ip: %v", ip)
}

func TestRelease(t *testing.T) {
	// 在192.168.0.0/24网段中释放方才分配的IP
	ip, ipnet, _ := net.ParseCIDR("192.168.0.1/24")
	err := ipAllocator.Release(ipnet, &ip)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("ip is %s\n", ip)
}
