{
  "targets": [
    {
      "target_name": "speedy",
      "sources": ["binding.cc"],
      "include_dirs": [
        "/home/felipe/projects/go/api/node_modules/node-addon-api/",
        "."
      ],
      "dependencies": [
        "<!(node -p \"require('node-addon-api').gyp\")"
      ],
      "cflags!": ["-fno-exceptions"],
      "cflags_cc!": ["-fno-exceptions"],
      "defines": ["NAPI_DISABLE_CPP_EXCEPTIONS"]
    }
  ]
}

