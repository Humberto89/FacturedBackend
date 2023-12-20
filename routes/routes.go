package routes

import (
	"Go_Gin/controllers"

	"Go_Gin/repositories"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Allow-Methods", "Allow-Headers", "Expose-Headers"}
	config.ExposeHeaders = []string{"Content-Length"}

	r.Use(cors.New(config))

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
	r.GET("/descargar-pdf/:id", func(c *gin.Context) {
		// Obtener el _id desde los parámetros de la URL
		id := c.Param("id")

		// Llamar a la función GetPDFDataByID con el _id proporcionado
		err := repositories.GetPDFDataByID(id)
		if err != nil {
			// Manejar el error, por ejemplo, enviar una respuesta de error al cliente
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Si todo está bien, podrías enviar una respuesta de éxito al cliente si es necesario
		c.JSON(http.StatusOK, gin.H{"message": "Descarga exitosa"})
	})
	//Ruta de reportes
	r.GET("/municipio/:id", func(c *gin.Context) { controllers.GetMunicipioByID(c, db) })
	r.GET("/pais/:id", func(c *gin.Context) { controllers.GetPaisByID(c, db) })
	r.GET("/departamento/:id", func(c *gin.Context) { controllers.GetDepartamentoByID(c, db) })
	r.GET("/reportecompra", controllers.ReporteCompra)

	return r
}
