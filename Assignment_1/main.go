package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
}

var users []User
var warehouse []Item

type Item struct {
	name    string
	price   int
	id      int
	ratings []int
}

func serialize() {
	byteArray, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		fmt.Println("sdfs")
	}
	err_ := ioutil.WriteFile("test.json", byteArray, 0644)
	if err_ != nil {
		return
	}
}

func deserialize() {
	file, _ := ioutil.ReadFile("test.json")
	_ = json.Unmarshal([]byte(file), &users)
}
func (u User) SignUp() {
	for {
		fmt.Print("Enter an username.\n")
		log := ""
		fmt.Fscan(os.Stdin, &log)
		ok := true
		for _, us := range users {
			if log == us.Login {
				fmt.Print("This username already exist.\n")
				ok = false
				break
			}
		}
		if ok == true {
			if len(log) >= 1 {
				u.Login = log
				break
			} else {
				fmt.Println("\nUsername must be at least 2 characters\n")
			}
		}
	}

	for {
		fmt.Print("Create a Password.\n")
		pass := ""
		fmt.Fscan(os.Stdin, &pass)
		if len(pass) >= 3 {
			u.Password = pass
			break
		} else {
			fmt.Println("\nPassword must be at least 3 characters\n")
		}
	}
	users = append(users, u)
	fmt.Print("User successfully created .\n")
	serialize()
}

func (u User) SignIn() {
	deserialize()
	for {
		log, pass := "", ""
		fmt.Print("Enter a Login.\n")
		fmt.Fscan(os.Stdin, &log)

		fmt.Print("Enter a Password.\n")
		fmt.Fscan(os.Stdin, &pass)

		ok := false
		for _, us := range users {
			if log == us.Login && pass == us.Password {
				ok = true
				break
			}
		}

		if ok == true {
			fmt.Print("\nYou have successfully logged in to your account, select the action you want to perform \n")
			management()
		} else {
			fmt.Print("\nTry again! Invalid data \n")
			break
		}
	}

}

func (i Item) getRating() int {
	if len(i.ratings) == 0 {
		return 0
	}
	sum := 0
	for _, i := range i.ratings {
		sum += i
		//fmt.Print(i, " ")
	}
	return sum / (len(i.ratings))
}

func rate(id int, r int) bool {
	for i := 0; i < len(warehouse); i++ {
		if id == warehouse[i].id {
			warehouse[i].ratings = append(warehouse[i].ratings, r)
			return true
		}
	}
	return false
}

func findItem(name string) bool {
	ok := false
	for _, i := range warehouse {
		if name == i.name {
			fmt.Println("\nName :", i.name)
			fmt.Println("Price :", i.price, "$")
			fmt.Println("ID :", i.id)
			fmt.Println("Rating : ", i.getRating())
			ok = true
		}
	}
	return ok
}

func printAll(filt string, a int, b int) bool {
	if len(warehouse) == 0 {
		return false
	}
	if filt == "price" {
		for _, i := range warehouse {
			if i.price >= a && i.price <= b {
				fmt.Println("******************")
				fmt.Println("ID: ", i.id)
				fmt.Println("Name: ", i.name)
				fmt.Println("Price :", i.price, "$")
				fmt.Println("Rating: ", i.getRating())
			}
		}
	} else if filt == "rating" {
		for _, i := range warehouse {
			if i.getRating() >= a && i.getRating() <= b {
				fmt.Println("******************")
				fmt.Println("ID: ", i.id)
				fmt.Println("Name: ", i.name)
				fmt.Println("Price :", i.price, "$")
				fmt.Println("Rating: ", i.getRating())
			}
		}
	} else {
		for _, i := range warehouse {
			fmt.Println("******************")
			fmt.Println("ID: ", i.id)
			fmt.Println("Name: ", i.name)
			fmt.Println("Price :", i.price, "$")
			fmt.Println("Rating: ", i.getRating())
		}
	}

	return true
}

