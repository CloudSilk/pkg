package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `json:"id" gorm:"primarykey;size:36" copier:"-"`
	CreatedAt time.Time      `json:"createdAt" copier:"-"`
	UpdatedAt time.Time      `json:"updatedAt" gorm:"index" copier:"-"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index" copier:"-"`
}

func (u *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return
}

type TenantModel struct {
	Model
	TenantID string `json:"tenantID" gorm:"index;size:36"`
}

// Paging common input parameter structure
type PageInfo struct {
	PageIndex  int64  `json:"pageIndex" form:"pageIndex" uri:"pageIndex"`
	PageSize   int64  `json:"pageSize" form:"pageSize" uri:"pageSize"`
	Pages      int64  `json:"pages" form:"pages" uri:"pages"`
	Records    int64  `json:"records" form:"records" uri:"records"`
	OrderField string `json:"orderField" form:"orderField" uri:"orderField"`
	Desc       bool   `json:"desc" form:"desc" uri:"desc"`
	Total      int64  `json:"total" form:"total" uri:"total"`
	Current    int64  `json:"current" form:"current" uri:"current"`
}

func (p PageInfo) GetOrder(defaultOrderStr string) string {
	if p.OrderField != "" {
		if p.Desc {
			defaultOrderStr = p.OrderField + " desc"
		} else {
			defaultOrderStr = p.OrderField
		}
	}
	return defaultOrderStr
}

// Find by id structure
type GetById struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type IdsReq struct {
	IDs []string `json:"ids" form:"ids" uri:"ids"`
}

type Empty struct{}

type CommonResponse struct {
	PageInfo
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CommonRequest struct {
	PageInfo
	Token    string `json:"token"`
	TenantID string `json:"tenantID" form:"tenantID" uri:"tenantID"`
}

type CommonDetailResponse struct {
	CommonResponse
	Data interface{} `json:"data"`
}

type DeleteRequest struct {
	PageName string `json:"pageName"`
	ID       string `json:"id"`
}

type EnableRequest struct {
	PageName string `json:"pageName"`
	ID       string `json:"ID"`
	Enable   bool   `json:"enable"`
}

const (
	Success                   = 20000
	InternalServerError       = 50000
	BadRequest                = 40000
	Unauthorized              = 40001
	ErrRecordNotFound         = 40002
	UserNameOrPasswordIsWrong = 41001
	UserIsNotExist            = 41002
	NoPermission              = 41003
	TokenInvalid              = 41004
	TokenExpired              = 41005
	UserDisabled              = 41006
)
