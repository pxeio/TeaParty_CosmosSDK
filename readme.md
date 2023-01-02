# partychain
**partychain** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/TeaPartyCrypto/partychain@latest! | sudo bash
```
`TeaPartyCrypto/partychain` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


export CELO_RPC_1=https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e
export CELO_RPC_2=https://celo-alfajores.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e
export MO_RPC_1="https://rpc.mo-scout.com"
export MO_RPC_2="https://rpc.mo-scout.com"
export ETH_RPC_1="https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e"
export ETH_RPC_2="https://goerli.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e"
export POLY_RPC_1="https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e"
export POLY_RPC_2="https://polygon-mumbai.infura.io/v3/61979797a8bb4bfe9dddd4ff9675db7e"
export SOL_RPC_1=https://api.testnet.solana.com
export SOL_RPC_2=https://api.testnet.solana.com
export DEV=true
export WATCH=true
export PARTY_CHAIN_1="127.0.0.1:9090"
