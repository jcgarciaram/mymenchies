// API Route which will be prepended :) to all routes called
var api_route = "https://4l0rztm4pg.execute-api.us-east-1.amazonaws.com/prod/v1/api"

// While no real authentication is done, hard-code store_id
var store_id = "58d3da6503948f001e75d84c"


// Load when index.html is first open
window.onload = function () {

    // While no cookie set up is yet implemented, allow all requests
    Vue.http.headers.common['Authorization'] = 'allow';
    
    // Custom component for datepicker.
    Vue.component('datepicker', {
    template: '\
        <div class="input-group date"><input class="form-control datetimepicker"\
            ref="input"\
            v-bind:value="value"\
            data-date-format="LLL"\
            data-date-end-date="0d"\
            placeholder=""\
            type="text"  />\
            <span class="input-group-addon">\
              <span class="glyphicon glyphicon-calendar"></span>\
            </span></div>\
    ',
    props: {
        value: {
            type: String,
            default: moment().format('LLL')
        }
    },
    mounted: function() {
        let self = this;
        this.$nextTick(function() {
            $(this.$el).datetimepicker().on('dp.change', function(e) {
                //this.value = moment(e.date).format('MMMM Do, YYYY h:mm a')
                self.updateValue(moment(e.date).format('LLL'));
            });
        });
    },
    methods: {
        updateValue: function (value) {
            
            this.$emit('input', value);
        },
    }
    });
    
    
    // Vue component for the cake-table. This will contain all functions related to cakes
    new Vue({
        el: '#cake-table',
        delimiters: ['$$$', '$$$'],
        data: {
            cakes: [],
            lenCakes: 0,
            newCake: {},
            editCake: {},
            editCakeIndex: null,
        },

        // This is run whenever the page is loaded to make sure we have a current cake list
        created: function() {
      
            // Use the vue-resource $http client to fetch data from the /cakes route
            this.$http.get(api_route + '/stores/' + store_id + '/cakes').then(function(response) {
                this.cakes = response.body ? response.body : [];
                this.lenCakes = this.cakes.length;
            })

        },
      
      
     
      
        // Methods for all API calls
        methods: {
        
        
            // Create new cake.
            createCake: function() {
        
                // Not sure what this does :). Copied and pasted from website. Please investigate
                if (!$.trim(this.newCake)) {
                    this.newCake = {};
                    return
                }

          
                // Post the new cake to the /cakes route using the $http client
                this.$http.post(api_route + '/stores/' + store_id + '/cakes', this.newCake).then((response) => {
        
                    // If API returns with OK status, add the cake to the cakes array
                    if (response.status == 200) {
                        console.log("cake has been added!");
                        
                        this.newCake.cake_pk = parseInt(response.body)
                        
                        this.newCake.prepared = 0;
                        this.newCake.decorated = 0;
                        this.cakes = insertCake(this.newCake, this.cakes);
                    }
                    
                    // Clear out newCake
                    this.newCake = {};
                    
                    // Update lenCakes
                    this.lenCakes = this.cakes.length;
        
                // Investigate
                }, (response) => {
                
                    
                    console.log(response.status);
        
                });
          
          
                // Reset the addCakeForm to be blank
                var frm = document.getElementsByName('addCakeForm')[0];
                frm.reset();
            
            
            },
            
            // Function to edit an already created cake
            editCakeFunction: function() {
                
                if (!this.editCake.prepared_timestamp) {
                    delete this.editCake.prepared_timestamp;
                }
                
                if (!this.editCake.decorated_timestamp) {
                    delete this.editCake.decorated_timestamp;
                }
                
                // Make API Call
                this.$http.put(api_route + '/stores/' + store_id + '/cakes/' + this.editCake.cake_pk, this.editCake).then((response) => {
                    
                    
                    
                    
                    
                    // cake has been updated!!
                    if (response.status == 200) {
                        console.log("cake has been edited!");
                        
                        // Update cake in local memory
                        this.cakes[this.editCakeIndex] = JSON.parse(JSON.stringify(this.editCake))
                        
                        this.editCake = {}
                        this.editCakeIndex = null
                        
                    }
                    
                // Not sure what this is doing :). Research please
                }, (response) => {
                    console.log(response.status);
                    
                });
        
            },
            
            // Function to delete a cake
            deleteCake: function(cakeid, index) {
            
                // Make API call
                this.$http.delete(api_route + '/stores/' + store_id + '/cakes/' + cakeid).then((response) => {
                    
                    // If respnse is OK
                    if (response.status == 200) {
                        console.log("cake ", cakeid, " deleted! index:", index);
                        this.cakes.splice(index, 1);
                        this.lenCakes = this.cakes.length;
                    }
                
                // Not sure what this is doing :). Research please 
                }, (response) => {
                
                    console.log(response);
                    
                })
            },
            
            
            // Function to mark a cake as prepared
            preparedCake: function(cakeid) {
            
                // Function to make API call
                this.$http.put(api_route + '/stores/' + store_id + '/cakes/' + cakeid + '/prepare').then((response) => {
                    
                    // If call is successfull
                    if (response.status == 200) {
                        console.log("cake has been prepared!");
                    }
                    
                }, (response) => {
                
                    console.log(response);
                    
                })
            
            },
        
            // Function to mark a cake as decorated
            decoratedCake: function(cakeid) {
            
                // Make API request
                this.$http.put(api_route + '/stores/' + store_id + '/cakes/' + cakeid + '/decorate').then((response) => {
                    
                    if (response.status == 200) {
                        console.log("cake has been decorated!");
                    }
                    
                    
                }, (response) => {
                
                    console.log(response)
                    
                })
            
            },
        
            // Function used to assign local variable editCake to the cake for which the edit button was just pushed. This will enable us to pre-populate the fields in the editCake Modal
            selectEdit: function(index) {
                this.editCakeIndex = index
                this.editCake = JSON.parse(JSON.stringify(this.cakes[index]))
    
            },
            
            // Function to clear the editCake variable in case Cancel is clicked in the editCake Modal
            dismissEditCake: function() {
                
                this.editCakeIndex = null
                this.editCake = {}
    
            },
        }
    });

}

// insertCake inserts a new cake into the cakes array into a correct sorted position based on the pickup_timestamp
function insertCake(element, array) {
    array.splice(locationOfCake(element, array), 0, element);
    return array;
}

// locationOfCake figures out the correct location of where to insert the cake based on the pickup_timestamp
function locationOfCake(element, array) {
    
    // If array is empty
    if (array.length == 0) {
        return 0
    } 
    
    // Define pivot as middle point
    var pivot = parseInt(array.length / 2);
    
    var elPickupTimestamp = moment(element.pickup_timestamp, "LLL")
    var pivotPickupTimestamp = moment(array[pivot].pickup_timestamp, "LLL")
    
    // If element is the same as the middle, insert right after
    if (pivotPickupTimestamp.isSame(elPickupTimestamp)) {
        return pivot+1;
    }
    
    // If element is less than middle element
    if (elPickupTimestamp.isBefore(pivotPickupTimestamp)) {
    
        return locationOfCake(element, array.slice(0, pivot));
    
    // If element is greater than middle element
    } else {
        return (pivot+1) + locationOfCake(element, array.slice(pivot+1, array.length));
    }
}


