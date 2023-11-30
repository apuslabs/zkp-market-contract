// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IERC20 {
    function transfer(address to, uint256 amount) external returns (bool);
}

contract MyContract {
    address private _tokenAddress;

    constructor(address tokenAddress) {
        _tokenAddress = tokenAddress;
    }

    function transferTokens(address to, uint256 amount) public {
        IERC20 token = IERC20(_tokenAddress);
        require(token.transfer(to, amount), "Transfer failed");
    }
}