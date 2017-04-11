package mymenchies_api

import (
    "time"
)


type UpdatedStruct struct{
    At          time.Time       `json:"at" dynamo:"at"`
    By          int             `json:"by" dynamo:"by"`
    Description string          `json:"description" dynamo:"description"`
} 