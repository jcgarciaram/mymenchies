package mymenchies_api

import (
    
    "github.com/jcgarciaram/general-api/apiutils"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/tmaiaroto/aegis/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "golang.org/x/crypto/bcrypt"
    "github.com/Sirupsen/logrus"
    "github.com/guregu/dynamo"
    
    "encoding/json"
    "net/url"
    "time"
    "fmt"
)



// PostUser
func PostUser(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    
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
    var userConfig User
    

    // Unmarshal body into userConfig struct defined above
    if err := json.Unmarshal(bodyByte, &userConfig); err != nil {
    
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error marshaling JSON to userConfig struct")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusUnprocessableEntity
        res.Body = "Error marshaling JSON to userConfig struct"
        return
    }

    // Hash password
    hashByte, err := bcrypt.GenerateFromPassword([]byte(userConfig.Password), 6)
    if err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error hashing password")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusInternalServerError
        res.Body = fmt.Sprintf("Error hashing password")
        return
    }
    
    // We will only store the hashed password
    userConfig.Password = string(hashByte)
    
    // Created , LastActivity
    userConfig.Created = time.Now()
    userConfig.LastActivity = time.Now()
    
    
    
    
    // Save user Config to DynamoDB
    db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
    table := db.Table("mymenchies_users")
    if err := table.Put(userConfig).If("attribute_not_exists('user_email')").Run(); err != nil {
        
        // If there was an error saving user, verify if user already exists.
        var tUser UserGet
        if getErr := table.Get("user_email", userConfig.UserEmail).One(&tUser); getErr == nil {
        
            res.Headers["Content-Type"] = "charset=UTF-8"
            res.StatusCode = StatusConflict
            res.Body = fmt.Sprintf("User already exists")
            return
            
        }
        
        // If the user did not exist, or there was an error retrieving the value from Dynamo, return original error
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error creating user")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusInternalServerError
        res.Body = fmt.Sprintf("Error creating user: %s", err.Error())
        return

    }

    res.StatusCode = StatusOK

}


// GetUser
func GetUser(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    userEmail := params.Get("useremail")
    
    db := dynamo.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
    table := db.Table("mymenchies_users")
    var user UserGet
    if err := table.Get("user_email", userEmail).One(&user); err != nil {
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error getting user")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusInternalServerError
        res.Body = fmt.Sprintf("Error getting user: %s", err.Error())
        return
    }

    retJson, _ := json.Marshal(user)
	res.Body = string(retJson)
    res.Headers["Content-Type"] = "application/json"
    
}

