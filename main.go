package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Arguments map[string]string

var (
	id        string
	item      string
	operation string
	filename  string
)

type user struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Perform(args Arguments, writer io.Writer) error {

	//Handling Errors
	notValid := validate(args)
	if notValid != nil {
		return notValid
	}

	//if file exists, if not creates a new file

	//Logic part
	switch args["operation"] {
	case "list":
		file, err := os.OpenFile(args["fileName"], os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil {
			return err
		}
		dataBytes, err := ioutil.ReadAll(file)
		fmt.Println(string(dataBytes))
		if err != nil {
			return err
		}
		_, err = writer.Write(dataBytes)
		if err != nil {
			return err
		}
	case "add":

		//TestAddOperation
		var u []user
		var newUser user

		file, err := os.OpenFile(args["fileName"], os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil {
			return err
		}
		//Converting json in string to struct
		err = json.Unmarshal([]byte(args["item"]), &newUser)
		if err != nil {
			return err
		}

		//append newUser to slice of structs
		u = append(u, newUser)

		//convert struct {u} to json
		data, err := json.Marshal(u)
		if err != nil {
			return err
		}

		//Adding slice of json items to the file
		file.Write(data)

	case "findById":

		var u []user
		file, err := os.OpenFile(args["fileName"], os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		defer file.Close()
		if err != nil {
			return err
		}

		data, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}
		err = json.Unmarshal(data, &u)
		if err != nil {
			return err
		}
		for _, val := range u {
			fmt.Println("struct - ", val)
			if val.Id == args["id"] {
				data, err := json.Marshal(val)
				if err != nil {
					return err
				}
				file.Write(data)
			}
		}
		//file1, err := os.OpenFile(args["fileName"], os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
		//defer file.Close()
		//if err != nil {
		//	return err
		//}

	case "remove":
		fmt.Println("in process...")
	}

	//Logic part
	return nil
}

func parseArgs() Arguments {
	args := Arguments{
		"operation": operation,
		"item":      item,
		"fileName":  filename,
		"id":        id,
	}
	flag.StringVar(&operation, "operation", "", "Console App")
	flag.StringVar(&filename, "fileName", "", "flag for file name")
	flag.StringVar(&item, "item", "", "item from json")
	flag.StringVar(&id, "id", "", "item id")

	flag.Parse()
	return args
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func validate(args Arguments) error {
	if args["operation"] == "" && args["fileName"] == filename && args["item"] == "" && args["id"] == "" {
		return errors.New("-operation flag has to be specified")
	}
	if args["operation"] == "list" && args["fileName"] == "" && args["item"] == "" && args["id"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "abcd" && args["fileName"] == filename && args["item"] == "" && args["id"] == "" {
		return errors.New("Operation abcd not allowed!")
	}
	if args["operation"] == "add" && args["fileName"] == "" && args["item"] == "" && args["id"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "add" && args["fileName"] == filename && args["item"] == "" && args["id"] == "" {
		return errors.New("-item flag has to be specified")
	}
	if args["operation"] == "findById" && args["fileName"] == "" && args["item"] == "" && args["id"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "findById" && args["fileName"] == filename && args["item"] == "" && args["id"] == "" {
		return errors.New("-id flag has to be specified")
	}
	if args["operation"] == "remove" && args["fileName"] == "" && args["item"] == "" && args["id"] == "" {
		return errors.New("-fileName flag has to be specified")
	}
	if args["operation"] == "remove" && args["fileName"] == filename && args["item"] == "" && args["id"] == "" {
		return errors.New("-id flag has to be specified")
	}
	//AddSameIdError

	return nil
}

//package main
//
//import (
//	"encoding/json"
//	"errors"
//	"flag"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"log"
//	"os"
//)
//
//type Arguments map[string]string
//
//var (
//	id        string
//	operation string // add, list, findById, remove
//	item      string // ‘{«id»: "1", «email»: «email@test.com», «age»: 23}’
//	filename  string // users.json
//)
//
//func init() {
//	flag.StringVar(&operation, "operation", "", "user CRUD operation")
//	flag.StringVar(&item, "item", "", "user item")
//	flag.StringVar(&filename, "fileName", "", "file for storing user items")
//	flag.StringVar(&id, "id", "", "user id")
//
//	flag.Parse()
//}
//
//type user struct {
//	id    string `json: "id"`
//	email string `json: "email"`
//	age   int    `json: "age"`
//}
//
//func main() {
//	err := Perform(parseArgs(), os.Stdout)
//	if err != nil {
//		panic(err)
//	}
//}
//
//func Perform(args Arguments, writer io.Writer) error {
//
//	// validate input
//	validationErr := validate(args)
//	if validationErr != nil {
//		return validationErr
//	}
//
//	// If the file doesn't exist, create it, or append to the file
//	f, err := os.OpenFile(args["fileName"], os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// logic
//	switch args["operation"] {
//	case "list":
//		data, err := ioutil.ReadAll(f)
//		if err != nil {
//			return err
//		}
//		_, err = writer.Write(data)
//		if err != nil {
//			return err
//		}
//
//	case "add":
//		var users []user
//		var newUser user
//
//		data, err := ioutil.ReadAll(f)
//		err = json.Unmarshal(data, &users)
//		if err != nil {
//			return err
//		}
//
//		err = json.Unmarshal([]byte(args["item"]), &newUser)
//		if err != nil {
//			return err
//		}
//
//		users = append(users, newUser)
//		_ = f.Truncate(0)
//		_, _ = f.Seek(0, 0)
//		newData, _ := json.Marshal(users)
//		_, err = f.Write(newData)
//
//	case "remove":
//		var users []user
//
//		data, err := ioutil.ReadAll(f)
//		err = json.Unmarshal(data, &users)
//		if err != nil {
//			return err
//		}
//
//		for index, user := range users {
//			if user.id == args["id"] {
//				users = append(users[:index], users[index+1:]...)
//			}
//		}
//
//		_ = f.Truncate(0)
//		_, _ = f.Seek(0, 0)
//		newData, err := json.Marshal(users)
//		if err != nil {
//			return err
//		}
//
//		_, err = f.Write(newData)
//		if err != nil {
//			return err
//		}
//	}
//
//	//if _, err := f.Write([]byte("appended some data\n")); err != nil {
//	//  f.Close() // ignore error; Write error takes precedence
//	//  log.Fatal(err)
//	//}
//
//	// end of logic
//
//	if err := f.Close(); err != nil {
//		log.Fatal(err)
//	}
//
//	return nil
//}
//
//func parseArgs() Arguments {
//	var arguments = Arguments{
//		"operation": operation,
//		"item":      item,
//		"fileName":  filename,
//		"id":        id,
//	}
//	return arguments
//}
//
//func validate(args Arguments) error {
//
//	var validOperations = []string{"add", "list", "findById", "remove"}
//
//	if args["operation"] == "" {
//		return errors.New("-operation flag has to be specified")
//	}
//
//	if !contains(args["operation"], validOperations) {
//		return errors.New(fmt.Sprintf("Operation %v not allowed!", args["operation"]))
//	}
//
//	if args["fileName"] == "" {
//		return errors.New("-fileName flag has to be specified")
//	}
//
//	if args["operation"] == "add" && args["item"] == "" {
//		return errors.New("-item flag has to be specified")
//	}
//
//	if args["operation"] == "remove" && args["id"] == "" {
//		return errors.New("-id flag has to be specified")
//	}
//
//	if args["operation"] == "findById" && args["id"] == "" {
//		return errors.New("-id flag has to be specified")
//	}
//
//	return nil
//}
//
//func contains(op string, ops []string) bool {
//	for _, operation := range ops {
//		if operation == op {
//			return true
//		}
//	}
//	return false
//}
