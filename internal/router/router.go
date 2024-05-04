package router

import (
	"github.com/gin-gonic/gin"
)

type Auth interface {
	HasPermission(roles ...string) gin.HandlerFunc
}

type Authorization interface {
	SignIn(*gin.Context)
}

type User interface {
	AdminGetUserList(*gin.Context)
	AdminGetUserDetail(*gin.Context)
	AdminCreateUser(*gin.Context)
	AdminUpdateUser(*gin.Context)
	AdminDeleteUser(*gin.Context)
}

type Item interface {
	AdminGetItemList(*gin.Context)
	AdminGetItemDetail(*gin.Context)
	AdminCreateItem(*gin.Context)
	AdminUpdateItem(*gin.Context)
	AdminDeleteItem(*gin.Context)
}
type Customer interface {
	AdminGetCustomerList(*gin.Context)
	AdminGetCustomerDetail(*gin.Context)
	AdminCreateCustomer(*gin.Context)
	AdminUpdateCustomer(*gin.Context)
	AdminDeleteCustomer(*gin.Context)
}
type Transaction interface {
	AdminGetTransactionList(*gin.Context)
	AdminGetTransactionDetail(*gin.Context)
	AdminCreateTransaction(*gin.Context)
	AdminUpdateTransaction(*gin.Context)
	AdminDeleteTransaction(*gin.Context)
}
type TransactionView interface {
	AdminGetTransactionViewList(*gin.Context)
	AdminGetTransactionViewDetail(*gin.Context)
	AdminCreateTransactionView(*gin.Context)
	AdminUpdateTransactionView(*gin.Context)
	AdminDeleteTransactionView(*gin.Context)
}
type Router struct {
	auth            Auth
	user            User
	authorization   Authorization
	item            Item
	customer        Customer
	transaction     Transaction
	transactionview TransactionView
}

func New(auth Auth, user User, authorization Authorization, item Item, customer Customer, transaction Transaction, transactionview TransactionView) *Router {
	return &Router{auth: auth,
		user:            user,
		authorization:   authorization,
		item:            item,
		customer:        customer,
		transaction:     transaction,
		transactionview: transactionview,
	}
}

func (r *Router) Init(port string) error {
	router := gin.Default()

	// gin engine
	router.Use(customCORSMiddleware())

	// auth
	router.POST("/api/v1/user/sign-in", r.authorization.SignIn)

	//user
	router.GET("/api/v1/admin/user/list", r.auth.HasPermission("Admin"), r.user.AdminGetUserList)
	router.GET("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminGetUserDetail)
	router.POST("/api/v1/admin/user/create", r.auth.HasPermission("Admin"), r.user.AdminCreateUser)
	router.PUT("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminUpdateUser)
	router.DELETE("/api/v1/admin/user/:id", r.auth.HasPermission("Admin"), r.user.AdminDeleteUser)

	//item
	router.GET("/api/v1/admin/item/list", r.auth.HasPermission("Admin"), r.item.AdminGetItemList)
	router.GET("/api/v1/admin/item/:id", r.auth.HasPermission("Admin"), r.item.AdminGetItemDetail)
	router.POST("/api/v1/admin/item/create", r.auth.HasPermission("Admin"), r.item.AdminCreateItem)
	router.PUT("/api/v1/admin/item/:id", r.auth.HasPermission("Admin"), r.item.AdminUpdateItem)
	router.DELETE("/api/v1/admin/item/:id", r.auth.HasPermission("Admin"), r.item.AdminDeleteItem)

	//customer
	router.GET("/api/v1/admin/customer/list", r.auth.HasPermission("Admin"), r.customer.AdminGetCustomerList)
	router.GET("/api/v1/admin/customer/:id", r.auth.HasPermission("Admin"), r.customer.AdminGetCustomerDetail)
	router.POST("/api/v1/admin/customer/create", r.auth.HasPermission("Admin"), r.customer.AdminCreateCustomer)
	router.PUT("/api/v1/admin/customer/:id", r.auth.HasPermission("Admin"), r.customer.AdminUpdateCustomer)
	router.DELETE("/api/v1/admin/customer/:id", r.auth.HasPermission("Admin"), r.customer.AdminDeleteCustomer)

	//transaction
	router.GET("/api/v1/admin/transaction/list", r.auth.HasPermission("Admin"), r.transaction.AdminGetTransactionList)
	router.GET("/api/v1/admin/transaction/:id", r.auth.HasPermission("Admin"), r.transaction.AdminGetTransactionDetail)
	router.POST("/api/v1/admin/transaction/create", r.auth.HasPermission("Admin"), r.transaction.AdminCreateTransaction)
	router.PUT("/api/v1/admin/transaction/:id", r.auth.HasPermission("Admin"), r.transaction.AdminUpdateTransaction)
	router.DELETE("/api/v1/admin/transaction/:id", r.auth.HasPermission("Admin"), r.transaction.AdminDeleteTransaction)

	//transactionview
	router.GET("/api/v1/admin/transactionview/list", r.auth.HasPermission("Admin"), r.transactionview.AdminGetTransactionViewList)
	router.GET("/api/v1/admin/transactionview/:id", r.auth.HasPermission("Admin"), r.transactionview.AdminGetTransactionViewDetail)
	router.POST("/api/v1/admin/transactionview/create", r.auth.HasPermission("Admin"), r.transactionview.AdminCreateTransactionView)
	router.PUT("/api/v1/admin/transactionview/:id", r.auth.HasPermission("Admin"), r.transactionview.AdminUpdateTransactionView)
	router.DELETE("/api/v1/admin/transactionview/:id", r.auth.HasPermission("Admin"), r.transactionview.AdminDeleteTransactionView)

	return router.Run(port)
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Max-Age", "3600")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)

			return
		}

		c.Next()
	}
}
