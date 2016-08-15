FROM centurylink/ca-certs
ADD bin/chieftan /

ENTRYPOINT /chieftan

ENV MONGODB_URL="mongodb://localhost:27017/chieftan"
ENV PORT=3000
EXPOSE $PORT

ARG VERSION="development"
LABEL VERSION=$VERSION

CMD ["server"]