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

type CommitteeGroup struct{}

func NewCommitteeGroup() *CommitteeGroup {
	m := &CommitteeGroup{}
	return m
}

// 规则申请 ruleApply(uint256) 返回 (bytes32,_isEnd)
func (c *CommitteeGroup) RuleApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		Rule *big.Int
	}

	err = abi.Methods["ruleApply"].Inputs.Unpack(&input, data)
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
	err = abi.Methods["ruleApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsChaingroup := &models.EvfsChaingroup{
		GroupId:         string("2"),
		PendingRule:     int32(input.Rule.Int64()),
		Name:            string("委员会"),
		TransactionHash: m.Hash,
	}
	return evfsChaingroup.UpdatePendingRule(res.IsEnd)
}

// 规则申请同意
func (c *CommitteeGroup) RuleAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
	evfsChaingroup := &models.EvfsChaingroup{
		GroupId:         string("2"),
		TransactionHash: m.Hash,
	}
	if res.IsEnd {
		return evfsChaingroup.UpdateRule(res.IsEnd)
	}
	return nil
}

// 成员申请
func (c *CommitteeGroup) MemberApply(m *entity.TransactionInfo, abi abi.ABI) error {
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
		Name            []byte
		Member          common.Address
		SubGroupAddress common.Address
		Op              *big.Int
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
		log.Errorf("hash=%s,result len=%d", m.Hash, len(data))
		return errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	err = abi.Methods["memberApply"].Outputs.Unpack(&res, data)
	if err != nil {
		log.Errorf("hash=%s, outputs unpack error=%s", m.Hash, err.Error())
		return err
	}
	evfsChaingroupMember := &models.EvfsChaingroupMember{
		MainChainGroupId: strings.ToLower(input.SubGroupAddress.Hex()),
		MemberAddress:    strings.ToLower(input.Member.Hex()),
	}
	memberInfo := &entity.MemberInfo{}
	if err := proto.Unmarshal(input.Name, memberInfo); err != nil {
		log.Infof("[analysisGenesis_committee] Failed to parse MemberInfo:", err)
		return err
	}
	evfsChaingroupMember.MemberName = memberInfo.Name
	if input.Op.Cmp(big.NewInt(constant.OP_REMOVE)) == 0 && res.IsEnd {
		return evfsChaingroupMember.Delete()
	} else {
		evfsChaingroupMember.ChaincommitteeId = res.ReqId.Hex()
		evfsChaingroupMember.JoinTime = m.Status.Timestamp / 1000
		evfsChaingroupMember.TransactionHash = m.Hash
		evfsChaingroupMember.Approve = constant.APPROVE_ING
		if res.IsEnd {
			evfsChaingroupMember.Approve = constant.APPROVE_SUCCESS
		}
		return evfsChaingroupMember.Add()
	}
}

// 成员申请同意
func (c *CommitteeGroup) MemberAgree(m *entity.TransactionInfo, abi abi.ABI) error {
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
		ReqId common.Hash
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
	evfsChaingroupMember := &models.EvfsChaingroupMember{
		ChaincommitteeId: input.ReqId.Hex(),
		TransactionHash:  m.Hash,
	}

	//同意
	if res.IsEnd {
		evfsChaingroupMember.Approve = constant.APPROVE_SUCCESS
		return evfsChaingroupMember.UpdateApprove()
	}
	return nil
}

// 解析创世合约信息--委员会相关
func (c *CommitteeGroup) analysisGenesis_committee(m *entity.TransactionInfo, abi abi.ABI) (int64, error) {
	codedata := constant.GetParamsHexStr(m.Body.CodeData)
	data, err := hex.DecodeString(codedata)
	if err != nil {
		log.Errorf("[analysisGenesis_committee] unknown code_data error=%s", err.Error())
		return constant.PARSING_DONE, err
	}
	if len(data)%32 != 0 {
		log.Errorf("[analysisGenesis_committee] hash=%s,code_data len=%d", m.Hash, len(data))
		return constant.PARSING_DONE, errors.New(fmt.Sprintf("hash=%s,code_data len=%d", m.Hash, len(data)))
	}
	var input struct {
		DefaultRule *big.Int         //委员会待办审批规则
		MemberNames [][]byte         //委员会人员名称
		Members     []common.Address //委员会成员--钱包地址
	}
	err = abi.Constructor.Inputs.Unpack(&input, data)
	if err != nil {
		log.Errorf("[analysisGenesis_committee] hash=%s, inputs unpack %s", m.Hash, err.Error())
		return constant.PARSING_DONE, err
	}
	for i := 0; i < len(input.Members); i++ {
		memberInfo := &entity.MemberInfo{}
		if err := proto.Unmarshal(input.MemberNames[i], memberInfo); err != nil {
			log.Infof("[analysisGenesis_committee] Failed to parse MemberInfo:", err)
			return constant.PARSING_DONE, err
		}
		evfsChaingroupMember := &models.EvfsChaingroupMember{
			MainChainGroupId: strings.ToLower("2"),
			MemberAddress:    strings.ToLower(input.Members[i].Hex()),
			ChaincommitteeId: constant.GenXid(),
			JoinTime:         m.Status.Timestamp / 1000,
			TransactionHash:  m.Hash,
			//创世块儿委员-直接通过
			Approve: constant.APPROVE_SUCCESS,
			//MemberName:strings.Trim(strings.TrimSpace(string(input.MemberNames[i])),"\u001b"),
		}
		evfsChaingroupMember.MemberName = memberInfo.Name
		evfsChaingroupMember.Add()
	}
	return constant.PARSING_DONE, nil
}
