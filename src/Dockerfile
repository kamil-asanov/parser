FROM golang:1.21-alpine
ENV LANGUAGE="en"
COPY . /opt/src/
WORKDIR /opt/src
RUN apk add --no-cache ca-certificates &&\
    chmod +x /opt/src
RUN go mod init parser/src
RUN go get github.com/Syfaro/telegram-bot-api && go get github.com/gocolly/colly && go get github.com/lib/pq
RUN go build 
EXPOSE 80/tcp
CMD [ "./src" ]