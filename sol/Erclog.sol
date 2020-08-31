pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./IERCLOG.sol";

contract Erclog is IERCLOG {

    mapping(string=>mapping(uint256=>string[])) msglogs;
    address public owner;
    address admin;
    constructor(address _owner) public {
        admin = msg.sender;
        owner = _owner;
    }
    
    
    modifier onlyowner() {
        require(owner == msg.sender || admin == msg.sender, "only owner can do this");
        _;
    }
    function updateOwner(address _owner) public onlyowner {
        owner = _owner;
    }
    function isEqual(string memory a, string memory b) public view returns (bool) {
        bytes32 ha = keccak256(abi.encode(a));
        bytes32 hb = keccak256(abi.encode(b));
        return ha == hb;
    }
    function pushLog(string userid, string jsonData, uint256 month) external onlyowner {
        require(month > 202007, "month must valid");
        require(!isEqual("", userid), "userid must valid");
        msglogs[userid][month].push(jsonData);
    }
    function queryLog(string userid, uint256 begin, uint256 end) external view returns(string[] memory) {
        return msglogs[userid][begin];
    }
    function queryLogByMongh(string userid, uint256 month) external view returns(string[] memory) {
        return msglogs[userid][month];
    }
    function addr() external view returns (address) {
        return address(this);
    }
}