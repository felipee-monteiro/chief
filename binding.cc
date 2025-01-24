#include <iostream>
#include <napi.h>
#include <libspeedy.h>

Napi::String InitDriver(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    char* result = Connect("sa", "Epilefac57#$!$24042002", "localhost", "sigma", 1433);

    return Napi::String::New(env, result);
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
    exports.Set(Napi::String::New(env, "connect"), Napi::Function::New(env, InitDriver));
    return exports;
}

NODE_API_MODULE(speedy, Init)
