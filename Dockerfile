FROM alpine

apk add --update tini
ENTRYPOINT ["/sbin/tini", "--"]

ADD bin/chieftan /bin/chieftan

ENV MONGODB_URL="mongodb://localhost:27017/chieftan"
ENV PORT=3000
EXPOSE $PORT

ARG VERSION="development"
LABEL VERSION=$VERSION

WORKDIR /
ENTRYPOINT /bin/chieftan
CMD ["chieftan","server"]