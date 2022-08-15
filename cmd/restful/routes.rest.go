package cmdRestFulRoutes

import (
	portsRestFul "togo/internal/ports/restful"
	portsRestFulTask "togo/internal/ports/restful/tasks"
	portsRestFulTodo "togo/internal/ports/restful/todos"
	portsRestFulUser "togo/internal/ports/restful/users"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
)

func NewRoot(db *pg.DB) *gin.Engine {
	// Init Router
	router := gin.Default()

	// Route Handlers / Endpoints
	// cmdRestFulRoutes.Routes(router)

	initRoutes(router, db)

	return router
}

// Only for restful API
func initRoutes(router *gin.Engine, db *pg.DB) {
	userPort := portsRestFulUser.NewUserPort(db)
	taskPort := portsRestFulTask.NewTaskPort(db)

	router.GET("/", portsRestFul.Welcome)

	// Health check routes
	router.GET("/api/ping", portsRestFul.HealthCheck)

	// Users routes
	router.GET("/api/users", userPort.GetAllUsers)
	router.POST("/api/user", userPort.CreateUser)
	router.GET("/api/user/:userId", userPort.GetSingleUser)
	router.PUT("/api/user/:userId", userPort.EditUser)
	router.DELETE("/api/user/:userId", userPort.DeleteUser)

	// Tasks routes
	router.GET("/api/tasks", taskPort.GetAllTasks)
	router.POST("/api/task", taskPort.CreateTask)
	router.GET("/api/task/:taskId", taskPort.GetSingleTask)
	router.DELETE("/api/task/:userId/:taskId", taskPort.DeleteTask)

	// Example routes
	// router.GET("/albums", servicesAlbums.GetAlbums)
	// router.GET("/albums/:id", servicesAlbums.GetAlbumByID)
	// router.POST("/albums", servicesAlbums.PostAlbums)

	router.GET("/todos", portsRestFulTodo.GetAllTodos)
	router.POST("/todo", portsRestFulTodo.CreateTodo)
	router.GET("/todo/:todoId", portsRestFulTodo.GetSingleTodo)
	router.PUT("/todo/:todoId", portsRestFulTodo.EditTodo)
	router.DELETE("/todo/:todoId", portsRestFulTodo.DeleteTodo)
	router.NoRoute(portsRestFul.NotFound)
}
