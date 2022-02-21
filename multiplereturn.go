package main

func getBookInfo() (string, int) {
    return "War and Peace", 1000
}

title, pages := getBookInfo()

// In this title will be a string with value "war and peace" and pages will be a int with value 1000