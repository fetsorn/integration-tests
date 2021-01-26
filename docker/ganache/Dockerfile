FROM node:12.6

RUN npm install --unsafe-perm=true --allow-root -g ganache-cli

CMD [ "ganache-cli", "-d", "-mnemonic", "m", "-h", "0.0.0.0" ]
