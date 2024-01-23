package service

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/models"
	"example.cn/grabblock/tools/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/protobuf/proto"
)

type DataStorageDomain struct{}

func NewDataStorageDomain() *DataStorageDomain {
	m := &DataStorageDomain{}
	return m
}

// 数据存管域-成员申请
func (c *DataStorageDomain) MemberApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address
		Name     []byte
		Member   common.Address
		Op       *big.Int //1 添加  2 移除
	}

	err = abi.Methods["memberApply"].Inputs.Unpack(&input, data)
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
		log.Errorf("hash=%s,result len=", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["memberApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorageMember := models.EvfsStorageMember{
		DomainId:        strings.ToLower(input.DomainId.Hex()),
		Address:         strings.ToLower(input.Member.Hex()),
		TransactionHash: m.Hash,
		ReqId:           res.ReqId.Hex(),
		JoinTime:        m.Status.Timestamp,
		Approve:         constant.APPROVE_ING,
		Op:              input.Op.Int64(),
	}
	memberInfo := &entity.MemberInfo{}
	if err := proto.Unmarshal(input.Name, memberInfo); err != nil {
		log.Infof("[datastorage_domain_memberapply] Failed to parse MemberInfo:", err)
		return err
	}
	evfsStorageMember.AddressName = memberInfo.Name
	if input.Op.Cmp(big.NewInt(constant.OP_REMOVE)) == 2 && res.IsEnd {
		//成员申请时直接同意
		return evfsStorageMember.Delete()
	} else {
		if res.IsEnd {
			evfsStorageMember.Approve = constant.APPROVE_SUCCESS
		}
		// 添加、删除成员入库数据，后续是否同意时处理
		return evfsStorageMember.Add()
	}
}

// 数据存管域-成员申请同意
func (c *DataStorageDomain) MemberAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address
		ReqId    common.Hash
	}

	err = abi.Methods["memberAgree"].Inputs.Unpack(&input, data)
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
	err = abi.Methods["memberAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorageMember := models.EvfsStorageMember{
		DomainId: strings.ToLower(input.DomainId.Hex()),
		ReqId:    input.ReqId.Hex(),
	}
	//同意
	if res.IsEnd {
		evfsStorageMember.Approve = constant.APPROVE_SUCCESS
		return evfsStorageMember.UpdateApprove()
	}
	return nil
}

// 规则申请 ruleApply(uint256) 返回 (bytes32,_isEnd)
func (c *DataStorageDomain) RuleApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		DomainId common.Address
		Rule     *big.Int
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
		log.Errorf("unknown result error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["ruleApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorage := &models.EvfsStorage{
		DomainId:            strings.ToLower(input.DomainId.Hex()),
		PendingRule:         int32(input.Rule.Int64()),
		RuleTransactionHash: m.Hash,
		RuleReqId:           res.ReqId.Hex(),
	}
	return evfsStorage.UpdatePendingRule(res.IsEnd)
}

// 规则申请同意
func (c *DataStorageDomain) RuleAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address
		ReqId    common.Hash
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
		log.Errorf("unknown result error=", err.Error())
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
	if res.IsEnd {
		evfsStorage := &models.EvfsStorage{
			RuleReqId:           input.ReqId.Hex(),
			DomainId:            strings.ToLower(input.DomainId.Hex()),
			RuleTransactionHash: m.Hash,
		}
		return evfsStorage.UpdateRule(res.IsEnd)
	}
	return nil
}

