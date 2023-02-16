package post

import (
	models "GoContractDeployment/models"
	pRepo "GoContractDeployment/repository"
	"context"
	"database/sql"
)

// NewSQLPostRepo retunrs implement of post repository interface
func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			panic(err)
		}
	}(rows)

	payload := make([]*models.Post, 0)
	for rows.Next() {
		data := new(models.Post)

		err := rows.Scan(
			&data.ID,
			&data.Opcode,
			&data.ContractName,
			&data.ContractAddr,
			&data.ContractHash,
			&data.GasUsed,
			&data.GasUST,
			&data.ChainId,
			&data.CreatedAt,
			&data.CurrentStatus,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, num int64) ([]*models.Post, error) {
	query := "Select id, title, content From posts limit ?"

	return m.fetch(ctx, query, num)
}

func (m *mysqlPostRepo) GetByID(ctx context.Context, id int64) (*models.Post, error) {
	query := "Select id, title, content From posts where id=?"

	rows, err := m.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	payload := &models.Post{}
	if len(rows) > 0 {
		payload = rows[0]
	} else {
		return nil, models.ErrNotFound
	}

	return payload, nil
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *models.Post) (int64, error) {
	query := "Insert posts SET title=?, content=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}

	res, err := stmt.ExecContext(ctx, p.ContractAddr, p.ChainId)
	defer func(stmt *sql.Stmt) {

		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}(stmt)

	if err != nil {
		return -1, err
	}

	return res.LastInsertId()
}

func (m *mysqlPostRepo) Update(ctx context.Context, p *models.Post) (*models.Post, error) {
	query := "Update posts set title=?, content=? where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	_, err = stmt.ExecContext(
		ctx,
		p.ContractAddr,
		p.ChainId,
		p.ID,
	)
	if err != nil {
		return nil, err
	}
	defer func(stmt *sql.Stmt) {
		if err := stmt.Close(); err != nil {
			panic(err)
		}
	}(stmt)

	return p, nil
}

func (m *mysqlPostRepo) Delete(ctx context.Context, id int64) (bool, error) {
	query := "Delete From posts Where id=?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return false, err
	}
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return false, err
	}
	return true, nil
}
