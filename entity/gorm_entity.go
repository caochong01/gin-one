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
	DiaryTitle  string
	DiaryDetail string
}

//func (diary *Diary) TableName() string {
//	return "diary"
//}

func (diary *Diary) TableName() string {
	return "diary"
}

type Comment struct {
	Model
	DiaryId        uint
	CommentContent string
}

//func (comment *Comment) TableName() string {
//	return "comment"
//}

type Praise struct {
	Model
	DiaryId   uint
	PraiseNum uint
}

//func (praise *Praise) TableName() string {
//	return "praise"
//}
