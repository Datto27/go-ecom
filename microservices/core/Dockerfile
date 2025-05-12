FROM golang:1.23.4-bullseye

WORKDIR /go/app

COPY Makefile ./

COPY . .

# Install make
RUN apt-get update && apt-get install -y make

# application api port
EXPOSE 4000

CMD ["make", "run", "start"]