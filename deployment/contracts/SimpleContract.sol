// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "hardhat/console.sol";

contract SimpleContract {
    uint256 totalMessages;

    constructor() {

    }

    function message() public {
        totalMessages += 1;
        console.log("%s has sent an message", msg.sender);
    }

    function getTotalMessages() public view returns (uint256){
        return totalMessages;
    }
}
