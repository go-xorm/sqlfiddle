package sqlfiddle

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const (
	Mysql5_6 = 9
)

/*
{
  "_id" : "9_ed3b0c",
  "short_code" : "ed3b0c",
  "schema_structure" : [ {
    "columns" : [ {
      "name" : "id",
      "type" : "INT(10)"
    }, {
      "name" : "name",
      "type" : "VARCHAR(8)"
    }, {
      "name" : "birthday",
      "type" : "DATETIME(19)"
    } ],
    "table_name" : "person",
    "table_type" : "TABLE"
  } ]
}
*/

type column struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type schemaStructure struct {
	Columns   []column `json:"columns"`
	TableName string   `json:"table_name"`
	TableType string   `json:"table_type"`
}

type response struct {
	Id         string            `json:"_id"`
	Code       string            `json:"short_code"`
	Structures []schemaStructure `json:"schema_structure"`
}

func CreateSchema(dbType int, sql string) (*response, error) {
	payload := map[string]interface{}{"statement_separator": ";", "db_type_id": dbType, "ddl": sql}
	bs, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req := http.NewRequest("POST", "http://sqlfiddle.com/backend/createSchema?_action=create", bytes.NewBuffer(bs))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res response
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
