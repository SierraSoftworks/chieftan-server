FROM scratch
ADD ca-certificates.crt /etc/ssl/certs
ADD chieftan /

ENTRYPOINT /chieftan

ENV MONGODB_URL="mongodb://localhost:27017/chieftan"
ENV PORT=3000
EXPOSE $PORT

CMD ["server"]