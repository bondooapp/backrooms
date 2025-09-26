package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Result
//
// Result of response.
type Result struct {
	HttpStatus   int         `json:"httpStatus"`
	Success      bool        `json:"success"`
	Message      string      `json:"message"`
	RedirectType string      `json:"redirectType"`
	ShowType     string      `json:"showType"`
	TraceID      string      `json:"traceId"`
	Data         interface{} `json:"data"`
}

// page
//
// pagination param of response.
type page struct {
	Data     interface{} `json:"data"`
	Total    int64       `json:"total"`
	PageSize int         `json:"pageSize"`
	Current  int         `json:"current"`
	Pages    int         `json:"pages"`
}

// CreateOK
//
// Create ok of response.
func CreateOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusCreated, "create success", data)
}

// UpdateOK
//
// Update ok of response.
func UpdateOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusOK, "update success", data)
}

// QueryOK
//
// Query ok of response.
func QueryOK(c *gin.Context, data interface{}) {
	ok(c, http.StatusOK, "query success", data)
}

// QueryPageOK
//
// Query page ok of response.
func QueryPageOK(c *gin.Context, current int, pageSize int, total int64, data interface{}) {
	ok(c, http.StatusOK, "query success", page{
		Data:     data,
		Total:    total,
		PageSize: pageSize,
		Current:  current,
		Pages:    (int(total) + pageSize - 1) / pageSize,
	})
}

// DeleteOK
//
// Delete ok of response.
func DeleteOK(c *gin.Context) {
	ok(c, http.StatusOK, "delete success", nil)
}

// Error
//
// Generate error result of response.
func Error(c *gin.Context, httpStatus int, message string) (result Result) {
	return Result{
		HttpStatus:   httpStatus,
		Success:      false,
		Message:      message,
		RedirectType: "",
		ShowType:     "",
		TraceID:      getTraceId(c),
		Data:         nil,
	}
}

// ok
//
// Generate ok result of response.
func ok(c *gin.Context, httpStatus int, message string, data interface{}) {
	c.JSON(httpStatus, Result{
		HttpStatus:   httpStatus,
		Success:      true,
		Message:      message,
		RedirectType: "",
		ShowType:     "",
		TraceID:      getTraceId(c),
		Data:         data,
	})
}

// getTraceId
//
// Get traceId from context.
func getTraceId(c *gin.Context) string {
	traceId := c.GetString("traceId")
	if traceId == "" {
		traceId = "no traceId"
	}
	return traceId
}
