package main
/**
* This is the new app
**/
import (
    //"crypto"
    
    //"crypto/md5"
    //"hash"
    "fmt"
    "log"
    //"os"
    //"io"
    //"math/big"
    "database/sql"
    "net"
    //"net/http"
	_ "github.com/go-sql-driver/mysql"
	
)
var db *sql.DB;

func checkError(err error){
    if err!=nil {
        log.Fatal(err)
    }
}



func m2mServer(){
    service := "0.0.0.0:3412"
    tcpAddr, err := net.ResolveTCPAddr("tcp", service)
    checkError(err)
    listener, err := net.ListenTCP("tcp", tcpAddr)
    checkError(err)
    for {
        conn, _ := listener.Accept()
        go handleNewIOTConnection(conn)
    }
}
/**
* This is where the connection is negotiated - protocol based on TLS
* If client is unregistered
* 1 - client hello: [ 2 byte version | 4 bytes of dateTime | 28bytes of random data | Session ID]
* 2 - server hello: [ 4 byte of dateTime | 28 byte random data | session ID | certificate]
* 3 - client authentication: [ certifcate  ] --> Valid Certificate = pending authorization and secrecy
* If registered 
*/

func handleNewIOTConnection(conn net.Conn){
}


func main(){
    //Lets open our connection
    db, err := sql.Open("mysql","root:webyi2014@tcp(127.0.0.1:3306)/root_install")
	
	if err != nil {
	    log.Fatal(err)
	} 
	//Test if DB is connected
	err = db.Ping()
	if err != nil {
	    log.Fatal(err)
	}
	fmt.Printf("Connected to DB!!\n")
	//Lets define our routes
    //http.HandleFunc("/",)
	defer db.Close()
}