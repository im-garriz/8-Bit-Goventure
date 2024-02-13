FROM ubuntu:latest

# Set environment variables for Go
ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH $GOPATH/bin:$GOROOT/bin:$PATH

RUN apt-get update && \
    apt-get install -y wget git

WORKDIR /tmp

RUN wget https://go.dev/dl/go1.21.6.linux-amd64.tar.gz
RUN rm -rf /usr/local/go 
RUN tar -C /usr/local -xzf go1.21.6.linux-amd64.tar.gz
# RUN mv go /usr/local && \
RUN rm go1.21.6.linux-amd64.tar.gz

WORKDIR $GOPATH
