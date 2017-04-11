package mymenchies_api


type EmployeeRole struct {
    EmployeeRoleId          int         `json:"employee_role_id"`
    EmployeeRoleName        string      `json:"employee_role_name"`
    Wage                    float64     `json:"wage"`
}