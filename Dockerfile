FROM node:14.8.0-alpine

ENV APP_ROOT /src

ENV PLATFORM_TYPE=docker

RUN mkdir ${APP_ROOT}
WORKDIR ${APP_ROOT}
ADD . ${APP_ROOT}

RUN npm install
RUN npm run build

# Expose the app port
EXPOSE 3000
ENV NODE_ENV=production
ENV HOST 0.0.0.0

CMD [ "npm", "start" ]