package dict

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"go-admin/app/admin/models/system"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/apis"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type SysDictData struct {
	apis.Api
}

// @Summary 字典数据列表
// @Description 获取JSON
// @Tags 字典数据
// @Param status query string false "status"
// @Param dictCode query string false "dictCode"
// @Param dictType query string false "dictType"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data [get]
// @Security Bearer
func (e *SysDictData) GetSysDictDataList(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := &dto.SysDictDataSearch{}

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	list := make([]system.SysDictData, 0)
	var count int64
	s := service.SysDictData{}
	s.MsgID = msgID
	s.Orm = db.Debug()
	err = s.GetPage(req, &list, &count)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.PageOK(c, list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// @Summary 通过编码获取字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Param dictCode path int true "字典编码"
// @Success 200 {object} app.Response "{"code": 200, "data": [...]}"
// @Router /api/v1/dict/data/{dictCode} [get]
// @Security Bearer
func (e *SysDictData) GetSysDictData(c *gin.Context) {
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//查看详情
	req := &dto.SysDictDataById{}
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object system.SysDictData

	s := service.SysDictData{}
	s.MsgID = msgID
	s.Orm = db
	err = s.Get(req, &object)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, object, "查看成功")
}

// @Summary 添加字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body system.DictType true "data"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data [post]
// @Security Bearer
func (e *SysDictData) InsertSysDictData(c *gin.Context) {
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//新增操作
	req := &dto.SysDictDataControl{}
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, _ := req.GenerateM()
	// 设置创建人
	object.SetCreateBy(tools.GetUserId(c))

	s := service.SysDictData{}
	s.Orm = db
	s.MsgID = msgID
	err = s.Insert(object.(*system.SysDictData))
	if err != nil {
		log.Error(err)
		e.Error(c, http.StatusInternalServerError, err, "创建失败")
		return
	}

	e.OK(c, object.GetId(), "创建成功")
}

// @Summary 修改字典数据
// @Description 获取JSON
// @Tags 字典数据
// @Accept  application/json
// @Product application/json
// @Param data body models.DictType true "body"
// @Success 200 {string} string	"{"code": 200, "message": "添加成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "添加失败"}"
// @Router /api/v1/dict/data/{dictCode} [put]
// @Security Bearer
func (e *SysDictData) UpdateSysDictData(c *gin.Context) {
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	req := &dto.SysDictDataControl{}
	//更新操作
	err = req.Bind(c)
	if err != nil {
		log.Errorf("msgID[%s] request validate error, %s", msgID, err.Error())
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	object, _ := req.GenerateM()
	object.SetUpdateBy(tools.GetUserId(c))

	s := service.SysDictData{}
	s.Orm = db
	s.MsgID = msgID
	err = s.Update(object.(*system.SysDictData))
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "更新成功")
}

// @Summary 删除字典数据
// @Description 删除数据
// @Tags 字典数据
// @Param dictCode path int true "dictCode"
// @Success 200 {string} string	"{"code": 200, "message": "删除成功"}"
// @Success 200 {string} string	"{"code": -1, "message": "删除失败"}"
// @Router /api/v1/dict/data/{dictCode} [delete]
func (e *SysDictData) DeleteSysDictData(c *gin.Context) {
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	msgID := tools.GenerateMsgIDFromContext(c)
	//删除操作
	req := new(dto.SysDictDataById)
	err = req.Bind(c)
	if err != nil {
		log.Errorf("MsgID[%s] Bind error: %s", msgID, err)
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}
	var object common.ActiveRecord
	object, err = req.GenerateM()
	if err != nil {
		e.Error(c, http.StatusInternalServerError, err, "模型生成失败")
		return
	}

	// 设置编辑人
	object.SetUpdateBy(tools.GetUserId(c))

	s := service.SysDictData{}
	s.Orm = db
	s.MsgID = msgID
	err = s.Remove(req, object.(*system.SysDictData))
	if err != nil {
		log.Error(err)
		return
	}
	e.OK(c, object.GetId(), "删除成功")
}

func (e *SysDictData) GetSysDictDataAll(c *gin.Context) {
	msgID := tools.GenerateMsgIDFromContext(c)
	db, err := tools.GetOrm(c)
	if err != nil {
		log.Error(err)
		return
	}

	req := &dto.SysDictDataSearch{}

	//查询列表
	err = req.Bind(c)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "参数验证失败")
		return
	}

	list := make([]system.SysDictData, 0)
	s := service.SysDictData{}
	s.MsgID = msgID
	s.Orm = db
	err = s.GetAll(req, &list)
	if err != nil {
		e.Error(c, http.StatusUnprocessableEntity, err, "查询失败")
		return
	}

	e.OK(c, list, "查询成功")
}