FROM node:latest

ADD . /app
WORKDIR /app
RUN yarn

ENTRYPOINT ["node", "/app"]