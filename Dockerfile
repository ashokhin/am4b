FROM python:3.11-slim-bookworm

RUN apt-get -y update && apt-get -y upgrade && apt-get -y install wget gnupg2
# Adding trusting keys to apt for repositories
RUN wget -q -O - https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -

# Adding Google Chrome to the repositories
RUN sh -c 'echo "deb [arch=amd64] http://dl.google.com/linux/chrome/deb/ stable main" >> /etc/apt/sources.list.d/google-chrome.list'

# Updating apt to see and install Google Chrome
RUN apt-get -y update
RUN apt-get install -y google-chrome-stable

COPY . /app
WORKDIR /app
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

ENV AM4_USERNAME=""
ENV AM4_PASSWORD=""

CMD python ./app.py --username="${AM4_USERNAME}" --password="${AM4_PASSWORD}"
