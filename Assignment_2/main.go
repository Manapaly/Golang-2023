package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func templating(w http.ResponseWriter, filename string, data interface{}) {
	t, _ := template.ParseFiles(filename)
	err := t.ExecuteTemplate(w, filename, data)
	if err != nil {
		return
	}
}

func getSignInPage(w http.ResponseWriter) {
	templating(w, "SignIn.html", nil)
}

func getSignUpPage(w http.ResponseWriter) {
	templating(w, "SignUp.html", nil)
}

func getMainPage(w http.ResponseWriter) {
	templating(w, "MainPage.html", nil)
}

func getUserPage(w http.ResponseWriter) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, GetListOfProducts())
}

func getUser(r *http.Request) User {
	FirstName := r.FormValue("firstName")
	SecondName := r.FormValue("secondName")
	Phone := r.FormValue("phoneNo")
	email := r.FormValue("email")
	password := r.FormValue("password")
	return User{Email: email, Password: password, FirsName: FirstName, LastName: SecondName, Phone: Phone}
}

func signInUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	ok := Admin.VerifyUser(newUser)
	t, _ := template.ParseFiles("SignIn.html")
	if ok {
		getUserPage(w)
		return
	} else {
		t.ExecuteTemplate(w, "SignIn.html", "User doesn't exist, Try Again!")
	}
	return

}

func signUpUser(w http.ResponseWriter, r *http.Request) {
	newUser := getUser(r)
	err := Admin.CreateUser(newUser)

	fileName := "SignUp.html"
	t, _ := template.ParseFiles(fileName)
	if err != nil {
		err := t.ExecuteTemplate(w, fileName, "New User Sign-up Failure, Try Again!")
		if err != nil {
			return
		}
		return
	}

	err = t.ExecuteTemplate(w, fileName, "New User Sign-up Success!")
	if err != nil {
		return
	}
	return
}

func getUserPageFilteredByName(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, SearchByName(r))
}

func getUserPageFilteredByPrice(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, SearchByPrice(r))
}

func getUserPageFilteredByRating(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, SearchByRating(r))
}

func getUserPageAfterRating(w http.ResponseWriter, r *http.Request) {
	Products = RateProduct(r)
	t, _ := template.ParseFiles("list-products.html")
	t.Execute(w, GetListOfProducts())
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/sign-in":
		signInUser(w, r)
	case "/sign-up":
		signUpUser(w, r)
	case "/sign-in-form":
		getSignInPage(w)
	case "/sign-up-form":
		getSignUpPage(w)
	case "/sign-in-show-all":
		getUserPage(w)
	case "/sign-in-search-by-name":
		getUserPageFilteredByName(w, r)
	case "/sign-in-search-by-price":
		getUserPageFilteredByPrice(w, r)
	case "/sign-in-search-by-rating":
		getUserPageFilteredByRating(w, r)
	case "/sign-in-rating":
		getUserPageAfterRating(w, r)
	default:
		getMainPage(w)
	}
}

func main() {
	fmt.Println("http://localhost:8080")
	http.HandleFunc("/", userHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
