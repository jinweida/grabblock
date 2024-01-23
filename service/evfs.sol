pragma solidity ^0.4.23;

library SafeMath {
  function mul(uint256 a, uint256 b) internal pure returns (uint256) {
    if (a == 0) {
      return 0;
    }
    uint256 c = a * b; 
    assert(c / a == b);  
    return c;
  }
  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b > 0); // Solidity automatically throws when dividing by 0
    uint256 c = a / b; 
    // assert(a == b * c + a % b); // There is no case in which this doesn't hold
    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    // assert(b <= a);
    require(b <= a, "sub error");
    return a - b;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    assert(c >= a);
    return c;
  }
}
contract ERC20 {
  // events
  event Approval(address indexed owner, address indexed agent, uint256 value);
  event Transfer(address indexed from, address indexed to, uint256 value);
  // public functions
  function totalSupply() public view returns (uint256);
  function balanceOf(address addr) public view returns (uint256);
  function transfer(address to, uint256 value) public returns (bool);
  // public functions
  function allowance(address owner, address agent) public view returns (uint256);
  function transferFrom(address from, address to, uint256 value) public returns (bool);
  function approve(address agent, uint256 value) public returns (bool);
}
contract LockIdGen {
    uint256  requestCount;
    constructor() public {
        requestCount = 0;
    }
    function generateLockId() internal returns (bytes32 lockId) {
        return keccak256(abi.encodePacked(blockhash(block.number-1), address(this), ++requestCount));
    }
}
interface iEvfsGroup {
    // 获取上级group地址
    function getParent() external returns (address);
    // 获取所有下级group地址
    function getSubs() external returns (address[]);
    // 是否允许下级加入
    function isAllowJoin() external returns (bool);
    // 加入
    function subJoin() external returns (bool);

    // 组功能
    // 获取当前组成员列表
    function getMembers() external returns (address[]);
    // 获取当前组配置
    function getConfig(bytes32 _key) external returns (bytes);

    // 组管理 成员申请,
    function memberApply(address _member, uint256 _op, address _subGroup) external returns (bytes32 _reqId, bool _isEnd);
    function memberAgree(bytes32 _reqId) external returns (bool);
    // 组管理 配置管理
    function configApply(bytes32 _configKey, bytes _configValue) external returns (bytes32 _reqId, bool _isEnd);
    function configAgree(bytes32 _reqId) external returns (bool);
    // 组管理 审批规则管理
    function ruleApply(uint256 _rule) external returns (bytes32 _reqId, bool _isEnd);
    function ruleAgree(bytes32 _reqId) external returns (bool);
    // 组管理 审批拒绝
    function disagree(bytes32 _reqId) external returns (bool);

    // 调用方必须是当前组的parent
    function addMember(address _member) external;
    function removeMember(address _member) external;
}
interface iEvfsDomain {
    function existsDomain(address _domainId) external returns (bool);
    function createDomain(address _domainId, address _pdomainId, bytes _name, address[] _members) external;
    function setStatus(address _domainId, uint256 _status) external;

