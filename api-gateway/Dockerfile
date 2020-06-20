FROM node:latest

ARG OTIS_HOME
ENV OTIS_HOME=${OTIS_HOME}

ENV OTIS_ENV production
RUN mkdir /app
WORKDIR /app


COPY . .

RUN npm install --production

CMD ["node", "."]