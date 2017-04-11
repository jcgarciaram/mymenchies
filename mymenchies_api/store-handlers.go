package mymenchies_api

import (
    
    "github.com/jcgarciaram/general-api/apiutils"
    
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/tmaiaroto/aegis/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/Sirupsen/logrus"
    "github.com/guregu/dynamo"
    "gopkg.in/mgo.v2/bson"
    
    "encoding/json"
    "strconv"
    "net/url"
    "time"
    "log"
    "fmt"
)

// PostStore
func PostStore(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Read body from request
    var bodyByte []byte
    if tBody, err := apiutils.GetBodyFromEvent(evt); err != nil {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusInternalServerError
        res.Body = err.Error()
        return
    } else {
        bodyByte = tBody
    }
    
 
    // Struct to unmarshal body of request into
    var storeConfig Store
    

    // Unmarshal body into storeConfig struct defined above
    if err := json.Unmarshal(bodyByte, &storeConfig); err != nil {
    
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error marshaling JSON to storeConfig struct")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusUnprocessableEntity
        res.Body = "Error marshaling JSON to storeConfig struct"
        return
    }
    
    storeConfig.StoreId = bson.NewObjectId().Hex()
    
    // Created , LastActivity
    storeConfig.Created = time.Now()
    storeConfig.LastActivity = time.Now()
    
    
    // Save user Config to DynamoDB
    db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
    table := db.Table("mymenchies_stores")
    
    for {
        if err := table.Put(storeConfig).If("attribute_not_exists('store_id')").Run(); err != nil {
            
            // If there was an error saving store, verify if storeObjectId already exists.
            var tStore Store
            if getErr := table.Get("store_id", storeConfig.StoreId).One(&tStore); getErr == nil {
            
                storeConfig.StoreId = bson.NewObjectId().Hex()
                continue
            }
            
            // If the user did not exist, or there was an error retrieving the value from Dynamo, return original error
            logrus.WithFields(logrus.Fields{
                "err": err,
            }).Warn("Error creating store")   
            
            res.Headers["Content-Type"] = "charset=UTF-8"
            res.StatusCode = StatusInternalServerError
            res.Body = fmt.Sprintf("Error creating store: %s", err.Error())
            return

        }
        break
    }
    
    
    // Get all queries to create schema in MySQL database
    createSchemaQueries := getCreateSchemaQueries(storeConfig.StoreId)
    
    log.Println("About to run all create schema queries")
    
    // Run create schema / table queries
    getLastInsertId := false
    _, _, errStr, httpResponse := apiutils.RunUpsertQueries(createSchemaQueries, getLastInsertId)
    if httpResponse != 0 {
    
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = fmt.Sprintf(errStr)
        return
        
    }
    
    retJson, _ := json.Marshal(storeConfig)
	res.Body = string(retJson)

    res.Headers["Content-Type"] = "application/json"
    
    
}


// GetStore
func GetStore(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    storeId := params.Get("store")

    
    db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
    table := db.Table("mymenchies_stores")
    var store Store
    if err := table.Get("store_id", storeId).One(&store); err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error getting store")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusInternalServerError
        res.Body = fmt.Sprintf("Error getting store: %s", err.Error())
        return
    }

    retJson, _ := json.Marshal(store)
	res.Body = string(retJson)
    
	res.Headers["Content-Type"] = "application/json"
    
}