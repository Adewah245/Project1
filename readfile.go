package main
import "os"

func readfile(s string)(string, error){
	data, error := os.ReadFile(s)
	if error != nil{
		return s, nil
	}
	return string(data), error
}
