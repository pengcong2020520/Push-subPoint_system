pragma solidity ^0.5.0;

interface IUSER {
    function register(string calldata userid, string calldata pass, uint8 ownerid) external returns (bool) ;
    function setPasswd(string calldata userid, string calldata oldPass, string calldata newPass, uint8 ownerid) external ;
    function login(string calldata userid, string calldata pass) external view returns(bool) ;
    event Owneruserlog(string userid, uint8 ownerid, string _type);
}