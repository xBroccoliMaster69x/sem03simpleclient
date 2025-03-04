package main

import (
	"net"
	"log"
	"os"
	"github.com/xBroccoliMaster69x/is105sem03/mycrypt"
)

func main() {
	conn, err := net.Dial("tcp", "172.17.0.4:8000")
	if err != nil {
		log.Fatal(err)
	}
    
	log.Println("os.Args[1] = ", os.Args[1])

	kryptertMelding := mycrypt.Krypter([]rune(os.Args[1]), mycrypt.ALF_SEM03, 4)
	log.Println("kryptert melding: ", string(kryptertMelding))
	_,err = conn.Write([]byte(string(kryptertMelding)))



	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	} 
	dekryptertMelding := mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
	log.Printf("reply from proxy: %s", string(dekryptertMelding))

}
