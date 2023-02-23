package create

import (
	"GoContractDeployment/models"
	"GoContractDeployment/repository"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
)

// NewSQLPostRepo 返回后存储库接口的实现
func NewSQLPostRepo(Conn *sql.DB) repository.PostRepo {
	return &MysqlPostRepo{
		Conn: Conn,
	}
}

type MysqlPostRepo struct {
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
//func (m *MysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) []*Post {
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

func (myRepo *MysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) []*models.DataPost {

	queryContext, err := myRepo.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("查询时异常")
	}
	//
	//payload := make([]*models.DataPost, 0)
	//for queryContext.Next() {
	//	data := &models.DataPost{}
	//	var createdAt []uint8
	//	err := queryContext.Scan(
	//		&data.ID,
	//		&data.Opcode,
	//		&data.ContractName,
	//		&data.ContractAddr,
	//		&data.ContractHash,
	//		&data.GasUsed,
	//		&data.GasUST,
	//		&data.ChainId,
	//		&createdAt,
	//		&data.CurrentStatus,
	//	)
	//	if err != nil {
	//		log.Println("转为实体类时异常", err)
	//	}
	//	createdTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	//	if err != nil {
	//		log.Println("解析时间戳时异常", err)
	//	}
	//	data.CreatedAt = createdTime
	//	payload = append(payload, data)
	//}
	payload := dealWith(queryContext)
	return payload
}

func (myRepo *MysqlPostRepo) AddJob(ctx context.Context, p []models.ReceivePost) string {
	//query := "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"

	args := make([]string, len(p))
	for i := 0; i < len(p); i++ {
		_, err := myRepo.Conn.ExecContext(ctx, models.InsertIntoJob, p[i].Opcode, p[i].ContractName, p[i].ChainId)

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

func (myRepo *MysqlPostRepo) Operate(ctx context.Context, status int64) []*models.DataPost {
	//query := "SELECT * FROM go_test_db WHERE current_status=?"

	post := myRepo.fetch(ctx, models.SelectOperate, status)

	return post
}

func (myRepo *MysqlPostRepo) GetOne() (*models.DataPost, error) {
	//query := "SELECT * FROM go_test_db WHERE current_status=0 LIMIT 1"

	queryContext, _ := myRepo.Conn.Query(models.SelectGetOne)

	post := dealWith(queryContext)
	if len(post) != 0 {
		return post[0], nil
	}

	return new(models.DataPost), errors.New("数据为空")
}

func (myRepo *MysqlPostRepo) UpdateTask(which string, dataPost models.DataPost) string {
	switch {
	case which == models.UpdateTaskOne:
		//query := "UPDATE go_test_db SET =?, =? WHERE id=?"
		stmt, err := myRepo.Conn.Prepare("UPDATE go_test_db SET contract_address=?, contract_hash=? ,gas_used=? ,gas_usdt=?, current_status=? WHERE id=?")

		if err != nil {
			panic(err.Error())
		}

		result, err := stmt.Exec(dataPost.ContractAddr, dataPost.ContractHash, dataPost.GasUsed, dataPost.GasUST, dataPost.CurrentStatus, dataPost.ID)
		if err != nil {
			panic(err.Error())
		}

		_, err = result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
	}

	return ""
}

// dealWith 处理为对象
func dealWith(queryContext *sql.Rows) []*models.DataPost {

	payload := make([]*models.DataPost, 0)

	if queryContext != nil {

		for queryContext.Next() {

			data := &models.DataPost{}

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
			if len(createdAt) > 0 {
				createdTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
				if err != nil {
					log.Println("解析时间戳时异常", err)
				}
				data.CreatedAt = createdTime
			}

			payload = append(payload, data)
		}
	}
	return payload
}
