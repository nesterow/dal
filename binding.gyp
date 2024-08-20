{
  'targets': [
    {
      'target_name': 'dal',
      'sources': [
        'cgo/dal.h',
        'cgo/dal.cc'
      ],
      'libraries': [ '../cgo/dal.a' ],
      'defines': [ 'NAPI_DISABLE_CPP_EXCEPTIONS' ],
      "include_dirs": [
        "<!@(node -p \"require('node-addon-api').include\")"
      ],
    }
  ]
}