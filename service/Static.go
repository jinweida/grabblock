package service

import (
	"time"

	"example.cn/grabblock/models"
	"example.cn/grabblock/tools/log"
)

/*
*

	 各种统计数据
	1、biz 业务统计
	2、sys 系统统计
*/
type Static struct {
	BizInfoStatic models.EvfsBizStatic
	SysInfoStatic models.EvfsSystemStatic
}

/*
*

	创建操作对象
*/
func NewStatic() *Static {
	return &Static{}
}

/*
*

	统计业务域下合同调用次数--evfs_biz_contract_record 表使用
*/
func (t *Static) StaticBizContractCallTimes() {
	start := time.Now()
	data := models.GetEvfsBizContractRecord()
	mainTransaction := models.MainTransaction{}
	for _, m := range data {
		//统计合约使用次数
		count := mainTransaction.CountToAddress(m.ContractAddress)
		//更新合约使用次数
		m.UpdateAddressCount(m.ContractAddress, count)
	}
	log.Infof("parsing contract time cost=%v", time.Since(start))
}

/*
*

	统计业务域下各种统计信息--evfs_biz_static 表使用
*/
func (t *Static) StaticBiz() {
	start := time.Now()
	t.BizInfoStatic = models.EvfsBizStatic{}
	data := models.GetEvfsBizStatic()
	//遍历业务域id，将统计信息更新入库
	for _, m := range data {
		t.BizInfoStatic.BizId = m.BizId
		t.BizInfoStatic.BizName = m.BizName
		t.staticBizFileSizeAndCount(m.BizId)
		t.staticBizDataSizeAndCount(m.BizId)
		t.staticBizContractCount(m.BizId)
		t.staticBizOnlineUserCount(m.BizId)
		t.staticBizOrgCount(m.BizId)
		t.staticBizSysCount(m.BizId)
		t.BizInfoStatic.Save()
	}
	log.Infof("StaticBiz time cost=%v", time.Since(start))
}

/*
*

	统计业务域某系统下---各种统计信息--evfs_system_static 表使用
*/
func (t *Static) StaticBizSys() {
	start := time.Now()
	t.SysInfoStatic = models.EvfsSystemStatic{}
	data := models.GetEvfsSystemStatic()
	//遍历系统id，将统计信息更新入库
	for _, m := range data {
		t.SysInfoStatic.SysId = m.SysId
		t.SysInfoStatic.SysName = m.SysName
		t.SysInfoStatic.BizId = m.BizId
		t.SysInfoStatic.OrgId = m.OrgId
		t.SysInfoStatic.OrgName = m.OrgName
		t.staticSysFileSizeAndCount(m.BizId, m.SysId)
		t.staticSysDataSizeAndCount(m.BizId, m.SysId)
		t.staticSysOnlineUserCount(m.BizId, m.SysId)
		t.SysInfoStatic.Save()
	}
	log.Infof("StaticBizSys time cost=%v", time.Since(start))
}

//--------------业务域下各种统计信息---------------------

// 统计业务域下--非结构化数据容量、数量---evfs_biz_static
func (t *Static) staticBizFileSizeAndCount(bizid string) {
	start := time.Now()
	evfsStorageFile := models.EvfsStorageFile{}
	size, count := evfsStorageFile.StaticSizeAndCount(bizid)
	t.BizInfoStatic.OnlineFileSize = size
	t.BizInfoStatic.OnlineFileCount = count
	log.Infof("staticBizFileSizeAndCount time cost=%v", time.Since(start))
}

// 统计业务域下--结构化数据容量、数量---evfs_biz_static
func (t *Static) staticBizDataSizeAndCount(bizid string) {
	start := time.Now()
	evfsStorageData := models.EvfsStorageData{}
	size, count := evfsStorageData.StaticSizeAndCount(bizid)
	t.BizInfoStatic.OnlineDataSize = size
	t.BizInfoStatic.OnlineDataCount = count
	log.Infof("staticBizDataSizeAndCount time cost=%v", time.Since(start))
}

// 统计业务域下--合约数量---evfs_biz_static
func (t *Static) staticBizContractCount(bizid string) {
	start := time.Now()
	evfsBizContractRecord := models.EvfsBizContractRecord{}
	count := evfsBizContractRecord.Count(bizid)
	t.BizInfoStatic.ContractCount = int64(count)
	log.Infof("staticBizContractCount time cost=%v", time.Since(start))
}

// 统计业务域下--已上链用户---evfs_biz_static
func (t *Static) staticBizOnlineUserCount(bizid string) {
	start := time.Now()
	mainAccountTransaction := models.MainAccountTransaction{}
	t.BizInfoStatic.OnlineUserCount = mainAccountTransaction.CountOnlineUser(bizid)
	log.Infof("staticBizOnlineUserCount time cost=%v", time.Since(start))
}

// 统计业务域下----企业统计---evfs_biz_static
func (t *Static) staticBizOrgCount(bizid string) {
	start := time.Now()
	evfsSystem := models.EvfsSystem{}
	t.BizInfoStatic.OrgCount = evfsSystem.CountOrg(bizid)
	log.Infof("staticBizOrgCount time cost=%v", time.Since(start))
}

// 统计业务域下--系统统计---evfs_biz_static
func (t *Static) staticBizSysCount(bizid string) {
	start := time.Now()
	evfsSystem := models.EvfsSystem{}
	t.BizInfoStatic.SysCount = evfsSystem.CountSys(bizid)
	log.Infof("staticBizSysCount time cost=%v", time.Since(start))
}

//--------------业务域下系统信息统计---------------------

// 统计业务域下某系统--非结构化数据容量、数量---evfs_system_static
func (t *Static) staticSysFileSizeAndCount(bizid string, sysid string) {
	start := time.Now()
	evfsStorageFile := models.EvfsStorageFile{}
	size, count := evfsStorageFile.StaticSizeAndCountInsys(bizid, sysid)
	t.SysInfoStatic.OnlineFileSize = size
	t.SysInfoStatic.OnlineFileCount = count
	log.Infof("staticSysFileSizeAndCount time cost=%v", time.Since(start))
}

// 统计业务域下某系统--结构化数据容量、数量---evfs_system_static
func (t *Static) staticSysDataSizeAndCount(bizid string, sysid string) {
	start := time.Now()
	evfsStorageData := models.EvfsStorageData{}
	size, count := evfsStorageData.StaticSizeAndCountInsys(bizid, sysid)
	t.SysInfoStatic.OnlineDataSize = size
	t.SysInfoStatic.OnlineDataCount = count
	log.Infof("staticSysDataSizeAndCount time cost=%v", time.Since(start))
}

// 统计业务域下某系统--已上链用户---evfs_system_static
func (t *Static) staticSysOnlineUserCount(bizid string, sysid string) {
	start := time.Now()
	mainAccountTransaction := models.MainAccountTransaction{}
	t.SysInfoStatic.OnlineUserCount = mainAccountTransaction.CountOnlineUserInsys(bizid, sysid)
	log.Infof("staticSysOnlineUserCount time cost=%v", time.Since(start))
}
