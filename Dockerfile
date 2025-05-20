FROM debian:stable-slim

RUN apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y install wget gnupg2 && \
    wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add - && \
    sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list' && \
    apt-get -y update && \
    apt-get install -y google-chrome-stable

RUN mkdir -p /app/bin ; mkdir -p /app/conf

COPY ./bin /app/bin/

EXPOSE 9150/tcp

VOLUME /app/conf

CMD ["/app/bin/ambot", "--app.config", "/app/conf/config.yaml"]
