package cmdRestFulRoutes

import (
	portsRestFul "togo/internal/ports/restful"
	portsRestFulTask "togo/internal/ports/restful/tasks"
	portsRestFulTodo "togo/internal/ports/restful/todos"
	portsRestFulUser "togo/internal/ports/restful/users"

	"github.com/gin-gonic/gin"
)

// Only for restful API
func Routes(router *gin.Engine) {
	router.GET("/", portsRestFul.Welcome)

	// Health check routes
	router.GET("/api/ping", portsRestFul.HealthCheck)

	// Users routes
	router.GET("/api/users", portsRestFulUser.GetAllUsers)
	router.POST("/api/user", portsRestFulUser.CreateUser)
	router.GET("/api/user/:userId", portsRestFulUser.GetSingleUser)
	router.PUT("/api/user/:userId", portsRestFulUser.EditUser)
	router.DELETE("/api/user/:userId", portsRestFulUser.DeleteUser)

	// Tasks routes
	router.GET("/api/tasks", portsRestFulTask.GetAllTasks)
	router.POST("/api/task", portsRestFulTask.CreateUser)
	router.GET("/api/task/:taskId", portsRestFulTask.GetSingleUser)
	router.PUT("/api/task/:taskId", portsRestFulTask.EditUser)
	router.DELETE("/api/task/:taskId", portsRestFulTask.DeleteUser)

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
