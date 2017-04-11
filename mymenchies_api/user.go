package mymenchies_api

import (
    "time"
)

type User struct {
    UserEmail           string         `json:"user_email" dynamo:"user_email"`
    Password            string         `json:"password" dynamo:"password"`
    Store               string         `json:"store_id" dynamo:"store_id"`
    Role                string         `json:"role" dynamo:"role"`
    FirstName           string         `json:"first_name" dynamo:"first_name"`
    MiddleName          string         `json:"middle_name" dynamo:"middle_name"`
    LastName            string         `json:"last_name" dynamo:"last_name"`
    SecondLastName      string         `json:"second_last_name" dynamo:"second_last_name"`
    Created             time.Time       `json:"created" dynamo:"created"`
    Updated             []UpdatedStruct `json:"updated" dynamo:"updated"`
    LastActivity        time.Time       `json:"last_activity" dynamo:"last_activity"`
}

type UserGet struct {
    UserEmail           string         `json:"user_email" dynamo:"user_email"`
    Store               string         `json:"store_id" dynamo:"store_id"`
    Role                string         `json:"role" dynamo:"role"`
    FirstName           string         `json:"first_name" dynamo:"first_name"`
    MiddleName          string         `json:"middle_name" dynamo:"middle_name"`
    LastName            string         `json:"last_name" dynamo:"last_name"`
    SecondLastName      string         `json:"second_last_name" dynamo:"second_last_name"`
    Created             time.Time       `json:"created" dynamo:"created"`
    Updated             []UpdatedStruct `json:"updated" dynamo:"updated"`
    LastActivity        time.Time       `json:"last_activity" dynamo:"last_activity"`
}