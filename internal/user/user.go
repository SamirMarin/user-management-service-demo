package user

import (
	"github.com/SamirMarin/user-management-service/internal/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
	awsDynamoDb "github.com/aws/aws-sdk-go/service/dynamodb"
	awsDynamoDbAttribute "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"time"
)

type User struct {
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Firstname  string     `json:"firstname"`
	Lastname   string     `json:"lastname"`
	Dob        time.Time  `json:"dob"`
	Membership Membership `json:"membership"`
}
type Membership struct {
	Kind    string    `json:"kind"`
	Owner   bool      `json:"owner"`
	Joined  time.Time `json:"joined"`
	Renewal time.Time `json:"renewal"`
}

var tableName = "User"

func (u *User) CreateUser() error {
	dynamoDbClient := dynamodb.NewClient(tableName)
	err := dynamoDbClient.StoreItem(u)
	if err != nil {
		return err
	}
	return nil
}
func (u *User) GetUser() error {
	dynamoDbClient := dynamodb.NewClient(tableName)
	err, getItemOutput := dynamoDbClient.GetItem(u)
	if err != nil {
		return err
	}
	err = awsDynamoDbAttribute.UnmarshalMap(getItemOutput.Item, u)
	return nil
}

func (u *User) ToDynamoDbAttribute() map[string]*awsDynamoDb.AttributeValue {
	return map[string]*awsDynamoDb.AttributeValue{
		"Username": {
			S: aws.String(u.Username),
		},
		"Email": {
			S: aws.String(u.Email),
		},
		"Firstname": {
			S: aws.String(u.Firstname),
		},
		"Lastname": {
			S: aws.String(u.Lastname),
		},
		"Dob": {
			S: aws.String(u.Dob.Format(time.RFC3339)),
		},
		"Membership": {
			M: map[string]*awsDynamoDb.AttributeValue{
				"Kind": {
					S: aws.String(u.Membership.Kind),
				},
				"Owner": {
					BOOL: aws.Bool(u.Membership.Owner),
				},
				"Joined": {
					S: aws.String(u.Membership.Joined.Format(time.RFC3339)),
				},
				"Renewal": {
					S: aws.String(u.Membership.Renewal.Format(time.RFC3339)),
				},
			},
		},
	}
}

func (u *User) ToDynamoDbItemInput() *awsDynamoDb.GetItemInput {
	return &awsDynamoDb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*awsDynamoDb.AttributeValue{
			"Username": {
				S: aws.String(u.Username),
			},
		},
	}
}
