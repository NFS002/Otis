# Use project home directory or git repo url as build context
#
# Build from git url:
#        docker build -f=service/merchant/Dockerfile -t=otis.merchant.service https://gitlab.com/otis_team/backend
# Build from project home directory:
#        docker build -f=service/merchant/Dockerfile -t=otis.merchant.service .
#
# Run with default service-config.py:
#        docker run --env-file=.env -v=`pwd`/certs:/app/certs --expose=3005 -p=0.0.0.0:3005:3005 otis.merchant.service
# Run with specific service-config.py:
#       docker run --env-file=.env -v=`pwd`/certs:/app/certs -v=`pwd`/service/merchant/service-config.py:/app/merchant-service/service-config.py --expose=3005 -p=0.0.0.0:3005:3005 otis.merchant.service
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