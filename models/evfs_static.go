package models

type EvfsStatic struct {
	Id             int32 `gorm:"column:id;primary_key" json:"id"`
	AccountCount   int32 `gorm:"column:account_count" json:"account_count"`       //主节点数
	ResourceCount  int32 `gorm:"column:resource_count" json:"resource_count"`     //资源总数
	DomainCount    int32 `gorm:"column:domain_count" json:"domain_count"`         //数据存储域总数
	BizCount       int32 `gorm:"column:biz_count" json:"biz_count"`               //业务域总数
	ClientCount    int32 `gorm:"column:client_count" json:"client_count"`         //前置节点总数
	BizSystemCount int32 `gorm:"column:biz_system_count" json:"biz_system_count"` //接入系统总数
	OrgCount       int32 `gorm:"column:org_count" json:"org_count"`               //企业总数
	UserCount      int32 `gorm:"user_count" json:"user_count"`                    //上链用户总数

	VideoTypeCount int32 `gorm:"video_type_count" json:"video_type_count"` //视频文件总数
	SoundTypeCount int32 `gorm:"sound_type_count" json:"sound_type_count"` //音频文件总数
	PicTypeCount   int32 `gorm:"pic_type_count" json:"pic_type_count"`     //图片文件总数
	OtherTypeCount int32 `gorm:"other_type_count" json:"other_type_count"` //其他文件总数
	FileTypeCount  int32 `gorm:"other_type_count" json:"file_type_count"` //文件总数

	FileUserdSize int32 `gorm:"file_userd_size" json:"file_userd_size"` //文件存储已用
	FileTotalSize int32 `gorm:"file_total_size" json:"file_total_size"` //文件存储总量

	StructUserdSize int32 `gorm:"struct_userd_size" json:"struct_userd_size"` //结构数据存储总量
	StructTotalSize int32 `gorm:"struct_total_size" json:"struct_total_size"` //结构数据存储总量

	FileUpSize  int32 `gorm:"file_up_size" json:"file_up_size"`   //上链文件总量
	FileUpCount int32 `gorm:"file_up_count" json:"file_up_count"` //已上链(个)

	StructUpSize  int32 `gorm:"struct_up_size" json:"struct_up_size"`   //上链结构数据总量
	StructUpCount int32 `gorm:"struct_up_count" json:"struct_up_count"` //已上链(条)

	CommitteedCount int32 `gorm:"committeed_count" json:"committeed_count"` //联盟委员个数

	// 所有主节点-非结构化统计
	//MainNodeUnstructuredTotalSize  int32 `gorm:"main_node_unstructured_total_size" json:"struct_up_size"`   //所有主节点'非'结构化总容量
	//MainNodeUnstructuredUsedSize int32 `gorm:"main_node_unstructured_used_size" json:"struct_up_count"` //所有主节点'非'结构化-已用容量
	//MainNodeUnstructuredUsedNum int32 `gorm:"main_node_unstructured_used_num" json:"struct_up_count"` //所有主节点'非'结构化-已用入库条数
	// 所有主节点-结构化统计
	//MainNodestructuredTotalSize  int32 `gorm:"main_node_structured_total_size" json:"struct_up_size"`   //所有主节点结构化总容量
	//MainNodestructuredUsedSize int32 `gorm:"main_node_structured_used_size" json:"struct_up_count"` //所有主节点结构化-已用容量
	//MainNodestructuredUsedNum int32 `gorm:"main_node_structured_used_num" json:"struct_up_count"` //所有主节点结构化-已用入库条数

}
