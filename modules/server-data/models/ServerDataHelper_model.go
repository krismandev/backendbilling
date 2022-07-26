package models

import (
	"billingdashboard/connections"
	"billingdashboard/modules/server-data/datastruct"
	"strings"

	log "github.com/sirupsen/logrus"
)

func GetSingleServerData(serverDataID string, conn *connections.Connections, req datastruct.ServerDataRequest) (map[string]string, error) {
	var result map[string]string
	var err error

	// -- EXAMPLE
	// if len(server-dataID) == 0 {
	// 	server-dataID = req.ServerDataID
	// }
	// query := "SELECT server-dataid, server-dataname FROM server-data WHERE server-dataid = ?"
	// results, _, err := conn.DBAppConn.SelectQueryByFieldNameSlice(query, server-dataID)
	// if err != nil {
	// 	return result, err
	// }

	// // convert from []map[string]string to single map[string]string
	// for _, res := range results {
	// 	result = res
	// 	break
	// }
	return result, err
}

func LoadServerData(conn *connections.Connections, req datastruct.ServerDataRequest) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	var resultQuery []map[string]string
	var err error

	// runQuery := "SELECT server_data_id, server_id, server_account, item_id, account_id, external_smscount,external_transdate, external_transcount, invoice_id FROM server_data "
	runQuery := "SELECT LPAD(FLOOR(RAND() * 999999.99), 6, '0') as data_id,'11' as server_id,b.accountid as server_account,d.item_id,b.accountid as server_account,a.transdate,b.user_id,a.sender,a.operatorcode,c.category,sum(a.smscount),sum(a.transcount),null as price,now() as created_date FROM `bulksms_log_daily` a left join smsgw.smsgw_user b on a.userid=b.user_id left join smsgw.user_to_application c on b.user_id=c.user_id  left join dbbilling.item d on d.operator=a.operatorcode and d.route=c.category  and c.application_id=1 and d.category='USAGE' left join dbbilling.server_account e on e.serveraccount=b.accountid and e.server_id=@serverid where ((resultcode IN (0,1,2,3,5,4002,4003,4004,4005,4006,4007,4008, 4009,4010,4068) and smscid not in (35,36,143,144)) or (resultcode IN (0,1,2,3,5) and smscid in (35,36,143,144))) and c.application_id=1 and d.category='USAGE' and date_format(a.transdate,'%Y%m') = '202204' group by b.accountid,d.item_id,e.account_id,a.transdate,b.user_name,a.sender,a.operatorcode,c.category"

	resultQuery, num, err := conn.DBDashbConn.SelectQueryByFieldNameSlice(runQuery)

	log.Info("LihatNum", num)
	log.Info("LihatData", resultQuery)
	bulkInserQuery := "INSERT IGNORE INTO server_data(server_data_id,server_id,server_account,item_id,account_id,external_transdate,external_user_id,external_sender,external_operatorcode,external_route,external_smscount,external_transcount,external_price,created_date) VALUES "
	var paramsBulkInsert []interface{}
	var stringGroup []string

	for _, each := range resultQuery {
		partquery := "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		paramsBulkInsert = append(paramsBulkInsert, each["data_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_account"])
		paramsBulkInsert = append(paramsBulkInsert, each["item_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["server_account"])
		paramsBulkInsert = append(paramsBulkInsert, each["transdate"])
		paramsBulkInsert = append(paramsBulkInsert, each["user_id"])
		paramsBulkInsert = append(paramsBulkInsert, each["sender"])
		paramsBulkInsert = append(paramsBulkInsert, each["operatorcode"])
		paramsBulkInsert = append(paramsBulkInsert, each["category"])
		paramsBulkInsert = append(paramsBulkInsert, each["sum(a.smscount)"])
		paramsBulkInsert = append(paramsBulkInsert, each["sum(a.transcount)"])
		paramsBulkInsert = append(paramsBulkInsert, each["price"])
		paramsBulkInsert = append(paramsBulkInsert, each["created_date"])
		stringGroup = append(stringGroup, partquery)
	}

	final_query := bulkInserQuery + strings.Join(stringGroup, ", ")
	log.Info("FinalQuery", final_query)

	_, _, errInsert := conn.DBAppConn.Exec(final_query, paramsBulkInsert...)
	if errInsert != nil {
		return nil, err
	}

	return result, err
}

func CheckServerDataExists(serverDataID string, conn *connections.Connections) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, server-dataID)

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount == 0 {
	// 	return errors.New("ServerData ID is not exists")
	// }
	return nil
}

func CheckServerDataDuplicate(exceptID string, conn *connections.Connections, req datastruct.ServerDataRequest) error {
	// var param []interface{}
	// qry := "SELECT COUNT(server-dataid) FROM server-data WHERE server-dataid = ?"
	// param = append(param, req.ServerDataID)
	// if len(exceptID) > 0 {
	// 	qry += " AND server-dataid <> ?"
	// 	param = append(param, exceptID)
	// }

	// cnt, _ := conn.DBAppConn.GetFirstData(qry, param...)
	// datacount, _ := strconv.Atoi(cnt)
	// if datacount > 0 {
	// 	return errors.New("ServerData ID is already exists. Please use another ServerData ID")
	// }
	return nil
}
