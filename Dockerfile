# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

RUN mkdir /home/app

COPY . /home/app/

WORKDIR /home/app

RUN chmod +x install.sh

RUN ./install.sh

RUN go build  -o hts .

CMD [ "/home/app/hts" ]
