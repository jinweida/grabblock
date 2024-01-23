package service

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"strings"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/models"
	"example.cn/grabblock/tools/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
)

type AdminGroup struct{}

func NewAdminGroup() *AdminGroup {
	m := &AdminGroup{}
	return m
}

// 节点申请
func (c *AdminGroup) ChainNodeApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		OrgAddress common.Address
		OrgName    []byte
		NodeAddr   common.Address
		Info       []byte
		Amount     *big.Int
		NodeType   *big.Int
		Op         *big.Int // 1 添加  2 删除
	}
	err = abi.Methods["chainNodeApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		ReqId common.Hash
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["chainNodeApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	nodeInfo := &entity.NodeInfo{}
	if err := proto.Unmarshal(input.Info, nodeInfo); err != nil {
		log.Infof("[ChainNodeApply] Failed to parse nodeInfo:", err)
		return err
	}
	orgInfo := &entity.OrgInfo{}
	if err := proto.Unmarshal(input.OrgName, orgInfo); err != nil {
		log.Infof("[ChainNodeApply] Failed to parse OrgInfos:", err)
		return err
	}
	evfsNode := &models.EvfsNode{
		ChainnodeId:     res.ReqId.Hex(),
		OrgId:           strings.ToLower(input.OrgAddress.Hex()),
		OrgName:         strings.TrimSpace(string(input.OrgName)),
		NodeAddress:     strings.ToLower(input.NodeAddr.Hex()),
		NodeInfo:        nodeInfo.Name,
		NodeType:        int32(input.NodeType.Int64()), //1 主节点  2 资源节点  3 前置节点
		Approve:         constant.APPROVE_ING,
		TransactionHash: m.Hash,
		Op:              input.Op.Int64(),
	}
	//1添加
	if int32(input.Op.Int64()) == 1 {
		//申请即同意
		if res.IsEnd {
			evfsNode.Approve = constant.APPROVE_SUCCESS
		}
		return evfsNode.Add()
	} else {
		if res.IsEnd {
			return evfsNode.Delete()
		} else {
			return evfsNode.Add()
		}
	}
}

// 节点申请同意
func (c *AdminGroup) ChainNodeAgree(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		ReqId common.Hash
	}

	err = abi.Methods["chainNodeAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["chainNodeAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsNode := &models.EvfsNode{
		ChainnodeId:     input.ReqId.Hex(),
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsNode.Approve = constant.APPROVE_SUCCESS
		return evfsNode.UpdateApprove()
	}
	return nil
}

// 规则申请 ruleApply(uint256) 返回 (bytes32,_isEnd)
func (c *AdminGroup) RuleApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		Rule *big.Int
	}

	err = abi.Methods["ruleApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		ReqId common.Hash
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["ruleApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("unknown outputs unpack error=", err.Error())
		return err
	}
	evfsChaingroup := &models.EvfsChaingroup{
		GroupId:         string("1"),
		PendingRule:     int32(input.Rule.Int64()),
		Name:            string("管理员"),
		TransactionHash: m.Hash,
	}
	return evfsChaingroup.UpdatePendingRule(res.IsEnd)
}

// 规则申请同意
func (c *AdminGroup) RuleAgree(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		ReqId common.Hash
	}

	err = abi.Methods["ruleAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["ruleAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsChaingroup := &models.EvfsChaingroup{
		GroupId:         string("1"), //1 为管理员
		TransactionHash: m.Hash,
	}
	if res.IsEnd {
		return evfsChaingroup.UpdateRule(res.IsEnd)
	}
	return nil
}

