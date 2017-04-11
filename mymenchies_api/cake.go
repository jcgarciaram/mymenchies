package mymenchies_api

import (
    "time"
    
    "github.com/jcgarciaram/general-api/apiutils"
)

type Cake struct {
    CakePk                  int                 `json:"cake_pk"`
    GuestName               string              `json:"guest_name"`
    GuestEmail              string              `json:"guest_email"`
    GuestPhone              string              `json:"guest_phone"`
    Message                 string              `json:"message"`
    PickupTimestamp         apiutils.CustomTime `json:"pickup_timestamp"`
    CakeFlavor              string              `json:"cake_flavor"`
    FrozenYogurtFlavor1     string              `json:"froyo_flavor_1"`
    FrozenYogurtFlavor2     string              `json:"froyo_flavor_2"`
    Toppings                string              `json:"toppings"`
    Sauce                   string              `json:"sauce"`
    Decorations             string              `json:"decorations"`
    Picture                 string              `json:"picture"`
    DesignNumber            string              `json:"design_number"`
    PreparedTimestamp       time.Time           `json:"prepared_timestamp"`
    Prepared                int                 `json:"prepared"`
    DecoratedTimestamp      time.Time           `json:"decorated_timestamp"`
    Decorated               int                 `json:"decorated"`
}