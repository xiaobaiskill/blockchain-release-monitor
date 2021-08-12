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