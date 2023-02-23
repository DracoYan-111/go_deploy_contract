package models

const (
	UpdateTaskOne = "update one"
)

const (
	InsertIntoJob = "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"
	SelectOperate = "SELECT * FROM go_test_db WHERE current_status=?"
	SelectGetOne  = "SELECT * FROM go_test_db WHERE current_status=0 LIMIT 1"
)
