package models

const (
	UpdateTaskOne = "update one"
	TaskOneSql    = "UPDATE go_test_db SET contract_address=?, contract_hash=? ,gas_used=? ,gas_usdt=?, current_status=? WHERE id=?"
)

const (
	InsertIntoJob  = "INSERT INTO go_test_db (opcode, contract_name, chain_id) VALUES (?, ?, ?)"
	SelectOperate  = "SELECT * FROM go_test_db WHERE current_status=1"
	SelectGetOne   = "SELECT * FROM go_test_db WHERE current_status=0 LIMIT 1"
	UpdateStateOne = "UPDATE go_test_db SET current_status=2 WHERE id=?"
)
