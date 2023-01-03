FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nestjs:0.1.0 as develop

WORKDIR /webapi/

COPY webtools/ /webapi/

CMD [ "yarn", "run", "start:debug" ]

FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nestjs:0.1.0 as builder

WORKDIR /webapi/

COPY webtools/ /webapi/

RUN yarn run prebuild
RUN yarn run build

FROM ghcr.io/ms-naughty-ghost/dev-docker-image-nodejs:0.1.0 as production

WORKDIR /webapi/

COPY --from=builder /webapi/dist/ /webapi/dist/

CMD [ "node", "dist/main" ]