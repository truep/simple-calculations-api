FROM golang:latest as builder

ENV APP_HOME /app

WORKDIR $APP_HOME
COPY . .

RUN make

FROM alpine:3.14 


COPY --from=builder /app/bin/calc-api ./

CMD [ "./calc-api" ]
