FROM maven:3.8.6-openjdk-8 AS MAVEN_BUILD

COPY ./ ./

RUN pwd

RUN ls -al ./

RUN mvn clean package assembly:single

FROM openjdk:8-jre-slim-buster

COPY --from=MAVEN_BUILD /target/am4bot-jar-with-dependencies.jar /app/am4bot.jar

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

COPY src/main/resources/log4j2.properties /app

WORKDIR /app

VOLUME [ "/app" ]

ENV AM4_USERNAME=""
ENV AM4_PASSWORD=""
ENV FUEL_GOOD_PRICE=500
ENV CO2_GOOD_PRICE=120
ENV FUEL_BUDGET_PERCENT=70
ENV MAINTANANCE_BUDGET_PERCENT=50
ENV MARKETING_BUDGET_PERCENT=70
ENV AIRCRAFT_WEAR_PERCENT=80
ENV AIRCRAFT_MAX_HOURS_TO_ACHECK=24
ENV RUN_MODE="once"
ENV SERVICE_SLEEP_SEC=300
ENV SCANNER_FILE="am4scanner.csv"

CMD java -cp am4bot.jar -Dlog4j.configurationFile=log4j2.properties com.ashokhin.am4bot.App
