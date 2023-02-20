package post

import (
	models "GoContractDeployment/models"
	pRepo "GoContractDeployment/repository"
	"context"
	"database/sql"
	"log"
	"time"
)

// NewSQLPostRepo 返回后存储库接口的实现
func NewSQLPostRepo(Conn *sql.DB) pRepo.PostRepo {
	return &mysqlPostRepo{
		Conn: Conn,
	}
}

type mysqlPostRepo struct {
	Conn *sql.DB
}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {

	queryContext, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	var payload []*models.Post
	var created string

	for queryContext.Next() {
		data := &models.Post{}

		err = queryContext.Scan(
			&data.ID,
			&data.Opcode,
			&data.ContractName,
			&data.ContractAddr,
			&data.ContractHash,
			&data.GasUsed,
			&data.GasUST,
			&data.ChainId,
			&created,
			&data.CurrentStatus,
		)

		data.CreatedAt, err = time.Parse("2006-01-02 15:04:05", created)

		if err != nil {
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}
func (m *mysqlPostRepo) AddJob(ctx context.Context, p models.Post) string {
	return "s"
}

func (m *mysqlPostRepo) Create(ctx context.Context, p *models.ReceivePost) string {
	query := "Insert posts SET title=?, content=?"

	log.Println(p.Opcode, p.ChainId, p.ContractName)

	_, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
	}
	//
	//res, err := stmt.ExecContext(ctx, p.Title, p.Content)
	//defer stmt.Close()
	//
	//if err != nil {
	//	return -1, err
	//}

	return "res.LastInsertId()"
}

func (m *mysqlPostRepo) Fetch(ctx context.Context, id int64) *models.Post {
	query := "Select * From go_test_db where id=?"
	post, err := m.fetch(ctx, query, id)

	if err != nil || len(post) > 1 {
		log.Fatal("====通过ID查询异常====", err)
	}
	return post[0]
}
