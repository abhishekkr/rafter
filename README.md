
## Rafter

> a go package to enable any service utilize the distributed power of Raft & Gossip
>
> inspired from [ysf/raftsample](https://github.com/yusufsyaifudin/raft-sample.git)


### Build

```
mkdir -p temp
go build -o temp/cmd-client cmd/client/cmd_client.go
go build -o temp/cmd-server cmd/server/cmd_server.go
```

### Usage

* run a server A with raft id `alpha` at `localhost:6660` and tcp server at `localhost:6661`

```
RAFTER_NODE_PORT=6660 RAFTER_SERVER_PORT=6661 RAFTER_NODE_ID=alpha RAFTER_VOLUME_DIR=/tmp/raft-badger-alpha $(dirname $0)/cmd-server
```

* run a server B with raft id `beta` at `localhost:6670` and tcp server at `localhost:6671`

```
RAFTER_NODE_PORT=6670 RAFTER_SERVER_PORT=6671 RAFTER_NODE_ID=beta RAFTER_VOLUME_DIR=/tmp/raft-badger-beta $(dirname $0)/cmd-server
```

* run the client to talk to `alpha` node

```
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /stats
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /join -id beta -address 127.0.0.1:6670
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /leave -id beta -address 127.0.0.1:6670
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /stats
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /payload -payload-file ./go.mod
RAFTER_NODE_ID=alpha RAFTER_TARGET_PORT=6661 temp/cmd-client -action /stats
```

---
