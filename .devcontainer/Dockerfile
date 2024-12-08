FROM docker:26.0.0-dind

RUN apk update && apk add curl

COPY --from=golang:1.23-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
