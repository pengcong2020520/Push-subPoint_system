pragma solidity ^0.5.0;

import "./IUSER.sol";

contract User {
    mapping(string=>string) Passwds;
    mapping(uint8=>address) owners;
    mapping(address=>bool) isOwnerExist;
    address admin;
    constructor() public {
        admin = msg.sender;
    }
    

    
    modifier onlyadmin() {
        require(admin == msg.sender);
        _;
    }
    
    modifier onlyowner(uint8 ownerid) {
        require(msg.sender == owners[ownerid] || admin == msg.sender, "only owner can do this");
        _;
    }
    
    function setOwner(uint8 ownerid, address owner) external onlyadmin {
        require((owners[ownerid] == address(0)), "ownerid must be empty");
        require(!isOwnerExist[owner], "address must not be owner");
        owners[ownerid] = owner;
        isOwnerExist[owner] = true;
    }
    
    function updateOwner(uint8 ownerid, address owner) external onlyowner(ownerid) {
        require((owners[ownerid] != address(0)), "ownerid must not be empty");
        require(!isOwnerExist[owner], "address must not be owner");        
        owners[ownerid] = owner;
        isOwnerExist[owner] = false;
    }
    
    function isEqual(string memory a, string memory b) internal view returns (bool) {
        bytes32 ha = keccak256(abi.encode(a));
        bytes32 hb = keccak256(abi.encode(b));
        return ha == hb;
    }

    //true present userid valid; false present userid invalid
    function isUseridValid(string calldata userid) external view returns (bool) {
        return !isEqual(Passwds[userid], "");
    }
    
    function owneridValid(uint8 ownerid) external view returns (bool) {
        return owners[ownerid] != address(0);
    }
    
    event Owneruserlog(string userid, uint8 ownerid, string _type);
    
    function register(string calldata userid, string calldata pass, uint8 ownerid) external onlyowner(ownerid) returns (bool) {
        require(isEqual(Passwds[userid], ""), "user must not register");
        require(!isEqual(pass, ""), "password must not empty");
        Passwds[userid] = pass;
        
        emit Owneruserlog(userid, ownerid, "register");
	    return true;
    }

    function setPasswd(string calldata userid, string calldata oldPass, string calldata newPass, uint8 ownerid) external onlyowner(ownerid) returns (bool) {
        require(isEqual(Passwds[userid], oldPass), "user or passwd error");
        Passwds[userid] = newPass;
        emit Owneruserlog(userid, ownerid, "setPasswd");
	    return true;
    }

    

    function login(string calldata userid, string calldata pass) external view returns(bool)  {
        return isEqual(Passwds[userid], pass);
    }

    function addr() external view returns (address) {
        return address(this);
    }

    

}

