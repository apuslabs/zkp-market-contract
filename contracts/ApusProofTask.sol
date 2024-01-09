// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "./market.sol";
import "./token.sol";

import "./ApusData.sol";
import "./console.sol";
import "./taskData.sol";


contract ApusProofTask is TaskData {

    constructor(address _marketAddr, address _tokenAddr) TaskData(_marketAddr, _tokenAddr) {
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


    function hasResource() public view returns(uint256){
        return market.marketCapacity();
    }
} 
