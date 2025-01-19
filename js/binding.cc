#include <napi.h>
#include <dlfcn.h>
#include <iostream>

extern "C" {
    const char* HelloGo(); // Declara a função Go que será chamada
}

Napi::String HelloGoWrapper(const Napi::CallbackInfo& info) {
    // Tenta carregar a biblioteca Go
    void* handle = dlopen("./libhello.so", RTLD_LAZY);
    if (!handle) {
        std::cerr << "Erro ao carregar a biblioteca Go: " << dlerror() << std::endl;
        return Napi::String::New(info.Env(), "Erro ao carregar a biblioteca Go");
    }

    // Chama a função Go depois de carregar a biblioteca
    const char* result = HelloGo();  // Chama a função Go
    dlclose(handle);  // Fecha a biblioteca

    return Napi::String::New(info.Env(), result);  // Retorna o resultado para Node.js
}

Napi::Object Init(Napi::Env env, Napi::Object exports) {
    // Registra a função HelloGoWrapper em Node.js
    exports.Set("helloGo", Napi::Function::New(env, HelloGoWrapper));
    return exports;
}

NODE_API_MODULE(libhello, Init) 
