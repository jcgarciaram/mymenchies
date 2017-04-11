package mymenchies_api

import (
    "github.com/jcgarciaram/general-api/apiutils"
    
    "time"
    "fmt"
)

type Store struct {
    StoreId         string             `json:"store_id" dynamo:"store_id"`
    StoreName       string             `json:"store_name" dynamo:"store_name"`
    StoreEmail      string             `json:"store_email" dynamo:"store_email"`
    ContactUser     string             `json:"contact_user" dynamo:"contact_user"`
    Created         time.Time       `json:"created" dynamo:"created"`              
    Updated         []UpdatedStruct `json:"updated" bson:"updated"`
    LastActivity     time.Time       `json:"last_activity" dynamo:"last_activity"`
}



// getCreateSchemaQueries returns all the queries for a 
func getCreateSchemaQueries(storeId string) []apiutils.UpsertQuery {

    // Build query to run in MySQL
    createQueries := []apiutils.UpsertQuery{
        {
            Query: fmt.Sprintf("DROP SCHEMA IF EXISTS `%s`", storeId),
            Parameters: []interface{}{},
        },
        {
            Query: fmt.Sprintf("CREATE SCHEMA `%s`", storeId),
            Parameters: []interface{}{},
        },
        {
            Query: fmt.Sprintf("GRANT SELECT ON `%s`.* TO 'mymenchiesread'", storeId),
            Parameters: []interface{}{},
        },
        {
            // cake
            Query: fmt.Sprintf(
            "CREATE TABLE `%s`.`cake` " +
            "(`cake_pk` INT NOT NULL AUTO_INCREMENT COMMENT ''," +
            "`guest_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`guest_email` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`guest_phone` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`message` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`pickup_date` TIMESTAMP NOT NULL COMMENT ''," +
            "`cake_flavor` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`froyo_flavor_1` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`froyo_flavor_2` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`toppings` VARCHAR(200) DEFAULT NULL COMMENT ''," +
            "`sauce` VARCHAR(200) DEFAULT NULL COMMENT ''," +
            "`decorations` VARCHAR(200) DEFAULT NULL COMMENT ''," +
            "`picture` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`design_number` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`prepared_timestamp` TIMESTAMP DEFAULT NULL COMMENT ''," +
            "`prepared` TINYINT(1) DEFAULT 0 COMMENT ''," +
            "`decorated_timestamp` TIMESTAMP DEFAULT NULL COMMENT ''," +
            "`decorated` TINYINT(1) DEFAULT 0 COMMENT ''," +
            "`last_updated_by` VARCHAR(45) NOT NULL COMMENT ''," +
            "`last_updated_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT ''," +
            "`created_by` VARCHAR(45) NOT NULL COMMENT ''," +
            "`created_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT ''," +
            "PRIMARY KEY (`cake_pk`)  COMMENT '')",
            
            storeId),
            
            Parameters: []interface{}{},
        },
        
        {
            // employee
            Query: fmt.Sprintf(
            "CREATE TABLE `%s`.`employee` " +
            "(`employee_pk` INT NOT NULL AUTO_INCREMENT COMMENT ''," +
            "`first_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`middle_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`last_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`second_last_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`display_name` VARCHAR(45) NOT NULL COMMENT ''," +
            "`phone_number` VARCHAR(45) NOT NULL COMMENT ''," +
            "`employee_role_id` INT NOT NULL COMMENT ''," +
            "`active` TINYINT(1) NOT NULL COMMENT ''," +
            "`last_updated_by` VARCHAR(45) NOT NULL COMMENT ''," +
            "`last_updated_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT ''," +
            "`created_by` VARCHAR(45) NOT NULL COMMENT ''," +
            "`created_timestamp` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT ''," +
            "PRIMARY KEY (employee_pk)  COMMENT '')",
            
            storeId),
            
            Parameters: []interface{}{},
        },
        
        {
            // role
            Query: fmt.Sprintf(
            "CREATE TABLE `%s`.`employee_role` " +
            "(`employee_role_id` INT NOT NULL AUTO_INCREMENT COMMENT ''," +
            "`employee_role_name` VARCHAR(45) DEFAULT NULL COMMENT ''," +
            "`wage` DECIMAL DEFAULT NULL COMMENT ''," +
            "PRIMARY KEY (employee_role_id)  COMMENT '')",
            
            storeId),
            
            Parameters: []interface{}{},
        },
    }
    
    return createQueries
}