func management() {
	fmt.Println("\n1. Поиск товара")
	fmt.Println("2. Посмотреть весь ассортимент")
	fmt.Println("3. Дать оценку товару")
	fmt.Println("4. Выйти с аккаунта")

	for {
		fmt.Print("\nEnter your choice : ")
		choice := 0
		fmt.Fscan(os.Stdin, &choice)
		valid := false

		switch choice {
		case 1:
			fmt.Print("\nEnter name of item : ")
			name := ""
			fmt.Fscan(os.Stdin, &name)
			if findItem(name) {
				valid = true
				break
			} else {
				fmt.Println("\nItem with this name not in stock")
			}
		case 2:
			fmt.Println("\nВыберите тип фильтрации: ")
			fmt.Println("1. По цене: ")
			fmt.Println("2. По рейтингу: ")
			fmt.Println("3. Вывесьти весь ассортимент по дефолту: ")
			fmt.Print("\nEnter your choice : ")
			ch := 0
			a := 0
			b := 0
			fmt.Fscan(os.Stdin, &ch)
			filt := ""
			if ch == 1 {
				filt = "price"
				fmt.Println("\nВведите минимальную цену: ")
				fmt.Fscan(os.Stdin, &a)
				fmt.Println("\nВведите максимальную цену: ")
				fmt.Fscan(os.Stdin, &b)
			} else if ch == 2 {
				filt = "rating"
				fmt.Println("\nВведите минимальный рейтинг: ")
				fmt.Fscan(os.Stdin, &a)
				fmt.Println("\nВведите максимальную рейтинг: ")
				fmt.Fscan(os.Stdin, &b)
			} else {
				filt = "all"
			}
			if printAll(filt, a, b) {
				valid = true
				break
			} else {
				fmt.Println("\nNo items in stock")
			}
		case 3:
			fmt.Print("\nEnter id of item : ")
			id := 0
			fmt.Fscan(os.Stdin, &id)
			fmt.Print("\nEnter rate which you want to add [0/10] : ")
			rat := 0
			fmt.Fscan(os.Stdin, &rat)
			if rate(id, rat) {
				valid = true
				break
			} else {
				fmt.Println("\nItem with this id not in stock")
			}

		case 4:
			Choices()
			valid = true
			break

		case 5:
			// Choices()
			valid = true
			break

		}
		if valid {
			management()
		} else {
			fmt.Println("\nInvalid input! Try Again..")
			management()
		}

	}

}

func Choices() {
	fmt.Print("\n*******************************************************\n")
	fmt.Print("************************SHOP***************************\n")
	fmt.Println("*******************************************************")

	for {
		fmt.Println("\n1. Registration")
		fmt.Println("2. Login ")
		fmt.Println("3. Exit")
		fmt.Println("Enter choice : ")

		c := 0

		fmt.Fscan(os.Stdin, &c)
		u := User{}
		switch c {
		case 1:
			u.SignUp()
			// b=false;
			break

		case 2:
			u.SignIn()
			// b=false;
			break
		case 3:
			fmt.Println("\n*********Thanks**********")
			os.Exit(3)
			break

		default:
			fmt.Println("invalid input... Try again")
			continue
		}

	}
}

func main() {
	warehouse = append(warehouse, Item{"Iphone", 390, 1, []int{}},
		Item{"Samsung A6", 290, 2, []int{}}, Item{"Nokia 3310", 1200, 3, []int{}},
		Item{"Iphone", 340, 4, []int{}}, Item{"Samsung S21", 300, 5, []int{}},
		Item{"Asus v6", 500, 6, []int{}}, Item{"Asus TUF", 103, 7, []int{}},
		Item{"Charger", 40, 8, []int{}}, Item{"Mouse", 50, 9, []int{}},
		Item{"Keyboard", 60, 10, []int{}}, Item{"USB", 20, 11, []int{}},
	)
	//users := map[string]User{}
	////users = append(users, User{"log", "pass"})
	//users["sdf"] = User{"Ab", "sdf"}
	//byteArray, err := json.MarshalIndent(users, "", " ")
	//if err != nil {
	//	fmt.Println("sdfs")
	//}
	//err_ := ioutil.WriteFile("test.json", byteArray, 0644)
	//if err_ != nil {
	//	return
	//}
	//fmt.Println(len(warehouse))
	Choices()
}
