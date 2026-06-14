# Golang中使用grpc

## 1. 必须安装的两个插件

请在你的终端（Terminal）中分别运行以下两条命令：

```bash
# 插件 1：用于生成基础的 Protobuf 序列化/结构体代码 (生成 xxx.pb.go)
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# 插件 2：用于生成 gRPC 的客户端/服务端接口代码 (生成 xxx_grpc.pb.go)
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

---

## 2. 核心大前提：你本地必须有 `protoc` 编译器

光有这两个 Go 插件是不够的，因为插件无法独立运行。它们必须依赖 Google 官方的 **`protoc` 编译器（Protocol Buffer Compiler）**。

如果你本地还没有安装 `protoc`，可以根据你的操作系统进行安装：

* **Windows (使用 Scoop 或 Chocolatey 快捷安装):**
```bash
scoop install protobuf
# 或者使用 choco
choco install protoc

```


*(或者去 [Protobuf GitHub Releases](https://github.com/protocolbuffers/protobuf/releases) 下载 `protoc-xxx-win64.zip`，解压后把里面的 `bin` 目录路径配置到 Windows 的系统环境变量 `Path` 中)*
* **Mac (使用 Homebrew):**
```bash
brew install protobuf
```


* **Linux (Ubuntu/Debian):**
```bash
sudo apt install -y protobuf-compiler
```



---

## 🔍 怎么检查是否全部安装成功？

安装完成后，你可以通过检查它们各自的版本号，来确保所有工具都已经正确配置到系统的环境变量中：

```bash
protoc --version          # 检查编译器
protoc-gen-go --version   # 检查 Go 结构体插件
protoc-gen-go-grpc -version # 检查 gRPC 插件

```

如果这三个命令都能正常输出版本号，你就可以顺利执行 `protoc --go_out=...` 命令来生成 gRPC 代码了！


# 生成 go grpc代码
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protoapi.proto
```