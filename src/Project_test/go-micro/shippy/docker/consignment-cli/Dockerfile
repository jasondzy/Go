FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY ./consignment-cli  /app/
COPY ./consignment.json /app/

ENTRYPOINT [ "./consignment-cli" ]

