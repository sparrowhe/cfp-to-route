package handler

import (
	"github.com/labstack/echo/v4"
)

type StatusType struct {
	Status int
	Code   int
	Msg    string
}

var (
	StatusNormal        = StatusType{200, 10001, "Ok"}
	StatusUnderMaintain = StatusType{503, -1001, "Under maintain"}
	//StatusInternalError         = StatusType{500, -1101, "Internal error"}
	//StatusServiceError          = StatusType{500, -1102, "Service error"}
	//StatusExternalError         = StatusType{500, -1103, "External error"}
	StatusInvalidData           = StatusType{422, -1201, "Parameter invalid"}
	StatusTooMayRequest         = StatusType{429, -1205, "Too Many Request"}
	StatusResourceNotFound      = StatusType{404, -1202, "Resource not found"}
	StatusLogicalError          = StatusType{400, -1203, "Logical error"}
	StatusResourceExists        = StatusType{400, -1204, "Resource exists"}
	StatusNotAuthorized         = StatusType{401, -1301, "Not authorized"}
	StatusPermissionDenied      = StatusType{403, -1302, "Permission denied"}
	StatusUnreasonableOperation = StatusType{400, -1303, "Unreasonable operation"}
)

func Error(statusType StatusType, extraMsg string) StatusType {
	return StatusType{
		Status: statusType.Status,
		Code:   statusType.Code,
		Msg:    extraMsg,
	}
}

type Responser struct {
	Context echo.Context
}

func NewResponser(context echo.Context) Responser {
	return Responser{Context: context}
}
func (r Responser) Ok() error {
	return r.Context.JSON(StatusNormal.Status, echo.Map{
		"code": StatusNormal.Code,
	})
}
func (r Responser) Data(data interface{}) error {
	return r.Context.JSON(StatusNormal.Status, data)
}
func (r Responser) Error(code StatusType) error {
	return r.Context.JSON(code.Status, echo.Map{
		"code":    code.Code,
		"message": code.Msg,
	})
}
func ErrorResponse(c echo.Context, code StatusType) error {
	return c.JSON(code.Status, echo.Map{
		"code":    code.Code,
		"message": code.Msg,
	})
}
func Response(c echo.Context) error {
	return c.JSON(StatusNormal.Status, echo.Map{
		"code": StatusNormal.Code,
	})
}
func DataResponse(c echo.Context, data interface{}) error {
	return c.JSON(StatusNormal.Status, data)
}
