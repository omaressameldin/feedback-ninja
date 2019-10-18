ARG WORKDIR=/usr/src/app
ARG APP_SRC=./app
ARG BUILD_FILE=feedback-ninja

FROM golang:1.12.4-alpine
ARG BUILD_FILE
ARG WORKDIR
ARG APP_SRC

WORKDIR $WORKDIR

COPY ./go.mod .
COPY ./go.sum .

RUN apk add git &&\
    apk --update add ca-certificates &&\
    go mod download

COPY ${APP_SRC}/. ${APP_SRC}/.
RUN CGO_ENABLED=0 GOOS=linux go test ./... &&\
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $BUILD_FILE ${APP_SRC}/.


FROM bash:4.3.48
ARG BUILD_FILE
ENV BUILD_FILE $BUILD_FILE
ARG WORKDIR

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=0 ${WORKDIR}/${BUILD_FILE} $BUILD_FILE

RUN addgroup -S appuser && adduser -S appuser -G appuser -u 1000 &&\
    chown -R appuser $BUILD_FILE
USER appuser

CMD ./$BUILD_FILE
