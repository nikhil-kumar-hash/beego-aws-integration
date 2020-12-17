package main

import (
	"context"
	"fmt"

	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type user_request struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "admin:Nikhilo1!@(testdb.c6fq0vlfh1fb.us-east-2.rds.amazonaws.com:3306)/testdb")
}

func insert(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	var user_req user_request
	json.Unmarshal([]byte(request.Body), &user_req)
	fmt.Print(request.Body)
	fmt.Print(request)
	o := orm.NewOrm()

	user := new(User)

	user.Name = user_req.Username
	user.Password = user_req.Password
	fmt.Println(user_req.Username)
	fmt.Println(user_req.Password)
	fmt.Println(o.Insert(user))
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}
