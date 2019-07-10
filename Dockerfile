FROM golang:latest
COPY . /app
WORKDIR /app
ENV GOPROXY=https://goproxy.io
RUN [ "go", "build"]
EXPOSE 8080
CMD [ "./go-forum" ]