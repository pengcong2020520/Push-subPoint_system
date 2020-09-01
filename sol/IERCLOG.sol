pragma solidity ^0.5.0;
pragma experimental ABIEncoderV2;
interface IERCLOG {
    function pushLog(uint8 ownerid, string calldata userid, string calldata jsonData, uint256 month) external;
    function queryLog(uint8 ownerid, string calldata userid, uint256 begin, uint256 end) external view returns(string[] memory);
}