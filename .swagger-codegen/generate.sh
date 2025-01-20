##!/bin/sh
if [ -x "$(command -v swagger-codegen)" ]; then
    codegen_dir=$(dirname "$0")
    swagger-codegen generate \
        --lang go \
        --input-spec https://developer.shortcut.com/api/rest/v3/shortcut.swagger.json \
        --output generated \
        --config $codegen_dir/config.json
else
    echo "Please install swagger-codegen: https://github.com/swagger-api/swagger-codegen | brew install swagger-codegen"
    exit 1
fi
