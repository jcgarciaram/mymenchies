package mymenchies_api

import (
    "github.com/jcgarciaram/general-api/apiutils"
    "github.com/tmaiaroto/aegis/lambda"
    "github.com/Sirupsen/logrus"
    
    "encoding/json"
    "strconv"
    "net/url"
    "fmt"
)



// PostEmployeeRole new employee role for a store
func PostEmployeeRole(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Get parameters from URL request
    storeId := params.Get("store")
    
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

    // Struct to unmarshal body of request
    var er EmployeeRole

    
    // Unmarshal body into Cake struct defined above
    if err := json.Unmarshal(bodyByte, &er); err != nil {
    
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error marshaling JSON to Cake struct")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusUnprocessableEntity
        res.Body = "Error marshaling JSON to EmployeeRole struct"
        return
    }
    
    
    
    // Query to insert cake
    query := fmt.Sprintf(
        "INSERT INTO `%s`.`employee_role` " + 
        "(`employee_role_name`," +
        "`wage`) " +
        "VALUES (?,?)", 
        
        storeId)
    
    parameters := []interface{}{
        er.EmployeeRoleName,
        er.Wage,
    }
 
    
    // Build query to run in MySQL
    upsertQueries := []apiutils.UpsertQuery{
        {
            Query: query,
            Parameters: parameters,
        },

    }
    
    // Run queries
    getLastInsertId := true
    lastInsertId, _, errStr, httpResponse := apiutils.RunUpsertQueries(upsertQueries, getLastInsertId)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }
    res.Body = strconv.Itoa(lastInsertId)
    res.StatusCode = StatusOK
}


// GetEmployeeRoles retrieves all employee roles for a single store
func GetEmployeeRoles(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    // Get parameters from URL request
    storeId := params.Get("store")

    // Query to get all cakes
    query := "SELECT * from `employee_role` order by `wage` asc"
    
    // Run query from MySQL
    getTotalCount := false
    rowMapSlice, _, errStr, httpResponse := apiutils.RunSelectQuery(storeId, query, []interface{}{}, getTotalCount)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }

    // If no employee-roles are found, return StatusNoContent
    if len(rowMapSlice) == 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusNoContent
        res.Body = ""
        return
    }

    // Marshal response and return
    retJson, _ := json.Marshal(rowMapSlice)
	res.Body = string(retJson)
    
	res.Headers["Content-Type"] = "application/json"
    
}
