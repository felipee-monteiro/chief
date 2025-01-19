#include <napi.h>
#include <dlfcn.h>
#include <iostream>

extern "C" {
    const char* HelloGo();
}

Napi::String HelloGoWrapper(const Napi::CallbackInfo& info) {
    void* handle = dlopen("./libhello.so", RTLD_LAZY);
    if (!handle) {
        std::cerr << "Erro ao carregar a biblioteca Go: " << dlerror() << '\n';
        return Napi::String::New(info.Env(), "Erro ao carregar a biblioteca Go");
    }

    const char* result = HelloGo();  
    dlclose(handle);

    return Napi::String::New(info.Env(), result);
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
    exports.Set("helloGo", Napi::Function::New(env, HelloGoWrapper));
    return exports;
}

NODE_API_MODULE(binding, Init) 
