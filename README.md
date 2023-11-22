# go-grpc-tutorial

### PREREQUISITES

- Installed Go, e.g. for macOS `brew install go` or use [`gvm`](https://github.com/moovweb/gvm)
- Installed `protoc` (for macOS: `brew install protobuf`)
- Installed `protoc-gen-go` (`go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28`)
- Installed `protoc-gen-go-grpc` (`go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2`)
- Setup environment variable `PATH` to include `$GOPATH/bin`
- Recommend editor [Visual Studio Code](https://code.visualstudio.com/) with [`vscode-proto3` extension](https://marketplace.visualstudio.com/items?itemName=zxh404.vscode-proto3) to compile protobufs on save (or Ctrl+P > compile to compile manually)

### USAGE

- start server in one terminal
    ```command
    $ go run main.go -server
    2023/11/22 23:28:57 [server] listening on tcp://127.0.0.1:9991
    ```
- start client in another terminal
    ```command
    % go run main.go 'Wow, the echo is amazing here!'
    2023/11/22 23:29:16 [client] connecting to tcp://127.0.0.1:9991
    2023/11/22 23:29:16 [client] sending request: data="Wow, the echo is amazing here!"
    2023/11/22 23:29:16 [client] received response: data="Wow, the echo is amazing here!", status=OK
    ```
- observe received messages in server terminal
    ```command
    % go run main.go -server
    2023/11/22 23:28:57 [server] listening on tcp://127.0.0.1:9991
    2023/11/22 23:29:16 [server] received request: data="Wow, the echo is amazing here!"
    2023/11/22 23:29:16 [server] sending response: data="Wow, the echo is amazing here!", status=OK
    ```