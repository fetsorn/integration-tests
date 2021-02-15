# Build Geth in a stock Go builder container
FROM golang:1.14-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers git

RUN mkdir /heco
COPY ./heco /heco
WORKDIR /heco
RUN make geth

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /heco/build/bin/geth /usr/local/bin/

EXPOSE 8545 8546 8547 30303 30303/udp
COPY genesis.json .
RUN geth --datadir test-chain-dir init genesis.json
CMD [ "geth", "--datadir", "test-chain-dir", "--rpc", "--rpcaddr", "0.0.0.0", "--rpcvhosts", "*", "--rpccorsdomain", "*", "--nodiscover", "--maxpeers", "0"] 
