1.使用libp2p构建一host
    对于主机的概念:
    对于大多数应用程序,主机是开始使用的基本构建块.
    主机是一个抽象,它管理群上的服务.它提供了一个干净的接口来连接到给定远程对等点上的服务.
 1>创建带有默认配置的host,参考host.DefaultHost()
 2>如果希望配置有更多控制,可以向构造函数指定一些选项.有关构造函数支持的所有配置
 的完整列表参考:options
    type Config struct {
        PeerKey crypto.PrivKey

        Transports         []TptC
        Muxers             []MsMuxC
        SecurityTransports []MsSecC
        Insecure           bool
        Protector          pnet.Protector

        Relay     bool
        RelayOpts []circuit.RelayOpt

        ListenAddrs  []ma.Multiaddr
         bhost.AddrsFactory
        Filters      *filter.Filters
AddrsFactory
        ConnManager ifconnmgr.ConnManager
        NATManager  NATManagerC
        Peerstore   pstore.Peerstore
        Reporter    metrics.Reporter
    }