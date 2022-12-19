<h1>🚀 Echo Playground!</h1>

<p>This repository is nothing because is just my playground to learn to build a REST API on Echo</p>

<h2>Reference</h2>
<a href="https://echo.labstack.com/">https://echo.labstack.com/</a>

<h2>Installation</h2>

```bash
go get github.com/labstack/echo/v4
```

<h2>Quick Start</h2>

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf(err.Error())
	}

	appUrl := fmt.Sprintf("%s:%s", os.Getenv("ECHO_HOST"), os.Getenv("ECHO_PORT"))

	e := echo.New()
	// Routes
	e.Logger.Fatal(e.Start(appUrl))
}
```

<h2>API List</h2>
<ul>
    <li>Find Books</li>
    <li>Insert Books</li>
</ul>
