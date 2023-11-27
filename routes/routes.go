package routes

import (
	"Go_Gin/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(r *gin.Engine, collection *mongo.Collection, db *gorm.DB) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Allow-Methods", "Allow-Headers", "Expose-Headers"}
	config.ExposeHeaders = []string{"Content-Length"}

	r.Use(cors.New(config))

	//Ruta para recepcion DTE
	r.POST("/upload", func(c *gin.Context) { controllers.FilesUpload(c, collection) })
	//=================================================================//

	// Rutas para el CRUD de Inscripcion de provedores
	r.GET("/formulario", func(c *gin.Context) { controllers.GetFormulario(c, db) })
	r.GET("/formulario/:id", func(c *gin.Context) { controllers.GetFormularioByID(c, db) })
	r.POST("/formulario", func(c *gin.Context) { controllers.CreateFormulario(c, db) })
	r.PUT("/formulario/:id", func(c *gin.Context) { controllers.UpdateFormulario(c, db) })
	r.DELETE("/formulario/:id", func(c *gin.Context) { controllers.DeleteFormulario(c, db) })
	// En tus rutas de servidor
	r.GET("/paises", func(c *gin.Context) { controllers.GetPaises(c, db) })
	r.GET("/departamentos/:paisID", func(c *gin.Context) { controllers.GetDepartamentos(c, db) })
	r.GET("/municipios/:departamentoID", func(c *gin.Context) { controllers.GetMunicipios(c, db) })
	//=================================================================//
	//Ruta para la busqueda de DTE
	r.POST("/busqueda", func(c *gin.Context) { controllers.GetDTEs(c) })

	return r
}
