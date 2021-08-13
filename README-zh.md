blockchain release monitor
---
blockchain release monitor 是区块链项目版本更新的监控通知系统。

### 监控的公链项目
* Etherum
* Etherum2.0
* Binance Smart Chain(bsc)
* Huobi ECO Chain (heco)
* OKexChain
* polygon
* Polkadot Relay Chain(polkadot)
* avalanche

### 如何使用
* vim .config.json
```
{
  "rest_listen": ":80",
  "grpc_listen": ":50051",
  "notification": {
    "work_wechat": "XXXX",
    "slack": "XXXX"
  }
}
```

* up
```
docker-compose up -d
```

### 监控通知群

请同意一下方式进入监控通知群：

* slack: [邀请连接](https://join.slack.com/t/perception-networkhq/shared_invite/zt-tq6mee6b-ETN3P3kP28BRpn3BHsSCJw)