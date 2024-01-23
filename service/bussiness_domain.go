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

type BussinessDomain struct{}

func NewBussinessDomain() *BussinessDomain {
	m := &BussinessDomain{}
	return m
}

// 成员申请
func (c *BussinessDomain) MemberApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address //业务域ID
		Name     []byte         //成员名称
		Member   common.Address //成员地址
		Op       *big.Int       //1 添加  2移除
	}

	err = abi.Methods["memberApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return nil
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
	err = abi.Methods["memberApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsBizMember := models.EvfsBizMember{
		BizId:           input.DomainId.Hex(),
		Address:         input.Member.Hex(),
		TransactionHash: m.Hash,
		ReqId:           res.ReqId.Hex(),
		JoinTime:        time.Unix(m.Status.Timestamp/1000, 0),
		Approve:         constant.APPROVE_ING,
		Op:              input.Op.Int64(),
	}
	if res.IsEnd {
		evfsBizMember.Approve = constant.APPROVE_SUCCESS
	}
	return evfsBizMember.Add()
}

// 成员申请同意
func (c *BussinessDomain) MemberAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		DomainId common.Address
		ReqId    common.Hash
	}

	err = abi.Methods["memberAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", err.Error())
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
	err = abi.Methods["memberAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsStorageMember := models.EvfsBizMember{
		BizId: input.DomainId.Hex(),
		ReqId: input.ReqId.Hex(),
	}
	//同意
	if res.IsEnd {
		evfsStorageMember.Approve = constant.APPROVE_SUCCESS
		return evfsStorageMember.UpdateApprove()
	}
	return nil
}

// 规则申请
func (c *BussinessDomain) RuleApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
	evfsBiz := &models.EvfsBiz{
		BizId:               input.DomainId.Hex(),
		PendingRule:         int32(input.Rule.Int64()),
		RuleTransactionHash: m.Hash,
		RuleReqId:           res.ReqId.Hex(),
	}
	return evfsBiz.UpdatePendingRule(res.IsEnd)
}

// 规则申请同意
func (c *BussinessDomain) RuleAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		evfsBiz := &models.EvfsBiz{
			RuleReqId:           input.ReqId.Hex(),
			DomainId:            input.DomainId.Hex(),
			RuleTransactionHash: m.Hash,
		}
		return evfsBiz.UpdateRule(res.IsEnd)
	}
	return nil
}

// 业务系统申请
func (c *BussinessDomain) BizSystemApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return err
	}
	var input struct {
		OrgAddress common.Address //企业地址
		OrgName    []byte         //企业信息
		DomainId   common.Address //业务域地址
		BizSystem  common.Address //业务域下系统地址
		Name       []byte         //业务系统信息
		Op         *big.Int       // 操作 1 添加  2 移除
	}
	err = abi.Methods["bizSystemApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", err.Error())
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
	err = abi.Methods["bizSystemApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}

	orginfo := &entity.OrgInfo{}
	if err := proto.Unmarshal(input.OrgName, orginfo); err != nil {
		log.Infof("[bizSystemApply] Failed to parse orginfo:", err)
		return err
	}

	bizSystemInfo := &entity.BizSystemInfo{}
	if err := proto.Unmarshal(input.Name, bizSystemInfo); err != nil {
		log.Infof("[bizSystemApply] Failed to parse BizSystemInfo:", err)
		return err
	}
	jsons, errs := json.Marshal(bizSystemInfo) //转换成JSON返回的是byte[]
	if errs != nil {
		return errs
	}
	evfsSystem := models.EvfsSystem{
		ReqId:           res.ReqId.Hex(),                        //待办ID
		BizId:           strings.ToLower(input.DomainId.Hex()),  //业务域ID
		SysId:           strings.ToLower(input.BizSystem.Hex()), //系统ID
		SysInfo:         string(jsons),                          //系统信息
		SysName:         bizSystemInfo.SysName,                  //系统名称
		OrgId:           input.OrgAddress.Hex(),                 //企业ID
		OrgName:         orginfo.OrgName,                        //企业名称
		Approve:         constant.APPROVE_ING,
		TransactionHash: m.Hash,
	}

	if res.IsEnd {
		evfsSystem.Approve = constant.APPROVE_SUCCESS
	}
	//申请即同意
	return evfsSystem.Add(res.IsEnd)
}

// 业务系统申请同意
func (c *BussinessDomain) BizSystemAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		ReqId    common.Hash    //待办ID
		DomainId common.Address //业务域ID
	}

	err = abi.Methods["bizSystemAgree"].Inputs.Unpack(&input, data)
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
	err = abi.Methods["bizSystemAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsSystem := &models.EvfsSystem{
		ReqId:           input.ReqId.Hex(),
		BizId:           input.DomainId.Hex(),
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsSystem.Approve = constant.APPROVE_SUCCESS
		return evfsSystem.UpdateApprove()
	}
	return nil
}

