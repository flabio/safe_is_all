package utils

const (
	ID                    = "id"
	MESSAGE               = "message"
	STATUS                = "status"
	DATA                  = "data"
	CREATED               = "was successfully created"
	UPDATED               = "was updated successfully"
	REMOVED               = "was successfully removed"
	ERROR_QUERY           = "error query, please try again later"
	ERROR_CREATE          = "error creating"
	ERROR_PARSING_BODY    = "error parsing body"
	ERROR_UPDATE          = "error updating"
	ERROR_DELETE          = "error deleting"
	ERROR_REQUIRED_FIELDS = " is required."
	EMPTY                 = ""
)
const (
	DB_DIFF_ID        = "id <>?"
	DB_EQUAL_ID       = "id=?"
	DB_EQUAL_NAME     = "name =?"
	DB_ORDER_DESC     = "id desc"
	DB_EQUAL_EMAIL_ID = "email=?"
)

const (
	ID_NO_EXIST         = "The id not exists"
	NAME_ALREADY_EXIST  = "Name already exists"
	EMAIL_ALREADY_EXIST = "Email already exists"
)
const (
	NAME      = "name"
	SCHOOL_ID = "school_id"
	ACTIVE    = "active"
)
const (
	NAME_IS_FIELD_REQUIRED      = "The field name is required."
	SCHOOL_ID_IS_FIELD_REQUIRED = "The field school_id is required."
)

const (
	NAME_IS_REQUIRED      = "The name is required."
	SCHOOL_ID_IS_REQUIRED = "The id school is required."
)

const (
	AUTHORIZATION string = "Authorization"
	BEARER        string = "Bearer "
	TOKEN_INVALID string = "Token not provided or invalid"
)
