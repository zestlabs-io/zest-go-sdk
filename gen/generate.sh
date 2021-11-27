#!/bin/bash
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
PDIR=${DIR}/..

rm -rf ${PDIR}/src/openapi
mkdir -p ${PDIR}/src/openapi

# Copy swagger definitions
if [ "$1" != "nocopy" ]; then
  cp ${PDIR}/../appscape/api/swagger/v1/appscape.swagger.json ${DIR}/.
  cp ${PDIR}/../auth-service/api/swagger/v1/service.swagger.json ${DIR}/auth.swagger.json
  cp ${PDIR}/../data-distribution/functions/api/swagger/v1/functions.swagger.json ${DIR}/.
  cp ${PDIR}/../data-distribution/distr-config/api/swagger/v1/distr-config.swagger.json ${DIR}/.
  cp ${PDIR}/../data-distribution/data/api/swagger/v1/data.swagger.json ${DIR}/.
fi

# Swagger combine
# npm install -g swagger-combine
# swagger-combine config.json -o fullSchema.json
npm install swagger-merger -g
swagger-merger -i ${DIR}/fullRef.json -o ${DIR}/fullSchema.json


# docker pull quay.io/goswagger/swagger
# docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger generate client -f ./gen/fullSchema.json -t api

if [ ! -f /usr/local/bin/swagger ]; then
  download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url')
  sudo curl -o /usr/local/bin/swagger -L'#' "$download_url"
  sudo chmod +x /usr/local/bin/swagger
fi

swagger generate client -f ./gen/fullSchema.json -t api
# Generate
# openapi-generator generate --remove-operation-id-prefix -i ${DIR}/fullSchema.json -g go -o ${PDIR}/openapi
# rm -rf ${PDIR}/openapi/go.mod ${PDIR}/openapi/go.sum