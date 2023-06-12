package main

import (
	"fmt"
	"golang.org/x/net/context"
	"log"

	rpc "github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc/imservice"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var (
	rdb = &RedisClient{} //gets address of the redis client struct
)

func main() {
	ctx := context.Background() // https://www.digitalocean.com/community/tutorials/how-to-use-contexts-in-go

	e := rdb.InitClient(ctx, "redis:6379", "")
	//address must be the redis server and not hostname
	if e != nil {
		errMsg := fmt.Sprintf("Redis client unable to boot up, err: %v", e)
		log.Fatal(errMsg)
	}

	r, err := etcd.NewEtcdRegistry([]string{"etcd:2379"}) // r should not be reused.
	if err != nil {
		log.Fatal(err)
	}

	svr := rpc.NewServer(new(IMServiceImpl), server.WithRegistry(r), server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: "demo.rpc.server",
	}))

	err = svr.Run()
	if err != nil {
		log.Println(err.Error())
	}
}
