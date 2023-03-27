package api

import (
	"encoding/json"
	"net/http"

	"github.com/kurakura967/go-openapi-demo/api/generated/openapi"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) GetUserById(w http.ResponseWriter, r *http.Request, userId openapi.UserId) {
	var res openapi.User
	var age int64 = 10
	var gender string = "1"
	var id int64 = userId
	var name string = "taro"

	res = openapi.User{
		Age:    &age,
		Gender: &gender,
		Id:     &id,
		Name:   &name,
	}

	w.Header().Add("Content-Type", "application/json")
	var err error
	err = json.NewEncoder(w).Encode(res);
	if err != nil {
		WriteInternalErrorResponse(w, r, err)
	}
}

func (s *Server) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	var res []openapi.User

	var age int64 = 10
	var id1 int64 = 1
	var id2 int64 = 2

	res = []openapi.User{
		{
			Age:    &age,
			Gender: ToPtr("1"),
			Id:     &id1,
			Name:   ToPtr("taro"),
		},
		{
			Age:    &age,
			Gender: ToPtr("2"),
			Id:     &id2,
			Name:   ToPtr("teru"),
		},
	}

	w.Header().Add("Content-Type", "application/json")
	var err error
	err = json.NewEncoder(w).Encode(res);
	if err != nil {
		WriteInternalErrorResponse(w, r, err)
	}
}
