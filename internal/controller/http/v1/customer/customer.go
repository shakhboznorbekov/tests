package customer

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"github.com/tests/internal/pkg"
	customer "github.com/tests/internal/repository/postgres/customer"
	"github.com/tests/internal/service/request"
	"github.com/tests/internal/service/response"
)

type Controller struct {
	customer Customer
}

func NewController(customer Customer) *Controller {
	return &Controller{customer: customer}
}

func (cl Controller) AdminGetCustomerList(c *gin.Context) {
	var filter customer.Filter
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


	if len(fieldErrors) > 0 {
		response.RespondError(c, &pkg.Error{
			Err:    errors.New("invalid query"),
			Fields: fieldErrors,
			Status: http.StatusBadRequest,
		})
		return
	}

	data, count, er := cl.customer.AdminGetList(c, filter)
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

func (cl Controller) AdminGetCustomerDetail(c *gin.Context) {
	idParam, err := request.GetParam(c, reflect.String, "id")
	var id string
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if value, ok := idParam.(string); ok {
		id = value
	}

	data, er := cl.customer.AdminGetById(c, id)
	if er != nil {
		response.RespondError(c, er)

		return
	}

	response.Respond(c, gin.H{
		"status": true,
		"data":   data,
	})
}

func (cl Controller) AdminCreateCustomer(c *gin.Context) {
	var data customer.AdminCreateRequest

	er := request.BindFunc(c, &data)
	if er != nil {
		response.RespondError(c, er)

		return
	}

	detail, er := cl.customer.AdminCreate(c, data)
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

func (cl Controller) AdminUpdateCustomer(c *gin.Context) {
	idParam, err := request.GetParam(c, reflect.String, "id")
	var id string
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else if value, ok := idParam.(string); ok {
		id = value
	}
	var data customer.AdminUpdateRequest

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

	err2 := cl.customer.AdminUpdate(c, data)
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

func (cl Controller) AdminDeleteCustomer(c *gin.Context) {
	idParam, err1 := request.GetParam(c, reflect.String, "id")
	var id string
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1)
	} else if value, ok := idParam.(string); ok {
		id = value
	}

	err := cl.customer.AdminDelete(c, id, "Admin")
	if err != nil {
		response.RespondError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok!",
		"status":  true,
	})
}
