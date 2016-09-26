FROM alpine

RUN apk add --update tini
ENTRYPOINT ["/sbin/tini", "--"]

ADD bin/chieftan /bin/chieftan

ENV MONGODB_URL="mongodb://localhost:27017/chieftan"
ENV PORT=3000
EXPOSE $PORT

ARG VERSION="development"
LABEL VERSION=$VERSION

WORKDIR /bin
CMD ["chieftan","server", "--log-level", "DEBUG"]