{
  'targets': [
    {
      'target_name': 'dal',
      'sources': [
        'binding/dal.h',
        'binding/dal.cc'
      ],
      'libraries': [ '../binding/dal.a' ],
      'defines': [ 'NAPI_DISABLE_CPP_EXCEPTIONS' ],
      "include_dirs": [
        "<!@(node -p \"require('node-addon-api').include\")"
      ],
    }
  ]
}