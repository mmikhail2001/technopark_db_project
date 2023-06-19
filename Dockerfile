FROM golang:1.20-alpine3.18 as golang
COPY . /build
WORKDIR /build
RUN go build -o main cmd/app/app.go

FROM postgres as postgres

COPY --from=golang /build/main /build/main
COPY --from=golang /build/config/config.yml /build/config/config.yml
COPY --from=golang /build/db/db.sql /build/db/db.sql

COPY db/db.sql /docker-entrypoint-initdb.d/db.sql
COPY run.sh /docker-entrypoint-initdb.d/run.sh

RUN chmod 777 /docker-entrypoint-initdb.d/run.sh

ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB postgres

EXPOSE 5000

