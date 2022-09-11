package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/**
 * @author hongjun500
 * @date 2022/9/11 14:08
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 文章模型接口
 */

type Article struct {
	Model Model

	// tag_id字段加了一个索引
	TagId int    `json:"tag_id" gorm:"index"`
	Tag   Tag    `json:"tag"`
	Title string `json:"title"`

	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func (tag *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func ExistArticleById(id int) (exist bool) {
	var article Article
	db.Select("id").Where("id =? ", id).First(&article)
	exist = article.Model.ID > 0
	return
}

func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func ListArticles(pageNum, pageSize int, maps interface{}) (articles []Article) {
	db.Preloads("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ? ", id).First(&article)
	// gorm通过Related进行关联查询
	db.Model(&article).Related(&article.Tag)
	return
}

func EditaArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id =? ", id).Update(data)
	return true
}

func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagId:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id =? ", id).Delete(Article{})
	return true
}
