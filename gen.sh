#!/bin/bash
echo "Generating Java code..."

sudo rm -rf ./src


docker run --rm -w=/local -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate \
    -i nftinfo.yaml \
    -g go-echo-server \
    -o .

# docker run --rm -w=/local -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate \
#     -i https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.yaml \
#     -g go-echo-server \
#     -o ./out/go

# docker run --rm -w=/local -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate \
#     -i https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.yaml \
#     -g go-server \
#     -o ./out/go

# docker run --rm -w=/local -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate \
#     -i https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.yaml \
#     -g go \
#     -o ./out/go



# docker run --rm -w=/local -v "${PWD}:/local" openapitools/openapi-generator-cli:v6.5.0 generate \
#     -i https://raw.githubusercontent.com/openapitools/openapi-generator/master/modules/openapi-generator/src/test/resources/3_0/petstore.yaml \
#         https://raw.githubusercontent.com/OAI/OpenAPI-Specification/main/examples/v3.0/petstore.yaml
#     -g go \
#     -o ./out/go