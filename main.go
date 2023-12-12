package main

import (
    "fmt"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/storage/mysql"
)

func xyzHandler(c *fiber.Ctx) error {
    msg := fmt.Sprintf("%s, %s", c.Params("file"), c.Params("ext"))
    return c.SendString(msg)

}

func main() {
    fmt.Println("Hello,")

   // store := mysql.New(mysql.Config{
   // ConnectionURI:   "root@tcp(127.0.0.1:3306)/go-test",
   // Table: "person",
   // Reset:           false,
   // })

    app := fiber.New()

    // adding logger 
    app.Use(logger.New())

    app.Get("/abc", func(c *fiber.Ctx) error {
        return c.SendString("abc route")
    })

    app.Get("/xyz/:file.:ext", xyzHandler)

    app.Listen(":3000")
}
