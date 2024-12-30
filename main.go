package main

import (
    "database/sql"
    "log"
    "net/http"
    
    "crud-app/config"
    "crud-app/handler"
    "crud-app/repository"
    "crud-app/service"
    
    _ "github.com/lib/pq"
)

func main() {
    config, err := config.LoadConfig(".")
    if err != nil {
        log.Fatal("cannot load config:", err)
    }

    conn, err := sql.Open(config.DBDriver, config.DBSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    userRepo := repository.NewUserRepository(conn)
    userService := service.NewUserService(userRepo)
    userHandler := handler.NewUserHandler(userService)

    // Register routes
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodPost:
            userHandler.CreateUser(w, r)
        case http.MethodGet:
            if r.URL.Query().Get("id") != "" {
                userHandler.GetUser(w, r)
            } else {
                userHandler.ListUsers(w, r)
            }
        case http.MethodPut:
            userHandler.UpdateUser(w, r)
        case http.MethodDelete:
            userHandler.DeleteUser(w, r)
        default:
            http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        }
    })

    log.Printf("Server starting on %s", config.ServerAddress)
    err = http.ListenAndServe(config.ServerAddress, nil)
    if err != nil {
        log.Fatal("cannot start server:", err)
    }
}
