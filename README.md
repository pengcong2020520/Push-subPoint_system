## 基于ethereum超市推送-积分子系统
### 项目目的
* 基于ethereum的超市推送积分子系统主要是提供用户注册登录，积分消费充值以及转移，日志推送以及记录的服务接口，方便调用；
* 通过区块链技术可以完成积分的消费充值功能的实时清算，以及日志信息的实时记录；
* 确保了用户的体验，使得积分的消费以及充值或者积分规则都变得更为公开透明；
* 方便用户之间的积分交易流通，体现积分的真正价值；
* 通过区块链技术能够最大程度的保证积分的安全；
* 充分利用区块链技术优势中的去中心化、开放性、不可篡改的优势，推动了商业活动的发展。
### 项目介绍
本项目主要是基于Ethereum的智能合约对超市的积分进行存储，通过积分规则以及操作日志计算并生成区块链存储的积分。推送-积分子系统是报送系统的核心模块，主要负责连接区块链与报送主控系统，通过智能合约来维护区块链系统的数据，并通过HTTP接口与报送主控系统进行通信。
### 系统介绍
系统设计主要分为三部分：
* 智能合约部分
		负责积分、积分使用情况的存储；
		负责账户地址与密码的存储，这一部分以userid对应pwd的方式存储，使得权限的设计更为灵活
* 积分接口服务部分
		用户模块主要包括登录，注册，密码修改
		积分模块主要包括积分的使用、获得以及授权转让
		日志模块主要包括日志信息的记录、查询以及推送
* 数据库部分
		记录用户信息，积分使用详细情况

### 合约设计部分

合约的设计架构示意图如下。

数据合约 <--> 业务合约（多个）<--> 总控合约 <------> 管理员/用户

		数据合约部分主要包括用户数据、日志数据以及积分情况的存储；
		业务合约分为管理员权限、普通用户权限操作的业务；
		总控合约主要负责与管理员以及用户端进行交互，满足管理员以及用户的需求。总控合约即可以调用业务合约执行相应的业务，也可以直接调用数据合约，返回用户所需要的数据需求。

#### 积分合约
定义一个接口合约文件（相当于积分合约的一个标准）

* 查询token发行总量函数，相当于用户可以查询到总共发行的积分数值，增加积分系统的透明度，提高用户体验：
```
function totalSupply() external view returns (uint256);
```
* 查询userid对应的token余额函数，相当于用户的积分状态，提高积分系统的直观性：
```
function balanceOf(string calldata who) external  view returns (uint256);
```
* token转账函数，相当于用户积分的转移：
```
function transfer(uint8 ownerid,string calldata from, string calldata to, uint256 value) external returns (bool);
```
token转账交易函数，其中

	* `ownerid`代表的是执行token转账对应的管理员id;

	* `from`代表转账用户;

	* `to`代表被转账用户;

	* `value`代表转账数额.
值得说明的一点， 这里的token `transfer` 函数与其他token合约不同的地方在于有一个`ownerid`,这代表着该合约是的token转账只能由管理员操作，有效防止积分的随意炒作。

* 积分充值函数，即合约中的token挖矿函数：
```
function mint(uint8 ownerid, string calldata to, uint256 value) external returns (bool);
```
向指定的userid用户充值积分，积分可以是活动获得，或者消费获得等等，根据具体业务具体分析。

* 积分消费函数，即合约中的token销毁函数
```
function burn(uint8 ownerid, string calldata to, uint256 value) external returns (bool);
```
可以根据用户的积分使用情况，消除对应的积分。

* 管理员操作日志事件
```
event Ownerlog(uint8 ownerid, string from, string to, uint256  value, string _type);
```
管理员对应的对积分的操作，通过一个事件来反映，反正管理员操作的抵赖，实现透明度，保证用户积分的安全性。
可以有效对管理员进行一个监督与管理。

* 积分操作事件
```
event Transfer(string from, string to,uint256 value);
```
积分的转移，消费以及充值情况的事件，用于记录积分的使用情况。

