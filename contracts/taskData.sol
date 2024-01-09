
// SPDX-License-Identifier: MIT
pragma solidity >= 0.4.22 <0.9.0;

import "./market.sol";
import "./token.sol";

import "./ApusData.sol";

contract TaskData {

    ApusData.Task [] public tasks;
    mapping (uint256 => mapping (ApusData.TaskType => uint256))  _tasksIndex;

    ApusData.rewardInfo  emptyRewardInfo = ApusData.rewardInfo(address(0), 0);
    ApusData.Task emptyTask = ApusData.Task(0, 0, 0, address(0), new bytes(0), ApusData.TaskType.TaikoZKEvm, ApusData.TaskStatus.Posted, new bytes(0), emptyRewardInfo, 0, 0, 0);
    ApusData.ClientConfig emptyClientConfig = ApusData.ClientConfig(address(0), 0, "", 0, 0, 0, ApusData.ClientStatus.Running);

    Market public market;
    ERC20 public token;

    constructor(address _marketAddr, address _tokenAddr) {
        market = Market(_marketAddr);
        token = ERC20(_tokenAddr);
    }


    // 设置任务索引
    function setTaskIndex(uint256 id, ApusData.TaskType taskType, uint256 value) public {
        require(value != 0, "Value must be non-zero"); // 确保值不是0
        _tasksIndex[id][taskType] = value;
    }

    // 检查任务索引是否存在，并返回其值
    function isTaskIndexExists(uint256 id, ApusData.TaskType taskType) public view returns (bool, uint256) {
        uint256 value = _tasksIndex[id][taskType];
        bool exists = value != 0;
        return (exists, value);
    }

    function postTask(ApusData.TaskType _tp, uint256 uniqID, bytes calldata input, uint64 expiry, ApusData.rewardInfo memory ri) public {
        (bool exists, ) = isTaskIndexExists(uniqID, _tp);
        if (exists) {
            return ;
        }
        tasks.push(ApusData.Task(tasks.length + 1, 0, uniqID, address(0), input, _tp, ApusData.TaskStatus.Posted, new bytes(0), ri, expiry, block.timestamp, 0));
    }

    function dispatchTaskToClient(uint256 uniqID) public {
        // address prover, uint256 cid, 
        (bool exists, uint256 index) = isTaskIndexExists(uniqID, ApusData.TaskType.TaikoZKEvm);
        require(exists, "unknown uniqID"); // 如果不存在，抛出错误

        if (tasks[index]._stat == ApusData.TaskStatus.Posted) {
            tasks[index]._stat = ApusData.TaskStatus.Assigned;
            ApusData.ClientConfig memory cf;
            (, cf) = market.getLowestN();
            tasks[index].assigner = cf.owner;
            tasks[index].clientId = cf.id;
            market.dispatchTaskToClient(cf.owner, cf.id);
        }
    }

    function submitTask(ApusData.TaskType _tp, uint256 uniqID, bytes calldata result) public {
        // address prover, uint256 cid, 
        (bool exists, uint256 index) = isTaskIndexExists(uniqID, _tp);
        require(exists, "unknown uniqID"); // 如果不存在，抛出错误
        require(tasks[index]._stat == ApusData.TaskStatus.Assigned);
        tasks[index]._stat = ApusData.TaskStatus.Done;
        tasks[index].result = result;
        tasks[index].proveTime = block.timestamp;
        market.releaseTaskToClient(tasks[index].assigner, tasks[index].clientId);
        return ;
    }

    function slashTask (ApusData.TaskType _tp, uint256 uniqID) public {
        (bool exists, uint256 index) = isTaskIndexExists(uniqID, _tp);
        require(exists, "unknown uniqID"); // 如果不存在，抛出错误
        require(tasks[index]._stat == ApusData.TaskStatus.Assigned);
        tasks[index]._stat = ApusData.TaskStatus.Slashed;
        market.releaseTaskToClient(tasks[index].assigner, tasks[index].clientId);
        return ;
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
} 
