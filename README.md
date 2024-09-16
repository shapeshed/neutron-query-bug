go build main.go
./main

```sh
panic: invalid Go type math.Int for field neutron.dex.QueryEstimatePlaceLimitOrderRequest.amount_in

goroutine 1 [running]:
google.golang.org/protobuf/internal/impl.newSingularConverter({0x170a778, 0x143c540}, {0x1709e68, 0xc0003cede0})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/convert.go:142 +0x9ad
google.golang.org/protobuf/internal/impl.NewConverter({0x170a778, 0x143c540}, {0x1709e68, 0xc0003cede0})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/convert.go:60 +0x8f
google.golang.org/protobuf/internal/impl.fieldInfoForScalar({0x1709e68, 0xc0003cede0}, {{0x12f318a, 0x8}, {0x0, 0x0}, {0x170a778, 0x143c540}, {0x12f3193, 0x7d}, ...}, ...)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message_reflect_field.go:270 +0x193
google.golang.org/protobuf/internal/impl.(*MessageInfo).makeKnownFieldsFunc(0xc0001fe000, {0xffffffffffffffff, {0x0, 0x0}, 0xffffffffffffffff, {0x0, 0x0}, 0xffffffffffffffff, {0x0, 0x0}, ...})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message_reflect.go:80 +0x78a
google.golang.org/protobuf/internal/impl.(*MessageInfo).makeReflectFuncs(0xc0001fe000, {0x170a778, 0x13cbb20}, {0xffffffffffffffff, {0x0, 0x0}, 0xffffffffffffffff, {0x0, 0x0}, 0xffffffffffffffff, ...})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message_reflect.go:42 +0x58
google.golang.org/protobuf/internal/impl.(*MessageInfo).initOnce(0xc0001fe000)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message.go:90 +0x1b0
google.golang.org/protobuf/internal/impl.(*MessageInfo).init(...)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message.go:72
google.golang.org/protobuf/internal/impl.(*messageReflectWrapper).ProtoMethods(0x12a2600?)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/internal/impl/message_reflect_gen.go:162 +0x25
google.golang.org/protobuf/proto.protoMethods(...)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/proto/proto_methods.go:19
google.golang.org/protobuf/proto.MarshalOptions.size({{}, 0x90?, 0x9b?, 0x64?}, {0x16fda40, 0xc0004b2690})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/proto/size.go:33 +0x38
google.golang.org/protobuf/proto.MarshalOptions.Size({{}, 0xe0?, 0x2?, 0x41?}, {0x16dc2a0?, 0xc0004b2690?})
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/proto/size.go:26 +0x4c
google.golang.org/protobuf/proto.Size(...)
        /home/go/go/pkg/mod/google.golang.org/protobuf@v1.34.2/proto/size.go:16
google.golang.org/grpc/encoding/proto.(*codecV2).Marshal(0xc000731a28?, {0x14102e0?, 0xc000649b90?})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/encoding/proto/proto.go:49 +0x7a
google.golang.org/grpc.encode({0x77b9db667248?, 0x31b18e0?}, {0x14102e0?, 0xc000649b90?})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/rpc_util.go:694 +0x4a
google.golang.org/grpc.prepareMsg({0x14102e0?, 0xc000649b90?}, {0x77b9db667248?, 0x31b18e0?}, {0x0, 0x0}, {0x0, 0x0}, {0x16dfea0, 0xc00006ce80})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/stream.go:1830 +0xe5
google.golang.org/grpc.(*clientStream).SendMsg(0xc0001f98c0, {0x14102e0, 0xc000649b90})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/stream.go:906 +0xf1
google.golang.org/grpc.invoke({0x16ec210?, 0x31b18e0?}, {0x14858a8?, 0x12?}, {0x14102e0, 0xc000649b90}, {0x13c31a0, 0xc00071a4b0}, 0xc00071a4b0?, {0x0, ...})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/call.go:70 +0x9f
google.golang.org/grpc.(*ClientConn).Invoke(0xc0005ebc08, {0x16ec210?, 0x31b18e0?}, {0x14858a8?, 0xc000649b90?}, {0x14102e0?, 0xc000649b90?}, {0x13c31a0?, 0xc00071a4b0?}, {0x0, ...})
        /home/go/go/pkg/mod/google.golang.org/grpc@v1.66.2/call.go:37 +0x23f
github.com/neutron-org/neutron/v4/x/dex/types.(*queryClient).EstimatePlaceLimitOrder(0xc000731f18, {0x16ec210, 0x31b18e0}, 0xc000649b90, {0x0, 0x0, 0x0})
        /home/go/go/pkg/mod/github.com/neutron-org/neutron/v4@v4.2.3/x/dex/types/query.pb.go:2282 +0xc8
main.main()
        /home/go/neutron-query-bug/main.go:39 +0x1f9
````
