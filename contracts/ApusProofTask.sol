// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./market.sol";
import "./token.sol";

import "./ApusData.sol";
import "./console.sol";

contract ApusProofTask {

    ApusData.Task [] public tasks;

    Market private market;
    ERC20 private token;

    event eventPostTask(ApusData.TaskType _tp, uint256 taskId, bytes data);

    constructor(address _marketAddr, address _tokenAddr) {
        market = Market(_marketAddr);
        token = ERC20(_tokenAddr);
    }


    function getDailyTaskCount(uint numDays) public view returns (uint256[] memory) {
        require(numDays > 0, "Number of days must be greater than 0");
        uint256[] memory dailyTaskCount = new uint256[](numDays);
        uint256 today = block.timestamp / 1 days;

        for (uint i = 0; i < tasks.length; i++) {
            uint256 taskDay = tasks[i].assignTime / 1 days;
            uint256 dayDiff = today - taskDay;

            if (dayDiff < numDays) {
                dailyTaskCount[dayDiff] += 1;
            }
        }

        return dailyTaskCount;
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
        uint256 count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if (tasks[i]._stat == ApusData.TaskStatus.Done && tasks[i].proveTime > 0) {
                total += tasks[i].proveTime - tasks[i].assignTime;
                count++;
            }
        }
        return total / count;
    }

    function getProverTasks(address prover) public view returns(ApusData.Task[] memory){
        uint count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover) {
                count++;
            }
        }
        ApusData.Task[] memory proverTasks = new ApusData.Task[](count);
        count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover) {
                proverTasks[count] = tasks[i];
                count++;
            }
        }
        return proverTasks;
    }

    function getClientTasks(address prover, uint256 clientId) public view returns(ApusData.Task[] memory){
        uint count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover && tasks[i].clientId == clientId) {
                count++;
            }
        }
        ApusData.Task[] memory clientTasks = new ApusData.Task[](count);
        count = 0;
        for(uint i = 0; i < tasks.length; i++) {
            if(tasks[i].assigner == prover && tasks[i].clientId == clientId) {
                clientTasks[count] = tasks[i];
                count++;
            }
        }
        return clientTasks;
    }

    function slashTask (ApusData.TaskType _tp, uint256 uniqID) public {
        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i]._tp ==  _tp && tasks[i].uniqID == uniqID) {
                require(tasks[i]._stat == ApusData.TaskStatus.Assigned);
                tasks[i]._stat = ApusData.TaskStatus.Slashed;
                // tasks[i].proveTime = block.timestamp;
                // market.getProverConfig(tasks[i].assigner, tasks[i].clientId);
                market.releaseTaskToClient(tasks[i].assigner, tasks[i].clientId);
                return ;
            }
        }
    }
    function submitTask(ApusData.TaskType _tp, uint256 uniqID, bytes calldata result) public {
        for (uint256 i = 0; i < tasks.length; i++) {
            if (tasks[i]._tp ==  _tp && tasks[i].uniqID == uniqID) {
                require(tasks[i]._stat == ApusData.TaskStatus.Assigned);
                tasks[i]._stat = ApusData.TaskStatus.Done;
                tasks[i].result = result;
                tasks[i].proveTime = block.timestamp;
                // market.getProverConfig(tasks[i].assigner, tasks[i].clientId);
                market.releaseTaskToClient(tasks[i].assigner, tasks[i].clientId);
                return ;
            }
        }
    }

    ApusData.rewardInfo  emptyRewardInfo = ApusData.rewardInfo(address(0), 0);
    ApusData.Task emptyTask = ApusData.Task(0, 0, 0, address(0), new bytes(0), ApusData.TaskType.TaikoZKEvm, ApusData.TaskStatus.Posted, new bytes(0), emptyRewardInfo, 0, 0, 0);
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

        tasks.push(ApusData.Task(tasks.length + 1, 0, uniqID, address(0), input, _tp, ApusData.TaskStatus.Posted, new bytes(0), ri, expiry, block.timestamp, 0));
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

    function rewardTask(uint256 _taskID, uint256 amount, address _tokenAddress) public payable {
        require(tasks[_taskID - 1]._stat == ApusData.TaskStatus.Done);
        tasks[_taskID - 1]._stat = ApusData.TaskStatus.Payed;
        tasks[_taskID - 1].reward.token = _tokenAddress;
        tasks[_taskID - 1].reward.amount = amount;

        if (_tokenAddress == address(0)) {
            require(msg.value == amount);
            payable(tasks[_taskID - 1].assigner).transfer(msg.value);
        } else {
            ERC20 _token = ERC20(_tokenAddress);
            console.log(_tokenAddress, amount);
            _token.transferFrom(msg.sender, address(this), amount);
            _token.transfer(payable(tasks[_taskID - 1].assigner), amount);
            console.log(_tokenAddress, amount, "done");
        }
        
        token.reward(tasks[_taskID - 1].assigner);
    }

    function assignTask(uint256 taskID) public {
    }

    function hasResource() public view returns(uint256){
        return market.marketCapacity();
    }
} 
