FROM golang:1.19

RUN mkdir -p /home/webapp
WORKDIR /home/webapp
COPY . /home/webapp
RUN go build -o app
CMD ./app
# CMD /bin/bash -c "while true; do sleep 1; done"
