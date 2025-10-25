FROM debian:stable-slim

RUN apt -y install wget gnupg2 && \
    wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | gpg --dearmor -o /usr/share/keyrings/google_keyring.gpg && \
    sh -c 'echo "deb [arch=amd64 signed-by=/usr/share/keyrings/google_keyring.gpg] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list' && \
    apt -y update && \
    apt install -y google-chrome-stable

RUN mkdir -p /app/bin ; mkdir -p /app/conf

COPY ./bin /app/bin/

EXPOSE 9150/tcp

VOLUME /app/conf

CMD ["/app/bin/ambot", "--app.config", "/app/conf/config.yaml"]
