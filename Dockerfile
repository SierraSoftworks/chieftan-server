FROM node:latest

MAINTAINER Benjamin Pannell <admin@sierrasoftworks.com>

RUN useradd --create-home --user-group --shell /bin/false app
ENV HOME=/home/app

COPY package.json $HOME/chieftan/package.json
RUN cd $HOME/chieftan && npm install && npm cache clean

ARG VERSION
LABEL version=${VERSION:-development}

ADD ./ $HOME/chieftan/
RUN cd $HOME/chieftan && npm run build
RUN chown -R app:app $HOME/chieftan

USER chieftan

HEALTHCHECK --interval=10s --timeout=5s --retries=3 \
    CMD curl -f http://localhost/api/v1/status || exit 2

ENV CHIEFTAN_MONGODB="mongodb://mongodb/chieftan"

WORKDIR $HOME/chieftan
ENTRYPOINT node
CMD ["dist/index.js"]