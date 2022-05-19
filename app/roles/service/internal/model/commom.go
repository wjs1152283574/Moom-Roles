package model

type Commom struct {
	CreatorID   int64  `gorm:"type:int(11);COMMENT:创建者ID"`
	CreatorName string `gorm:"type:varchar(50);COMMENT:创建者名称（列表冗余）"`

	UpdatedTime int64 `gorm:"type:bigint(20);COMMENT:最后修改时间"`
	CreatedTime int64 `gorm:"type:bigint(20);COMMENT:创建时间"`
	DeleteTime  int64 `gorm:"type:bigint(20);COMMENT:删除时间"`
}
