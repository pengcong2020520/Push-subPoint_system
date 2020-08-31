pragma solidity^0.7.0;

contract ping {
    address admin;
    string myMsg="pong2";

    constructor (address owner) {
        admin = owner;
    }

    function setMsg(string memory _msg) public {
        myMsg = _msg;
    }
    function getMsg() public view returns (string memory) {
        return myMsg;
    }
}