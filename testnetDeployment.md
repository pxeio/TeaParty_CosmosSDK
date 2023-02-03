
./partychaind init party 

./partychaind keys --keyring-backend file add mac  

- address: party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr
  name: mac
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1cpf1FuBy/6+3MxENXZzKVJo3ICjVrpa+HaikjCOALN"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

sunset sing canal frozen friend stone neither party own decorate spider network suit boost number opera dune word loud win similar knee flush special


./partychaind add-genesis-account party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr 5000000000000000000000stake 
./partychaind add-genesis-account party1wn9w5w3lx0q3r55xt80zkclzv3lzxlrf05d3qa 5000000000000000000000stake 

./partychaind gentx mac  100000000stake  --keyring-backend file
K!1poiv^F6fs8XtOdxwqRpp6

./partychaind collect-gentxs


// on another computer i ran 
- address: party1wn9w5w3lx0q3r55xt80zkclzv3lzxlrf05d3qa
  name: alice
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A937d10es42OTxzDE79cxbRIURx4nQo4K5AnhZ+yS0JV"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

nation rare split giggle parade grow above clutch various venue muffin drop exotic scare image favorite speak dolphin good eight tuition glimpse clown alone


Updated the `config.toml` on a peer line 212 seed nodes to connect to 
seeds = "1c82a67c512a7172e40cefab76f679907498cdeb@209.126.11.245:26656"

