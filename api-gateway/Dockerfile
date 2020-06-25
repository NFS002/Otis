# Use project home directory or git repo url as build context
#
# Build from git url:
#        docker build -f=api-gateway/Dockerfile -t=${OTIS_CI_REGISTRY}/otis_team/backend/api-gateway:${OTIS_DOCKER_IMAGETAG} https://gitlab.com/otis_team/backend.git
# Build from project home
#        docker build -f=api-gateway/Dockerfile -t=${OTIS_CI_REGISTRY}/otis_team/backend/api-gateway:${OTIS_DOCKER_IMAGETAG} ${OTIS_HOME}
#
# Run with default gateway-config.js:
#        docker run --env-file=.env -v=${OTIS_HOME}/certs:/app/certs --expose=3000 -p=0.0.0.0:3000:3000 ${OTIS_CI_REGISTRY}/otis_team/backend/api-gateway:${OTIS_DOCKER_IMAGETAG}
# Run with specific gateway-config.js:
#        docker run --env-file=.env -v=${OTIS_HOME}/certs:/app/certs -v=${OTIS_HOME}/api-gateway/gateway-config.js:/app/api-gateway/gateway-config.js --expose=3000 -p=0.0.0.0:3000:3000 ${OTIS_CI_REGISTRY}/otis_team/backend/api-gateway:${OTIS_DOCKER_IMAGETAG}
#

FROM node:latest

RUN groupadd -g 61000 docker
RUN useradd -g 61000 -l -M -s /bin/false -u 61000 docker

RUN mkdir -p /app/api-gateway /app/lib && chown -R docker:docker /app

WORKDIR /app/api-gateway


COPY --chown=docker:docker api-gateway/package*.json ./
RUN npm install --production

USER docker

COPY --chown=docker:docker ./api-gateway /app/api-gateway
COPY --chown=docker:docker ./lib /app/lib

CMD ["node", "index.js"]