// 解析业务域-新创建合约
func (c *BussinessDomain) CreateContractApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return err
	}
	var input struct {
		DomainId common.Address //业务域ID
		Code     []byte         //bincode
		Info     []byte         //合约信息--proto
	}
	err = abi.Methods["createContractApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", err.Error())
		return err
	}
	var res struct {
		ReqId           common.Hash
		IsEnd           bool
		ContractAddress common.Address
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
	err = abi.Methods["createContractApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	bizContractInfo := &entity.BizContractInfo{}
	if err := proto.Unmarshal(input.Info, bizContractInfo); err != nil {
		log.Infof("[createContractApply] Failed to parse BizContractInfo:", err)
		return err
	}
	jsons, errs := json.Marshal(bizContractInfo) //转换成JSON返回的是byte[]
	if errs != nil {
		return errs
	}
	evfsBizContractFlow := models.EvfsBizContractFlow{
		ReqId:           res.ReqId.Hex(),                       //待办ID
		BizId:           strings.ToLower(input.DomainId.Hex()), //业务域ID
		ContractInfo:    string(jsons),                         //合约信息
		ContractName:    bizContractInfo.Name,                  //合约名称
		ContractMark:    bizContractInfo.Remark,                //合约备注
		StatusInfo:      int64(0),                              //0 创建
		Approve:         constant.APPROVE_ING,
		TransactionHash: m.Hash,
	}
	if res.IsEnd {
		evfsBizContractFlow.Approve = constant.APPROVE_SUCCESS
		evfsBizContractFlow.ContractAddress = res.ContractAddress.Hex() //合约地址
	}
	//申请即同意
	return evfsBizContractFlow.CreateContractAdd(m.Accepttimestamp)
}

// 业务系统申请同意
func (c *BussinessDomain) CreateContractAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		BizId common.Address //业务域ID
		ReqId common.Hash    //待办ID
	}
	err = abi.Methods["createContractAgree"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", m.Hash, err.Error())
		return err
	}
	var res struct {
		IsEnd           bool
		ContractAddress common.Address //合约地址
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
	err = abi.Methods["createContractAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsBizContractFlow := models.EvfsBizContractFlow{
		ReqId:           input.ReqId.Hex(),                  //待办ID
		BizId:           strings.ToLower(input.BizId.Hex()), //业务域ID
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsBizContractFlow.Approve = constant.APPROVE_SUCCESS
		evfsBizContractFlow.ContractAddress = res.ContractAddress.Hex()
		return evfsBizContractFlow.CreateContractUpdateApprove(m.Accepttimestamp)
	}
	return nil
}

// 解析业务域-合约冻结、解冻操作
func (c *BussinessDomain) ContractEnableApply(m *entity.TransactionInfo, abi abi.ABI) error {
	data, err := hex.DecodeString(m.Body.CodeData[10:])
	if err != nil {
		log.Errorf("unknown code_data error=", err.Error())
		return err
	}
	if len(data)%32 != 0 {
		log.Errorf("hash=%s,code_data len=%d", m.Hash, len(data))
		return err
	}
	var input struct {
		BizId           common.Address //业务域ID
		ContractAddress common.Address //合约地址
		Op              *big.Int       //操作 1冻结 2解冻
	}
	err = abi.Methods["contractEnableApply"].Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("hash=%s, inputs unpack %s", err.Error())
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
	err = abi.Methods["contractEnableApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsBizContractFlow := models.EvfsBizContractFlow{
		ReqId:      res.ReqId.Hex(),                    //待办ID
		BizId:      strings.ToLower(input.BizId.Hex()), //业务域ID
		StatusInfo: input.Op.Int64(),                   //1 冻结  2 解冻
	}
	if res.IsEnd {
		evfsBizContractFlow.Approve = constant.APPROVE_SUCCESS
	}
	//申请即同意
	return evfsBizContractFlow.ContractFrozen(m.Accepttimestamp)
}

// 解析业务域-合约冻结、解冻同意操作
func (c *BussinessDomain) ContractEnableAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		BizId common.Address //业务域ID
		ReqId common.Hash    //待办ID
	}
	err = abi.Methods["contractEnableAgree"].Inputs.Unpack(&input, data)
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
	err = abi.Methods["contractEnableAgree"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsBizContractFlow := models.EvfsBizContractFlow{
		ReqId:           input.ReqId.Hex(),                  //待办ID
		BizId:           strings.ToLower(input.BizId.Hex()), //业务域ID
		TransactionHash: m.Hash,
	}
	//申请即同意
	if res.IsEnd {
		evfsBizContractFlow.Approve = constant.APPROVE_SUCCESS
		return evfsBizContractFlow.ContractFrozenUpdateApprove(m.Accepttimestamp)
	}
	return nil
}
