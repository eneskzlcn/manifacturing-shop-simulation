FROM    golang:1.18-alpine3.16
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go mod tidy -go=1.18

ENTRYPOINT ["go", "run", "./cmd/simulation/main.go"]