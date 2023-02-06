rm -rf ~/.partychain
rm -rf ./partychaind
go build .
./partychaind init party
./partychaind keys add alice  --keyring-backend test
./partychaind add-genesis-account alice 5000000000000000000000stake
./partychaind gentx alice  100000000stake   --keyring-backend test
./partychaind collect-gentxs
./partychaind start --log_level error

