SELECT 
    c.id, 
    c.name, 
    (SELECT COUNT(*) FROM products p WHERE p.category_id = c.id) AS product_count,
    (SELECT COUNT(*) FROM orders o WHERE o.category_id = c.id) AS order_count,
    (SELECT COUNT(*) FROM vendors v WHERE v.category_id = c.id) AS vendor_count
FROM categories c;

-----------------
SELECT 
    c.id, 
    c.name, 
    COUNT(DISTINCT p.id) AS product_count,
    COUNT(DISTINCT o.id) AS order_count,
    COUNT(DISTINCT v.id) AS vendor_count
FROM categories c
LEFT JOIN products p ON p.category_id = c.id
LEFT JOIN orders o ON o.category_id = c.id
LEFT JOIN vendors v ON v.category_id = c.id
GROUP BY c.id, c.name;

---------------------


func (r *CategoryRepository) GetAllWithProducts(ctx context.Context) ([]models.CategoryWithProducts, error) {
    query := `
        SELECT 
            c.id, c.name,
            p.id AS product_id, p.name AS product_name,
            s.name AS supplier_name
        FROM categories c
        LEFT JOIN products p ON c.id = p.category_id
        LEFT JOIN suppliers s ON p.supplier_id = s.id;
    `
    rows, err := database.DB.QueryContext(ctx, query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var result []models.CategoryWithProducts
    for rows.Next() {
        var c models.CategoryWithProducts
        if err := rows.Scan(&c.ID, &c.Name, &c.ProductID, &c.ProductName, &c.SupplierName); err != nil {
            return nil, err
        }
        result = append(result, c)
    }
    return result, nil
}


---------------------------

package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"golang-crud/models"
)

type ShopRepository struct {
	DB *sql.DB
}

func (r *ShopRepository) GetShopFullData(ctx context.Context, shopID int64) (*models.ShopFullData, error) {
	query := `
	SELECT 
	    s.id AS shopId,
	    (SELECT JSON_ARRAYAGG(JSON_OBJECT(
	        'id', p.id,
	        'name', p.name,
	        'sku', p.sku,
	        'image', p.image,
	        'uomStr', p.uom,
	        'basePrice', p.base_price,
	        'discount', p.discount,
	        'maxQtyPerOrder', p.max_qty,
	        'discountedPrice', p.base_price - p.discount,
	        'variationCount', (SELECT COUNT(*) FROM variations v WHERE v.product_id = p.id),
	        'variations', (SELECT JSON_ARRAYAGG(JSON_OBJECT(
	            'id', v.id,
	            'name', v.name,
	            'sku', v.sku,
	            'image', v.image,
	            'basePrice', v.base_price,
	            'discount', v.discount,
	            'maxQtyPerOrder', v.max_qty,
	            'discountedPrice', v.base_price - v.discount,
	            'discountInPerc', TRUE
	        )) FROM variations v WHERE v.product_id = p.id),
	        'discountInPerc', TRUE,
	        'branchId', p.branch_id
	    )) FROM products p WHERE p.shop_id = s.id) AS products,
	    
	    (SELECT JSON_ARRAYAGG(JSON_OBJECT(
	        'id', c.id,
	        'name', c.name,
	        'tagIds', (SELECT JSON_ARRAYAGG(ct.tag_id) FROM category_tags ct WHERE ct.category_id = c.id),
	        'image', c.image
	    )) FROM categories c WHERE c.shop_id = s.id) AS categories,
	    
	    (SELECT JSON_ARRAYAGG(JSON_OBJECT(
	        'id', ch.id,
	        'type', ch.type,
	        'subType', ch.sub_type,
	        'title', ch.title,
	        'displayType', JSON_ARRAY(ch.display_type),
	        'hasOptions', ch.has_options,
	        'value', ch.value,
	        'requestParamName', ch.request_param_name,
	        'requestParamDataType', ch.request_param_data_type,
	        'options', (SELECT JSON_ARRAYAGG(JSON_OBJECT(
	            'id', o.id,
	            'name', o.name,
	            'requestParamName', o.request_param_name,
	            'isSelected', o.is_selected
	        )) FROM chip_options o WHERE o.chip_id = ch.id),
	        'range', JSON_ARRAY(ch.range_min, ch.range_max)
	    )) FROM chips ch WHERE ch.shop_id = s.id) AS chipList
	FROM shops s
	WHERE s.id = ?;
	`

	var jsonResult string
	err := r.DB.QueryRowContext(ctx, query, shopID).Scan(&jsonResult)
	if err != nil {
		return nil, err
	}

	var shopData models.ShopFullData
	if err := json.Unmarshal([]byte(jsonResult), &shopData); err != nil {
		return nil, err
	}

	return &shopData, nil
}
