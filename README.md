# wigfrid

wigfrid is the gRPC daemon of wendy panel, responsible for managing local docker containers.

## Dev

build program
```bash
make build os=linux arch=amd64
```

generate protobuf files
```bash
make proto
```

generate wire file
```bash
make wire
```