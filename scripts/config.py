# -*- coding:utf-8 -*-
import json
import os
import sys

# Localhost
url = 'http://123.60.220.55:8545'
chain_id = 1337
gas_limit = 3000000 # 您可能需要根据合约函数的复杂性和资源消耗进行调整
public_to_private_keys = {
    '0xD8f818a8680C0A1ba80Df6A4a0Dfcd9b35466715' :'0xc1226cc783efedaea3eb7a35cf56e76a4ff1d49d1760130919b8540c82ac6db1'
}

# # Sepolia
# url = 'https://eth-sepolia.g.alchemy.com/v2/j1yrdLvznv5AQ5NfphKOQZsFDU7-Jc8W'
# chain_id = 11155111
#
# # opBNB test_net
# url = 'https://opbnb-testnet-rpc.bnbchain.org'
# chain_id = 5611
#
# # Scroll Sepolia
# url = 'https://sepolia-rpc.scroll.io'
# chain_id = 534351

# Taiko Jolnir
url = 'https://rpc.jolnir.taiko.xyz'
chain_id = 167007
public_to_private_keys = {
    '0xD275E84eb1967f6DA2c475BbA1D312775ECEC21C' :'54bc1f7294e6378af87c8cb734ee98d516e450ac7259002a9c72faff569ef94c'
}


class _role:
    _contract_owner = '0xD275E84eb1967f6DA2c475BbA1D312775ECEC21C'
    _provider = '0xD275E84eb1967f6DA2c475BbA1D312775ECEC21C'
    _user = '0xD275E84eb1967f6DA2c475BbA1D312775ECEC21C'

    @classmethod
    def private_key(cls, public_key):
        return public_to_private_keys.get(public_key, None)

    @property
    def contract_owner(self):
        return type("owner", (), dict(public_key=self._contract_owner, private_key=self.private_key(self._contract_owner)))

    @property
    def provider(self):
        return type("provider", (), dict(public_key=self._provider, private_key=self.private_key(self._provider)))

    @property
    def user(self):
        return type("user", (), dict(public_key=self._user, private_key=self.private_key(self._user)))


role = _role()


def get_config(fileName):
    return json.load(open(os.path.join(os.getcwd(), "build/contract_address", fileName)))['address'], json.load(open(os.path.join(os.getcwd(), "build/contracts", fileName)))['abi']

market_contract_address, market_abi = get_config("Market.json")
apus_task_address, apus_task_abi = get_config('ApusProofTask.json')
token_abi = json.load(open(os.path.join(os.getcwd(), "build/contracts", 'ERC20.json')))['abi']
print("market address", market_contract_address)

# market_contract_address = env['APUS_TASK_CONTRACT_ADDRESS']

__all__ = ['role', 'url', 'chain_id', 'gas_limit', 'market_contract_address', 'market_abi', 'apus_task_address', 'apus_task_abi', 'token_abi']

if __name__ == '__main__':
    print(role.provider.public_key)
    print(role.provider.private_key)
