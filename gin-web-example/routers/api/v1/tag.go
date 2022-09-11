package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/hongjun500/Go-master/gin-web-example/models"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/e"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/setting"
	"github.com/hongjun500/Go-master/gin-web-example/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

/**
 * @author hongjun500
 * @date 2022/9/10 10:57
 * @tool ThinkPadX1隐士
 * Created with GoLand 2021.2
 * Description: 文章标签服务
 */

// GetTags 获取文章多个标签
func GetTags(context *gin.Context) {
	// 通过名称查询标签
	var name = context.Query("name")
	var maps = make(map[string]interface{})
	var data = make(map[string]interface{})
	if name != "" {
		maps["name"] = name
	}
	var state = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(context), setting.PageSize, maps)
	data["total"] = models.GetTagsTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsgByCode(code),
		"data": data,
	})
}

func AddTag(context *gin.Context) {
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy := context.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	var data interface{}
	// 校验无误
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			// 执行入库操作
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
		data = "ok"
	} else {
		data = valid.Errors
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsgByCode(code),
		"data": data,
	})
}

func EditTag(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()
	name := context.Query("name")
	modifiedBy := context.Query("modified_by")

	valid := validation.Validation{}

	state := -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = name
			}
			// 执行入库操作
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsgByCode(code),
		"data": make(map[string]string),
	})
}

func DeleteTag(context *gin.Context) {
	id := com.StrTo(context.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			// 执行删除操作
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsgByCode(code),
		"data": make(map[string]string),
	})
}
