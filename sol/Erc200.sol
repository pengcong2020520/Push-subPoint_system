pragma solidity ^0.5.0;

import "./IERC200.sol";
import "./SafeMath.sol";

contract Erc200 is IERC200 {
    uint256 total;
    string symbol;
    using SafeMath for uint256;
    mapping(uint8=>address) owner;
    address admin;//
    mapping(string=>uint256) balances;
    using SafeMath for uint256;
    constructor(string memory sym, address _admin) public {
        symbol = sym;
        admin = _admin;
        total = 0;
    }
    
    //  admin 
    modifier onlyadmin() {
        require(admin == msg.sender, "onlyadmin can do this");
        _;
    }
    
    // owner
    modifier onlyowner(uint8 ownerid)  {
        require(owner[ownerid] == msg.sender || admin == msg.sender, "only owner can do this");
        _;
    }
    
    //only admin can update owner
    function updateOwner(uint8 ownerid, address _owner) public onlyadmin {
        owner[ownerid] = _owner;
    }
    
    function isEqual(string memory a, string memory b) public pure returns (bool) {
        bytes32 ha = keccak256(abi.encode(a));
        bytes32 hb = keccak256(abi.encode(b));
        return ha == hb;
    }
    // 查询发行量
	function totalSupply() external view returns (uint256)	{
        return total;
	}
	// 查询余额
	function balanceOf(string calldata who)  external view returns (uint256)
	{
	    return balances[who];
	}
	
	// 记录owner的operater log
	//event ownerlog(uint8 _ownerid, string _from, string _to, uint256 _value, string _type);
	// 转账
	function transfer(uint8 ownerid,string calldata from, string calldata to, uint256 value) onlyowner(ownerid) external returns (bool){
	    require(!isEqual(from, ""), "owner must exists");
	    require(!isEqual(to, ""), "to must exists");
	    require(balances[from] >= value, "owner'value must enough");
	    require(value > 0, "value must bigger than 0");
	    balances[from] = balances[from].sub(value);
	    balances[to] = balances[to].add(value);
	    emit Transfer(from, to, value);
	    emit Ownerlog(ownerid, from, to, value, "transfer");
	}
	
	// 消费 
    function burn(uint8 ownerid, string calldata to, uint256 value) external onlyowner(ownerid) returns (bool) {
        require(owner[ownerid] != address(0), "owner must exists");
        require(!isEqual(to, ""), "to must exists");
        require(value > 0, "value must bigger than 0");
        total = total.sub(value);
        balances[to] = balances[to].sub(value);
        emit Transfer("", to, 0-value);
        emit Ownerlog(ownerid, "", to, value, "burn");
    }

	// 挖矿
	function mint(uint8 ownerid, string calldata to, uint256 value) external onlyowner(ownerid) returns (bool) {
	    require(owner[ownerid] != address(0), "owner must exists");
	    require(!isEqual(to, ""), "owner must exists");
	    require(value > 0, "value must bigger than 0");
	    total = total.add(value);
	    balances[to] = balances[to].add(value);
	    emit Transfer("", to, value);
	    emit Ownerlog(ownerid, "", to, value, "mint");
	    //emit msglog(total, balances[to]);
	}
    function addr() external view returns (address) {
        return address(this);
    }
}