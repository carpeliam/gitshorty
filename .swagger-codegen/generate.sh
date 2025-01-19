##!/bin/sh
echo $0
codegen_dir=$(dirname "$0")
swagger-codegen generate --lang go --input-spec https://developer.shortcut.com/api/rest/v3/shortcut.swagger.json --output sc --config $codegen_dir/config.json
