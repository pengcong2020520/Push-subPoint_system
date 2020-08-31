pragma solidity ^0.5.0;

interface  IERC200 {
	// 查询发行量
	function totalSupply() external view returns (uint256);
	// 查询余额
	function balanceOf(string calldata who) external  view returns (uint256) ;
	// 转账
	function transfer(uint8 ownerid,string calldata from, string calldata to, uint256 value) external returns (bool);
	// 挖矿
	function mint(uint8 ownerid, string calldata to, uint256 value) external returns (bool);
	//  销毁
    function burn(uint8 ownerid, string calldata to, uint256 value) external returns (bool);
    // Ownerlog  事件  
    event Ownerlog(uint8 ownerid, string from, string to, uint256  value, string _type);
	// Transfer事件
	event Transfer(string from, string to,uint256 value);
}