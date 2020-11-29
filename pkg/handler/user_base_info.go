package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/nizonglonggit/userapisrv/pkg/define"
	"github.com/nizonglonggit/userapisrv/pkg/model/psql"
	"net/http"
	"time"
)

func UserRegister(c *gin.Context) {
	var req define.RegisterUserReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusOK, result(define.CodeParamError, define.ErrParam+err.Error(), nil))
		return
	}

	user := define.User{
		NickName:        req.NickName,
		Gender:          req.Gender,
		Email:           req.Email,
		HeadPortraitUrl: req.HeadPortraitUrl,
	}

	// pwd
	user.PassWord = "md5:" + req.PassWord

	if req.HeadPortraitUrl == "" {
		user.HeadPortraitUrl = "test head url"
	}

	user.RegDate = time.Now().Format(define.YMDLayout)

	err := psql.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusOK, result(define.CodeDBError, define.ErrDB+err.Error(), nil))
		return
	}

	c.JSON(http.StatusOK, result(define.CodeSuccess, define.StatusOk, nil))
}

func ListAllUser(c *gin.Context) {
	users, err := psql.ListAllUsers()
	if err != nil {
		c.JSON(http.StatusOK, result(define.CodeDBError, define.ErrDB+err.Error(), nil))
		return
	}
	c.JSON(http.StatusOK, result(define.CodeSuccess, define.StatusOk, users))
}
