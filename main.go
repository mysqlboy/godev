package main 

// import "fmt"
 
import (
	"os"
	"fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
	// pass := "seven761"
	// user := "root"
	db, e := sql.Open("mysql", "root:seven761@tcp/")
	if( e != nil){
  		fmt.Print(e)
  		os.Exit(1)
 	}

	sqlcmd := "select id,time,info from information_schema.processlist where time > 10"
	st,err := db.Prepare(sqlcmd)
	if err != nil{
		fmt.Print(err);
		os.Exit(1)
	}

	rows, err := st.Query()
	if err != nil{
		fmt.Print(err);
		os.Exit(1)
	}

	i := 0
	for rows.Next(){
		i++
		var info string
		err = rows.Scan(&info)
		fmt.Printf("Query: %s \n", info)
	}
	db.Close()
	}
	//con.close()

