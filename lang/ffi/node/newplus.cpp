#include <node.h>
#include <v8.h>
extern "C" {
#include "../newplus/plus.h"
}

void CurrentTimestamp(const v8::FunctionCallbackInfo<v8::Value>& args){
  args.GetReturnValue().Set(static_cast<double>(current_timestamp()));
}

void PlusOne(const v8::FunctionCallbackInfo<v8::Value>& args){
  args.GetReturnValue().Set(plusone(args[0]->IntegerValue(args.GetIsolate()->GetCurrentContext()).FromJust()));
}

void init(v8::Local<v8::Object> exports){
  NODE_SET_METHOD(exports, "current_timestamp", CurrentTimestamp);
  NODE_SET_METHOD(exports, "plusone", PlusOne);
}

NODE_MODULE(binding, init);