FROM golang:latest 
RUN mkdir /app 
ADD . /app/
WORKDIR /app
RUN make dev-deps
RUN apt-get update && apt-get -y install netcat-openbsd ;
RUN alias nc='netcat'
CMD ["./scripts/start.sh"]