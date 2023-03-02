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

// NewSQLPostRepo An implementation of the repository interface after returning
func NewSQLPostRepo(Conn *sql.DB) repository.PostRepo {
	return &MysqlPostRepo{
		Conn: Conn,
	}
}

type MysqlPostRepo struct {
	Conn *sql.DB
}

func (myRepo *MysqlPostRepo) fetch(ctx context.Context, query string, args ...interface{}) []*models.DataPost {

	queryContext, err := myRepo.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Println("PostMysql:查询时异常")
	}
	payload := dealWith(queryContext)
	return payload
}

func (myRepo *MysqlPostRepo) AddJob(ctx context.Context, p []models.ReceivePost) string {
	//query := "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"

	args := make([]string, len(p))
	for i := 0; i < len(p); i++ {
		_, err := myRepo.Conn.ExecContext(ctx, models.InsertIntoJob, p[i].Opcode, p[i].ContractName, p[i].ChainId)

		if err != nil {
			log.Println("<==== PostMysql:插入数据异常 ====>", err)
			continue
		}
		args[i] = p[i].Opcode
	}
	log.Println("<==== PostMysql:插入数据成功 ====>", args)

	return fmt.Sprintf("%v", args)
}

func (myRepo *MysqlPostRepo) Operate() ([]*models.DataPost, error) {
	//query := "SELECT * FROM go_test_db WHERE current_status=2"

	queryContext, _ := myRepo.Conn.Query(models.SelectOperate)

	post := dealWith(queryContext)
	if len(post) != 0 {
		return post, nil
	}
	var nilPost []*models.DataPost

	return nilPost, errors.New("Operate:数据为空")
}

func (myRepo *MysqlPostRepo) GetOne() (*models.DataPost, error) {
	//query := "SELECT * FROM go_test_db WHERE current_status=0 LIMIT 1"

	queryContext, _ := myRepo.Conn.Query(models.SelectGetOne)

	post := dealWith(queryContext)
	if len(post) != 0 {
		return post[0], nil
	}

	return new(models.DataPost), errors.New("GetOne:数据为空")
}

func (myRepo *MysqlPostRepo) UpdateTask(which string, dataPost models.DataPost) string {
	switch {
	case which == models.UpdateTaskOne:
		//query := "UPDATE go_test_db SET contract_address=?, contract_hash=? ,gas_used=? ,gas_usdt=?, current_status=? WHERE id=?"
		stmt, err := myRepo.Conn.Prepare(models.TaskOneSql)
		if err != nil {
			panic(err.Error())
		}
		result, err := stmt.Exec(dataPost.ContractAddr, dataPost.ContractHash, dataPost.GasUsed, dataPost.GasUST, dataPost.CurrentStatus, dataPost.ID)
		if err != nil {
			panic(err.Error())
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		log.Println("rowsAffected", rowsAffected)
	}

	return ""
}

func (myRepo *MysqlPostRepo) UpdateState(idArray []int64) string {
	//query := "UPDATE go_test_db SET current_status=2 WHERE id=?"

	for i := 0; i < len(idArray); i++ {
		stmt, err := myRepo.Conn.Prepare(models.UpdateStateOne)
		if err != nil {
			panic(err.Error())
		}
		result, err := stmt.Exec(idArray[i])
		if err != nil {
			panic(err.Error())
		}

		_, err = result.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
	}
	return "<++++ PostMysql:状态更新完成 ++++>"
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
				log.Println("PostMysql:转换为实体类时出现异常", err)
			}
			if len(createdAt) > 0 {
				createdTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
				if err != nil {
					log.Println("PostMysql:解析时间戳时出现异常", err)
				}
				data.CreatedAt = createdTime
			}

			payload = append(payload, data)
		}
	}
	return payload
}
