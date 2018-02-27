package main

import "github.com/kataras/iris"

func main() {
  app := iris.Default()

  app.Handle("GET", "/", func(ctx iris.Context) {
    ctx.HTML("Hello Go World!\n")
  })

  app.Run(iris.Addr(":10080"))
}
