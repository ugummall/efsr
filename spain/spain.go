package main

import (
	"fmt"
	"os"
	"database/sql"
    _ "github.com/go-sql-driver/mysql" // or the driver of your choice
    efsr_sqltocsv "github.com/ugummall/efsr/sqltocsv"
)

func SpainFunc() string {
	return "this is SpainFunc"
}

func main() {

	var file_name string
	var deal_number string

	if len(os.Args) > 1 {
		deal_number = os.Args[1]
		fmt.Println("deal-number is: ", deal_number)
		if len(os.Args) > 2 {
			file_name = os.Args[2]
			fmt.Println("file-name is: ", file_name)
		} else {
			fmt.Println("it seems file-name is missing.")
			fmt.Println("command: spain <deal-number> <file-name>")	
			return
		}
	} else {
		fmt.Println("it seems deal-number is missing.")
		fmt.Println("command: spain <deal-number> <file-name>")
		return
	}

	db, err := sql.Open("mysql", "efsrsys:Efsrsys3k#8@tcp(bannertest.effiser.net:3305)/banner")
    
    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }
    
    // defer the close till after the main function has finished
    // executing 
    defer db.Close()

    fmt.Println("successfully connected to mysql db.")

	// we're assuming you've setup your sql.DB etc elsewhere
	rows, _ := db.Query("SELECT 'SPA_JOB_INSTANCE_ID_101' JOB_INSTANCE_ID, 'EFSR_SPAS' JOB_INTG_NAME, CONCAT_WS('', sysdate(), ':', v.DH_DEAL_ID, ':', v.DL_LINE_ID) as UNIQUE_KEY, 'EFSR_DEV' SOURCE_SYSTEM, v.*, 'C' CREATED_BY, DATE_FORMAT(sysdate(), '%d-%b-%y') CREATION_DATE, 'S' INTERFACE_STATUS FROM ukr_spa_in_v v WHERE 1=1 AND v.dh_deal_number = ?", deal_number )

	arr := efsr_sqltocsv.WriteFile(file_name, rows)
	fmt.Println("generating the file with the filename ", file_name)
	if arr != nil {
		panic(err)
	}
}
