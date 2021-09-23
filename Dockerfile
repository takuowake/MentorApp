FROM golang:latest

RUN mkdir /go/src/MentorApp

WORKDIR /go/src/MentorApp

ADD . /go/src/MentorApp

RUN GO111MODULE=off go get -u github.com/oxequa/realize
CMD ["realize", "start"]
