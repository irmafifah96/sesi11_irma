package models

type AddUserRequest struct {
	Id			string `json:id`
	Name		string `json:"name"`
	Age 		string `json:"age"`
	Email   	string `json:"email"`
	Password	string `json:"password"`
	Gender		string `json:"gender"`
	Role		string `json:"role"`
}
type AddUserResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
