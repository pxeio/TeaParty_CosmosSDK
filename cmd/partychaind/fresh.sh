rm -rf ~/.partychain
./partychaind init party 
./partychaind keys add alice  --keyring-backend test 
./partychaind add-genesis-account alice 5000000000000000000000stake 
./partychaind gentx alice  100000000stake   --keyring-backend test
./partychaind collect-gentxs
go run .  start --log_level error
