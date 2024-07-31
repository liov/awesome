http一般返回我们约定这样
```json
{
  "code": 0,
  "msg": "",
  "data": {}
}
```

grpc的错误status这样
```go

type Status = status.Status

// status.Status
type Status struct {
s *spb.Status
}

// spb.Status
type Status struct {
    state         protoimpl.MessageState
    sizeCache     protoimpl.SizeCache
    unknownFields protoimpl.UnknownFields
    
    // The status code, which should be an enum value of
    // [google.rpc.Code][google.rpc.Code].
    Code int32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
    // A developer-facing error message, which should be in English. Any
    // user-facing error message should be localized and sent in the
    // [google.rpc.Status.details][google.rpc.Status.details] field, or localized
    // by the client.
    Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
    // A list of messages that carry the error details.  There is a common set of
    // message types for APIs to use.
    Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}
```

这里的Details是个数组，并且并不是我们认为的data,而是错误的详情