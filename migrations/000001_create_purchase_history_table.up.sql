CREATE TABLE purchase_history (
id INT AUTO_INCREMENT PRIMARY KEY,
res_id INT,
menu_id int,
user_id int,
name VARCHAR(255) NULL,
price VARCHAR(255) NOT NULL DEFAULT "10.00",
-- FOREIGN KEY (res_id) REFERENCES restaurant(id)
-- FOREIGN KEY (menu_id) REFERENCES restaurant_menu(id)
-- FOREIGN KEY (user_id) REFERENCES users(id)
);