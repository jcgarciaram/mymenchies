// API Route which will be prepended :) to all routes called
var api_route = "https://4l0rztm4pg.execute-api.us-east-1.amazonaws.com/prod/v1/api"

// While no real authentication is done, hard-code store_id
var store_id = "58d3da6503948f001e75d84c"


function addLoadEvent(func) {
    var oldonload = window.onload;
    if (typeof window.onload != 'function') {
        window.onload = func;
    } else {
        window.onload = function() {
            if (oldonload) {
                oldonload();
            }
            func();
        }
    }
}

addLoadEvent(function() {


    // While no cookie set up is yet implemented, allow all requests
    Vue.http.headers.common['Authorization'] = 'allow';
    
    
    // Vue component for the employee-table. This will contain all functions related to cakes
    new Vue({
        el: '#employee-table',
        delimiters: ['$$$', '$$$'],
        data: {
            employeeRoles: [],
            employees: [],
            numEmployees: 0,
            newEmployee: {},
        },

        // This is run whenever the page is loaded to make sure we have a current cake list
        created: function() {
        
            // Use the vue-resource $http client to fetch data from the /employee_roles route
            this.$http.get(api_route + '/stores/' + store_id + '/employee_roles').then(function(response) {
                this.employeeRoles = response.body ? response.body : [];
            })
        },
      
      
     
      
        // Methods for all API calls
        methods: {
        
        
            // Create new employee.
            createEmployee: function() {
        
                // Not sure what this does :). Copied and pasted from website. Please investigate
                if (!$.trim(this.newEmployee)) {
                    this.newEmployee = {};
                    return
                }

          
                // Post the new employee to the /employees route using the $http client
                this.$http.post(api_route + '/stores/' + store_id + '/employees', this.newEmployee).then((response) => {
        
                    // If API returns with OK status, add the employee to the cakes array
                    if (response.status == 200) {
                        console.log("employee has been added!");
                        
                        this.newEmployee.employee_pk = parseInt(response.body)
                        
                        this.newEmployee.active = 1;
                        this.employees = insertEmployee(this.newEmployee, this.employees);
                    }
                    
                    // Clear out newEmployee
                    this.newEmployee = {};
                    
                    // Update numEmployees
                    this.numEmployees = this.employees.length;
        
                // Investigate
                }, (response) => {
                
                    
                    console.log(response.status);
        
                });
          
          
                // Reset the addCakeForm to be blank
                var frm = document.getElementsByName('addEmployeeForm')[0];
                frm.reset();
            
            
            },
            
            
        }
    });

})

// insertEmployee inserts a new employee into the employees array into end of employees
function insertEmployee(element, array) {
    array.splice(array.length, 0, element);
    return array;
}

