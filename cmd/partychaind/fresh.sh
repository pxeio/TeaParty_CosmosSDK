rm -rf ~/.partychain
rm -rf ./partychaind
go build .
./partychaind init party
./partychaind keys add alice  --keyring-backend test
./partychaind add-genesis-account alice 500000000stake
./partychaind gentx alice  100000000stake   --keyring-backend test
./partychaind collect-gentxs
sed -i '117s/false/true/' /Users/jeffreynaef/.partychain/config/app.toml
sed -i '138s/false/true/' /Users/jeffreynaef/.partychain/config/app.toml
./partychaind start --log_level error

