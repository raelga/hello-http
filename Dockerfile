FROM golang:1.12 as build
LABEL maintainer="Rael Garcia <rael@rael.io>"
WORKDIR /build
COPY . .
ENV GOOS linux
RUN go build -a -o app .

FROM scratch
COPY --from=build /build/app /
CMD [ "./app" ]