FROM node:18-alpine AS builder

WORKDIR /app

COPY . .

RUN yarn install && \
    yarn build

FROM node:18-alpine

WORKDIR /app

COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json /app/yarn.lock ./

RUN yarn install --production && \
    yarn cache clean

EXPOSE 3000

CMD ["node", "build/index.js"]