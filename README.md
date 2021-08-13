blockchain release monitor
---
blockchain release monitor is a monitoring and notification system for version updates of blockchain projects.
See the [中文文档](./README-zh.md) for Chinese readme.


### Monitored blockchain projects
* Etherum
* Etherum2.0
* Binance Smart Chain(bsc)
* Huobi ECO Chain (heco)
* OKexChain
* polygon
* Polkadot Relay Chain(polkadot)
* avalanche

### Getting Started
* vim .config.json
```
{
  "rest_listen": ":80",
  "grpc_listen": ":50051",
  "notification": {
    "slack": "slack webhook url"
  }
}
```

* up
```
docker-compose up -d
```

### notification 

Please use the following method to enter the notification group:

- Slack: [invite link](https://join.slack.com/t/perception-networkhq/shared_invite/zt-tq6mee6b-ETN3P3kP28BRpn3BHsSCJw)
