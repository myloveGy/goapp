package controllers

import (
	"project/models"
	"strings"

	"github.com/astaxie/beego"
)

// 表格返回数据
type DataTable struct {
	Echo  int         `json:"sEcho"`
	Count int64       `json:"iTotalRecords"`
	Total int64       `json:"iTotalDisplayRecords"`
	Data  interface{} `json:"aaData"`
}

// 后台控制器
type AdminController struct {
	BaseController
	SearchMap func() map[string]string
}

// 前置操作
func (this *AdminController) Prepare() {
	// 判断用户是否已经登录, 没有登录返回到登录页面
	//	if !this.isLogin("admin") {
	//		this.Redirect("/admin", 302)
	//	}

	// 使用的布局
	this.Data["admin"] = this.U
	this.Layout = "layout/admin.tpl"
}

// 查询方法
func (this *AdminController) Query(search map[string]string) models.Query {
	query := new(models.Query)

	// 处理默认查询信息
	query.Table = search["Table"]
	query.Order = search["orderBy"]
	query.IStart, _ = this.GetInt64("iDisplayStart")
	query.ILength, _ = this.GetInt64("iDisplayLength")
	query.Where = make(map[string]interface{})

	// 判断排序字段
	if order := this.GetString("params[orderBy]"); order != "" {
		query.Order = order
		delete(this.Ctx.Request.PostForm, "params[orderBy]")
	}

	// 判断排序方式
	sType := this.GetString("sSortDir_0")
	if sType != "" {
		query.Order = strings.TrimLeft(query.Order, "-")
		if sType == "desc" {
			query.Order = "-" + query.Order
		}

		delete(this.Ctx.Request.PostForm, "sSortDir_0")
	}

	// 判断查询信息
	if request := this.Ctx.Request.PostForm; request != nil {
		// 取出其他查询条件
		for k, v := range request {
			if strings.HasPrefix(k, "params[") {
				key := strings.Trim(strings.Trim(strings.Trim(k, "params"), "]"), "[")
				if tmp, ok := search[key]; ok {
					query.Where[tmp] = v[0]
				}
			}
		}
	}

	return *query
}

// 公共的查询数据的方法
func (this *AdminController) BaseSearch(arr interface{}, search map[string]string) {
	// 定义返回数据
	var data DataTable
	var err error
	this.E = ArrError{Status: 0, Msg: "服务器繁忙，请稍后再试...", Data: nil}

	// 处理查询数据信息
	query := this.Query(search)
	beego.Alert(query)

	// 查询数据
	data.Total, data.Count, err = models.FindAll(arr, query)
	if err == nil {
		this.E.Status = 1
		this.E.Msg = "Success"
		data.Data = arr
		this.E.Data = data
	}

	// 返回数据
	this.AjaxReturn()
}

// 公共的编辑的方法
func (this *AdminController) BaseUpdate(object interface{}, table string) {
	// 初始化返回信息
	this.E = ArrError{Status: 0, Msg: "请求数据为空", Data: nil}
	// 获取请求信息
	if actionType := this.GetString("actionType"); actionType != "" {
		this.E.Msg = "请求类型错误"
		this.E.Data = actionType
		// 判断请求类型
		if actionType == "insert" || actionType == "update" || actionType == "delete" {
			bTrue := true
			if actionType == "update" {
				// 修改数据需要先查询数据
				if id, err := this.GetInt64("id"); err == nil {
					if err := models.One(object, models.QueryOther{Table: table, Where: map[string]interface{}{"id": id}}); err != nil {
						this.E.Msg = err.Error()
						bTrue = false
					}
				} else {
					bTrue = false
				}
			}

			// 其他数据的处理
			if bTrue {
				this.E.Msg = "格式化数据出现错误"
				if err := this.ParseForm(object); err == nil {
					// 根据类型做出相应的处理
					switch actionType {
					case "insert": // 新增数据
						_, err = models.Insert(object)
					case "update": // 修改数据
						_, err = models.Update(object)
					case "delete": // 删除数据
						_, err = models.Delete(object)
					}

					// 判断返回数据
					if err == nil {
						this.E.Status = 1
						this.E.Msg = "恭喜你, 操作成功 ^.^"
						this.E.Data = object
					} else {
						this.E.Msg = "抱歉！执行该操作出现错误 Error：" + err.Error()
					}
				}
			}
		}
	}

	this.AjaxReturn()
}
