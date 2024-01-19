package routes

import (
	"Go_Gin/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "IdentifierEmp", "Authorization", "Access-Control-Allow-Origin"}
	// Habilita el uso de credenciales en las solicitudes CORS (si es necesario)
	config.AllowCredentials = true

	r.Use(cors.New(config))

	// r.OPTIONS("/*any", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Origin", "*")
	// 	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	// 	c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, IdentifierEmp, Authorization")
	// 	c.Header("Access-Control-Allow-Credentials", "true")

	// 	c.Status(204)
	// })

	//Ruta para recepcion DTE
	r.POST("/upload", func(c *gin.Context) { controllers.FilesUpload(c) })
	//=================================================================//

	// Rutas para el CRUD de Inscripcion de provedores
	r.GET("/formulario", func(c *gin.Context) { controllers.GetFormulario(c, db) })
	r.GET("/formulario/:id", func(c *gin.Context) { controllers.GetFormularioByID(c, db) })
	r.POST("/formulario", func(c *gin.Context) { controllers.CreateFormulario(c, db) })
	r.PUT("/formulario/:id", func(c *gin.Context) { controllers.UpdateFormulario(c, db) })
	r.DELETE("/formulario/:id", func(c *gin.Context) { controllers.DeleteFormulario(c, db) })
	r.GET("/paises", func(c *gin.Context) { controllers.GetPaises(c, db) })
	r.GET("/departamentos/:paisID", func(c *gin.Context) { controllers.GetDepartamentos(c, db) })
	r.GET("/municipios/:departamentoID", func(c *gin.Context) { controllers.GetMunicipios(c, db) })
	r.GET("/checkIdentification", func(c *gin.Context) { controllers.CheckIdentificationExists(c, db) })
	//=================================================================//
	//Ruta para la busqueda de DTE
	r.GET("/busqueda", func(c *gin.Context) { controllers.GetDTEs(c, db) })
	//Ruta de reportes
	r.GET("/municipio/:id", func(c *gin.Context) { controllers.GetMunicipioByID(c, db) })
	r.GET("/pais/:id", func(c *gin.Context) { controllers.GetPaisByID(c, db) })
	r.GET("/departamento/:id", func(c *gin.Context) { controllers.GetDepartamentoByID(c, db) })
	r.GET("/reporteanexo", controllers.ReporteAnexo)

	//Ruta para descargar pdf
	r.GET("/pdfdataget/:id", func(c *gin.Context) { controllers.PdfDataGet(c) })

	return r
}
