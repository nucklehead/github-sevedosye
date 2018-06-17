FROM golang:latest
RUN mkdir /app
RUN chmod a+rw /app
ADD . /app/
WORKDIR /app
RUN go get ./... || exit 0
RUN go build -o github-sevedosye .
CMD ["sh", "-c", "/app/github-sevedosye -p ${PORT} -t ${GITHUB_TOKEN}"]
