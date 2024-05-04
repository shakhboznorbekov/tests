package transactionview

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/tests/internal/entity"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/pkg/repository/postgres"
)

type Repository struct {
	*postgres.Database
}

func NewRepository(postgresDB *postgres.Database) *Repository {
	return &Repository{postgresDB}
}

func (r Repository) AdminGetList(ctx context.Context, filter Filter) ([]AdminGetListResponse, int, *pkg.Error) {
	whereQuery := "WHERE deleted_at IS NULL"
	var limitQuery, offsetQuery string

	if filter.CustomerID != nil {
		whereQuery += fmt.Sprintf(" AND customer_id = %s", *filter.CustomerID)
	}
	if filter.ItemID != nil {
		whereQuery += fmt.Sprintf(" AND item_id = %s", *filter.ItemID)
	}

	if filter.Limit != nil {
		limitQuery = fmt.Sprintf("LIMIT %d", *filter.Limit)
	}

	if filter.Offset != nil {
		offsetQuery = fmt.Sprintf("OFFSET %d", *filter.Offset)
	}
	query := fmt.Sprintf(`
		SELECT
			id,
			customer_id,
			customer_name,
			qty,
			price,
			amount,
			customer_id,
			item_id
		FROM
		    transactionviews
		%s %s %s
	`, whereQuery, limitQuery, offsetQuery)

	rows, er := r.QueryContext(ctx, query)
	if er != nil {
		return nil, 0, &pkg.Error{
			Err:    pkg.WrapError(er, "selecting transactionviews list"),
			Status: http.StatusInternalServerError,
		}
	}
	var list []AdminGetListResponse
	for rows.Next() {
		var detail AdminGetListResponse
		if er = rows.Scan(&detail.Id, &detail.QTY, &detail.CustomerName, &detail.ItemName, &detail.Price, &detail.Amount, &detail.CustomerID, &detail.ItemID); er != nil {
			return nil, 0, &pkg.Error{
				Err:    pkg.WrapError(er, "scanning transactionviews"),
				Status: http.StatusInternalServerError,
			}
		}
		list = append(list, detail)
	}
	countQuery := fmt.Sprintf(`
	SELECT
	    COUNT(*)
	FROM
		transactionviews
	%s
	`, whereQuery)
	countRows, er := r.QueryContext(ctx, countQuery)
	if er != nil {
		return nil, 0, &pkg.Error{
			Err:    pkg.WrapError(er, "selecting transactionviews count"),
			Status: http.StatusInternalServerError,
		}
	}
	count := 0

	for countRows.Next() {
		if er = countRows.Scan(&count); er != nil {
			return nil, 0, &pkg.Error{
				Err:    pkg.WrapError(er, "scanning transactionviews count"),
				Status: http.StatusInternalServerError,
			}
		}
	}
	return list, count, nil
}

func (r Repository) AdminGetById(ctx context.Context, id string) (AdminGetDetail, *pkg.Error) {
	var detail AdminGetDetail

	err := r.NewSelect().Model(&detail).Where("id = ?", id).Scan(ctx)
	if err != nil {
		return AdminGetDetail{}, &pkg.Error{
			Err:    err,
			Status: http.StatusInternalServerError,
		}
	}
	return detail, nil
}

func (r Repository) AdminCreate(ctx context.Context, request AdminCreateRequest) (AdminCreateResponse, *pkg.Error) {
	var response AdminCreateResponse

	dataCtx, er := r.CheckCtx(ctx)
	if er != nil {
		return AdminCreateResponse{}, er
	}
	if err := r.ValidateStruct(&request, "CustomerID", "ItemID"); err != nil {
		return AdminCreateResponse{}, err
	}

	response.Id = uuid.NewString()
	response.CustomerName = request.CustomerName
	response.ItemName = request.ItemName
	response.QTY = request.QTY
	response.Price = request.Price
	response.Amount = request.Amount
	response.CustomerID = request.CustomerID
	response.ItemID = request.ItemID
	response.CreatedBy = &dataCtx.UserId
	response.CreatedAt = time.Now()
	err := r.ManualInsert(ctx, &response, "AdminCreate transactionviews")
	if err != nil {
		return AdminCreateResponse{}, err
	}

	return response, nil
}

func (r Repository) AdminUpdate(ctx context.Context, request AdminUpdateRequest) *pkg.Error {
	userData, err := r.AdminGetById(ctx, request.Id)
	if err != nil {
		return err
	}
	dataCtx, er := r.CheckCtx(ctx)
	if er != nil {
		return er
	}
	q := r.NewUpdate().Table("transactionviews").Where("deleted_at is null AND id = ?", request.Id)

	if request.CustomerName != nil {
		q.Set("customer_name = ?", request.CustomerName)

	}
	if request.ItemName != nil {
		q.Set("item_name = ?", request.ItemName)

	}
	if request.QTY != nil {
		q.Set("qty = ?", request.QTY)

	}
	if request.Price != nil {
		q.Set("price = ?", request.Price)

	}
	if request.Amount != nil {
		q.Set("amount = ?", request.Amount)

	}
	if request.CustomerID != nil {
		q.Set("costumer_id = ?", request.CustomerID)

	}
	if request.ItemID != nil {
		q.Set("item_id = ?", request.ItemID)

	}
	q.Set("updated_at = ?", time.Now())
	q.Set("updated_by = ?", dataCtx.UserId)

	_, err1 := q.Exec(ctx)
	if err1 != nil {
		return &pkg.Error{
			Err:    pkg.WrapError(err1, "updating transactionviews"),
			Status: http.StatusInternalServerError,
		}
	}
	newUpdateData, err := r.AdminGetById(ctx, request.Id)
	if err != nil {
		return err
	}
	updateData := map[string]interface{}{
		"oldData": userData,
		"newData": newUpdateData,
	}

	var loggerData entity.LogCreateDto
	loggerData.Action = "AdminUpdateAll"
	loggerData.Method = "PUT"
	loggerData.Data = updateData
	err2 := r.LogCreate(ctx, loggerData)
	if err2 != nil {
		return err2
	}
	return nil
}

func (r Repository) AdminDelete(ctx context.Context, id, username string) *pkg.Error {

	return r.DeleteRow(ctx, "transactionviews", id, username)
}
