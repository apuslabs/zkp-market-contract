import json
import sys
import random
import time

import web3
import argparse
from scripts.config import role
from scripts.conn import *

none_address = '0x0000000000000000000000000000000000000000'


def gen_client_config(owner, client_id, url, max_instance, min_fee):
    return {
        'owner': owner,
        'id': client_id,  # 这是一个示例ID
        'url': url,  # 这是一个示例URL
        'minFee': min_fee,
        'maxZkEvmInstance': max_instance,
        'curInstance': 0,  # 这是一个示例的当前实例数
        'stat': 0
    }


class ContractLib:

    def join_market(self, prover, client_config):
        tx_hash = transaction(prover, market_contract.functions.joinMarket(client_config))
        return tx_hash

    def getLowestN(self):
        return market_contract.functions.getLowestN().call()

    def getProverConfig(self, addr, cid):
        return market_contract.functions.getProverConfig(addr, cid).call()

    def close_prover_client(self, prover, client_id, addr):

        tx_hash = transaction(prover, market_contract.functions.offlineClient(addr, client_id))
        return tx_hash

    def post_task(self, user, uniq_id):
        return transaction(user, apus_task_contract.functions.postTask(0, uniq_id, b"hello world! input", int(time.time()) + 90 * 60, dict(token=none_address, amount=10)))

    def dispatchTaskToClient(self, user, block_id):
        tx_hash2 = transaction(user, apus_task_contract.functions.dispatchTaskToClient(block_id))
        return tx_hash2

    def get_task(self, block_id):
        return apus_task_contract.functions.getTask(0, block_id).call()

    def get_client_by_index(self, index):
        return market_contract.functions.clients(index).call()

    def get_task_by_index(self, index):
        return apus_task_contract.functions.tasks(index).call()

    def submit_task(self, user, task_id):
        tx_hash2 = transaction(user, apus_task_contract.functions.submitTask(0, task_id, b"8b6ffb96d2377872afd4998fb8329183abb8a321bde906059c8fa4643040728a"))
        return tx_hash2

    def reward_task(self, user, tid, token_address, value):
        value = int(value * 10 ** 18)
        f = apus_task_contract.functions
        values = dict(value=value)
        if token_address != none_address:
            print("--------approve-----------")
            print(token_approve(user, token_address, apus_task_address, value)['status'])

        f = f.rewardTask(int(tid), value, token_address)
        tx_hash2 = transaction(user, f, **values)
        return tx_hash2
    
    def sync_task(self, user, clientId, uniqID, rewardAmount, assignTime, proveTime, _stat, _assigner, _rewardRecipient):
        tx = transaction(user, apus_task_contract.functions.syncTask(clientId, uniqID, _assigner, _stat, rewardAmount, _rewardRecipient, assignTime, proveTime))

        return tx

    def market_dispatch(self, user, addr, cid):
        return transaction(user, market_contract.functions.dispatchTaskToClient(addr, cid))

    def has_resource(self):
        return apus_task_contract.functions.hasResource().call()



connector = ContractLib()


def creat_client_2():
    # provider 0xC2600C80Beb521CC4E2f1b40B9D169c46E391390
    client_config = gen_client_config(role.provider.public_key, 22, 'http://ec2-18-209-35-10.compute-1.amazonaws.com', 1, 10)

    print("-" * 10, "加入market client", "-" * 10)
    tx = connector.join_market(role.provider, client_config)
    print(tx['status'])

def create_client():
    servers = [
     ['0x74e2A876b86d7C345765d97D0E46FF9B84575F6a', 11261823312923665969, 'http://117.187.208.214:9000', 10, 1, 0, 1]]

    # provider 0xC2600C80Beb521CC4E2f1b40B9D169c46E391390
    for server in servers:
        client_config = gen_client_config(server[0], server[1], server[2], 1, 10)

        print("-" * 10, "加入market client", "-" * 10)
        tx = connector.join_market(role.provider, client_config)
        print(tx['status'])

        print("-" * 10, "获取价格最低client", "-" * 10)
        result = connector.getLowestN()
        print(result)


task_id = 1344635


def post_task():
    print("-" * 10, "推送task", "-" * 10)
    tx = connector.post_task(role.provider, task_id)
    print(tx)
    print(tx['status'])


def get_task():
    print("-" * 10, "获取task & client", "-" * 10)
    result = connector.get_task(task_id)
    print(result)


