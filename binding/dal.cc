#include <napi.h>
#include "dal.h"

static void _InitSQLite(const Napi::CallbackInfo& args) {
  Napi::Buffer<uint8_t> buf = args[0].As<Napi::Buffer<uint8_t>>();
  GoString charstr = {reinterpret_cast<char *>(buf.Data()), long(buf.Length())};
  InitSQLite(charstr);
}

static Napi::Value Handle(const Napi::CallbackInfo& args) {
  Napi::Buffer<uint8_t> buf = args[0].As<Napi::Buffer<uint8_t>>();
  GoSlice input = {reinterpret_cast<void *>(buf.Data()), long(buf.Length()), long(buf.Length())};
  GoSlice result = HandleQuery(input);
  return Napi::Buffer<char>::Copy(args.Env(), reinterpret_cast<char *>(result.data), result.len);
}


static Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports["InitSQLite"] = Napi::Function::New(env, _InitSQLite);
  exports["Handle"] = Napi::Function::New(env, Handle);
  return exports;
}

NODE_API_MODULE(NODE_GYP_MODULE_NAME, Init)

