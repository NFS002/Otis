# Use project home directory or git repo url as build context
#
# Build from git url:
#        docker build -f=service/merchant/Dockerfile -t=${OTIS_CI_REGISTRY}/otis_team/backend/service/merchant:${OTIS_DOCKER_IMAGETAG} https://gitlab.com/otis_team/backend.git
# Build from project home directory:
#        docker build -f=service/merchant/Dockerfile -t=${OTIS_CI_REGISTRY}/otis_team/backend/service/merchant:${OTIS_DOCKER_IMAGETAG} ${OTIS_HOME}
#
# Run with default service-config.py:
#        docker run --env-file=.env -v=${OTIS_HOME}/certs:/app/certs --expose=3005 -p=0.0.0.0:3005:3005 ${OTIS_CI_REGISTRY}/otis_team/backend/service/merchant:${OTIS_DOCKER_IMAGETAG}
# Run with specific service-config.py:
#        docker run --env-file=.env -v=${OTIS_HOME}/certs:/app/certs -v=${OTIS_HOME}/service/merchant/service-config.py:/app/merchant-service/service-config.py --expose=3005 -p=0.0.0.0:3005:3005 ${OTIS_CI_REGISTRY}/otis_team/backend/service/merchant:${OTIS_DOCKER_IMAGETAG}
#

FROM python:3.6

RUN groupadd -g 61000 docker
RUN useradd -g 61000 -l -M -s /bin/false -u 61000 docker

RUN mkdir -p /app/service/merchant /app/lib && chown -R docker:docker /app

WORKDIR /app/service/merchant

COPY --chown=docker:docker service/merchant/requirements.txt ./
RUN pip install -r requirements.txt

USER docker

COPY --chown=docker:docker ./service/merchant /app/service/merchant
COPY --chown=docker:docker ./lib /app/lib

CMD ["python3", "."]