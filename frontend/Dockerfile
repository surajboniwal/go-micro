# Stage 1
FROM node:gallium-slim as builder

WORKDIR /app

COPY ./package.json ./

RUN yarn install

COPY . ./

RUN yarn build

# Stage 2
FROM nginx:mainline-alpine-slim

WORKDIR /usr/share/nginx/html

RUN rm -rf ./*

COPY --from=builder /app/build .

ENTRYPOINT ["nginx", "-g", "daemon off;"]