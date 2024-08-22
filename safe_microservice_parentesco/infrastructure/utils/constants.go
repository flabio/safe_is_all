package utils

const (
	ID                 string = "id"
	MESSAGE            string = "message"
	STATUS             string = "status"
	DATA               string = "data"
	USER               string = "user"
	CREATED            string = "was successfully created"
	UPDATED            string = "was updated successfully"
	REMOVED            string = "was successfully removed"
	ERROR_QUERY        string = "error query, please try again later"
	ERROR_CREATE       string = "error creating"
	ERROR_PARSING_BODY string = "error parsing body"
	ERROR_UPDATE       string = "error updating"
	ERROR_DELETE       string = "error deleting"
)
const (
	DB_DIFF_ID       string = "id <>?"
	DB_EQUAL_ID      string = "id=?"
	DB_EQUAL_USER_ID string = "user_id =?"
	DB_NAME          string = "name =?"
	DB_ORDER_DESC    string = "id desc"
	DB_TABLE_NAME    string = "Parentesco"
)

const (
	ID_NO_EXIST        string = "The id not exists"
	NAME_ALREADY_EXIST string = "name already exists"
)
const (
	AUTHORIZATION string = "Authorization"
	BEARER        string = "Bearer "
	TOKEN_INVALID string = "Token not provided or invalid"
)

const (
	EMPTY  = ""
	NAME   = "name"
	ACTIVE = "active"
)
const (
	NAME_IS_REQUIRED         = "The name is required."
	NAME_FIEDL_IS_REQUIRED   = "The field name is required."
	ACTIVE_FIELD_IS_REQUIRED = "The field active is required."
)
