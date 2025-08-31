package repository

import (
	"fmt"
	"golang-interview/database"
	"golang-interview/models"
	"log"
)

type RestaurantRepository struct{}

/*

ID    int     `json:"id"`
	ResId    int     `json:"resId"`
	MenuId    int     `json:"menuId"`
	UserId    int     `json:"userId"`
	Name  string  `json:"name"`
	Price  float64  `json:"price"`
*/
// Get all restaurant
// func (r *RestaurantRepository) GetAllPurchaseHistory(ctx context.Context,req models.PurchaseHistory) ([]models.PurchaseHistory, error) {
// 	uId := req.UserId
// 	rows, err := database.DB.QueryContext(ctx, `SELECT purchase_history.price, restaurant.name as restaurantname, restaurant_menu.name as menuname, users.name as username 
// 	FROM purchase_history LEFT JOIN restaurant ON restaurant.id = purchase_history.res_id
// 	LEFT JOIN restaurant_menu ON restaurant_menu.id = restaurant_menu.id
// 	LEFT JOIN users ON users.id = purchase_history.user_id
// 	WHERE purchase_history.user_id =? `,uId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var restaurant []models.PurchaseHistory
// 	for rows.Next() {
// 		var c models.PurchaseHistory
// 		if err := rows.Scan(&c.ID, &c.Price); err != nil {
// 			continue
// 		}
// 		restaurant = append(restaurant, c)
// 	}
// 	// log.Println(restaurant)
// 	return restaurant, nil
// }

func (r *RestaurantRepository) GetAllPurchaseHistory(req models.PurchaseHistory) ([]models.PurchaseHistory, error) {

    rows, err := database.DB.Query(`SELECT purchase_history.id,restaurant_menu.name as menuname, purchase_history.price, restaurant.name AS restaurantname, users.name AS username, users.balance
    FROM purchase_history
    LEFT JOIN restaurant ON restaurant.id = purchase_history.res_id
    LEFT JOIN restaurant_menu ON purchase_history.res_id = restaurant_menu.id
    LEFT JOIN users ON users.id = purchase_history.user_id
    WHERE purchase_history.user_id = ?`,req.UserId)
    if err != nil {
		
        return nil, err
		
    }
	//  fmt.Println( "alamin",rows)
    defer rows.Close()
// log.Println(rows)
    var purchaseHistory []models.PurchaseHistory
    for rows.Next() {
        var c models.PurchaseHistory
        // Adjust scan to include the restaurantName and userName
        if err := rows.Scan(&c.ID,&c.MenuName,&c.Price,&c.RestaurantName,&c.UserName,&c.UserBalance); err != nil {
			fmt.Println("scan error",err.Error())
            continue // Optionally log the error here for debugging
        }
		log.Println(c.UserBalance)
        // Here you would set additional fields if available  
        purchaseHistory = append(purchaseHistory, c)
    }
	fmt.Println("alamin", purchaseHistory)
    return purchaseHistory, nil
}




// Purchase Restaurant
func (r *RestaurantRepository) PurchaseOrder(c models.PurchaseHistory) (int64, error) {
	var arrayData models.PurchaseHistory
   if err  := database.DB.QueryRow(`SELECT balance FROM users WHERE id = ?`,c.UserId).Scan(&arrayData.UserBalance); err !=nil {
	return 0, err
   }
   
   currentBalance := arrayData.UserBalance  - c.Price
   _, err := database.DB.Exec("UPDATE users SET balance = ? WHERE id = ?", currentBalance, c.UserId)
	if err != nil {
		return 0, err
	}
	
// fmt.Println(arrayData.UserBalance)
   

	resesponse, error := database.DB.Exec("INSERT INTO purchase_history (res_id, menu_id,user_id,price) VALUES (?,?,?,?)", c.ResId, c.MenuId,c.UserId,c.Price)
	if error != nil {
		return 0, err
	}
	return resesponse.LastInsertId()
}
