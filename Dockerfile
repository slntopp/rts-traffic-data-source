FROM node:latest

ADD . /app
WORKDIR /app
RUN yarn

LABEL org.opencontainers.image.source https://github.com/slntopp/rts-traffic-data-source
ENTRYPOINT ["node", "/app"]