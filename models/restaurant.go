package models

type Restaurant struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Balance float64 `json:"balance"`
}

type RestaurantMenu struct {
	ID    int     `json:"id"`
	ResId    int     `json:"resId"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type User struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Email  string  `json:"email"`
}

type PurchaseHistory struct {
	ID    int     `json:"id" db:"id"`
	ResId    int     `json:"resId" db:"res_id"`
	MenuId    int     `json:"menuId" db:"menu_id"`
	UserId    int     `json:"userId" db:"user_id"`
	Name  string  `json:"name" db:"name"`
	Price  float64  `json:"price" db:"price"`
	RestaurantName string `json:"restaurantName" db:"resturantname"`
    UserName  string  `json:"userName" db:"username"`
    UserBalance  float64  `json:"userBalance" db:"balance"`
	MenuName string `json:"menuName" db:"menuname"`
}




