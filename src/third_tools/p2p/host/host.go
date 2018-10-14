package host

import (
	"context"
	"crypto/rand"
	"fmt"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-crypto"
)

//创建一个默认配置的host
//The host is an abstraction that manages services on top
//of  a swam.It provides a clean interface to  connect to
//a service on a give remote peer.
func DefaultHost(){

	//上下文控制节点的生命周期
	//The context governs the lifetime of the libp2p node
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//构建一个默认配置的简单主机
	//To construct a simple host with all the default settings,
	//Just use 'New' method
	h,err := libp2p.New(ctx)
	if err != nil{
		panic(err)
	}

	fmt.Printf("hello world, my hosts ID is %s \n", h.ID())
}

//If you want more control over the configuration,you can specify some
//options to the constructor.For a full list of all the
//configuration supported by the constructor see: options.go

//In this snippet we generate our own ID and specified on which
//address and we want to listen:
//指定我们自己的id和监听端口
func OwnerHost(){
	//Set your own keypair
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	priv, _, err := crypto.GenerateEd25519Key(rand.Reader)
	if err != nil {
		panic(err)
	}
	h2, err := libp2p.New(ctx,
			//use your own created keypair
			libp2p.Identity(priv),
			//set your own listen address
			//The config takes an array of addresses, specify as many as you want
			libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/9000"),
		)
	if err != nil{
		panic(err)
	}

	fmt.Printf("Hello world, my second hosts ID is %s \n", h2.ID())
	fmt.Printf("the host's addr: %s \n",h2.Addrs())
	fmt.Printf("the host's network: %v \n",h2.Network())
}
