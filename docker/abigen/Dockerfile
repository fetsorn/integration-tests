FROM golang
RUN git clone https://github.com/ethereum/go-ethereum.git
RUN cd go-ethereum/cmd/abigen; go build
RUN cp go-ethereum/cmd/abigen/abigen /usr/local/bin
RUN git clone https://github.com/ethereum/solidity.git
RUN apt update -y; apt install -y sudo
RUN cd solidity; ./scripts/install_deps.sh
RUN cd solidity/scripts; git checkout tags/v0.7.1; ./build.sh; mv /usr/local/bin/solc /usr/local/bin/solc7
RUN cd solidity/scripts; git checkout tags/v0.5.17; ./build.sh; mv /usr/local/bin/solc /usr/local/bin/solc5

