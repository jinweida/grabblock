package common

const (
	// 获取区块信息
	URL_BlOCK_INFO = "/fbs/bct/pbgbn.do"
	// 获取交易信息
	URL_TRANS_INFO = "/fbs/tct/pbgth.do"
	// 当前高度
	URL_BLOCK_STATE               = "/fbs/sio/pbcio.do"
	URL_TRANSFER_BALANCE          = "/fbs/act/pbgac.do"
	URL_TOKEN_BALANCE             = "/fbs/c20/pbqvalue.do"
	INNER_CODE_TYPE_TRANSFER      = 0  //普通交易
	INNER_CODE_TYPE_CVM           = 4  //CVM合约
	INNER_CODE_TYPE_CVM_CREATE    = 20 //创建合约
	INNER_CODE_TYPE_CVM_CALL      = 21 //调用合约
	INNER_CODE_TYPE_FILE          = 6  //文件
	INNER_CODE_TYPE_RC20          = 2  //token
	INNER_CODE_TYPE_RC20_CREATE   = 2  //token 创建
	INNER_CODE_TYPE_RC20_TRANSFER = 3  //token 转账
	RET_FAIL_CODE                 = -1
	RET_SUCCESS_CODE              = 1

	TRANSACTION_IN       = 0
	TRANSACTION_OUT      = 1
	TRANSACTION_OUT_TEXT = "out"
	TRANSACTION_IN_TEXT  = "in"

	OP_ADD    = 1 //操作类型：添加
	OP_REMOVE = 2 //操作类型：删除

	APPROVE_ING     = 0 //待审
	APPROVE_SUCCESS = 1 //通过

	NODE_TYPE_ACCOUNT = 1 //记账节点
	NODE_TYPE_SYNC    = 2 //只读节点
	NODE_TYPE_CLIENT  = 3 //前置节点

	GRABBLOCK_ON  = 0
	GRABBLOCK_OFF = 1

	TRANSACTION_SUCCESS = "0x01"

	PARSING_NOT         = 2
	PARSING_DONE        = 1
	PARSING_STATUS_FAIL = 0

	FILE_TYPE_VIDEO = 1
	FILE_TYPE_SOUND = 2
	FILE_TYPE_PIC   = 3
	FILE_TYPE_OTHER = 4
)
const (
	TimeTemplate = "2006-01-02 00:00:00"
)
const (
	AdminGroup = `[
  {
    "constant": false,
    "inputs": [
      {
        "name": "_member",
        "type": "address"
      }
    ],
    "name": "removeMember",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "memberAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_orgAddress",
        "type": "address"
      },
      {
        "name": "_orgName",
        "type": "bytes"
      },
      {
        "name": "_nodeAddr",
        "type": "address"
      },
      {
        "name": "_info",
        "type": "bytes"
      },
      {
        "name": "_amount",
        "type": "uint256"
      },
      {
        "name": "_nodeType",
        "type": "uint256"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "chainNodeApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_dsId",
        "type": "address"
      },
      {
        "name": "_amount",
        "type": "uint256"
      }
    ],
    "name": "storageLicenseApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "getMember",
    "outputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "storageLicenseAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "ruleAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "setStorageDomain",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_memberNames",
        "type": "bytes"
      },
      {
        "name": "_members",
        "type": "address[]"
      },
      {
        "name": "_amount",
        "type": "uint256"
      },
      {
        "name": "_rule",
        "type": "uint256"
      }
    ],
    "name": "dataStorageDomainApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      },
      {
        "name": "_dsId",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_member",
        "type": "address"
      },
      {
        "name": "_op",
        "type": "uint256"
      },
      {
        "name": "_subGroupAddress",
        "type": "address"
      }
    ],
    "name": "memberApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_key",
        "type": "bytes32"
      }
    ],
    "name": "getConfig",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "performMemberApply",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "chainNodeAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_configKey",
        "type": "bytes32"
      },
      {
        "name": "_configValue",
        "type": "bytes"
      }
    ],
    "name": "configApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getParent",
    "outputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getMembers",
    "outputs": [
      {
        "name": "",
        "type": "address[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getSubs",
    "outputs": [
      {
        "name": "",
        "type": "address[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "dataStorageDomainAgree",
    "outputs": [
      {
        "name": "_isCreated",
        "type": "bool"
      },
      {
        "name": "_dsId",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_rule",
        "type": "uint256"
      }
    ],
    "name": "ruleApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_member",
        "type": "address"
      }
    ],
    "name": "addMember",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "subJoin",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "disagree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "configAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "isAllowJoin",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "memberNames",
        "type": "bytes[]"
      },
      {
        "name": "members",
        "type": "address[]"
      },
      {
        "name": "_tokens",
        "type": "address[]"
      },
      {
        "name": "_parent",
        "type": "address"
      },
      {
        "name": "_rule",
        "type": "uint256"
      },
      {
        "name": "accountNodesName",
        "type": "bytes[]"
      },
      {
        "name": "accountNodes",
        "type": "address[]"
      },
      {
        "name": "orgInfos",
        "type": "bytes[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`
	BussinessDomain = `[
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_code",
        "type": "bytes"
      },
      {
        "name": "_info",
        "type": "bytes"
      }
    ],
    "name": "createContractApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      },
      {
        "name": "_contractAddress",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      }
    ],
    "name": "existsDomain",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_bizSystem",
        "type": "address"
      },
      {
        "name": "_data",
        "type": "bytes"
      },
      {
        "name": "_refId",
        "type": "bytes32"
      }
    ],
    "name": "setDataStorage",
    "outputs": [
      {
        "name": "_dataId",
        "type": "bytes32"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_pdomainId",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_members",
        "type": "address[]"
      },
      {
        "name": "_storage",
        "type": "uint256"
      },
      {
        "name": "_rule",
        "type": "uint256"
      },
      {
        "name": "_isEnableManage",
        "type": "bool"
      },
      {
        "name": "_isEnableDelFile",
        "type": "bool"
      }
    ],
    "name": "createDomain",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "memberAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "configAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "dataOrFileEnableAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "contractEnableAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_configKey",
        "type": "bytes32"
      },
      {
        "name": "_configValue",
        "type": "bytes"
      }
    ],
    "name": "configApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_status",
        "type": "uint256"
      }
    ],
    "name": "setStatus",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "ruleAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_rule",
        "type": "uint256"
      }
    ],
    "name": "ruleApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_contractAddress",
        "type": "address"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "contractEnableApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_newManager",
        "type": "address"
      }
    ],
    "name": "setManager",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "bizSystemAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "disagree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "createContractAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      },
      {
        "name": "_contractAddress",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_orgAddress",
        "type": "address"
      },
      {
        "name": "_orgName",
        "type": "bytes"
      },
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_bizSystem",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "bizSystemApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_member",
        "type": "address"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "memberApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_dataOrFileHash",
        "type": "bytes32"
      },
      {
        "name": "_op",
        "type": "uint256"
      },
      {
        "name": "_dataOrFile",
        "type": "uint256"
      }
    ],
    "name": "dataOrFileEnableApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "_parent",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`
	CommitteeGroup = `[
  {
    "constant": false,
    "inputs": [
      {
        "name": "_member",
        "type": "address"
      }
    ],
    "name": "removeMember",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "memberAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "getMember",
    "outputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "ruleAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_member",
        "type": "address"
      },
      {
        "name": "_op",
        "type": "uint256"
      },
      {
        "name": "_subGroupAddress",
        "type": "address"
      }
    ],
    "name": "memberApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_key",
        "type": "bytes32"
      }
    ],
    "name": "getConfig",
    "outputs": [
      {
        "name": "",
        "type": "bytes"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "performMemberApply",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_configKey",
        "type": "bytes32"
      },
      {
        "name": "_configValue",
        "type": "bytes"
      }
    ],
    "name": "configApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getParent",
    "outputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getMembers",
    "outputs": [
      {
        "name": "",
        "type": "address[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "getSubs",
    "outputs": [
      {
        "name": "",
        "type": "address[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_rule",
        "type": "uint256"
      }
    ],
    "name": "ruleApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_member",
        "type": "address"
      }
    ],
    "name": "addMember",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "subJoin",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "disagree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "configAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [],
    "name": "isAllowJoin",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "_defaultRule",
        "type": "uint256"
      },
      {
        "name": "memberNames",
        "type": "bytes[]"
      },
      {
        "name": "members",
        "type": "address[]"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`
	DataStorageDomain = `[
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      }
    ],
    "name": "existsDomain",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_pdomainId",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_members",
        "type": "address[]"
      },
      {
        "name": "_storage",
        "type": "uint256"
      },
      {
        "name": "_rule",
        "type": "uint256"
      },
      {
        "name": "_isEnableManage",
        "type": "bool"
      },
      {
        "name": "_isEnableDelFile",
        "type": "bool"
      }
    ],
    "name": "createDomain",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "memberAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "configAgree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_configKey",
        "type": "bytes32"
      },
      {
        "name": "_configValue",
        "type": "bytes"
      }
    ],
    "name": "configApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_addr",
        "type": "address"
      }
    ],
    "name": "setBusinessDomain",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_status",
        "type": "uint256"
      }
    ],
    "name": "setStatus",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "ruleAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_rule",
        "type": "uint256"
      }
    ],
    "name": "ruleApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "businessDomainAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      },
      {
        "name": "_bizId",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_nodeAddress",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_url",
        "type": "bytes"
      },
      {
        "name": "_amount",
        "type": "uint256"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "dataStorageNodeApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_names",
        "type": "bytes"
      },
      {
        "name": "_members",
        "type": "address[]"
      },
      {
        "name": "_op",
        "type": "uint256[]"
      }
    ],
    "name": "businessDomainApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      },
      {
        "name": "_bizId",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "dataStorageNodeAgree",
    "outputs": [
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_newManager",
        "type": "address"
      }
    ],
    "name": "setManager",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_reqId",
        "type": "bytes32"
      }
    ],
    "name": "disagree",
    "outputs": [
      {
        "name": "",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "_domainId",
        "type": "address"
      },
      {
        "name": "_name",
        "type": "bytes"
      },
      {
        "name": "_member",
        "type": "address"
      },
      {
        "name": "_op",
        "type": "uint256"
      }
    ],
    "name": "memberApply",
    "outputs": [
      {
        "name": "_reqId",
        "type": "bytes32"
      },
      {
        "name": "_isEnd",
        "type": "bool"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "_parent",
        "type": "address"
      },
      {
        "name": "_dataStorageTokenAddr",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  }
]`
)
