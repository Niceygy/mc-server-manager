FROM amazoncorretto:17.0.7-alpine
WORKDIR /var/mc/server/
RUN apk update
RUN apk add curl
EXPOSE 25565
CMD ["/var/mc/server/start.sh"]