    function ruleApply(address _domainId, uint256 _rule) external returns (bytes32 _reqId, bool _isEnd);
    function ruleAgree(address _domainId, bytes32 _reqId) public returns (bool _isEnd);
    function memberApply(address _domainId, address _member, uint256 _op) external returns (bytes32 _reqId, bool _isEnd);
    function memberAgree(address _domainId, bytes32 _reqId) public returns (bool);
}
contract EvfsGroup is iEvfsGroup, LockIdGen {
    using SafeMath for uint256;
    struct MemberRequest {
        address member;
        uint256 op;
        address subGroupAddress;
    }
    struct ConfigRequest {
        bytes32 configKey;
        bytes   configValue;
    }
    struct ConfirmResult {
        uint256 confirmResult;
        uint256 confirmOn;
        uint256 total;
        uint256 agrees;
        uint256 disagrees;
        uint256 ruleBase;
        uint256 rule;
    }
    mapping (bytes32 => MemberRequest)                  memberReqs;
    mapping (bytes32 => ConfigRequest)                  configReqs;
    mapping (bytes32 => ConfirmResult)                  confirmResults;

    uint256                                             rule;
    uint256                                             ruleBase;
    bool                                                allowJoin;
    address                                             parentGroupAddress;
    mapping (address => address)                        subGroups;

    mapping (bytes32 => bytes)                          groupConfigs;

    mapping (bytes32 => uint256)                        ruleReqs;

    mapping (bytes32 => mapping (address => address))   todosById;
    mapping (address => mapping (bytes32 => uint256))   todosHistoryAddress;
    

    bool                                                _canMemberApply;
    mapping (address => address)                        groupMembers;
    address[]                                           groupMembersArray;
    address                                             manager;

    mapping (bytes32 => address[])                      addressIndexByTodo;
    mapping (address => bytes32[])                      todoIndexByAddress;

    constructor (
        address []members
    ) 
    LockIdGen() 
    public {
        manager = msg.sender;
        uint256 numMans = members.length;
        for (uint256 i = 0; i < numMans; i++) {
          address pto = members[i];
          require(pto != address(0) && groupMembers[pto] == address(0x0), "invalidate parameters");
          groupMembers[pto] = pto;
          groupMembersArray.push(pto);
        }

        rule = uint256(200);
        ruleBase = uint256(300);
        allowJoin = true;
    }
    modifier onlyManager {
        require(msg.sender == manager, "only manager");
        _;
    }
    modifier onlyMember {
        require(msg.sender == groupMembers[msg.sender], "not in group");
        _;
    }
    modifier onlyParent {
        require(msg.sender == parentGroupAddress, "operation not allowed");
        _;
    }
    modifier canMemberApply(address _member) {
        require (_canMemberApply, "cannot apply");
        // require(_subGroupAddress == address(0) || subGroups[_subGroupAddress] == _subGroupAddress, "cannot apply");
        require(groupMembers[_member] == address(0), "member exists");
        _;
    }
    modifier canConfigApply(bytes32 _configKey) {
        require(msg.sender == groupMembers[msg.sender], "not in group");
        _;
    }
    modifier canRuleApply(uint256 _rule) {
        require(_rule <= ruleBase, "invalidate rule");
        require(msg.sender == groupMembers[msg.sender], "not in group");
        _;
    }
    modifier canConfirm(bytes32 _reqId) {
        require(confirmResults[_reqId].confirmResult == uint256(0), "request finished");
        require(todosById[_reqId][msg.sender] == msg.sender, "cannot confirm");
        require(todosHistoryAddress[msg.sender][_reqId] == 0, "already confirmed");
        _;
    }
    function getParent() external returns (address) {
        return parentGroupAddress;
    }
    function getSubs() external returns (address[]) {
        return groupMembersArray;
    }
    function isAllowJoin() external returns (bool) {
        return allowJoin;
    }
    function subJoin() external returns (bool) {
        require(allowJoin, "not allowed");
        require(subGroups[msg.sender] == address(0), "already join");
        subGroups[msg.sender] = msg.sender;
    }
    function getMembers() external returns (address[]) {
        return groupMembersArray;
    }
    function getConfig(bytes32 _key) external returns (bytes) {
        require(_key.length > 0, "invalidate key");
        return groupConfigs[_key];
    }
    function memberApply(address _member, uint256 _op, address _subGroupAddress) external onlyMember canMemberApply(_member) returns (bytes32 _reqId, bool _isEnd) {
        require(_member != address(0) , "0");
        _isEnd = false;

        _reqId = generateLockId();
        memberReqs[_reqId] = MemberRequest({
            member: _member,
            op: _op,
            subGroupAddress: _subGroupAddress
        });

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;

            addressIndexByTodo[_reqId].push(member);
            todoIndexByAddress[member].push(_reqId);
        }

        if (groupMembers[msg.sender] == msg.sender) {
            _isEnd = memberAgree(_reqId);
        }
    }
    function clear(address _addr, bytes32 _reqId) internal {
        uint256 len1 = addressIndexByTodo[_reqId].length;

        for (uint256 j=0; j<len1; j++) {
            if (addressIndexByTodo[_reqId][j] == _addr) {
                // 移位，删除元素
                for (uint k = j; k<len1-1; k++){
                    addressIndexByTodo[_reqId][k] = addressIndexByTodo[_reqId][k+1];
                }
                delete addressIndexByTodo[_reqId][len1-1];
                addressIndexByTodo[_reqId].length--;
                break;
            }
        }

        len1 = todoIndexByAddress[_addr].length;
        for (j=0; j<len1; j++) {
            if (todoIndexByAddress[_addr][j] == _reqId) {
                // 移位，删除元素
                for (k = j; k<len1-1; k++){
                    todoIndexByAddress[_addr][k] = todoIndexByAddress[_addr][k+1];
                }
                delete todoIndexByAddress[_addr][len1-1];
                todoIndexByAddress[_addr].length--;
                break;
            }
        }
    }
    function clearAll(bytes32 _reqId) internal {
        address[] _addrs = addressIndexByTodo[_reqId];
        uint256 len = _addrs.length;
        for (uint256 i=0; i<len; i++) {
            uint256 len1 = todoIndexByAddress[_addrs[i]].length;

            for (uint256 j=0; j<len1; j++) {
                if (todoIndexByAddress[_addrs[i]][j] == _reqId) {
                    // 移位，删除元素
                    for (uint k = j; k<len1-1; k++){
                        todoIndexByAddress[_addrs[i]][k] = todoIndexByAddress[_addrs[i]][k+1];
                    }
                    delete todoIndexByAddress[_addrs[i]][len1-1];
                    todoIndexByAddress[_addrs[i]].length--;
                    break;
                }
            }
        }

        delete addressIndexByTodo[_reqId];
    }
    function memberAgree(bytes32 _reqId) public canConfirm(_reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if ((_confirmResult.agrees.mul(1000).div(_confirmResult.total)) >= (_confirmResult.rule.mul(1000).div(_confirmResult.ruleBase))) {
            // approved
            performMemberApplySelf(_reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);

            clearAll(_reqId);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
    function disagree(bytes32 _reqId) external canConfirm(_reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.disagrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(3);

        if (_confirmResult.disagrees.mul(1000).div(_confirmResult.total) > _confirmResult.ruleBase.sub(_confirmResult.rule).mul(1000).div(_confirmResult.ruleBase)) {
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(2);

            clearAll(_reqId);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
    function performMemberApply(bytes32 _reqId) external onlyParent {
        performMemberApplySelf(_reqId);
    }
    function configApply(bytes32 _configKey, bytes _configValue) external canConfigApply(_configKey) returns (bytes32 _reqId, bool _isEnd) {
        require(_configValue.length > 0 && _configKey.length > 0, "0");
        _isEnd = false;

        _reqId = generateLockId();
        configReqs[_reqId] = ConfigRequest({
            configKey: _configKey,
            configValue: _configValue
        });

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });
        
        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;

            // 如果是成员发起的申请，默认自动审批通过
            if (member == msg.sender) {
                _isEnd = configAgree(_reqId);
            }
        }
    }
    function configAgree(bytes32 _reqId) public canConfirm(_reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            groupConfigs[configReqs[_reqId].configKey] = configReqs[_reqId].configValue;
            clearAll(_reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
    function ruleApply(uint256 _rule) external canRuleApply(_rule) returns (bytes32 _reqId, bool _isEnd) {
        _reqId = generateLockId();
        ruleReqs[_reqId] = _rule;

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;

            // 如果是成员发起的申请，默认自动审批通过
            if (member == msg.sender) {
                _isEnd = ruleAgree(_reqId);
            }
        }
    }
    function ruleAgree(bytes32 _reqId) public canConfirm(_reqId) returns (bool _isEnd) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            rule = ruleReqs[_reqId];
            clearAll(_reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isEnd = _confirmResult.confirmResult != uint256(0);
    }
    function performMemberApplySelf(bytes32 _reqId) internal {
        MemberRequest memory memberRequest = memberReqs[_reqId];
        iEvfsGroup _sugGroup;
        address _member = memberRequest.member;
        if (memberRequest.op == uint256(1)) {
            if (memberRequest.subGroupAddress == address(0)) {
                require(groupMembers[_member] == address(0), "duplicate member");
                groupMembers[_member] = _member;
                groupMembersArray.push(_member);
            } else {
                _sugGroup = iEvfsGroup(memberRequest.subGroupAddress);
                _sugGroup.addMember(_member);
            }
        } else if (memberRequest.op == uint256(2)) {
            if (memberRequest.subGroupAddress == address(0)) {
                groupMembers[_member] = address(0);
                uint256 len = groupMembersArray.length;
                for (uint256 i=0; i<len; i++) {
                    if (groupMembersArray[i] == _member) {
                        // 移位，删除元素
                        for (uint j = i; j<len-1; j++){
                            groupMembersArray[j] = groupMembersArray[j+1];
                        }
                        delete groupMembersArray[len-1];
                        groupMembersArray.length--;
                        break;
                    }
                }
            } else {
                _sugGroup = iEvfsGroup(memberRequest.subGroupAddress);
                _sugGroup.removeMember(_member);
            }
        } else {
            // todo other
        }
    }
    function addMember(address _member) external onlyParent {
        require(groupMembers[_member] == address(0), "duplicate member");
        groupMembers[_member] = _member;
        groupMembersArray.push(_member);
    }
    function removeMember(address _member) external onlyParent {
        groupMembers[_member] = address(0);
        uint256 len = groupMembersArray.length;
        for (uint256 i=0; i<len; i++) {
            if (groupMembersArray[i] == _member) {
                // 移位，删除元素
                for (uint j = i; j<len-1; j++){
                    groupMembersArray[j] = groupMembersArray[j+1];
                }
                delete groupMembersArray[len-1];
                groupMembersArray.length--;
                break;
            }
        }
    }
}
contract CommitteeGroup is EvfsGroup {
    constructor (
        address[] members
    ) 
    EvfsGroup(members)
    public {
        _canMemberApply = true;
    }
}
contract AdminGroup is EvfsGroup {
    struct ChainNodeReq {
        address     addr;
        uint256     amount;
        uint256     nodeType;
        uint256     op;
        bytes       info;
        address     orgAddress;
        bytes       orgName;
    }
    struct DataStorageDomainReq {
        bytes       name;
        address[]   members;
        uint256     amount;
        address     orgAddress;
        bytes       orgName;
    }
    struct StorageLicenseReq {
        address     dsId;
        uint256     amount;
    }
    struct ChainNode {
        address     addr;
        uint256     nodeType;
        bytes       info;
        uint256     status;
    }

    ERC20                                       _dataStorageToken;
    DataStorageDomain                           _dataStorageDomain;

    mapping (bytes32 => DataStorageDomainReq)   _dataStorageDomainReqs;
    mapping (bytes32 => StorageLicenseReq)      _storageLicenseReqs;

    ERC20                                       _accountNodeLicenseToken;
    ERC20                                       _syncNodeLicenseToken;
    ERC20                                       _clientLicenseToken;

    mapping (address => ChainNode)              _chainAccountNode;
    address[]                                   _chainAccountNodeArray;
    mapping (address => ChainNode)              _chainSyncNode;
    address[]                                   _chainSyncNodeArray;
    mapping (bytes32 => ChainNodeReq)           _chainNodeReqs;
    mapping (address => ChainNode)              _chainClient;
    address[]                                   _chainClientArray;

    mapping (address => bytes)                  _orgs;
    mapping (address => address)                _orgRelations;

    constructor (
        address[] members,
        address _dataStorageTokenAddr,
        address _accountNodeLicenseTokenAddr,
        address _syncNodeLicenseTokenAddr,
        address _clientLicenseTokenAddr,
        address _parent
    ) 
    EvfsGroup(members) 
    public {
        _dataStorageToken = ERC20(_dataStorageTokenAddr);
        _accountNodeLicenseToken = ERC20(_accountNodeLicenseTokenAddr);
        _syncNodeLicenseToken = ERC20(_syncNodeLicenseTokenAddr);
        _clientLicenseToken = ERC20(_clientLicenseTokenAddr);
        parentGroupAddress = _parent;
        _canMemberApply = false;
    }

    // 设置存管域合约
    function setStorageDomain(address _addr) external onlyManager returns (bool){
        _dataStorageDomain = DataStorageDomain(_addr);
        return true;
    }
    // 创建存管域
    function dataStorageDomainApply(address _orgAddress, bytes _orgName, bytes _name, address[] _members, uint256 _amount) external onlyMember returns (bytes32 _reqId, bool _isEnd, address _dsId) {
        _reqId = generateLockId();

        _dataStorageDomainReqs[_reqId] = DataStorageDomainReq({
            name: _name,
            members: _members,
            amount: _amount,
            orgAddress: _orgAddress,
            orgName: _orgName
        });

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;

            addressIndexByTodo[_reqId].push(member);
            todoIndexByAddress[member].push(_reqId);
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (groupMembers[msg.sender] == msg.sender) {
            (_isEnd, _dsId) = dataStorageDomainAgree(_reqId);
        }
    }
    function dataStorageDomainAgree(bytes32 _reqId) public onlyMember canConfirm(_reqId) returns (bool _isCreated, address _dsId) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            DataStorageDomainReq memory _dsReq = _dataStorageDomainReqs[_reqId];
            _dsId = address(generateLockId());
            _dataStorageDomain.createDomain(_dsId, address(0), _dsReq.name, _dsReq.members);
            _dataStorageToken.transferFrom(address(this), _dsId, _dsReq.amount);

            _orgRelations[_dsId] = _dsReq.orgAddress;
            if (_orgs[_dsReq.orgAddress].length == uint256(0)) {
               _orgs[_dsReq.orgAddress] = _dsReq.orgName;
            }
            clearAll(_reqId);

            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isCreated = _confirmResult.confirmResult != uint256(0);

    }
    // 发放容量许可
    function storageLicenseApply(address _dsId, uint256 _amount) external onlyMember returns (bytes32 _reqId, bool _isEnd) {
        // todo must exsits
        _reqId = generateLockId();
        _storageLicenseReqs[_reqId] = StorageLicenseReq({
            dsId: _dsId,
            amount: _amount
        });

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (groupMembers[msg.sender] == msg.sender) {
            _isEnd = storageLicenseAgree(_reqId);
        }
    }
    function storageLicenseAgree(bytes32 _reqId) public onlyMember canConfirm(_reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            _dataStorageToken.transferFrom(address(this), _storageLicenseReqs[_reqId].dsId ,_storageLicenseReqs[_reqId].amount);

            clearAll(_reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
    function chainNodeApply(address _orgAddress, bytes _orgName, address _nodeAddr,bytes _info, uint256 _amount, uint256 _nodeType, uint256 _op) external returns (bytes32 _reqId, bool _isEnd) {
        _reqId = generateLockId();
        _chainNodeReqs[_reqId] = ChainNodeReq({
            addr: _nodeAddr,
            amount: _amount,
            nodeType: _nodeType,
            op: _op,
            info: _info,
            orgAddress: _orgAddress,
            orgName: _orgName
        });

        uint256 len = groupMembersArray.length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: rule
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = groupMembersArray[i];
            todosById[_reqId][member] = member;
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (groupMembers[msg.sender] == msg.sender) {
            _isEnd = chainNodeAgree(_reqId);
        }
    }
    function chainNodeAgree(bytes32 _reqId) public onlyMember canConfirm(_reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            ChainNodeReq memory _chainNodeReq = _chainNodeReqs[_reqId];
            if (_chainNodeReq.op == uint256(1)) {
                if (_orgs[_chainNodeReq.orgAddress].length == uint256(0)) {
                    _orgs[_chainNodeReq.orgAddress] = _chainNodeReq.orgName;
                }
                _orgRelations[_chainNodeReq.addr] = _chainNodeReq.orgAddress;

                // new
                if (_chainNodeReq.nodeType == uint256(1)) {
                    // account node
                    _accountNodeLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainAccountNode[_chainNodeReq.addr] = ChainNode({
                        addr: _chainNodeReq.addr,
                        nodeType: _chainNodeReq.nodeType,
                        info: _chainNodeReq.info,
                        status: uint256(1)
                    });
                    _chainAccountNodeArray.push(_chainNodeReq.addr);
                } else if (_chainNodeReq.nodeType == uint256(2)) {
                    // sync node
                    _syncNodeLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainSyncNode[_chainNodeReq.addr] = ChainNode({
                        addr: _chainNodeReq.addr,
                        nodeType: _chainNodeReq.nodeType,
                        info: _chainNodeReq.info,
                        status: uint256(1)
                    });
                    _chainSyncNodeArray.push(_chainNodeReq.addr);
                } else {
                    _clientLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainClient[_chainNodeReq.addr] = ChainNode({
                        addr: _chainNodeReq.addr,
                        nodeType: _chainNodeReq.nodeType,
                        info: _chainNodeReq.info,
                        status: uint256(1)
                    });
                    _chainClientArray.push(_chainNodeReq.addr);
                }
            } else {
                _orgRelations[_chainNodeReq.addr] = address(0);
                if (_chainNodeReq.nodeType == uint256(1)) {
                    // account node
                    _accountNodeLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainAccountNode[_chainNodeReq.addr].status = uint256(2);
                } else if (_chainNodeReq.nodeType == uint256(2)) {
                    // sync node
                    _syncNodeLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainSyncNode[_chainNodeReq.addr].status = uint256(2);
                } else {
                    _clientLicenseToken.transferFrom(address(this), _chainNodeReq.addr, _chainNodeReq.amount);
                    _chainClient[_chainNodeReq.addr].status = uint256(2);
                }
            }
            clearAll(_reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
}
contract EvfsDomain is iEvfsDomain, LockIdGen {
    using SafeMath for uint256;
    struct DomainItem {
        address     id;
        address     parentId;
        bytes       name;
        uint256     status;
    }
    struct MemberRequest {
        address member;
        uint256 op;
        address domainId;
    }
    struct ConfirmResult {
        uint256 confirmResult;
        uint256 confirmOn;
        uint256 total;
        uint256 agrees;
        uint256 disagrees;
        uint256 ruleBase;
        uint256 rule;
    }
    
    uint256                                                     ruleBase;
    // ERC20                                               _managementToken;
    address                                                     _managementAddress;
    mapping (address => uint256)                                _domainStatus;
    mapping (address => DomainItem)                             _domains;

    mapping (address => mapping(address => address))            _domainMembers;
    mapping (address => address[])                              _domainMemberArray;

    mapping (bytes32 => ConfirmResult)                          confirmResults;
    mapping (address => mapping (bytes32 => mapping (address => address)))           todosById;
    mapping (address => mapping (address => mapping (bytes32 => uint256)))           todosHistoryAddress;

    address                                                     manager;
    mapping (address => mapping (bytes32 => uint256))           ruleReqs;
    mapping (address => uint256)                                _domainRules;
    mapping (bytes32 => MemberRequest)                          memberReqs;

    mapping (address => mapping (bytes32 => address[]))         addressIndexByTodo;
    mapping (address => mapping (address => bytes32[]))         todoIndexByAddress;

    mapping (address => bytes)                                  _orgs;
    mapping (address => address)                                _orgRelations;

    constructor (
        address _parent
    )
    public {
        _managementAddress = _parent;
        manager = msg.sender;
        // _managementToken = ERC20(_tokenAddr);
        ruleBase = 300;
    }

    modifier onlyManager {
        require(msg.sender == manager, "only manager");
        _;
    }
    modifier onlyParent {
        require (msg.sender == _managementAddress, "unAuthorization");
        _;
    }
    modifier onlyMember(address _domainId) {
        require (msg.sender == _domainMembers[_domainId][msg.sender], "unAuthorization");
        _;
    }
    modifier canMemberApply(address _domainId, address _member) {
        // require (_canMemberApply, "cannot apply");
        // require(_subGroupAddress == address(0) || subGroups[_subGroupAddress] == _subGroupAddress, "cannot apply");
        require(_domainMembers[_domainId][_member] == address(0), "member exists");
        _;
    }
    modifier canRuleApply(address _domainId, uint256 _rule) {
        require(_rule <= ruleBase, "invalidate rule");
        require(msg.sender == _domainMembers[_domainId][msg.sender], "not in group");
        _;
    }
    modifier canConfirm(address _domainId, bytes32 _reqId) {
        require(confirmResults[_reqId].confirmResult == uint256(0), "request finished");
        require(todosById[_domainId][_reqId][msg.sender] == msg.sender, "cannot confirm");
        require(todosHistoryAddress[_domainId][msg.sender][_reqId] == 0, "already confirmed");
        _;
    }
    function setManager(address _newManager) external onlyManager {
        manager = _newManager;
    }
    function createDomain(address _domainId, address _pdomainId, bytes _name, address[] _members) external onlyParent {
        require (_domains[_domainId].id == address(0), "domain exists");

        _domains[_domainId] = DomainItem({
            id: _domainId,
            parentId: _pdomainId,
            name: _name,
            status: uint256(1)
        });
        
        _domainRules[_domainId] = 200;
        
        uint256 len = _members.length;
        for (uint256 i=0; i<len; i++) {
            _domainMembers[_domainId][_members[i]] = _members[i];
            _domainMemberArray[_domainId].push(_members[i]);
        }
    }
    function existsDomain(address _domainId) external onlyParent returns (bool) {
        return _domains[_domainId].id == _domainId;
    }
    function disagree(address _domainId, bytes32 _reqId) external canConfirm(_domainId, _reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.disagrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(3);

        if (_confirmResult.disagrees.mul(1000).div(_confirmResult.total) > _confirmResult.ruleBase.sub(_confirmResult.rule).mul(1000).div(_confirmResult.ruleBase)) {
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(2);

            clearAll(_domainId, _reqId);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
    function setStatus(address _domainId, uint256 _status) external onlyParent {
        require(_domains[_domainId].id == _domainId, "not exists");
        _domains[_domainId].status = _status;
    }
    function clear(address _domainId, address _addr, bytes32 _reqId) internal {
        uint256 len1 = addressIndexByTodo[_domainId][_reqId].length;
        for (uint256 j=0; j<len1; j++) {
            if (addressIndexByTodo[_domainId][_reqId][j] == _addr) {
                // 移位，删除元素
                for (uint k = j; k<len1-1; k++){
                    addressIndexByTodo[_domainId][_reqId][k] = addressIndexByTodo[_domainId][_reqId][k+1];
                }
                delete addressIndexByTodo[_domainId][_reqId][len1-1];
                addressIndexByTodo[_domainId][_reqId].length--;
                break;
            }
        }

        len1 = todoIndexByAddress[_domainId][_addr].length;
        for (j=0; j<len1; j++) {
            if (todoIndexByAddress[_domainId][_addr][j] == _reqId) {
                // 移位，删除元素
                for (k = j; k<len1-1; k++){
                    todoIndexByAddress[_domainId][_addr][k] = todoIndexByAddress[_domainId][_addr][k+1];
                }
                delete todoIndexByAddress[_domainId][_addr][len1-1];
                todoIndexByAddress[_domainId][_addr].length--;
                break;
            }
        }
    }
    function clearAll(address _domainId, bytes32 _reqId) internal {
        address[] _addrs = addressIndexByTodo[_domainId][_reqId];
        

        uint256 len = _addrs.length;
        for (uint256 i=0; i<len; i++) {
            uint256 len1 = todoIndexByAddress[_domainId][_addrs[i]].length;

            for (uint256 j=0; j<len1; j++) {
                if (todoIndexByAddress[_domainId][_addrs[i]][j] == _reqId) {
                    // 移位，删除元素
                    for (uint k = j; k<len-1; k++){
                        todoIndexByAddress[_domainId][_addrs[i]][k] = todoIndexByAddress[_domainId][_addrs[i]][k+1];
                    }
                    delete todoIndexByAddress[_domainId][_addrs[i]][len1-1];
                    todoIndexByAddress[_domainId][_addrs[i]].length--;
                    break;
                }
            }
        }

        delete addressIndexByTodo[_domainId][_reqId];
    }
    function ruleApply(address _domainId, uint256 _rule) external canRuleApply(_domainId, _rule) onlyMember(_domainId) returns (bytes32 _reqId, bool _isEnd) {
        _reqId = generateLockId();
        ruleReqs[_domainId][_reqId] = _rule;

        uint256 len = _domainMemberArray[_domainId].length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: _domainRules[_domainId]
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = _domainMemberArray[_domainId][i];
            todosById[_domainId][_reqId][member] = member;

            addressIndexByTodo[_domainId][_reqId].push(member);
            todoIndexByAddress[_domainId][member].push(_reqId);
        }

        if (_domainMembers[_domainId][msg.sender] == msg.sender) {
            _isEnd = ruleAgree(_domainId, _reqId);
        }
    }
    function ruleAgree(address _domainId, bytes32 _reqId) public canConfirm(_domainId, _reqId) returns (bool _isEnd) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            _domainRules[_domainId] = ruleReqs[_domainId][_reqId];
            
            clearAll(_domainId, _reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isEnd = _confirmResult.confirmResult != uint256(0);
    }
    function memberApply(address _domainId, address _member, uint256 _op) external onlyMember(_domainId) canMemberApply(_domainId, _member) returns (bytes32 _reqId, bool _isEnd) {
        require(_member != address(0) , "0");
        _isEnd = false;

        _reqId = generateLockId();
        memberReqs[_reqId] = MemberRequest({
            member: _member,
            op: _op,
            domainId: _domainId
        });

        uint256 len = _domainMemberArray[_domainId].length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: _domainRules[_domainId]
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = _domainMemberArray[_domainId][i];
            todosById[_domainId][_reqId][member] = member;
        }

        if (_domainMembers[_domainId][msg.sender] == msg.sender) {
            _isEnd = memberAgree(_domainId, _reqId);
        }
    }
    function memberAgree(address _domainId, bytes32 _reqId) public canConfirm(_domainId, _reqId) returns (bool) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(2);
        if ((_confirmResult.agrees.mul(1000).div(_confirmResult.total)) >= (_confirmResult.rule.mul(1000).div(_confirmResult.ruleBase))) {
            // approved
            MemberRequest memory memberRequest = memberReqs[_reqId];
            address _member = memberRequest.member;
            if (memberRequest.op == uint256(1)) {
                require(_domainMembers[memberRequest.domainId][_member] == address(0), "duplicate member");
                _domainMembers[memberRequest.domainId][_member] = _member;
                _domainMemberArray[memberRequest.domainId].push(_member);
            } else if (memberRequest.op == uint256(2)) {
                _domainMembers[memberRequest.domainId][_member] = address(0);
                uint256 len = _domainMemberArray[memberRequest.domainId].length;
                for (uint256 i=0; i<len; i++) {
                    if (_domainMemberArray[memberRequest.domainId][i] == _member) {
                        for (uint j = i; j<len-1; j++){
                            _domainMemberArray[memberRequest.domainId][j] = _domainMemberArray[memberRequest.domainId][j+1];
                        }
                        delete _domainMemberArray[memberRequest.domainId][len-1];
                        _domainMemberArray[memberRequest.domainId].length--;
                        break;
                    }
                }
            } else {
                // todo other
            }

            clearAll(_domainId, _reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        return _confirmResult.confirmResult != uint256(0);
    }
}
contract DataStorageDomain is EvfsDomain {
    struct BusinessDomainReq {
        address     domainId;
        bytes       name;
        address[]   members;
        uint256     op;
        uint256     status;
        address     orgAddress;
        bytes       orgName;
    }
    struct DataStorageNodeReq {
        address     nodeAddr;
        bytes       name;
        uint256     amount;
        uint256     op;
        bytes       url;
    }
    struct DataStroageNode {
        address     nodeAddr;
        bytes       name;
        uint256     status;
        bytes       url;
    }
    
    BusinessDomain                                      _businessDomain;
    mapping (address => mapping (bytes32 => BusinessDomainReq))              _businessDomainReqs;

    mapping (address => mapping (bytes32 => DataStorageNodeReq))    _dataStorageNodeReqs;
    mapping (address => mapping (address => DataStroageNode))       _dataStorageNodes;
    mapping (address => address[])                                  _dataStorageNodeArray;

    ERC20                                       _dataStorageToken;

    constructor (
        address _parent,
        address _dataStorageTokenAddr
    ) EvfsDomain(_parent)
    public {
        _dataStorageToken = ERC20(_dataStorageTokenAddr);
    }

    function setBusinessDomain(address _addr) public onlyManager returns (bool) {
        _businessDomain = BusinessDomain(_addr);
        return true;
    }

    // 业务域
    function businessDomainApply(address _orgAddress, bytes _orgName, address _domainId, bytes _name, address[] _members, uint256 _op) external onlyMember(_domainId) returns (bytes32 _reqId, bool _isEnd, address _bizId) {
        _reqId = generateLockId();

        _businessDomainReqs[_domainId][_reqId] = BusinessDomainReq({
            domainId: _domainId,
            name: _name,
            members: _members,
            op: _op,
            status: uint256(0),
            orgAddress: _orgAddress,
            orgName: _orgName
        });

        uint256 len = _domainMemberArray[_domainId].length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: _domainRules[_domainId]
        });
        (_isEnd,_bizId) = assignBusinessApplyTask(_domainId, _reqId, len);
    }
    
    function assignBusinessApplyTask(address _domainId, bytes32 _reqId, uint256 _len) internal returns (bool _isEnd, address _bizId){
        // 分配任务到当前组成员
        for (uint256 i=0; i<_len; i++) {
            address member = _domainMemberArray[_domainId][i];
            todosById[_domainId][_reqId][member] = member;
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (_domainMembers[_domainId][msg.sender] == msg.sender) {
            (_isEnd,_bizId) = businessDomainAgree(_domainId, _reqId);
        }
    }

    function businessDomainAgree(address _domainId, bytes32 _reqId) public onlyMember(_domainId) canConfirm(_domainId, _reqId) returns (bool _isEnd, address _bizId) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            BusinessDomainReq memory _bizDomainReq = _businessDomainReqs[_domainId][_reqId];
            if (_bizDomainReq.op == uint256(1)) {
                _bizId = address(generateLockId());
                _businessDomain.createDomain(_bizId, _domainId, _bizDomainReq.name, _bizDomainReq.members);

                if (_orgs[_bizDomainReq.orgAddress].length == uint256(0)) {
                    _orgs[_bizDomainReq.orgAddress] = _bizDomainReq.orgName;
                }
                _orgRelations[_bizId] = _bizDomainReq.orgAddress;
            } else {
                // remove
                _businessDomain.setStatus(_bizDomainReq.domainId, _bizDomainReq.status);
                _orgRelations[_bizDomainReq.domainId] = address(0);
            }
            clearAll(_domainId, _reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isEnd = _confirmResult.confirmResult != uint256(0);
    }
    
    function dataStorageNodeApply(address _domainId, address _nodeAddress, bytes _name, bytes _url, uint256 _amount, uint256 _op) external returns (bytes32 _reqId, bool _isEnd) {
        _reqId = generateLockId();

        _dataStorageNodeReqs[_domainId][_reqId] = DataStorageNodeReq({
            nodeAddr: _nodeAddress,
            amount: _amount,
            op: _op,
            name: _name,
            url: _url
        });

        uint256 _len = _domainMemberArray[_domainId].length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: _len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: _domainRules[_domainId]
        });
        
        _isEnd = assignDataStorageNodeApplyTask(_domainId, _reqId, _len);
    }
    
    function assignDataStorageNodeApplyTask(address _domainId, bytes32 _reqId, uint256 _len) internal returns (bool _isEnd){
        // 分配任务到当前组成员
        for (uint256 i=0; i<_len; i++) {
            address member = _domainMemberArray[_domainId][i];
            todosById[_domainId][_reqId][member] = member;
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (_domainMembers[_domainId][msg.sender] == msg.sender) {
            _isEnd = dataStorageNodeAgree(_domainId, _reqId);
        }
    }

    function dataStorageNodeAgree(address _domainId, bytes32 _reqId) public onlyMember(_domainId) canConfirm(_domainId, _reqId) returns (bool _isEnd) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            DataStorageNodeReq memory _nodeReq = _dataStorageNodeReqs[_domainId][_reqId];
            if (_nodeReq.op == uint256(1)) {
                _dataStorageNodes[_domainId][_nodeReq.nodeAddr] = DataStroageNode({
                    nodeAddr: _nodeReq.nodeAddr,
                    name: _nodeReq.name,
                    status: _nodeReq.op,
                    url: _nodeReq.url
                });
                _dataStorageNodeArray[_domainId].push(_nodeReq.nodeAddr);

                if (_nodeReq.amount > 0) {
                    _dataStorageToken.transferFrom(address(this), _nodeReq.nodeAddr, _nodeReq.amount);
                }
            } else {
                // remove
                _dataStorageNodes[_domainId][_nodeReq.nodeAddr].status = _nodeReq.op;
            }
            clearAll(_domainId, _reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isEnd = _confirmResult.confirmResult != uint256(0);
    }
}
contract BusinessDomain is EvfsDomain {
    struct BizSystemReq {
        // address domainId;
        address sysAddress;
        bytes   name;
        uint256 op;
        address orgAddress;
        bytes   orgName;
    }
    struct BizSystem {
        address addr;
        bytes name;
        uint256 status;
    }

    mapping (bytes32 => BizSystemReq)                       _bizSystemReqs;
    mapping (address => mapping (address => BizSystem))     _bizSystems;
    mapping (address => address[])                          _bizSystemArray;

    mapping (address => mapping (bytes32 => bytes))         _datas;
    mapping (address => mapping (bytes32 => bytes32))       _dataRels;

    constructor (
        address _parent
    ) 
    EvfsDomain(_parent)
    public { }

    modifier onlyBizSystem(address _domainId, address _bizSystem) {
        require(_bizSystem == _bizSystems[_domainId][_bizSystem].addr, "only bizSystem");
        require(uint256(1) == _bizSystems[_domainId][_bizSystem].status, "bizSystem disabled");
        _;
    }

    function bizSystemApply(address _orgAddress, bytes _orgName, address _domainId, address _bizSystem, bytes _name, uint256 _op) external onlyMember(_domainId) returns (bytes32 _reqId, bool _isEnd) {
        require(_bizSystems[_domainId][_bizSystem].addr == address(0), "system exists");
        require(_name.length > 0, "invalidate name");
        
        _reqId = generateLockId();

        _bizSystemReqs[_reqId] = BizSystemReq({
            // domainId: _domainId,
            sysAddress: _bizSystem,
            name: _name,
            op: _op,
            orgAddress: _orgAddress,
            orgName: _orgName
        });

        uint256 len = _domainMemberArray[_domainId].length;
        confirmResults[_reqId] = ConfirmResult({
            confirmResult: 0,
            confirmOn: 0,
            total: len,
            agrees: 0,
            disagrees: 0,
            ruleBase: ruleBase,
            rule: _domainRules[_domainId]
        });

        // 分配任务到当前组成员
        for (uint256 i=0; i<len; i++) {
            address member = _domainMemberArray[_domainId][i];
            todosById[_domainId][_reqId][member] = member;
        }

        // 如果是成员发起的申请，默认自动审批通过
        if (_domainMembers[_domainId][msg.sender] == msg.sender) {
            _isEnd = bizSystemAgree(_domainId, _reqId);
        }
    }

    function bizSystemAgree(address _domainId, bytes32 _reqId) public onlyMember(_domainId) canConfirm(_domainId, _reqId) returns (bool _isEnd) {
        ConfirmResult memory _confirmResult = confirmResults[_reqId];
        _confirmResult.agrees += 1;
        todosHistoryAddress[_domainId][msg.sender][_reqId] = uint256(2);
        if (_confirmResult.agrees.mul(1000).div(_confirmResult.total) >= _confirmResult.rule.mul(1000).div(_confirmResult.ruleBase)) {
            // approved
            BizSystemReq memory _bizSysReq = _bizSystemReqs[_reqId];
            if (_bizSysReq.op == uint256(1)) {
                if (_orgs[_bizSysReq.orgAddress].length == uint256(0)) {
                    _orgs[_bizSysReq.orgAddress] = _bizSysReq.orgName;
                }
                _orgRelations[_bizSysReq.sysAddress] = _bizSysReq.orgAddress;

                _bizSystems[_domainId][_bizSysReq.sysAddress] = BizSystem({
                    addr: _bizSysReq.sysAddress,
                    name: _bizSysReq.name,
                    status: uint256(1)
                });
                
                _bizSystemArray[_domainId].push(_bizSysReq.sysAddress);
            } else {
                address _delSysAddr = _bizSysReq.sysAddress;
                _bizSystems[_domainId][_bizSysReq.sysAddress].status = uint256(2);
                _orgRelations[_delSysAddr] = address(0);

                uint256 len = _bizSystemArray[_domainId].length;
                for (uint256 i=0; i<len; i++) {
                    if (_bizSystemArray[_domainId][i] == _delSysAddr) {
                        // 移位，删除元素
                        for (uint j = i; j<len-1; j++){
                            _bizSystemArray[_domainId][j] = _bizSystemArray[_domainId][j+1];
                        }
                        delete _bizSystemArray[_domainId][len-1];
                        _bizSystemArray[_domainId].length--;
                        break;
                    }
                }
            }
            clearAll(_domainId, _reqId);
            _confirmResult.confirmOn = block.number;
            _confirmResult.confirmResult = uint256(1);
        } else {
            clear(_domainId, msg.sender, _reqId);
        }
        confirmResults[_reqId] = _confirmResult;
        _isEnd = _confirmResult.confirmResult != uint256(0);
    }

    function setDataStorage(address _domainId, address _bizSystem, bytes _data, bytes32 _refId) external onlyBizSystem(_domainId, _bizSystem) returns (bytes32 _dataId) {
        _dataId = generateLockId();
        _datas[_domainId][_dataId] = _data;
        _dataRels[_domainId][_dataId] = _refId;
    }
}
