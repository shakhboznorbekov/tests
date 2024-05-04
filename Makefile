run:
	go run cmd/main.go


swag-gen:
	swag init -g routes/router.go -o controllers/docs

