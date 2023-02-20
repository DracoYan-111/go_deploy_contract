package create

import (
	models "GoContractDeployment/models"
	pRepo "GoContractDeployment/repository"
	"context"
	"database/sql"
	"fmt"
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

//// Post 数据库的信息结构
//type Post struct {
//	ID            int64     `json:"id"`
//	Opcode        string    `json:"opcode"`
//	ContractName  string    `json:"contract_name"`
//	ContractAddr  string    `json:"contract_address"`
//	ContractHash  string    `json:"contract_hash"`
//	GasUsed       int64     `json:"gas_price"`
//	GasUST        int64     `json:"gas_usdt"`
//	ChainId       int64     `json:"chain_id"`
//	CreatedAt     time.Time `json:"created_at"`
//	CurrentStatus int64     `json:"current_status"`
//}
//
//func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) []*Post {
//
//	queryContext, err := m.Conn.QueryContext(ctx, query, args...)
//	if err != nil {
//		log.Println("查询时异常")
//	}
//
//	payload := make([]*Post, 0)
//	for queryContext.Next() {
//		data := &Post{}
//		err := queryContext.Scan(
//			&data.ID,
//			&data.Opcode,
//			&data.ContractName,
//			&data.ContractAddr,
//			&data.ContractHash,
//			&data.GasUsed,
//			&data.GasUST,
//			&data.ChainId,
//			&data.CreatedAt,
//			&data.CurrentStatus,
//		)
//		if err != nil {
//			log.Println("转为实体类时异常", err)
//		}
//		payload = append(payload, data)
//	}
//	return payload
//}

func (m *mysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) []*models.Post {

	queryContext, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("查询时异常")
	}

	payload := make([]*models.Post, 0)
	for queryContext.Next() {
		data := &models.Post{}
		var createdAt []uint8
		err := queryContext.Scan(
			&data.ID,
			&data.Opcode,
			&data.ContractName,
			&data.ContractAddr,
			&data.ContractHash,
			&data.GasUsed,
			&data.GasUST,
			&data.ChainId,
			&createdAt,
			&data.CurrentStatus,
		)
		if err != nil {
			log.Println("转为实体类时异常", err)
		}
		createdTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			log.Println("解析时间戳时异常", err)
		}
		data.CreatedAt = createdTime
		payload = append(payload, data)
	}
	return payload
}

func (m *mysqlPostRepo) AddJob(ctx context.Context, p []models.ReceivePost) string {
	query := "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"

	args := make([]string, len(p))
	for i := 0; i < len(p); i++ {
		_, err := m.Conn.ExecContext(ctx, query, p[i].Opcode, p[i].ContractName, p[i].ChainId)

		if err != nil {
			log.Println("====插入数据异常====", err)
			continue
		}
		// 获取最新Id
		//id, err := res.LastInsertId()
		//if err != nil {
		//	log.Println("====获取最新ID异常====", err)
		//	continuer
		//} else {
		//	log.Println("++++获取最新ID成功++++")
		//}
		args[i] = p[i].Opcode
	}
	log.Println("成功插入数据", args)

	return fmt.Sprintf("%v", args)
}

func (m *mysqlPostRepo) Operate(ctx context.Context, status int64) []*models.Post {
	query := "SELECT * FROM go_test_db WHERE current_status=?"
	post := m.fetch(ctx, query, status)

	return post
}
