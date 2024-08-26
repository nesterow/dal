#include <napi.h>
#include "dal.h"

static void _InitSQLite(const Napi::CallbackInfo& args) {
  Napi::Buffer<uint8_t> buf = args[0].As<Napi::Buffer<uint8_t>>();
  GoString charstr = {reinterpret_cast<char *>(buf.Data()), long(buf.Length())};
  InitSQLite(charstr);
}

static Napi::Object RowIterator(const Napi::CallbackInfo& args) {
  Napi::Env env = args.Env();
  Napi::Object it = Napi::Object::New( env );
  Napi::Buffer<uint8_t> buf = args[0].As<Napi::Buffer<uint8_t>>();
  GoSlice input = {reinterpret_cast<void *>(buf.Data()), long(buf.Length()), long(buf.Length())};
  int iter = CreateRowIterator(input);
  auto next_row = [=](const Napi::CallbackInfo& a){
    void* next = NextRow(iter);
    if (next == nullptr) {
      FreeIter(iter);
      return Napi::Buffer<uint8_t>::New(env, 0);
    }
    Napi::Buffer<uint8_t> val = Napi::Buffer<uint8_t>::Copy(env, reinterpret_cast<uint8_t *>(next), GetLen(iter));
    free(next);
    return val;
  };
  auto free_iter = [=](const Napi::CallbackInfo& a){
    FreeIter(iter);
  };
  it.Set("next", Napi::Function::New(env, next_row));
  it.Set("free", Napi::Function::New(env, free_iter));
  return it;
}


static Napi::Object Init(Napi::Env env, Napi::Object exports) {
  exports["InitSQLite"] = Napi::Function::New(env, _InitSQLite);
  exports["RowIterator"] = Napi::Function::New(env, RowIterator);
  return exports;
}

NODE_API_MODULE(NODE_GYP_MODULE_NAME, Init)

