FROM golang:1.18.3-alpine

WORKDIR /usr/app/src

COPY . .

RUN go install
RUN go build -o main .

CMD [ "/usr/app/src/main" ]