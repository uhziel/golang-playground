package main

import "log"

func main() {
	log.Println("-- log const --")
	log.Println("log.Ldate: ", log.Ldate)
	log.Println("log.Ltime: ", log.Ltime)
	log.Println("log.Lmicroseconds: ", log.Lmicroseconds)
	log.Println("log.Llongfile: ", log.Llongfile)
	log.Println("log.Lshortfile: ", log.Lshortfile)
	log.Println("log.LUTC: ", log.LUTC)
	log.Println("log.Lmsgprefix: ", log.Lmsgprefix)
	log.Println("log.LstdFlags: ", log.LstdFlags) // 它是默认值 Ldate | Ltime
	log.Println()

	log.Println("Flags:", log.Flags())
	log.Println("Prefix:", log.Prefix())
	log.Println("log1")
	log.Printf("%s", "log2")
	log.Println()

	log.Println("-- log prefix --")
	log.SetPrefix("zhulei ")
	log.Println("log3")
	log.Println()

	log.Println("-- log flags --")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile | log.Lmsgprefix)
	log.Println("log4")
}
