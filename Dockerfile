FROM node:6.2-slim

MAINTAINER Benjamin Pannell <admin@sierrasoftworks.com>

RUN useradd --create-home --user-group --shell /bin/false app
ENV HOME=/home/app

COPY package.json $HOME/chieftan/package.json
RUN cd $HOME/chieftan && npm install --silent && npm cache clean --silent

ARG VERSION
LABEL version=${VERSION:-development}

ADD ./ $HOME/chieftan/
RUN cd $HOME/chieftan && npm run build --silent
RUN chown -R app:app $HOME/chieftan

USER app

ENV PORT=3000
EXPOSE $PORT

HEALTHCHECK --interval=10s --timeout=5s --retries=3 \
    CMD curl -f http://localhost:$PORT/api/v1/status || exit 2

ENV CHIEFTAN_MONGODB="mongodb://mongodb/chieftan"

WORKDIR $HOME/chieftan
ENTRYPOINT ["npm", "run"]
CMD ["start"]