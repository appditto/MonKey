FROM node:16-alpine

ENV APP_ROOT /src

ENV PLATFORM_TYPE=docker

RUN mkdir ${APP_ROOT}
WORKDIR ${APP_ROOT}
ADD . ${APP_ROOT}

RUN npm install
RUN npm run export

# Expose the app port
ENV NODE_ENV=production

CMD [ "npx", "serve", "__sapper__/export" ]
