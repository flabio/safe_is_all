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
	FIRST_NAME    string = "first_name"
	LAST_NAME     string = "last_name"
	PHONE         string = "phone"
	USER_ID       string = "user_id"
	PARENTESCO_ID string = "parentesco_id"
	ACTIVE        string = "active"
)
const (
	EMPTY                                  = ""
	FIRST_NAME_FIELD_IS_REQUIRED    string = "The first name field is required"
	LAST_NAME_FIELD_IS_REQUIRED     string = "The last name field is required"
	PHONE_FIELD_IS_REQUIRED         string = "The phone field is required"
	USER_ID_FIELD_IS_REQUIRED       string = "The user id field is required"
	PARENTESCO_ID_FIELD_IS_REQUIRED string = "The parentesco id field is required"
	ACTIVE_FIELD_IS_REQUIRED        string = "The active field is required"
)
const (
	FIRST_NAME_IS_REQUIRED    string = "The first name is required"
	LAST_NAME_IS_REQUIRED     string = "The last name is required"
	PHONE_IS_REQUIRED         string = "The phone is required"
	USER_ID_IS_REQUIRED       string = "The user ID is required"
	PARENTESCO_ID_IS_REQUIRED string = "The parentest is required"
	ACTIVE_IS_REQUIRED        string = "The active is required"
)
