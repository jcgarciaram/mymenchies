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



// GetCakes retrieves all cakes for a single store
func GetCakes(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    // Get parameters from URL request
    storeId := params.Get("store")

    // Query to get all cakes
    query := "SELECT * from `cake` order by `pickup_timestamp` asc"
    
    // Run query from MySQL
    getTotalCount := false
    rowMapSlice, _, errStr, httpResponse := apiutils.RunSelectQuery(storeId, query, []interface{}{}, getTotalCount)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }

    // If no cakes are found, return StatusNoContent
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

// GetCake retrieves a particular cake from MySQL and returns all details
func GetCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {
    
    // Get parameters from URL request
    storeId := params.Get("store")
    cakeId := params.Get("cake")

    // Query to get cake
    query := "SELECT * from `%s`.`cake` where `cake_pk` = ?"
    
    // Run query from MySQL
    getTotalCount := false
    rowMapSlice, _, errStr, httpResponse := apiutils.RunSelectQuery(storeId, query, []interface{}{cakeId}, getTotalCount)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }

    // If no cakes are found, return StatusNoContent
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


// Post cake adds a new cake for a store
func PostCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

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
    var c Cake

    
    // Unmarshal body into Cake struct defined above
    if err := json.Unmarshal(bodyByte, &c); err != nil {
    
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error marshaling JSON to Cake struct")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusUnprocessableEntity
        res.Body = "Error marshaling JSON to Cake struct"
        return
    }
    
    
    
    // Query to insert cake
    query := fmt.Sprintf(
        "INSERT INTO `%s`.`cake` " + 
        "(`guest_name`," +
        "`guest_email`," +
        "`guest_phone`," +
        "`message`," +
        "`pickup_timestamp`," +
        "`cake_flavor`," +
        "`froyo_flavor_1`," +
        "`froyo_flavor_2`," +
        "`toppings`," +
        "`sauce`," +
        "`decorations`," +
        "`picture`," +
        "`design_number`," +
        "`last_updated_by`," +
        "`created_by`) " +
        "VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)", 
        
        storeId)
    
    parameters := []interface{}{
        c.GuestName,
        c.GuestEmail,
        c.GuestPhone,
        c.Message,
        c.PickupTimestamp.String(),
        c.CakeFlavor,
        c.FrozenYogurtFlavor1,
        c.FrozenYogurtFlavor2,
        c.Toppings,
        c.Sauce,
        c.Decorations,
        c.Picture,
        c.DesignNumber,
        UserEmail,
        UserEmail,
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


// Delete cake deletes a cake for a store
func DeleteCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Get parameters from URL request
    storeId := params.Get("store")
    cakeId := params.Get("cake")
    
    
    
    // Query to delete cake
    query := fmt.Sprintf("DELETE FROM `%s`.`cake` WHERE `cake_pk` = ?", storeId) 
    parameters := []interface{}{cakeId}
    
    // Build query to run in MySQL
    upsertQueries := []apiutils.UpsertQuery{
        {
            Query: query,
            Parameters: parameters,
        },

    }
    
    // Run queries
    getLastInsertId := false
    _, _, errStr, httpResponse := apiutils.RunUpsertQueries(upsertQueries, getLastInsertId)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }
    
    res.StatusCode = StatusOK

}


// Toggles prepared field between 0 and 1 for a cake
func PrepareCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Get parameters from URL request
    storeId := params.Get("store")
    cakeId := params.Get("cake")
    
    
    
    // Query to run
    query := fmt.Sprintf("UPDATE `%s`.`cake` SET `prepared` = CASE WHEN `prepared` = 0 THEN 1 ELSE 0 END, `prepared_timestamp` = CASE WHEN `prepared` = 1 THEN CURRENT_TIMESTAMP ELSE NULL END WHERE `cake_pk` = ?", storeId) 
    parameters := []interface{}{cakeId}
    
    // Build query to run in MySQL
    upsertQueries := []apiutils.UpsertQuery{
        {
            Query: query,
            Parameters: parameters,
        },

    }
    
    // Run queries
    getLastInsertId := false
    _, _, errStr, httpResponse := apiutils.RunUpsertQueries(upsertQueries, getLastInsertId)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }
    
    res.StatusCode = StatusOK

}


// Toggles decorate field between 0 and 1 for a cake
func DecorateCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Get parameters from URL request
    storeId := params.Get("store")
    cakeId := params.Get("cake")
    
    
    
    // Query to run
    query := fmt.Sprintf("UPDATE `%s`.`cake` SET `decorated` = CASE WHEN `decorated` = 0 THEN 1 ELSE 0 END, `decorated_timestamp` = CASE WHEN `decorated` = 1 THEN CURRENT_TIMESTAMP ELSE NULL END WHERE `cake_pk` = ?", storeId) 
    parameters := []interface{}{cakeId}
    
    // Build query to run in MySQL
    upsertQueries := []apiutils.UpsertQuery{
        {
            Query: query,
            Parameters: parameters,
        },

    }
    
    // Run queries
    getLastInsertId := false
    _, _, errStr, httpResponse := apiutils.RunUpsertQueries(upsertQueries, getLastInsertId)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }
    
    res.StatusCode = StatusOK

}

// Edit cake edits a specific cake
func EditCake(ctx *lambda.Context, evt *lambda.Event, res *lambda.ProxyResponse, params url.Values) {

    // Get parameters from URL request
    storeId := params.Get("store")
    cakeId  := params.Get("cake")
    
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
    var c Cake

    
    // Unmarshal body into Cake struct defined above
    if err := json.Unmarshal(bodyByte, &c); err != nil {
    
        logrus.WithFields(logrus.Fields{
            "err": err,
        }).Warn("Error marshaling JSON to Cake struct")   
        
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = StatusUnprocessableEntity
        res.Body = "Error marshaling JSON to Cake struct"
        return
    }
    
    
    
    // Query to insert cake
    query := fmt.Sprintf(
        "UPDATE `%s`.`cake` SET " + 
        "`guest_name` = ?," +
        "`guest_email` = ?," +
        "`guest_phone` = ?," +
        "`message` = ?," +
        "`pickup_timestamp` = ?," +
        "`cake_flavor` = ?," +
        "`froyo_flavor_1` = ?," +
        "`froyo_flavor_2` = ?," +
        "`toppings` = ?," +
        "`sauce` = ?," +
        "`decorations` = ?," +
        "`picture` = ?," +
        "`design_number` = ?," +
        "`last_updated_by` = ? " +
        "WHERE cake_pk = ?",
        
        storeId)
    
    parameters := []interface{}{
        c.GuestName,
        c.GuestEmail,
        c.GuestPhone,
        c.Message,
        c.PickupTimestamp.String(),
        c.CakeFlavor,
        c.FrozenYogurtFlavor1,
        c.FrozenYogurtFlavor2,
        c.Toppings,
        c.Sauce,
        c.Decorations,
        c.Picture,
        c.DesignNumber,
        UserEmail,
        cakeId,
    }
 
    
    // Build query to run in MySQL
    upsertQueries := []apiutils.UpsertQuery{
        {
            Query: query,
            Parameters: parameters,
        },

    }
    
    // Run queries
    getLastInsertId := false
    _, _, errStr, httpResponse := apiutils.RunUpsertQueries(upsertQueries, getLastInsertId)
    if httpResponse != 0 {
        res.Headers["Content-Type"] = "charset=UTF-8"
        res.StatusCode = strconv.Itoa(httpResponse)
        res.Body = errStr
        return
    }

    res.StatusCode = StatusOK
}

