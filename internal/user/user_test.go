package user

import (
	"github.com/aws/aws-sdk-go/aws"
	awsDynamoDb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestToDynamoDbAttribute(t *testing.T) {
	user := &User{
		Username:  "testUser",
		Email:     "testEmail",
		Firstname: "testFirstname",
		Lastname:  "testLastname",
		Dob:       time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC),
		Membership: Membership{
			Kind:    "testKind",
			Owner:   true,
			Joined:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			Renewal: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	dynamodbAttribute := user.ToDynamoDbAttribute()
	expectedDynamodbAttribute := map[string]*awsDynamoDb.AttributeValue{
		"Username": {
			S: &user.Username,
		},
		"Email": {
			S: &user.Email,
		},
		"Firstname": {
			S: &user.Firstname,
		},
		"Lastname": {
			S: &user.Lastname,
		},
		"Dob": {
			S: aws.String(user.Dob.Format(time.RFC3339)),
		},
		"Membership": {
			M: map[string]*awsDynamoDb.AttributeValue{
				"Kind": {
					S: &user.Membership.Kind,
				},
				"Owner": {
					BOOL: &user.Membership.Owner,
				},
				"Joined": {
					S: aws.String(user.Membership.Joined.Format(time.RFC3339)),
				},
				"Renewal": {
					S: aws.String(user.Membership.Renewal.Format(time.RFC3339)),
				},
			},
		},
	}
	assert.Equal(t, expectedDynamodbAttribute, dynamodbAttribute)
}

func TestToDynamoDbItemInput(t *testing.T) {
	user := &User{
		Username: "testUser",
	}
	dynamodbItemInput := user.ToDynamoDbItemInput()
	expectedDynamodbItemInput := &awsDynamoDb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*awsDynamoDb.AttributeValue{
			"Username": {
				S: &user.Username,
			},
		},
	}
	assert.Equal(t, expectedDynamodbItemInput, dynamodbItemInput)
}

func testCreateGetUser(t *testing.T) {
	creatUser := &User{
		Username:  "testUser",
		Email:     "testEmail",
		Firstname: "testFirstname",
		Lastname:  "testLastname",
		Dob:       time.Date(1991, time.January, 1, 0, 0, 0, 0, time.UTC),
		Membership: Membership{
			Kind:    "testKind",
			Owner:   true,
			Joined:  time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC),
			Renewal: time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	err := creatUser.CreateUser()
	assert.Nil(t, err)
	getUser := &User{
		Username: "testUser",
	}
	err = getUser.GetUser()
	assert.Nil(t, err)
	assert.Equal(t, creatUser, getUser)
}
