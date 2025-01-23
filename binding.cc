#include <napi.h>
#include <iostream>
#include <dlfcn.h>

typedef const char* (*ConnectFunc)();

Napi::String CallGoFunction(const Napi::CallbackInfo& info) {
    Napi::Env env = info.Env();

    void* handle = dlopen("./libspeedy.so", RTLD_LAZY);
    
    if (!handle) {
        throw Napi::Error::New(env, "Falha ao carregar a biblioteca Go");
    }

    ConnectFunc helloWorld = (ConnectFunc)dlsym(handle, "Connect");
    
    if (!helloWorld) {
        dlclose(handle);
        throw Napi::Error::New(env, "Função não encontrada na biblioteca Go");
    }

    const char* result = helloWorld();

    dlclose(handle);

    return Napi::String::New(env, result);
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
    exports.Set(Napi::String::New(env, "connect"), Napi::Function::New(env, CallGoFunction));
    return exports;
}

NODE_API_MODULE(speedy, Init)
