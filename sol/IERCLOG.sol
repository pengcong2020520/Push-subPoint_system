pragma solidity^0.4.25;
pragma experimental ABIEncoderV2;
interface IERCLOG {
    
    function pushLog(string userid, string jsonData, uint256 month) external;
    function queryLog(string userid, uint256 begin, uint256 end) external view returns(string[] memory);
    function queryLogByMongh(string userid, uint256 month) external view returns(string[] memory);
}