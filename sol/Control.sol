pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;

import "./Erclog.sol";
import "./Erc200.sol";
import "./User.sol";

contract Control {
    address admin;
    User user;
    Erc200 token;
    Erclog erclog;
    constructor(string _admin, string memory sym) public {
        admin = _admin;
        user = new User();
        token = new Erc200(sym);
        erclog = new Erclog();
    }
    
    function useridValid(string userid) external view returns (bool) {
        return user.isUseridValid(userid);
    }
    
    modifier onlyadmin() {
        require(admin == msg.sender, "only owner can do this");
        _;
    }
    
    modifier onlyUserOwner(string ownerid) {
        require(user.checkid(ownerid) == msg.sender || admin == msg.sender, "userid must be owner to this");
        _;
    }
    modifier onlyErc200Owner(string ownerid) {
        require(token.checkid(ownerid) == msg.sender || admin == msg.sender, "userid must be owner to this");
        _;
    }
    modifier onlyErclogOwner(string ownerid) {
        require(erclog.checkid(ownerid) == msg.sender || admin == msg.sender, "userid must be owner to this");
        _;
    }    
    //设置 user owner log 的 owner
    function setOwner(uint8 itype, uint8 ownerid, address owner) external view returns (address) {
        if(itype == 1) {
            return user.setOwner(ownerid, owner);
        }
        if(itype == 2) {
            return erc200.setOwner(ownerid, owner);
        }
        if(itype == 3) {
            return erclog.setOwner(ownerid, owner);
        }
    }
    // change owner 
    function updateOwner(uint8 itype, uint8 ownerid, address owner) onlyowner public {
        if(itype == 1) {
            return user.updateOwner(ownerid, owner);
        }
        if(itype == 2) {
            return erc200.updateOwner(ownerid, owner);
        }
        if(itype == 3) {
            return erclog.updateOwner(ownerid, owner);
        }
    }
    
    //user - call
    function register(string userid, string pass) external onlyUserOwner(ownerid) {
        user.register(userid, pass);
    }
    
    function setPasswd(string userid, string oldPass, string newPass) external onlyUserOwner(ownerid) {
        user.setPasswd(userid, oldPass, newPass);
    }
    
    function login(string userid, string pass) external view returns (bool) {
        return user.login(userid, pass);
    }
    //token - call 
    // 查询发行量
	function totalSupply() external view returns (uint256) {
	    return token.totalSupply();
	}
	// 查询余额
	function balanceOf(string who) external view  returns (uint256)  {
	    return token.balanceOf(who);
	}
	// 转账
	function transfer(string owner, string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool) {
	    token.transfer(owner, to, value);
	}
	// 挖矿 
	function mint(string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool) {
	    token.mint(to, value);
	}
	// 销毁 
	function burn(string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool) {
	    token.burn(to, value);
	}
    //log - call 
    function pushLog(string userid, string jsonData, uint256 month) external onlyErclogOwner(ownerid) {
        erclog.pushLog(userid, jsonData, month);
    }
    function queryLogByMongh(string userid, uint256 month) external view returns(string[] memory) {
        return erclog.queryLogByMongh(userid, month);
    }
    function getAddr(uint8 itype) external view returns (address) {
        if(itype == 1) {
            return user.addr();
        }
        if(itype == 2) {
            return token.addr();
        }
        if(itype == 3) {
            return erclog.addr();
        }
    }
}