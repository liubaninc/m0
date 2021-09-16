FROM nginx
VOLUME /tmp1
ENV LANG en_US.UTF-8
RUN apt-get -y update && apt-get -y install nginx 
# ADD default /etc/nginx/sites-enabled/
ADD default.conf /etc/nginx/conf.d/
# ADD nginx.conf /etc/nginx/
# ADD ./build.tar.gz /usr/share/nginx/m0-wallet/
ADD ./build.tar.gz /usr/share/nginx/html/
EXPOSE 80
ENTRYPOINT  ["nginx","-g","daemon off;"]



