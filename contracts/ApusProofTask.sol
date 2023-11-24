// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./market.sol";
import "./token.sol";

import "./ApusData.sol";

contract ApusProofTask {

    ApusData.Task [] public tasks;

    Market private market;
    ERC20 private token;

    event eventPostTask(ApusData.TaskType _tp, uint256 taskId, bytes data);

    constructor(address _marketAddr, address _tokenAddr) {
        market = Market(_marketAddr);
        token = ERC20(_tokenAddr);
    }

    function getTaskCount() public view returns(uint256){
        return tasks.length;
    }

    function getAssignedTaskCount() public view returns(uint256){
        uint256 count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i]._stat == ApusData.TaskStatus.Assigned) {
                count += 1;
            }
        }
        return count;
    }

    function getLatestTaskId() public view returns(uint256){
        return tasks[tasks.length - 1].id;
    }

    function getAvgReward() public view returns(uint256){
        uint256 total = 0;
        for(uint i = 0; i < tasks.length; i++) {
            total += tasks[i].reward.amount;
        }
        return total / tasks.length;
    }

    function getAvgProofTime() public view returns(uint256){
        uint256 total = 0;
        for(uint i = 0; i < tasks.length; i++) {
            total += tasks[i].proveTime - tasks[i].assignTime;
        }
        return total / tasks.length;
    }

    function getProverTasks(address prover) public view returns(ApusData.Task[] memory){
        ApusData.Task[] memory proverTasks;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover) {
                proverTasks.push(tasks[i]);
            }
        }
        return proverTasks;
    }

    function getClientTasks(address prover, uint256 clientId) public view returns(ApusData.Task[] memory){
        ApusData.Task[] memory clientTasks;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover && tasks[i].clientId == clientId) {
                clientTasks.push(tasks[i]);
            }
        }
        return clientTasks;
    }

    function submitTask(ApusData.TaskType _tp, uint256 uniqID, bytes calldata result) public {
        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i]._tp ==  _tp && tasks[i].uniqID == uniqID) {
                require(tasks[i]._stat == ApusData.TaskStatus.Assigned);
                tasks[i]._stat = ApusData.TaskStatus.Done;
                tasks[i].result = result;
                task[i].proveTime = block.timestamp;
                market.getProverConfig(tasks[i].assigner, tasks[i].clientId);
                market.releaseTaskToClient(tasks[i].assigner, tasks[i].clientId);
                token.reward(tasks[i].assigner);
                return ;
            }
        }
    }

    ApusData.rewardInfo  emptyRewardInfo = ApusData.rewardInfo(address(0), 0);
    ApusData.Task emptyTask = ApusData.Task(0, 0, 0, address(0), new bytes(0), ApusData.TaskType.TaikoZKEvm, ApusData.TaskStatus.Posted, new bytes(0), emptyRewardInfo, 0);
    ApusData.ClientConfig emptyClientConfig = ApusData.ClientConfig(address(0), 0, "", 0, 0, 0, ApusData.ClientStatus.Running);

    function getTask(ApusData.TaskType _tp, uint256 uniqID) public view returns(ApusData.Task memory, ApusData.ClientConfig memory){
        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i]._tp ==  _tp && tasks[i].uniqID == uniqID) {
                if (tasks[i]._stat != ApusData.TaskStatus.Posted) {
                    ApusData.ClientConfig memory cf = market.getProverConfig(tasks[i].assigner, tasks[i].clientId);
                    return (tasks[i], cf);
                }
                return (tasks[i], emptyClientConfig);
            }
        }
        return (emptyTask, emptyClientConfig);
    }

    function postTask(ApusData.TaskType _tp, uint256 uniqID, bytes calldata input, uint64 expiry, ApusData.rewardInfo memory ri) public {
        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i]._tp ==  _tp && tasks[i].uniqID == uniqID) {
                return ;
            }
        }

        tasks.push(ApusData.Task(tasks.length + 1, 0, uniqID, address(0), input, _tp, ApusData.TaskStatus.Posted, new bytes(0), ri, expiry));
        emit eventPostTask(_tp, uniqID, input);
    }

    function dispatchTaskToClient(uint256 taskID) public {
        // address prover, uint256 cid, 

        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i].uniqID == taskID) {
                if (tasks[i]._stat == ApusData.TaskStatus.Posted) {
                    tasks[i]._stat = ApusData.TaskStatus.Assigned;
                    ApusData.ClientConfig memory cf;
                    (, cf) = market.getLowestN();

                    tasks[i].assigner = cf.owner;
                    tasks[i].clientId = cf.id;
                    market.dispatchTaskToClient(cf.owner, cf.id);
                }
            }
        }
    }

    function assignTask(uint256 taskID) public {
    }

    function hasResource() public view returns(uint256){
        return market.marketCapacity();
    }
} 
