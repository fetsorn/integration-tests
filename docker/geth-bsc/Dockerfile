# for manual build (not used with the docker image)
FROM golang

RUN git clone https://github.com/binance-chain/bsc
WORKDIR /bsc
COPY . /bsc
RUN make geth

CMD [ "./build/bin/geth", "--datadir", "test-chain-dir", "--rpc", "--dev", "--rpcaddr", "0.0.0.0"]
