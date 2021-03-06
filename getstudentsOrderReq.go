package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func studentsOrderReq(c *gin.Context) {

	// reqBody := OrderRequested{}

	// // reqBody.Approve_grant = "pending"

	// sqlStatatement := `SELECT "id", "book_id", "issue_date", "approve_grant", "order_id" FROM "public"."students_order_detail"
	// where "approve_grant" = 'pending';`

	// row := DB.QueryRow(sqlStatatement)

	// err := row.Scan(&reqBody.Id, &reqBody.Book_id, &reqBody.Issue_date, &reqBody.Approve_grant, &reqBody.Order_ID)

	// fmt.Println("reqBody.Order_ID of pending : ", reqBody.Order_ID)

	// if err != nil {
	// 	//log.Fatal("ppt", err)
	// 	res := gin.H{
	// 		"error":  "unable to view order list",
	// 		"status": "there is no any pending requests",
	// 	}
	// 	c.JSON(http.StatusBadRequest, res)
	// 	// c.Abort()
	// 	// return

	// } else {

	// 	res := gin.H{
	// 		"status":           "success",
	// 		"pending_order_id": reqBody.Order_ID,
	// 	}
	// 	c.JSON(http.StatusOK, res)

	// }

	var grant_status string

	grant_status = "pending"

	total_rows := count_rows(grant_status)

	if total_rows > 0 {

		users := []OrderApprove{}

		// var approved_orders_id int

		sqlStatatement1 := (`select order_id from students_order_detail where approve_grant = 'pending'`)

		rows, err := DB.Query(sqlStatatement1)

		if err != nil {
			log.Println("Failed to execute query: ", err)
			return
		}

		defer rows.Close()
		user := OrderApprove{}
		for rows.Next() {
			rows.Scan(&user.Order_ID)
			users = append(users, user)
		}

		res := gin.H{
			"pending status": users,
		}

		c.JSON(http.StatusOK, res)

		return

		// users := []User{}
		// userSQL := "SELECT id, name, email,user_id, phone, city,  password FROM users"

	}
	// return

	// sqlStatatement2 := `SELECT "order_id" FROM "public"."students_order_detail" where "approve_grant" = 'approved';`

	// row2 := DB.QueryRow(sqlStatatement2)

	// // fmt.Println("check getUserByEmail", reqBody)

	// err2 := row2.Scan(&reqBody.Order_ID)

	// fmt.Println("reqBody.Order_ID", reqBody.Order_ID)

	grant_status = "approved"

	total_rows1 := count_rows(grant_status)

	if total_rows1 > 0 {

		users := []OrderApprove{}

		// var approved_orders_id int

		sqlStatatement2 := (`select order_id from students_order_detail where approve_grant = 'approved'`)

		rows, err := DB.Query(sqlStatatement2)

		if err != nil {
			log.Println("Failed to execute query: ", err)
			return
		}

		defer rows.Close()
		user := OrderApprove{}
		for rows.Next() {
			rows.Scan(&user.Order_ID)
			users = append(users, user)
		}

		res := gin.H{
			"approved status": users,
		}

		c.JSON(http.StatusOK, res)

		return

		// users := []User{}
		// userSQL := "SELECT id, name, email,user_id, phone, city,  password FROM users"

	}

	// var all_order_id []int

	// all_order_id = append(all_order_id, reqBody.Order_ID)

	// // if err2 != nil {
	// // 	// log.Fatal(err2)
	// // 	res := gin.H{
	// // 		"still pending order id": all_order_id,
	// // 		"error":                  "no any approved order list",
	// // 	}
	// // 	c.JSON(http.StatusOK, res)
	// // } else {
	// res := gin.H{
	// 	"status":                 "success",
	// 	"approved order_id list": all_order_id,
	// }
	// c.JSON(http.StatusOK, res)
	// // }

	// return

}

func count_rows(grant_status string) int {
	var count int

	err := DB.QueryRow("SELECT COUNT(*) FROM students_order_detail where approve_grant = $1", grant_status).Scan(&count)
	switch {
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Println("Number of rows are", count)
	}
	return count
}
