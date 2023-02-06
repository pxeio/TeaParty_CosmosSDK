Cosmos faucet
https://github.com/tendermint/faucetcurl https://get.starport.network/faucet! | bash 

export PATH=/home/jeff/partychain:$PATH
export PATH=/usr/local/bin:$PATH
export ACCOUNT_NAME=alice
export DENOMS=stake
export CREDIT_AMOUNT=100
export NODE='http://localhost:26657'
export KEYRING_BACKEND=test
export HOME=~/.partychain
nohup faucet --cli-name partychaind --denoms stake --credit-amount 100 --account-name alice --node http://localhost:26657 --keyring-backend test &



curl -X POST -d '{"address": "party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr", "coins": ["10stake"]}' http://localhost:8000

curl -X POST -d '{"address": "party1hedrjppvwuzvpxj2w5rp6xg3gqc0mx7klk8dwl", "coins": ["1000000stake"]}' http://144.126.148.104:8000

