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

// PostEmployee new employee for a store
func PostEmployee(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

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
    var e Employee

    
    // Unmarshal body into Cake struct defined above
    if err := json.Unmarshal(bodyByte, &e); err != nil {
    
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
        "INSERT INTO `%s`.`employee` " + 
        "(`first_name`," +
        "`middle_name`," +
        "`last_name`," +
        "`second_last_name`," +
        "`display_name`," +
        "`phone_number`," +
        "`employee_role_id`," +
        "`active`) " +
        "VALUES (?,?,?,?,?,?,?,?)", 
        
        storeId)
    
    parameters := []interface{}{
        e.FirstName,
        e.MiddleName,
        e.LastName,
        e.SecondLastName,
        e.DisplayName,
        e.PhoneNumber,
        e.EmployeeRoleId,
        1,
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


// GetEmployees retrieves all employees for a single store
func GetEmployees(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    // Get parameters from URL request
    storeId := params.Get("store")

    // Query to get all employees
    query := "SELECT " +
    "e.`employee_pk`," +
    "e.`first_name`," +
    "e.`middle_name`," +
    "e.`last_name`," +
    "e.`second_last_name`," +
    "e.`display_name`," +
    "e.`phone_number`," +
    "e.`employee_role_id`," +
    "er.`employee_role_name`," +
    "er.`wage`," +
    "e.`active` " +
    "from `employee` e " +
    "join `employee_role` er on e.`employee_role_id` = er.`employee_role_id` " +
    "where `active` = 1"
    
    // Run query from MySQL
    getTotalCount := false
    rowMapSlice, _, errStr, httpResponse := apiutils.RunSelectQuery(storeId, query, []interface{}{}, getTotalCount)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }

    // If no employees are found, return StatusNoContent
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