// 存管域申请
func (c *AdminGroup) DataStorageDomainApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	if len(m.Status.Result) == 0 {
		return errors.New(fmt.Sprintf("[DataStorageDomainApply]  hash=%s,m.Status.Result len=%d", m.Hash, len(data)))
	}
	var input struct {
		Name        []byte           //存管域名称
		MemberNames []byte           //存管域-管理员名称
		Members     []common.Address //存管域-管理员地址
		Amount      *big.Int         //容量--忽略
		Rule        *big.Int         //规则 默认200
	}
	err = abi.Methods["dataStorageDomainApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}
	var res struct {
		ReqId common.Hash
		DsId  common.Address
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	//json.NewEncoder(os.Stdout).Encode(res)
	err = abi.Methods["dataStorageDomainApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	dsDomainInfo := &entity.DSDomainInfo{}
	if err := proto.Unmarshal(input.Name, dsDomainInfo); err != nil {
		log.Infof("[DataStorageDomainApply] Failed to parse MemberInfo:", err)
		return err
	}
	evfsStorage := models.EvfsStorage{
		ReqId:       res.ReqId.Hex(),
		DomainId:    strings.ToLower(res.DsId.Hex()),
		StorageName: dsDomainInfo.Name,
		//OrgId:           strings.ToLower(input.OrgAddress.Hex()),
		//OrgName:         strings.TrimSpace(string(input.OrgName)),
		Approve:         constant.APPROVE_ING,
		CapacitySize:    input.Amount.Int64(),
		TransactionHash: m.Hash,
		Members:         make([]models.EvfsStorageMember, 0),
		AcceptTime:      m.Accepttimestamp,
	}
	memberInfo := &entity.MemberInfos{}
	if err := proto.Unmarshal(input.MemberNames, memberInfo); err != nil {
		log.Infof("[DataStorageDomainApply] Failed to parse MemberInfo:", err)
		return err
	}
	for i := 0; i < len(input.Members); i++ {
		mem := models.EvfsStorageMember{
			AddressName:     memberInfo.Name[i].Name,
			Address:         strings.ToLower(input.Members[i].Hex()),
			DomainId:        strings.ToLower(res.DsId.Hex()),
			ReqId:           res.ReqId.Hex(),
			TransactionHash: m.Hash,
			Name:            dsDomainInfo.Name,
			JoinTime:        m.Status.Timestamp,
			Approve:         constant.APPROVE_ING,
			Op:              int64(1),
		}
		if res.IsEnd {
			mem.Approve = constant.APPROVE_SUCCESS
		}
		evfsStorage.Members = append(evfsStorage.Members, mem)
	}
	//申请即同意
	if res.IsEnd {
		evfsStorage.Approve = constant.APPROVE_SUCCESS
		evfsStorage.AcceptTime = m.Status.Timestamp
	}
	return evfsStorage.Add()
}

// 存管域申请同意
func (c *AdminGroup) DataStorageDomainAgree(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		ReqId common.Hash
	}

	err = abi.Methods["dataStorageDomainAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		IsCreated bool
		DsId      common.Address
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["dataStorageDomainAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorage := models.EvfsStorage{
		ReqId:           input.ReqId.Hex(),
		DomainId:        strings.ToLower(res.DsId.Hex()),
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsCreated {
		evfsStorage.Approve = constant.APPROVE_SUCCESS
		return evfsStorage.UpdateApprove()
	}
	return nil
}

// 存储License申请
func (c *AdminGroup) StorageLicenseApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		DsId   common.Address
		Amount *big.Int
	}

	err = abi.Methods["storageLicenseApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		ReqId common.Hash
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["storageLicenseApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorage := models.EvfsStorage{
		DomainId:            strings.ToLower(input.DsId.Hex()),
		SizeTransactionHash: m.Hash,
		PendingCapacitySize: input.Amount.Int64(),
		CapacityReqId:       res.ReqId.Hex(),
		AcceptTime:          m.Status.Timestamp,
	}
	//申请即同意
	return evfsStorage.UpdatePendingSize(res.IsEnd)
}

// 存储License申请同意
func (c *AdminGroup) StorageLicenseAgree(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		ReqId common.Hash
	}

	err = abi.Methods["storageLicenseAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		IsEnd bool
	}
	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["storageLicenseAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	if res.IsEnd {
		evfsStorage := models.EvfsStorage{
			CapacityReqId:       input.ReqId.Hex(),
			RuleTransactionHash: m.Hash,
		}
		return evfsStorage.UpdateSize()
	}
	return nil
}

// 解析创世合约信息--管理员相关
func (c *AdminGroup) analysisGenesis_admingrop(m *entity.TransactionInfo, abi abi.ABI) (int64, error) {
	codedata := constant.GetParamsHexStr(m.Body.CodeData)
	data, err := hex.DecodeString(codedata)
	if err != nil {
		log.Errorf("unknown code_data error=%s", err.Error())
		return constant.PARSING_DONE, err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return constant.PARSING_DONE, errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	//委员会创建了管理员--所以有父。父（Parent）为委员会地址，可以不入库。
	//（1）管理员信息入--evfs_chaingroup_member
	//（2）主节点地址入-evfs-node
	//（3）入企业信息--evfs-org
	var input struct {
		MemberNames      [][]byte         //管理员成员姓名--proto
		Members          []common.Address //管理员成员--钱包地址
		Tokens           []common.Address //4 token地址（类似代币）
		Parent           common.Address   //委员会合约地址
		Rule             *big.Int         //管理员待办审批规则 0 100（3/1） 200（3/2） 300（all）
		AccountNodesName [][]byte         //主节点信息（proto）
		AccountNodes     []common.Address //主节点地址（记账节点）
		OrgInfos         [][]byte         //proto 组织信息（企业）
	}
	err = abi.Constructor.Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return constant.PARSING_DONE, err
	}
	//evfs_chaingroup_member
	for i := 0; i < len(input.Members); i++ {
		memberInfo := &entity.MemberInfo{}
		if err := proto.Unmarshal(input.MemberNames[i], memberInfo); err != nil {
			log.Infof("[analysisGenesis_admingrop] Failed to parse MemberInfo:", err)
			return constant.PARSING_DONE, err
		}
		evfsChaingroupMember := &models.EvfsChaingroupMember{
			MainChainGroupId: strings.ToLower("1"),
			MemberAddress:    strings.ToLower(input.Members[i].Hex()),
			ChaincommitteeId: constant.GenXid(),
			JoinTime:         m.Status.Timestamp / 1000,
			TransactionHash:  m.Hash,
			//创世块儿管理员-直接通过
			Approve: constant.APPROVE_SUCCESS,
		}
		evfsChaingroupMember.MemberName = memberInfo.Name
		evfsChaingroupMember.Add()
	}
	// evfs-node 主节点信息
	for i := 0; i < len(input.AccountNodes); i++ {
		nodeInfo := &entity.NodeInfo{}
		if err := proto.Unmarshal(input.AccountNodesName[i], nodeInfo); err != nil {
			log.Infof("[analysisGenesis_admingrop] Failed to parse OrgInfos:", err)
			return constant.PARSING_DONE, err
		}
		orgInfo := &entity.OrgInfo{}
		if err := proto.Unmarshal(input.OrgInfos[i], orgInfo); err != nil {
			log.Infof("[analysisGenesis_admingrop] Failed to parse OrgInfos:", err)
			return constant.PARSING_DONE, err
		}
		evfsNode := &models.EvfsNode{
			ChainnodeId:     constant.GenXid(),
			OrgId:           orgInfo.OrgAddress,
			OrgName:         strings.TrimSpace(orgInfo.OrgName),
			NodeAddress:     strings.ToLower(input.AccountNodes[i].Hex()),
			NodeInfo:        nodeInfo.Name,
			NodeType:        int32(1),
			Approve:         constant.APPROVE_SUCCESS,
			TransactionHash: m.Hash,
		}
		evfsNode.Add()
	}
	return constant.PARSING_DONE, nil
}
