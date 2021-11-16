/*
@Time   : 2021/10/18 18:13
@Author : ckx0709
@Remark :
*/
package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	//获取ctx
	ctx := context.Background()

	//cli客户端对象
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	//获取容器id 这个其实docker ps那个命令，不过我们只需要容器id
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	//遍历获取到的容器
	for _, container := range containers {
		fmt.Println("--------容器ID-------")
		fmt.Println(container.ID)
		fmt.Println(container.Image)
		fmt.Println(container.ImageID)
		fmt.Println("根据容器id获取容器的stats")
		//通过cli的ContainerStats方法可以获取到 docker stats命令的详细信息，其实是一个容器监控的方法
		//这个方法返回了容器使用CPU 内存 网络 磁盘等诸多信息
		containerStats, err := cli.ContainerStats(ctx, container.ID, false)
		if err != nil {
			panic(err)
		}
		/**
		ContainerStats的返回的结构如下 注意这个Body的类型是io.ReadCloser 好奇怪的类型 下面我们给他转成json
		type ContainerStats struct {
			Body   io.ReadCloser `json:"body"`
			OSType string        `json:"ostype"`
		}
		*/
		fmt.Println(containerStats)
		fmt.Println("containerStats.Body的内容是: ", containerStats.Body)
		buf := new(bytes.Buffer)
		//io.ReadCloser 转换成 Buffer 然后转换成json字符串
		buf.ReadFrom(containerStats.Body)
		newStr := buf.String()
		fmt.Printf(newStr)
	}
}
