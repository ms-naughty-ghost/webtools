FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nestjs:0.1.0 as develop

WORKDIR /webapi/

COPY webtools/ /webapi/

RUN yarn install

CMD [ "yarn", "run", "start:debug" ]

FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nestjs:0.1.0 as builder

WORKDIR /webapi/

COPY webtools/ /webapi/
RUN yarn install
RUN yarn run prebuild
RUN yarn run build

FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nodejs:0.1.0 as production

WORKDIR /webapi/

COPY --from=builder /webapi/dist/ /webapi/dist/
COPY webtools/package.json /webapi/package.json
COPY webtools/yarn.lock /webapi/yarn.lock
RUN yarn install --production

EXPOSE 80

CMD [ "node", "dist/main" ]