本合约的一个设计亮点：

1. 引入管理员id，使得合约可以经过多个管理员进行操作管控一份合约；
2. 通过由一个admin来管控管理员id，实现了企业级有效的分级管理制度，使得用户以及admin可以对管理员进行监督；
3. 积分的操作情况，均需要通过管理员来进行操控，合理控制了滥用积分，也算对用户积分使用的一个合理监管；
4. 积分的使用更为透明，灵活，利于开展更多的企业级商业活动。

#### 用户合约
* **数据结构** 
定义三个hash表，用来存储用户数据，管理员owner数据以及管理员账户是否存在：
``` solidity
 mapping(string=>string) Passwds;
 mapping(uint8=>address) owners;
 mapping(address=>bool) isOwnerExist;
```
其中hash表`isOwnerExist`主要是用于防止账户重复设置管理员，以及admin授权时出现混乱。

* **功能函数**
	* 注册功能
	```solidity
	function register(string calldata userid, string calldata pass, uint8 ownerid) external returns (bool) ;
	```
	* 登录功能
	```solidity
	function login(string calldata userid, string calldata pass) external view returns(bool) ;
	```
	* 修改密码功能
	```solidity
	function setPasswd(string calldata userid, string calldata oldPass, string calldata newPass, uint8 ownerid) external ;
	```
    在上述的功能函数中，userid代表的是用户的唯一标识，ownerid代表的是管理员账户id，用户的注册和修改密码功能都需要owner的参与，owner作为用户的服务者且也是用户账户的管理者。通过区块链的方式，使得owner并不能直接找到账户的密码，同时用户可以通过事件订阅来监督owner的操作是否符合自己的预期，透明性的增加让用户有着极佳的使用体验。所以需要定义一个owner操作日志事件。
	* owner操作日志事件
	```solidity
	event Owneruserlog(string userid, uint8 ownerid, string _type);
	```

#### 日志合约
日志合约方面主要用来推送和上传用户的积分操作日志
* 上传日志函数功能
```solidity
function pushLog(uint8 ownerid, string calldata userid, string calldata jsonData, uint256 month) external;
```
在设计中，采用按月存储的方式，简化存储在区块链上资源的数量，有效缩小了区块链存储的压力

* 查询日志功能
```solidity
 function queryLog(uint8 ownerid, string calldata userid, uint256 begin, uint256 end) external view returns(string[] memory);
```
这里的设计也是按月查询，如果需要具体化到查询的时间，最好可以通过后端进行一些简单的处理即可，保持合约的简洁性。

#### 总控合约
总控合约主要负责用一个admin账户来控制owner的权限，有利于分级化管理。
* 设置owner功能的实现函数
```
function setOwner(uint8 itype, uint8 ownerid, address owner) external onlyadmin ;
```
只需要通过一个类别编号，就可以轻松用一个函数来操控多个模块合约。

* 更新owner功能的实现
更新owner可以通过原owner本人操作，也可以由admin进行操作，符合分级管理的习惯。
```solidity
function updateOwner(uint8 itype, uint8 ownerid, address owner) public;
```
与设置owner基本类似
* 调用其他合约的功能模块以及调用其他合约查看信息
```solidity
function register(string userid, string pass) external onlyUserOwner(ownerid)；
function setPasswd(string userid, string oldPass, string newPass) external onlyUserOwner(ownerid)；
function login(string userid, string pass) external view returns (bool)；
```
用户登录注册以及修改密码功能:
登录是通过调用user合约进行数据匹配查看；
注册和修改密码功能需要管理员权限进行。
```solidity
	// 转账
	function transfer(string owner, string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool);
	// 挖矿 
	function mint(string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool);
	// 销毁 
	function burn(string to, uint256 value) external onlyErc200Owner(ownerid) returns (bool);
```
积分的转移，消费或或者充值功能。

总控合约主要负责与用户端的交互，获取用户的需求，然后调用相应的合约进行数据的CRUD。
















