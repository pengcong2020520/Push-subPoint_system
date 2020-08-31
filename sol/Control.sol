pragma solidity ^0.4.25;
pragma experimental ABIEncoderV2;

import "./Erclog.sol";
import "./Erc200.sol";
import "./User.sol";

contract Control {
    address owner;
    User user;
    Erc200 token;
    Erclog erclog;
    constructor(address _onwer, string memory sym) public {
        owner = _onwer;
        user = new User(_onwer);
        token = new Erc200(sym, _onwer);
        erclog = new Erclog(owner);

    }
    
    function useridValid(string userid) external view returns (bool) {
        return user.useridValid(userid);
    }
    
    modifier onlyowner() {
        require(owner == msg.sender, "only owner can do this");
        _;
    }
    
    function updateOwner(address _onwer) onlyowner public {
        owner = _onwer;
        user.updateOwner(_onwer);
        token.updateOwner(_onwer);
        erclog.updateOwner(_onwer);
    }
    
    function upgradeUser(address useraddr) external onlyowner {
        user = User(useraddr);
    }
    
    function upgradeErc200(address ercaddr) external onlyowner {
        token = Erc200(ercaddr);
    }
    
    function upgradeErclog(address logaddr) external onlyowner {
        erclog = Erclog(logaddr);
    }
    
    //user - call
    function register(string userid, string pass) external {
        user.register(userid, pass);
    }
    
    function setPasswd(string userid, string oldPass, string newPass) external {
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
	function transfer(string owner, string to, uint256 value) external  returns (bool) {
	    token.transfer(owner, to, value);
	}
	// 转账
	function mint(string to, uint256 value) external returns (bool) {
	    token.mint(to, value);
	}
    //log - call 
    function pushLog(string userid, string jsonData, uint256 month) external  {
        erclog.pushLog(userid, jsonData, month);
    }
    function queryLog(string userid, uint256 begin, uint256 end) external view returns(string[] memory) {
        return erclog.queryLog(userid, begin, end);
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