// 业务域申请
func (c *DataStorageDomain) BusinessDomainApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address   //数据存管域地址
		Name     []byte           //proto BizDomainInfo 业务域名称
		Names    []byte           //proto MemberInfos   业务域管理员名字
		Members  []common.Address //业务域管理员地址
		Op       []*big.Int       //  数组----【0】-- 1 添加 2移除   （ 暂时只有添加）
	}

	err = abi.Methods["businessDomainApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		ReqId common.Hash
		IsEnd bool
		BizId common.Address
	}

	data, err = hex.DecodeString(m.Status.Result[2:])
	if err != nil {
		log.Errorf("unknown result error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return err
	}
	err = abi.Methods["businessDomainApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	bizDomainInfo := &entity.BizDomainInfo{}
	if err := proto.Unmarshal(input.Name, bizDomainInfo); err != nil {
		log.Infof("[analysisGenesis_admingrop] Failed to parse MemberInfo:", err)
		return err
	}
	memberInfos := &entity.MemberInfos{}
	if err := proto.Unmarshal(input.Names, memberInfos); err != nil {
		log.Infof("[analysisGenesis_admingrop] Failed to parse MemberInfo:", err)
		return err
	}
	evfsBiz := models.EvfsBiz{
		ReqId:           res.ReqId.Hex(),
		BizId:           strings.ToLower(res.BizId.Hex()),
		BizName:         bizDomainInfo.Name,
		DomainId:        strings.ToLower(input.DomainId.Hex()),
		Approve:         constant.APPROVE_ING,
		TransactionHash: m.Hash,
		Members:         make([]models.EvfsBizMember, 0),
	}
	for i := 0; i < len(input.Members); i++ {
		mem := models.EvfsBizMember{
			BizId:           strings.ToLower(res.BizId.Hex()),
			ReqId:           res.ReqId.Hex(),
			Address:         strings.ToLower(input.Members[i].Hex()),
			Name:            memberInfos.Name[i].Name,
			TransactionHash: m.Hash,
			JoinTime:        time.Unix(m.Status.Timestamp, 0),
			Approve:         constant.APPROVE_ING,
			Op:              int64(1),
		}
		if res.IsEnd {
			evfsBiz.Approve = constant.APPROVE_SUCCESS
		}
		evfsBiz.Members = append(evfsBiz.Members, mem)
	}
	if res.IsEnd {
		evfsBiz.Approve = constant.APPROVE_SUCCESS
	}
	//申请即同意
	return evfsBiz.Add(res.IsEnd)
}

// 业务域申请同意
func (c *DataStorageDomain) BusinessDomainAgree(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		DomainId common.Address //存管域ID
		ReqId    common.Hash    //待办ID
	}

	err = abi.Methods["businessDomainAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}

	var res struct {
		IsEnd bool           //是否通过
		BizId common.Address //业务域地址
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
	err = abi.Methods["businessDomainAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsBiz := models.EvfsBiz{
		DomainId:        strings.ToLower(input.DomainId.Hex()),
		ReqId:           input.ReqId.Hex(),
		BizId:           strings.ToLower(res.BizId.Hex()),
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsBiz.Approve = constant.APPROVE_SUCCESS
		//更新状态
		evfsBiz.UpdateApprove()
		//查找数据获取bizname
		findinfo := evfsBiz.UseBizidFind(evfsBiz.BizId)
		//如果同意-入库统计信息（用于后续数据统计）
		evfsBizStatic := models.EvfsBizStatic{
			BizId:   findinfo.BizId,
			BizName: findinfo.BizName,
		}
		return evfsBizStatic.Add()
	}
	return nil
}

// 存储节点申请
func (c *DataStorageDomain) DataStorageNodeApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId    common.Address // 存管域ID
		NodeAddress common.Address //节点地址
		Name        []byte         //节点信息
		Url         []byte         //节点地址
		Amount      *big.Int       //容量
		Op          *big.Int       //1:添加 2:移除
	}

	err = abi.Methods["dataStorageNodeApply"].Inputs.Unpack(&input, data)
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
	err = abi.Methods["dataStorageNodeApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	resourceNodeInfo := &entity.ResourceNodeInfo{}
	if err := proto.Unmarshal(input.Name, resourceNodeInfo); err != nil {
		log.Infof("[dataStorageNodeApply] Failed to parse resourceNodeInfo:", err)
		return err
	}
	jsons, errs := json.Marshal(resourceNodeInfo) //转换成JSON返回的是byte[]
	if errs != nil {
		return errs
	}
	evfsResourcenode := models.EvfsResourcenode{
		ReqId:           res.ReqId.Hex(),
		DomainId:        strings.ToLower(input.DomainId.Hex()),
		NodeAddress:     strings.ToLower(input.NodeAddress.Hex()),
		NodeName:        resourceNodeInfo.NodeName,
		NodeInfo:        string(jsons),
		Url:             string(input.Url),
		Approve:         constant.APPROVE_ING,
		Cpu:             resourceNodeInfo.Cpu,
		Disk:            resourceNodeInfo.Disk,
		Memory:          resourceNodeInfo.Memory,
		Bandwidth:       resourceNodeInfo.Bandwidth,
		OrgId:           resourceNodeInfo.OrgAddress,
		OrgName:         resourceNodeInfo.OrgName,
		TransactionHash: m.Hash,
		Op:              input.Op.Int64(),
	}
	if res.IsEnd {
		evfsResourcenode.Approve = constant.APPROVE_SUCCESS
	}
	//1添加
	if int32(input.Op.Int64()) == 1 {
		//申请即同意
		if res.IsEnd {
			evfsResourcenode.Approve = constant.APPROVE_SUCCESS
		}
		return evfsResourcenode.Add()
	} else {
		if res.IsEnd {
			//直接删除
			return evfsResourcenode.Delete()
		} else {
			//需要后续同意删除
			return evfsResourcenode.Add()
		}
	}
}

// 存储节点申请同意
func (c *DataStorageDomain) DataStorageNodeAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		ReqId    common.Hash
		DomainId common.Address
	}

	err = abi.Methods["dataStorageNodeAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", err.Error())
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
	err = abi.Methods["dataStorageNodeAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsResourcenode := &models.EvfsResourcenode{
		ReqId:           input.ReqId.Hex(),
		DomainId:        strings.ToLower(input.DomainId.Hex()),
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsResourcenode.Approve = constant.APPROVE_SUCCESS
		return evfsResourcenode.UpdateApprove()
	}
	return nil
}
