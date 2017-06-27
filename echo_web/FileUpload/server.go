package main

import (
	"os"

	"io"

	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func upload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	// 上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// 创建一个新文件来保存
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()
	if err != nil {
		return err
	}

	// 把上传的文件拷贝给刚创建的新文件
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>文件 %s 上传成功! name=%s email=%s</p>", file.Filename, name, email))

}

func main() {
	e := echo.New()

	e.Use(middleware.Logger(), middleware.Recover())

	e.Static("/", "public")
	e.POST("/upload", upload)

	e.Logger.Fatal(e.Start(":1323"))
}
