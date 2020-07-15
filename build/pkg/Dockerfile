FROM golang:latest as builder
RUN mkdir yaya
COPY ./ /yaya
#RUN cd /yaya \
#    & GOPROXY="https://athens.azurefd.net" \
#    & go mod tidy
WORKDIR /yaya/cmd
RUN  GOOS=linux GOARCH=amd64 go build -o yaya .


FROM alpine:latest
EXPOSE 8080
COPY --from=builder /yaya/cmd/yaya /opt/
RUN ls -a /opt

ENTRYPOINT ["/bin/sh" ,"/opt/yaya"]