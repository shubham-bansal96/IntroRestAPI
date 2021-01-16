package model

import(
	"testing"
	"net/http"
)

func TestGetUserNotFound(t *testing.T){
	user, err:= GetUser(0)

	if user != nil {
		t.Error("we were not expecting a user with id 0")
	}

	if err==nil{
		t.Error("we were expecting an error when user id is 0")
	}

	if err.StatusCode != http.StatusNotFound{
		t.Error("we were expecting 404 when user is not found")
	}
}

// we have two returns for GetUer function in User_dao so we need to handle test cases for both returns
func TestGetUserNoError(t *testing.T){
	user, err:= GetUser(123)

	if user==nil{
		t.Error("We were expecting a user with id 123")
	}

	if err!=nil{
		t.Error("we were expecting an error when user id is 123")
	}
}