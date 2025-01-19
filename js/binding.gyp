{
  "targets": [
    {
      "target_name": "binding",
      "sources": ["binding.cc"],
      "dependencies": [
        "<!(node -e \"require('node-addon-api');\")"
      ]
    }
  ]
}

