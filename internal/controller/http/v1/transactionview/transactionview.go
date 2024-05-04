package transactionview

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/repository/postgres/transactionview"
	"github.com/tests/internal/service/request"
	"github.com/tests/internal/service/response"
)

type Controller struct {
	transactionview TransactionView
}

func NewController(transactionview TransactionView) *Controller {
	return &Controller{transactionview: transactionview}
}

func (cl Controller) AdminGetTransactionViewList(c *gin.Context) {
	var filter transactionview.Filter
	fieldErrors := make([]pkg.FieldError, 0)

	limit, err := request.GetQuery(c, reflect.Int, "limit")
	if err != nil {
		fieldErrors = append(fieldErrors, *err)
	} else if value, ok := limit.(*int); ok {
		filter.Limit = value
	}

	offset, err := request.GetQuery(c, reflect.Int, "offset")
	if err != nil {
		fieldErrors = append(fieldErrors, *err)
	} else if value, ok := offset.(*int); ok {
		filter.Offset = value
	}

	costumer_id, err := request.GetQuery(c, reflect.String, "costumer_id")
	if err != nil {
		fieldErrors = append(fieldErrors, *err)
	} else if value, ok := costumer_id.(*string); ok {
		filter.CustomerID = value
	}
	item_id, err := request.GetQuery(c, reflect.String, "item_id")
	if err != nil {
		fieldErrors = append(fieldErrors, *err)
	} else if value, ok := item_id.(*string); ok {
		filter.ItemID = value
	}

	if len(fieldErrors) > 0 {
		response.RespondError(c, &pkg.Error{
			Err:    errors.New("invalid query"),
			Fields: fieldErrors,
			Status: http.StatusBadRequest,
		})
		return
	}

	data, count, er := cl.transactionview.AdminGetList(c, filter)
	if er != nil {
		response.RespondError(c, er)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data": map[string]interface{}{
			"results": data,
			"count":   count,
		},
	})
}

func (cl Controller) AdminGetTransactionViewDetail(c *gin.Context) {
	idParam, err := request.GetParam(c, reflect.String, "id")
	var id string
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if value, ok := idParam.(string); ok {
		id = value
	}

	data, er := cl.transactionview.AdminGetById(c, id)
	if er != nil {
		response.RespondError(c, er)

		return
	}

	response.Respond(c, gin.H{
		"status": true,
		"data":   data,
	})
}

func (cl Controller) AdminCreateTransactionView(c *gin.Context) {
	var data transactionview.AdminCreateRequest

	er := request.BindFunc(c, &data)
	if er != nil {
		response.RespondError(c, er)

		return
	}

	detail, er := cl.transactionview.AdminCreate(c, data)
	if er != nil {
		response.RespondError(c, er)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
		"data":    detail,
	})
}

func (cl Controller) AdminUpdateTransactionView(c *gin.Context) {
	idParam, err := request.GetParam(c, reflect.String, "id")
	var id string
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if value, ok := idParam.(string); ok {
		id = value
	}
	var data transactionview.AdminUpdateRequest

	er := request.BindFunc(c, &data)
	if er != nil {
		c.JSON(er.Status, gin.H{
			"message": er.Err.Error(),
			"status":  false,
		})

		return
	}
	if data.Id == "" {
		data.Id = id
	}

	err2 := cl.transactionview.AdminUpdate(c, data)
	if err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err2.Err.Error(),
			"status":  false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
	})
}

func (cl Controller) AdminDeleteTransactionView(c *gin.Context) {
	idParam, err1 := request.GetParam(c, reflect.String, "id")
	var id string
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
	} else if value, ok := idParam.(string); ok {
		id = value
	}

	err := cl.transactionview.AdminDelete(c, id, "Admin")
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
	})
}
