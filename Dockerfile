FROM python:3.11-slim-bookworm

RUN apt-get -y update && \
    apt-get -y upgrade && \
    apt-get -y install wget gnupg2
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

ENV USERNAME=""
ENV PASSWORD=""
ENV FUEL_GOOD_PRICE=450
ENV CO2_GOOD_PRICE=120
ENV FUEL_BUDGET_PERCENT=70
ENV MAINTANANCE_BUDGET_PERCENT=50
ENV MARKETING_BUDGET_PERCENT=70
ENV AIRCRAFT_WEAR_PERCENT=30
ENV AIRCRAFT_MAX_HOURS_TO_ACHECK=12
ENV RUN_MODE="once"
ENV SERVICE_SLEEP_SEC=300


CMD python ./app.py \
    --username="${USERNAME}" \
    --password="${PASSWORD}" \
    --fuel-good-price=${FUEL_GOOD_PRICE} \
    --co2-good-price=${CO2_GOOD_PRICE} \
    --fuel-budget-percent=${FUEL_BUDGET_PERCENT} \
    --maintenance-budget-percent=${MAINTANANCE_BUDGET_PERCENT} \
    --marketing-budget-percent=${MARKETING_BUDGET_PERCENT} \
    --aircraft-wear-percent=${AIRCRAFT_WEAR_PERCENT} \
    --aircraft-max-hours-to-acheck=${AIRCRAFT_MAX_HOURS_TO_ACHECK} \
    --run-mode=${RUN_MODE} \
    --service-sleep-sec=${SERVICE_SLEEP_SEC}
