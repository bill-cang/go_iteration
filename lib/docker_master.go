/*
@Time   : 2021/10/18 18:11
@Author : ckx0709
@Remark :
*/
package lib

import (
	"github.com/google/gopacket/pcap"
)

func GetAllDevs() []pcap.Interface {
	devices, _ := pcap.FindAllDevs()
	return devices
}
