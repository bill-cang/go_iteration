/*
@Time   : 2021/10/22 15:20
@Author : ckx0709
@Remark :
*/
package main

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"log"
)

func main() {
	GetAll()
}

//获取所有(网络)设备
func GetAll() {
	// 得到所有的(网络)设备
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	// 打印设备信息
	fmt.Println("Devices found:")
	for _, device := range devices {
		fmt.Println("\nName: ", device.Name)
		fmt.Println("Description: ", device.Description)
		fmt.Println("Devices addresses: ", device.Description)
		for _, address := range device.Addresses {
			fmt.Println("- IP address: ", address.IP)
			fmt.Println("- Subnet mask: ", address.Netmask)
		}
	}
}
