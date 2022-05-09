/*
@Time   : 2022/3/1 17:34
@Author : ckx0709
@Remark :
*/
package entity

type CasRule struct {
	ID    int    `gorm:"column:id"`                         // 自增主键
	PType string `gorm:"size:128;uniqueIndex:unique_index"` // Policy Type - 用于区分 policy和 group(role)
	V0    string `gorm:"size:128;uniqueIndex:unique_index"` // subject
	V1    string `gorm:"size:128;uniqueIndex:unique_index"` // object
	V2    string `gorm:"size:128;uniqueIndex:unique_index"` // action
	V3    string `gorm:"size:128;uniqueIndex:unique_index"` // 这个和下面的字段无用，仅预留位置，如果你的不是
	V4    string `gorm:"size:128;uniqueIndex:unique_index"` // sub, obj, act的话才会用到
	V5    string `gorm:"size:128;uniqueIndex:unique_index"` // 如 sub, obj, act, suf就会用到 V3
}

func (r CasRule) TableName() string {
	return "t_casbin_rule"
}
