package main

import (

    api "github.com/jcgarciaram/general-api"
    "github.com/jcgarciaram/general-api/routes"
    mm "github.com/jcgarciaram/mymenchies/mymenchies_api"
    
)


func main() {
	r := routes.Routes{}
    
    // Append general_museum routes
    r.AppendRoutes(mm.GetRoutes())
    

    router := api.NewRouter(r)


    router.Listen()
    // router.Gateway()
    
}
