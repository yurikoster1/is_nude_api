#@version 3.7.5

FROM nginx:1.19.1-alpine
MAINTAINER Yuri Koster <yurikoster1@gmail.com>

EXPOSE 80
EXPOSE 443

#RUN rm -r /var/www/html
RUN rm /etc/nginx/conf.d/default.conf

RUN apk add --no-cache openssl bash

ENV DOCKERIZE_VERSION v0.6.1
RUN wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
    && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz


COPY ./entrypoint.sh /docker-entrypoint.sh
RUN chmod +x /docker-entrypoint.sh
ENTRYPOINT ["/docker-entrypoint.sh"]

STOPSIGNAL SIGTERM
CMD ["nginx", "-g", "daemon off;"]