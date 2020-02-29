FROM nginx:stable-alpine

# Sets which config file to use
ARG config

COPY $config /etc/nginx/conf.d/default.conf

#COPY ssl.pem /etc/nginx/ssl/
#COPY ssl.key /etc/nginx/ssl/

EXPOSE 80/tcp
EXPOSE 443/tcp

CMD ["/bin/sh", "-c", "exec nginx -g 'daemon off;';"]

WORKDIR /usr/share/nginx/html
