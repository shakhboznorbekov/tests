package user

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/tests/internal/entity"
	"github.com/tests/internal/pkg"
	"github.com/tests/internal/pkg/repository/postgres"
	"github.com/tests/internal/service/hash"
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

	if filter.FirstName != nil {
		firstName := strings.Replace(*filter.FirstName, " ", "", -1)
		whereQuery += fmt.Sprintf(" AND REPLACE(first_name, ' ', '') ilike '%s'", "%"+firstName+"%")
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
			username,
			first_name,
			last_name,
			status
		FROM
		    users
		%s %s %s
	`, whereQuery, limitQuery, offsetQuery)

	rows, er := r.QueryContext(ctx, query)
	if er != nil {
		return nil, 0, &pkg.Error{
			Err:    pkg.WrapError(er, "selecting user list"),
			Status: http.StatusInternalServerError,
		}
	}
	var list []AdminGetListResponse
	for rows.Next() {
		var detail AdminGetListResponse
		if er = rows.Scan(&detail.Id, &detail.Username, &detail.FirstName, &detail.LastName, &detail.Status); er != nil {
			return nil, 0, &pkg.Error{
				Err:    pkg.WrapError(er, "scanning user"),
				Status: http.StatusInternalServerError,
			}
		}
		list = append(list, detail)
	}
	countQuery := fmt.Sprintf(`
	SELECT
	COUNT(*)
	FROM
		users
	%s
	`, whereQuery)
	countRows, er := r.QueryContext(ctx, countQuery)
	if er != nil {
		return nil, 0, &pkg.Error{
			Err:    pkg.WrapError(er, "selecting user count"),
			Status: http.StatusInternalServerError,
		}
	}
	count := 0

	for countRows.Next() {
		if er = countRows.Scan(&count); er != nil {
			return nil, 0, &pkg.Error{
				Err:    pkg.WrapError(er, "scanning user count"),
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

func (r Repository) GetByFirstName(ctx context.Context, firstName string) (AdminGetDetail, *pkg.Error) {
	var detail AdminGetDetail
	err := r.NewSelect().Model(&detail).Where("first_name = ?", firstName).Scan(ctx)

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
	if err := r.ValidateStruct(&request, "FirstName", "LastName", "Username", "Status", "Gmail", "Password"); err != nil {
		return AdminCreateResponse{}, err
	}

	response.Id = uuid.NewString()
	response.FirstName = request.FirstName
	response.LastName = request.LastName
	response.Username = request.Username
	hashPassword, err2 := hash.HashPassword(request.Password)
	if err2 != nil {
		return AdminCreateResponse{}, &pkg.Error{
			Err:    err2,
			Status: http.StatusInternalServerError,
		}
	}
	response.Status = request.Status
	response.Password = hashPassword
	response.Gmail = request.Gmail
	response.CreatedBy = &dataCtx.UserId
	response.CreatedAt = time.Now()
	err := r.ManualInsert(ctx, &response, "AdminCreate")
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
	q := r.NewUpdate().Table("users").Where("deleted_at is null AND id = ?", request.Id)

	if request.FirstName != nil {
		q.Set("first_name = ?", request.FirstName)

	}
	if request.LastName != nil {
		q.Set("last_name = ?", request.LastName)

	}
	if request.Username != nil {
		q.Set("username = ?", request.Username)

	}
	if request.Password != nil {
		q.Set("password = ?", request.Password)

	}
	if request.Status != nil {
		q.Set("status = ?", request.Status)

	}
	if request.Gmail != nil {
		q.Set("gmail = ?", request.Gmail)

	}
	q.Set("updated_at = ?", time.Now())
	q.Set("updated_by = ?", dataCtx.UserId)

	_, err1 := q.Exec(ctx)
	if err1 != nil {
		return &pkg.Error{
			Err:    pkg.WrapError(err1, "updating user"),
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

	return r.DeleteRow(ctx, "users", id, username)
}
