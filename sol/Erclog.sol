pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
import "./IERCLOG.sol";
contract Erclog {
    mapping(uint8=>address) owners;
    mapping(address=>bool) isOwnerExist;
    address admin;
    mapping(string=>mapping(uint256=>string[])) msglogs;
    
    constructor () public {
        admin = msg.sender;
    }
    
    modifier onlyadmin() {
        require(admin == msg.sender, "only admin can do it!");
        _;
    }
    
    modifier onlyOwner(uint8 ownerid) {
        require(msg.sender == admin || owners[ownerid] == msg.sender, "only owner can do it");
        _;
    }
    
    function setOwner(uint8 ownerid, address owner) external onlyadmin {
        require(owners[ownerid] == address(0), "ownerid must be empty!");
        require(!isOwnerExist[owner], "owner must not be owner!");
        owners[ownerid] = owner;
        isOwnerExist[owner] = true; 
    }
    
    
    function updataOwner(uint8 ownerid, address owner) external onlyOwner(ownerid)  {
        require(!isOwnerExist[owner], "owner must not be owner");
        owners[ownerid] = owner;
        isOwnerExist[owner] = true;
        isOwnerExist[msg.sender] = false;
    }  
    
    function isEqual(string memory a, string memory b) public view returns (bool) {
        bytes32 ha = keccak256(abi.encode(a));
        bytes32 hb = keccak256(abi.encode(a));
        return ha == hb;
    }
    
    event PushLog(uint8 _ownerid, string _userid, string _jsonData, uint256 _month);
    
    function pushLog(uint8 ownerid, string calldata userid, string calldata jsonData, uint256 month) external onlyOwner(ownerid) {
        require(month > 202008, "month must be valid!");
        require(!isEqual(userid, ""), "userid must not be empty! ");
        msglogs[userid][month].push(jsonData);
        
        emit PushLog(ownerid, userid, jsonData, month);
    }
    
    function queryLogByMonth(string calldata userid, uint256 month) external view returns(string[] memory) {
        require(month > 202008, "month must be valid!");
        require(!isEqual(userid, ""));
        return msglogs[userid][month];
    }
    
}