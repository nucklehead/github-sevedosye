FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go get ./... || exit 0
RUN go build -o github-sevedosye .
CMD ["/app/github-sevedosye", "-p", "$PORT", "-t", "$GITHUB_TOKEN"]
