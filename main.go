package main

import (
    "fmt" 
    "github.com/kataras/iris/v12"
)


func main(){
    fmt.Println("|---------------------------|")
	fmt.Println("|----------admin------------|")
	fmt.Println("|---------------------------|")


    app := iris.Default()

	app.Get("/",func(ctx iris.Context){
		ctx.HTML("html")
	})

	app.Run(iris.Addr(":86"))
}
