package entity

import (
	"gorm.io/gorm"
	"time"
)

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Diary struct {
	Model
	DiaryTitle  string `form:"diaryTitle" json:"diaryTitle" binding:"required"`
	DiaryDetail string `form:"diaryDetail" json:"diaryDetail" binding:"required"`
}

//func (diary *Diary) TableName() string {
//	return "diary"
//}

type Comment struct {
	Model
	DiaryId        uint   `form:"diaryId" json:"diaryId" binding:"required"`
	CommentContent string `form:"commentContent" json:"commentContent" binding:"required"`
}

type Praise struct {
	Model
	DiaryId   uint `form:"diaryId" json:"diaryId" binding:"required"`
	PraiseNum uint `form:"praiseNum" json:"praiseNum" binding:"required"`
}
