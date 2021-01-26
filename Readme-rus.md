Пакет предназначен для интеграционного тестирования смарт-контрактов Nebula, IB-Port, LU-Port.

Требования к хосту: docker, docker-compose

Поддерживаемые сети:
- Ethereum testnet ROPSTEN
- Ethereum private net (поднимается в контейнере docker)
- TRON testnet SHASTA
- WAVES testnet "stagenet"
- WAVES private net (поднимается в контейнере docker)

Исходные коды контрактов должны быть помещены в соответствующую папку:
- solidity/0.7/Nebula|Token|IBPort|LUPort - для Ethereum testnet и privatenet
- solidity/tron/Nebula|Token|IBPort|LUPort - для TRON
- waves/scripts/gravity|nebula (TODO - сейчас требуется уже скомпилированный base64, компиляция из исходного кода - в прогрессе)

При использовании testnet может потребоваться предварительно профинансировать адреса тестовыми монетами:
- для WAVES testnet STAGENET: 3MUBYCv4iHmmCA16zZ8GzKTSc3Fj1umn8Jb
- для Ethereum testnet ROPSTEN: 0x90F8bf6A479f320ead074411a4B0e7944Ea8c9C1
- для TRON testnet SHASTA: THxQonMCidbEgQSFxWcEgkyw1YfdpiqGz6

Далее необходимо запустить bash-скрипт test-*.sh с соответствующим названием,
например, test-eth-private.sh для тестирования контрактов в ethereum private network

В пакете уже содержится скомпилированный код смарт-контрактов, при изменении исходного кода запустить перекомпиляцию recompile.sh
Скрипт компиляции произведет сборку в контейнере solidity compiler версий 5 и 7, а также утилиты abigen.
Это может занять продолжительное время при первом запуске.

