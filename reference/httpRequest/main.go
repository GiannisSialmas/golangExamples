package main

import (
	"net/http"
)

func main() {

	resp, err := http.Get("http://gobyexample.com")
	// if err != nil {
	// 	panic(err)
	// }
	// defer resp.Body.Close()
	// fmt.Println("Response status:", resp.Status)

	// scanner := bufio.NewScanner(resp.Body)
	// for i := 0; scanner.Scan() && i < 5; i++ {
	//     fmt.Println(scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	//     panic(err)
	// }

	// OR
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// log.Println(string(body))

}
