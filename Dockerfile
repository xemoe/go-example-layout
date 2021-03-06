FROM golang:1.15

WORKDIR /go/src/app

COPY . .
RUN go get -d -v ./...
RUN chmod -R 777 /go/pkg

ARG USER_ID
ARG GROUP_ID

RUN addgroup --gid $GROUP_ID user
RUN adduser --disabled-password --gecos '' --uid $USER_ID --gid $GROUP_ID user

USER user