def dispatch_task(tid):
    print("-" * 10, "分配机器", "-" * 10)
    tx = connector.dispatchTaskToClient(role.provider, tid)
    print(tx['status'])


def submit_task(tid):
    print("-" * 10, "提交任务", "-" * 10)
    tx = connector.submit_task(role.provider, tid)
    print(tx['status'])


def reward_task(tid, token_address, amount):
    print("-" * 10, "奖励任务", "-" * 10)
    tx = connector.reward_task(role.provider, tid, token_address, amount)
    print(tx)
    print(tx['status'])


def get_client_config(tid=None):
    if tid is None:
        tid = task_id
    # print("-" * 10, "获取client配置", "-" * 10)
    task, _ = connector.get_task(tid)
    result = connector.getProverConfig(task[3], task[1])
    print(result)

def sync_task( clientId, uniqID, rewardAmount, assignTime, proveTime, _stat, _assigner, _rewardRecipient):
    # print("-" * 10, "同步task", "-" * 10)
    tx = connector.sync_task(role.provider, clientId, uniqID, rewardAmount, assignTime, proveTime, _stat, _assigner, _rewardRecipient)
    print(tx['status'])

# def auto_init():
#     client_config = gen_client_config(env['PROVER_PUBLIC_KEY'], int(env['CLIENT_ID']), env['CLIENT_URL'], int(env['MAX_ZKEVM_INSTANCE']), int(env['MIN_FEE']))
#     print(client_config)
#     print("-" * 10, "加入market client", "-" * 10)
#     tx = connector.join_market(type("owner", (), dict(public_key=env['PROVER_PUBLIC_KEY'], private_key=env['PROVER_PRIVATE_KEY'])), client_config)
#     if tx['status'] != 1:
#         print("Join Market Failed")
#     else:
#         print("Join Market Success")


if __name__ == '__main__':
    # auto_init()
    # print(connector.market_dispatch(role.provider, '0x0000000000000000000000000000000000000000', 0)['status'])
    # create_client()
    # creat_client_2()
    # print(connector.close_prover_client(role.provider, 129k94572626014186902, '0x4DF42590c13b8110086A188b83a662DF0e6af1B8')['status'])
    # print(connector.has_resource())
    # for i in range(1):
    # for i in range(4):
    #     task_id = 2232 + i
    #     post_task()
    #     dispatch_task(task_id)
    # submit_task(2233)
    # reward_task(15, none_address, 0)
    
    # submit_task(2237)
    # reward_task(19, none_address, 0)
    # with open('task.record', 'r') as f:
    #     for i, line in enumerate(f):
    #         if i <= 2:
    #             continue
    #         line = json.loads(line)
    #         sync_task(line[1], line[2], line[8][1], line[10], line[11], line[6], line[3], line[8][0])
    #         print(i, line)            
    #         if i == 10:
    #             break
        
#     with open('task.record', 'r') as f:
#         for i, line in enumerate(f):
#             line = json.loads(line)
# # def sync_task( clientId, uniqID, rewardAmount, assignTime, proveTime, _stat, _assigner, _rewardRecipient):
#             print(line)
#             sync_task(16172934865820142997, 2236, line[8][1], line[10], line[11], 1, '0x74e2A876b86d7C345765d97D0E46FF9B84575F6a', line[8][0])
#             sync_task(16498400227274920934, 2237, line[8][1], line[10], line[11], 1, '0x863c9b8159B3F95687a600B1b21aE159618b31b1', line[8][0])
#             print(i, line)            
#             break
    
    # submit_task(1730403)
    # submit_task(1486673)
    # reward_task(task_id, '0x5A0c49a80Df506d46BF7A5F46d9339e9779e9664', 1.5)
    # print(connector.market_dispatch(role.provider, '0xC2600C80Beb521CC4E2f1b40B9D169c46E391390', 170073054677)['status'])
    # print(connector.has_resource())
    # v = connector.get_task_by_index(43)
    # print(v)

    # with open('task.record', 'w') as f:
    for index, i in enumerate(range(100000)):
        try:
            v = connector.get_task_by_index(i)
            print(index, ":", v[0], v[1], v[2], v[3], v[6])
        except Exception as e:
            break

    for i in range(100000):
        try:
            v = connector.get_client_by_index(i)
            # print(v)
            if v[6] in ('0', 0):
                print(v)
        except:
            break

    try:
        print(connector.getLowestN())
    except Exception as e:
        print("no lowest machine")
        pass
        # print(e)
    # get_task()
    # get_task()
    # get_task()
