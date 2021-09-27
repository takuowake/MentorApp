FROM golang:latest

RUN mkdir /go/src/MentorApp

WORKDIR /go/src/MentorApp

ADD . /go/src/MentorApp

# Perform a hot-reload
RUN GO111MODULE=off go get -u github.com/oxequa/realize
RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
CMD ["realize", "start"]

