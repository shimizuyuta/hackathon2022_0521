FROM node:14.17.0 as build

WORKDIR /myapp

COPY . .
RUN yarn install && \
    yarn build

FROM nginx:stable-alpine

COPY --from=build /myapp/build /var/www
COPY ./nginx /etc/nginx/conf.d

WORKDIR /usr/share/nginx/html

CMD ["nginx", "-g", "daemon off;"]