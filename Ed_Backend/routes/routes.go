// routes/routes.go
package routes

import (
	"Backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
)

func SetupRouter(r *gin.Engine, db *gorm.DB) *gin.Engine {

	// Configuracion middleware CORS
	//config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:5173"} // URL de frontend
	// config.AllowHeaders = []string{"Access-Control-Allow-Headers", "X-Requested-With", "Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
	// config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Allow-Methods", "Allow-Headers", "Expose-Headers"}
	config.ExposeHeaders = []string{"Content-Length"}
	// config.AllowCredentials = true

	r.Use(cors.New(config))
	// r.Use(cors.New(cors.Config{
	// 	//AllowOrigins:     []string{"http://localhost:5173"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-length"},
	// 	AllowCredentials: true,
	// 	AllowAllOrigins:  true,
	// }))

	// Rutas para el CRUD
	r.GET("/formulario", func(c *gin.Context) { controllers.GetFormulario(c, db) })
	r.GET("/formulario/:id", func(c *gin.Context) { controllers.GetFormularioByID(c, db) })
	r.POST("/formulario", func(c *gin.Context) { controllers.CreateFormulario(c, db) })
	r.PUT("/formulario/:id", func(c *gin.Context) { controllers.UpdateFormulario(c, db) })
	r.DELETE("/formulario/:id", func(c *gin.Context) { controllers.DeleteFormulario(c, db) })
	// En tus rutas de servidor
	r.GET("/paises", func(c *gin.Context) { controllers.GetPaises(c, db) })
	r.GET("/departamentos/:paisID", func(c *gin.Context) { controllers.GetDepartamentos(c, db) })
	r.GET("/municipios/:departamentoID", func(c *gin.Context) { controllers.GetMunicipios(c, db) })

	return r
}