{
  "genesis_time": "2023-02-03T15:39:17.252786356Z",
  "chain_id": "partychain",
  "initial_height": "1",
  "consensus_params": {
    "block": {
      "max_bytes": "22020096",
      "max_gas": "-1",
      "time_iota_ms": "1000"
    },
    "evidence": {
      "max_age_num_blocks": "100000",
      "max_age_duration": "172800000000000",
      "max_bytes": "1048576"
    },
    "validator": {
      "pub_key_types": [
        "ed25519"
      ]
    },
    "version": {}
  },
  "app_hash": "",
  "app_state": {
    "auth": {
      "params": {
        "max_memo_characters": "256",
        "tx_sig_limit": "7",
        "tx_size_cost_per_byte": "10",
        "sig_verify_cost_ed25519": "590",
        "sig_verify_cost_secp256k1": "1000"
      },
      "accounts": [
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "party1wn9w5w3lx0q3r55xt80zkclzv3lzxlrf05d3qa",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        },
        {
          "@type": "/cosmos.auth.v1beta1.BaseAccount",
          "address": "party1hedrjppvwuzvpxj2w5rp6xg3gqc0mx7klk8dwl",
          "pub_key": null,
          "account_number": "0",
          "sequence": "0"
        }
      ]
    },
    "authz": {
      "authorization": []
    },
    "bank": {
      "params": {
        "send_enabled": [],
        "default_send_enabled": true
      },
      "balances": [
        {
          "address": "party1wn9w5w3lx0q3r55xt80zkclzv3lzxlrf05d3qa",
          "coins": [
            {
              "denom": "stake",
              "amount": "5000000000000000000000"
            }
          ]
        },
        {
          "address": "party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr",
          "coins": [
            {
              "denom": "stake",
              "amount": "5000000000000000000000"
            }
          ]
        },
        {
          "address": "party1hedrjppvwuzvpxj2w5rp6xg3gqc0mx7klk8dwl",
          "coins": [
            {
              "denom": "stake",
              "amount": "5000000000000000000000"
            }
          ]
        }
      ],
      "supply": [],
      "denom_metadata": []
    },
    "capability": {
      "index": "1",
      "owners": []
    },
    "crisis": {
      "constant_fee": {
        "denom": "stake",
        "amount": "1000"
      }
    },
    "distribution": {
      "params": {
        "community_tax": "0.020000000000000000",
        "base_proposer_reward": "0.010000000000000000",
        "bonus_proposer_reward": "0.040000000000000000",
        "withdraw_addr_enabled": true
      },
      "fee_pool": {
        "community_pool": []
      },
      "delegator_withdraw_infos": [],
      "previous_proposer": "",
      "outstanding_rewards": [],
      "validator_accumulated_commissions": [],
      "validator_historical_rewards": [],
      "validator_current_rewards": [],
      "delegator_starting_infos": [],
      "validator_slash_events": []
    },
    "evidence": {
      "evidence": []
    },
    "feegrant": {
      "allowances": []
    },
    "genutil": {
      "gen_txs": [
        {
          "body": {
            "messages": [
              {
                "@type": "/cosmos.staking.v1beta1.MsgCreateValidator",
                "description": {
                  "moniker": "party",
                  "identity": "",
                  "website": "",
                  "security_contact": "",
                  "details": ""
                },
                "commission": {
                  "rate": "0.100000000000000000",
                  "max_rate": "0.200000000000000000",
                  "max_change_rate": "0.010000000000000000"
                },
                "min_self_delegation": "1",
                "delegator_address": "party14jxp6ffh7sr9mh5frekq93gkqy7tqknxqf2yhr",
                "validator_address": "partyvaloper14jxp6ffh7sr9mh5frekq93gkqy7tqknx0craq6",
                "pubkey": {
                  "@type": "/cosmos.crypto.ed25519.PubKey",
                  "key": "qP+7QxcIUyJU8NtwJ4ZwOHgfLVVbUJ9OpxxSIWlEh5o="
                },
                "value": {
                  "denom": "stake",
                  "amount": "100000000"
                }
              }
            ],
            "memo": "1c82a67c512a7172e40cefab76f679907498cdeb@209.126.11.245:26656",
            "timeout_height": "0",
            "extension_options": [],
            "non_critical_extension_options": []
          },
          "auth_info": {
            "signer_infos": [
              {
                "public_key": {
                  "@type": "/cosmos.crypto.secp256k1.PubKey",
                  "key": "A1cpf1FuBy/6+3MxENXZzKVJo3ICjVrpa+HaikjCOALN"
                },
                "mode_info": {
                  "single": {
                    "mode": "SIGN_MODE_DIRECT"
                  }
                },
                "sequence": "0"
              }
            ],
            "fee": {
              "amount": [],
              "gas_limit": "200000",
              "payer": "",
              "granter": ""
            },
            "tip": null
          },
          "signatures": [
            "+LpzA2mThKXRo0mzVTRewHbTaGmqqRu+NhMzY7jvYMJlb/ZvxQRdeAjQ7LkOG8r4LOS/Uq7uzt374W3Fj47B5g=="
          ]
        }
      ]
    },
    "gov": {
      "starting_proposal_id": "1",
      "deposits": [],
      "votes": [],
      "proposals": [],
      "deposit_params": {
        "min_deposit": [
          {
            "denom": "stake",
            "amount": "10000000"
          }
        ],
        "max_deposit_period": "172800s"
      },
      "voting_params": {
        "voting_period": "172800s"
      },
      "tally_params": {
        "quorum": "0.334000000000000000",
        "threshold": "0.500000000000000000",
        "veto_threshold": "0.334000000000000000"
      }
    },
    "group": {
      "group_seq": "0",
      "groups": [],
      "group_members": [],
      "group_policy_seq": "0",
      "group_policies": [],
      "proposal_seq": "0",
      "proposals": [],
      "votes": []
    },
    "ibc": {
      "client_genesis": {
        "clients": [],
        "clients_consensus": [],
        "clients_metadata": [],
        "params": {
          "allowed_clients": [
            "06-solomachine",
            "07-tendermint"
          ]
        },
        "create_localhost": false,
        "next_client_sequence": "0"
      },
      "connection_genesis": {
        "connections": [],
        "client_connection_paths": [],
        "next_connection_sequence": "0",
        "params": {
          "max_expected_time_per_block": "30000000000"
        }
      },
      "channel_genesis": {
        "channels": [],
        "acknowledgements": [],
        "commitments": [],
        "receipts": [],
        "send_sequences": [],
        "recv_sequences": [],
        "ack_sequences": [],
        "next_channel_sequence": "0"
      }
    },
    "interchainaccounts": {
      "controller_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "ports": [],
        "params": {
          "controller_enabled": true
        }
      },
      "host_genesis_state": {
        "active_channels": [],
        "interchain_accounts": [],
        "port": "icahost",
        "params": {
          "host_enabled": true,
          "allow_messages": []
        }
      }
    },
    "mint": {
      "minter": {
        "inflation": "0.130000000000000000",
        "annual_provisions": "0.000000000000000000"
      },
      "params": {
        "mint_denom": "stake",
        "inflation_rate_change": "0.130000000000000000",
        "inflation_max": "0.200000000000000000",
        "inflation_min": "0.070000000000000000",
        "goal_bonded": "0.670000000000000000",
        "blocks_per_year": "6311520"
      }
    },
    "params": null,
    "party": {
      "params": {},
      "tradeOrdersList": [],
      "pendingOrdersList": [],
      "ordersAwaitingFinalizerList": [],
      "ordersUnderWatchList": []
    },
    "slashing": {
      "params": {
        "signed_blocks_window": "100",
        "min_signed_per_window": "0.500000000000000000",
        "downtime_jail_duration": "600s",
        "slash_fraction_double_sign": "0.050000000000000000",
        "slash_fraction_downtime": "0.010000000000000000"
      },
      "signing_infos": [],
      "missed_blocks": []
    },
    "staking": {
      "params": {
        "unbonding_time": "1814400s",
        "max_validators": 100,
        "max_entries": 7,
        "historical_entries": 10000,
        "bond_denom": "stake",
        "min_commission_rate": "0.000000000000000000"
      },
      "last_total_power": "0",
      "last_validator_powers": [],
      "validators": [],
      "delegations": [],
      "unbonding_delegations": [],
      "redelegations": [],
      "exported": false
    },
    "transfer": {
      "port_id": "transfer",
      "denom_traces": [],
      "params": {
        "send_enabled": true,
        "receive_enabled": true
      }
    },
    "upgrade": {},
    "vesting": {}
  }
}