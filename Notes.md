# Go Backend with Gin - Quick Guide

## 1. Packages vs Modules

### Package

A **package** is a collection of `.go` files in one directory. It's for organizing code within your project.

```go
// file: user/user.go
package user  // All files in user/ folder must be package user

type User struct {
    ID   int
    Name string
}
```

### Module

A **module** is a versioned collection of packages managed by `go.mod` and `go.sum`. It's for dependency management.

```
go.mod file declares:
- Module name: github.com/yourname/myproject
- Go version required
- Dependencies and versions
```

| Aspect | Package                 | Module                         |
| ------ | ----------------------- | ------------------------------ |
| Scope  | Local code organization | Global dependency management   |
| File   | None                    | `go.mod`                       |
| Import | `import "myapp/user"`   | `import "github.com/user/pkg"` |

---

## 2. Understanding go.sum

**go.sum** is a lock file that stores cryptographic hashes of all dependencies to ensure:

- ✅ Everyone gets exact same versions
- ✅ Dependencies haven't been tampered with
- ✅ Reproducible builds

### Why Two Hashes Per Dependency?

```
github.com/gin-gonic/gin v1.9.0 h1:abc123=      # Package hash
github.com/gin-gonic/gin v1.9.0/go.mod h1:def456=  # go.mod hash
```

- First: Verifies the entire package
- Second: Verifies the package's dependencies

### Auto-Updated When You:

```bash
go get github.com/gin-gonic/gin      # Add dependency
go get -u ./...                      # Update all
go mod tidy                          # Clean up
```

### Best Practices:

- ✅ Always commit `go.mod` and `go.sum` to git
- ✅ Run `go mod tidy` before committing
- ❌ Never manually edit `go.sum`
- ❌ Never delete `go.sum`

---

## 3. Gin Framework Overview

**Gin** is a fast, lightweight web framework for building REST APIs.

**Why Gin?**

- ⚡ Blazing fast performance
- 📦 Minimal dependencies
- 🛣️ Simple routing
- 🔧 Built-in middleware support
- 🎯 Easy JSON handling

### Quick Start

```bash
go get github.com/gin-gonic/gin
```

```go
package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/hello", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Hello World"})
    })
    r.Run(":8080")
}
```

---

## 4. gin.Context Explained

**gin.Context** carries request/response data through your handler. It's like the "glue" between HTTP and your code.

### Common Methods

```go
func handler(c *gin.Context) {
    // 📥 GET DATA FROM REQUEST
    id := c.Param("id")                    // URL: /users/123
    role := c.Query("role")                // Query: ?role=admin
    token := c.GetHeader("Authorization")  // Headers
    c.GetJSON(&user)                       // Parse JSON body

    // 📤 SEND DATA IN RESPONSE
    c.JSON(200, gin.H{"key": "value"})     // Send JSON
    c.String(200, "Hello")                 // Send text
    c.File("image.png")                    // Send file
    c.Redirect(301, "/newpath")            // Redirect

    // ❌ ERROR HANDLING
    c.AbortWithStatusJSON(400, err)        // Stop & error
}
```

### Real Example

```go
func getUser(c *gin.Context) {
    id := c.Param("id")                    // /users/123
    role := c.Query("role")                // ?role=admin
    token := c.GetHeader("Authorization")

    c.JSON(200, gin.H{
        "user_id": id,
        "role": role,
        "token": token,
    })
}
```

---

## 5. gin.H Explained

**gin.H** is shorthand for `map[string]interface{}` - an easy way to create JSON responses.

```go
// Simple response
c.JSON(200, gin.H{"message": "Success", "id": 42})

// Nested data
c.JSON(200, gin.H{
    "status": "ok",
    "user": gin.H{
        "name": "Alice",
        "email": "alice@example.com",
    },
})

// Error response
c.JSON(400, gin.H{"error": "Invalid email"})
```

**Tip:** For complex/repeated responses, use **structs** instead:

```go
type UserResponse struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
}

user := UserResponse{ID: 1, Name: "Alice"}
c.JSON(200, user)
```

---

## 6. API Grouping

**Group** related endpoints to organize code and apply middleware (like auth) to specific routes.

### Basic Grouping

```go
func main() {
    r := gin.Default()

    // Users API
    users := r.Group("/users")
    {
        users.GET("", getAllUsers)          // GET /users
        users.GET(":id", getUser)           // GET /users/123
        users.POST("", createUser)          // POST /users
        users.PUT(":id", updateUser)        // PUT /users/123
        users.DELETE(":id", deleteUser)     // DELETE /users/123
    }

    // Products API
    products := r.Group("/products")
    {
        products.GET("", getAllProducts)    // GET /products
        products.GET(":id", getProduct)     // GET /products/456
        products.POST("", createProduct)    // POST /products
    }

    r.Run(":8080")
}
```

### With Middleware (Authentication)

```go
// Auth middleware
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token != "valid-token" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        c.Next()
    }
}

func main() {
    r := gin.Default()

    // Public endpoints
    public := r.Group("/api")
    {
        public.GET("/products", getAllProducts)
    }

    // Protected endpoints
    protected := r.Group("/api")
    protected.Use(authMiddleware())
    {
        protected.POST("/products", createProduct)
        protected.DELETE("/products/:id", deleteProduct)
    }

    r.Run(":8080")
}
```

### Nested Groups (Versioning)

```go
r := gin.Default()

v1 := r.Group("/api/v1")
{
    users := v1.Group("/users")
    {
        users.GET("", getAllUsers)
        users.POST("", createUser)

        // Nested: /api/v1/users/123/orders
        orders := users.Group(":userID/orders")
        {
            orders.GET("", getUserOrders)
            orders.POST("", createOrder)
        }
    }
}

v2 := r.Group("/api/v2")
{
    v2.GET("/users", getAllUsersV2)  // Improved version
}
```

### Project Structure

```
myapi/
├── main.go
├── controller/
│   ├── user.go
│   ├── product.go
│   └── order.go
├── middleware/
│   ├── auth.go
│   └── logger.go
└── router/
    └── routes.go
```

```go
// router/routes.go
func SetupRoutes(r *gin.Engine) {
    r.Use(middleware.Logger())

    v1 := r.Group("/api/v1")
    {
        // Users
        users := v1.Group("/users")
        {
            users.GET("", controller.GetAllUsers)
            users.POST("", controller.CreateUser)
        }

        // Products (admin only)
        products := v1.Group("/products")
        {
            products.GET("", controller.GetAllProducts)

            admin := products.Group("")
            admin.Use(middleware.AdminAuth())
            {
                admin.POST("", controller.CreateProduct)
                admin.DELETE(":id", controller.DeleteProduct)
            }
        }
    }
}
```

---

## Quick Reference

```bash
# Packages & Modules
go mod init myproject                 # Create module
go get github.com/user/package        # Add dependency
go get -u ./...                       # Update all dependencies
go mod tidy                           # Clean up go.mod & go.sum

# Testing
go test ./...                         # Run all tests
```

---

**Happy Coding! 🚀**
