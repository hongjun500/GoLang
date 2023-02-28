package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * @author hongjun500
 * @date 2022/9/10 11:26
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 文章标签模型接口
 */

type Tag struct {
	// 这里的Model相当于继承的意思
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

// GetTags 分页查询  ps：最后括号里表示显示地声明了返回值tags, 这个变量就可以直接在函数中使用，并且可以直接return
func GetTags(pageNum, pageSize int, maps interface{}) (tags []Tag) {
	// 同个models包下  db *gorm.DB是可以直接使用的
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagsV2 分页查询 ps：作用同上，推荐上面的写法
func GetTagsV2(pageNum, pageSize int, maps interface{}) []Tag {
	var tags []Tag
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return tags

}

func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	/*if tag.ID >0 {
		return true
	}
	return false*/
	return tag.ID > 0
}

func AddTag(name string, state int, createdBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	})
	return true
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}

func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})
	return true
}

func EditTag(id int, data interface{}) bool {
	db.Model(&Tag{}).Where("id = ?", id).Update(data)
	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
