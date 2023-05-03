FROM redis
COPY ./config/redis.conf /redis.conf
CMD ["redis-server", "/redis.conf"]
FROM loadimpact/k6
COPY ./scripts /scripts
ENV K6_OUT influxdb=http://influxdb:8086/k6
CMD ["k6", "run", "/scripts"]
FROM influxdb:1.8.10
ENV INFLUXDB_DB k6
ENV INFLUXDB_HTTP_MAX_BODY_SIZE 0
FROM grafana/grafana
ENV GF_AUTH_ANONYMOUS_ENABLED true
ENV GF_AUTH_ANONYMOUS_ORG_ROLE Admin
FROM mariadb
ENV MARIADB_ROOT_PASSWORD 1234
ENV MARIADB_DATABASE customer
COPY ./data /var/lib/
EXPOSE 6379 8086 3000 3306


FROM golang:1.20.3
WORKDIR /app
COPY . /app
RUN go clean --modcache
RUN go mod download
RUN go build -o main .

EXPOSE 5000

CMD ["./main"]


