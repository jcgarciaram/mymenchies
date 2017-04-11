package mymenchies_api


type Employee struct {
    EmployeePk              int         `json:"employee_pk"`
    FirstName               string      `json:"first_name"`
    MiddleName              string      `json:"middle_name"`
    LastName                string      `json:"last_name"`
    SecondLastName          string      `json:"second_last_name"`
    DisplayName             string      `json:"display_name"`
    PhoneNumber             string      `json:"phone_number"`
    EmployeeRoleId          int         `json:"employee_role_id"`
    EmployeeRoleName        string      `json:"employee_role_name"`
    Wage                    float64     `json:"wage"`
    Active                  int         `json:"active"`
}