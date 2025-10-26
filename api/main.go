package main

import (
	authcontroller "go-restapi-gin/controllers/authcontrollers"
	usercontroller "go-restapi-gin/controllers/usercontrollers"
	"go-restapi-gin/middleware"
	"go-restapi-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// auth
	r.POST("/api/register", authcontroller.Register)
	r.POST("/api/login", authcontroller.Login)
	r.POST("/api/logout", middleware.AuthRequired(), authcontroller.Logout)

	// endpoint public / orders
	r.GET("/api/orders", usercontroller.Users)
	r.GET("/api/orders/:id", usercontroller.Show)
	r.POST("/api/orders", usercontroller.Create)
	r.PUT("/api/orders/:id", usercontroller.Update)
	r.DELETE("/api/orders", usercontroller.Delete)

	// contoh protected route
	protected := r.Group("/api")
	protected.Use(middleware.AuthRequired())

	// Role
	protected.GET("/roles", usercontroller.IndexRole)
	protected.GET("/roles/:id", usercontroller.ShowRole)
	protected.POST("/roles", usercontroller.CreateRole)
	protected.PUT("/roles/:id", usercontroller.UpdateRole)
	protected.DELETE("/roles/:id", usercontroller.DeleteRole)

	// Menu
	protected.GET("/menus", usercontroller.IndexMenu)
	protected.GET("/menus/:id", usercontroller.ShowMenu)
	protected.POST("/menus", usercontroller.CreateMenu)
	protected.PUT("/menus/:id", usercontroller.UpdateMenu)
	protected.DELETE("/menus/:id", usercontroller.DeleteMenu)

	// Akses
	protected.GET("/akses", usercontroller.IndexAkses)
	protected.GET("/akses/:id", usercontroller.ShowAkses)
	protected.POST("/akses", usercontroller.CreateAkses)
	protected.PUT("/akses/:id", usercontroller.UpdateAkses)
	protected.DELETE("/akses/:id", usercontroller.DeleteAkses)

	// Contact Messages API
	r.GET("/contact_messages", usercontroller.IndexContact)
	r.GET("/contact_messages/:id", usercontroller.ShowContact)
	r.POST("/contact_messages", usercontroller.CreateContact)
	r.PUT("/contact_messages/:id", usercontroller.UpdateContact)
	r.DELETE("/contact_messages/:id", usercontroller.DeleteContact)

	// Company Profile API
	protected.GET("/company_profile", usercontroller.IndexCompany)
	protected.GET("/company_profile/:id", usercontroller.ShowCompany)
	protected.POST("/company_profile", usercontroller.CreateCompany)
	protected.PUT("/company_profile/:id", usercontroller.UpdateCompany)
	protected.DELETE("/company_profile/:id", usercontroller.DeleteCompany)

	//users
	protected.GET("/profile", func(c *gin.Context) {
		u, _ := c.Get("user")
		c.JSON(200, gin.H{"user": u})
	})

	// Users API
	protected.GET("/users", usercontroller.Users)
	protected.GET("/users/:id", usercontroller.Show)
	protected.POST("/users", usercontroller.Create)
	protected.PUT("/users/:id", usercontroller.Update)
	protected.DELETE("/users/:id", usercontroller.Delete)

	// Blog APIs (protected)
	protected.GET("/blog-posts", usercontroller.IndexBlogPosts)
	protected.GET("/blog-posts/:id", usercontroller.ShowBlogPost)
	protected.POST("/blog-posts", usercontroller.CreateBlogPost)
	protected.PUT("/blog-posts/:id", usercontroller.UpdateBlogPost)
	protected.DELETE("/blog-posts/:id", usercontroller.DeleteBlogPost)

	protected.GET("/blog-comments", usercontroller.IndexBlogComments)
	protected.GET("/blog-comments/:id", usercontroller.ShowBlogComment)
	protected.POST("/blog-comments", usercontroller.CreateBlogComment)
	protected.PUT("/blog-comments/:id", usercontroller.UpdateBlogComment)
	protected.DELETE("/blog-comments/:id", usercontroller.DeleteBlogComment)

	// Posts Editor API
	api := r.Group("/api")
	{
		api.GET("/posts", usercontroller.IndexPosts)
		api.GET("/posts/:id", usercontroller.ShowPost)
		api.POST("/posts", usercontroller.CreatePost)
		api.PUT("/posts/:id", usercontroller.UpdatePost)
		api.DELETE("/posts/:id", usercontroller.DeletePost)

		// FAQ routes
		api.GET("/faqs", usercontroller.IndexFaqs)
		api.GET("/faqs/:id", usercontroller.ShowFaq)
		api.POST("/faqs", usercontroller.CreateFaq)
		api.PUT("/faqs/:id", usercontroller.UpdateFaq)
		api.DELETE("/faqs/:id", usercontroller.DeleteFaq)


	}

	// Service API
	api.GET("/services", usercontroller.IndexServices)
	api.GET("/services/:id", usercontroller.ShowService)
	protected.POST("/services", usercontroller.CreateService)
	protected.PUT("/services/:id", usercontroller.UpdateService)
	protected.DELETE("/services/:id", usercontroller.DeleteService)

	// projects, project_images APIs (protected)
	protected.GET("/project-images", usercontroller.IndexProjectImages)
	protected.GET("/project-images/:id", usercontroller.ShowProjectImage)
	protected.POST("/project-images", usercontroller.CreateProjectImage)
	protected.PUT("/project-images/:id", usercontroller.UpdateProjectImage)
	protected.DELETE("/project-images/:id", usercontroller.DeleteProjectImage)

	// Template API (PUBLIC - untuk editor)
	// Daftar semua template
	api.GET("/templates", usercontroller.GetAllTemplates)

	// Mengambil template specific
	api.GET("/templates/users", usercontroller.GetUsersTemplate)
	api.GET("/templates/contact", usercontroller.GetContactTemplate)
	api.GET("/templates/about", usercontroller.GetAboutTemplate)

	// Dynamic template getter
	api.GET("/templates/:name", usercontroller.GetTemplateByName)

	// Save template sebagai post
	api.POST("/templates/save", usercontroller.SaveTemplateAsPost)

	// Gallery API
	protected.GET("/galleries", usercontroller.IndexGalleries)
	protected.GET("/galleries/:id", usercontroller.ShowGallery)
	protected.POST("/galleries", usercontroller.CreateGallery)
	protected.PUT("/galleries/:id", usercontroller.UpdateGallery)
	protected.DELETE("/galleries/:id", usercontroller.DeleteGallery)

	// Upload logo client (public, bisa juga diproteksi)
	r.POST("/api/upload-client-logo", usercontroller.UploadClientLogo)

	// Team Members API
	api.GET("/team_members", usercontroller.IndexTeamMembers)
	api.GET("/team_members/:id", usercontroller.ShowTeamMember)
	protected.POST("/team_members", usercontroller.CreateTeamMember)
	protected.PUT("/team_members/:id", usercontroller.UpdateTeamMember)
	protected.DELETE("/team_members/:id", usercontroller.DeleteTeamMember)

	// Clients API
	protected.GET("/clients", usercontroller.IndexClients)
	protected.GET("/clients/:id", usercontroller.ShowClient)
	protected.POST("/clients", usercontroller.CreateClient)
	protected.PUT("/clients/:id", usercontroller.UpdateClient)
	protected.DELETE("/clients/:id", usercontroller.DeleteClient)

	// Serve static files for uploads (logo, gallery, etc)
	r.Static("/uploads", "./uploads")

	protected.GET("/projects", usercontroller.IndexProjects)
	protected.GET("/projects/:id", usercontroller.ShowProject)
	protected.POST("/projects", usercontroller.CreateProject)
	protected.PUT("/projects/:id", usercontroller.UpdateProject)
	protected.DELETE("/projects/:id", usercontroller.DeleteProject)

	// === ROUTES EXPERIENCES ===
	r.GET("/api/experiences", usercontroller.IndexExperiences)
	r.GET("/api/experiences/:id", usercontroller.ShowExperience)
	r.POST("/api/experiences", usercontroller.CreateExperience)
	r.PUT("/api/experiences/:id", usercontroller.UpdateExperience)
	r.DELETE("/api/experiences/:id", usercontroller.DeleteExperience)

	r.GET("/api/pelatihan", usercontroller.IndexPelatihan)
    r.GET("/api/pelatihan/:id", usercontroller.ShowPelatihan)
    r.POST("/api/pelatihan", usercontroller.CreatePelatihan)
    r.PUT("/api/pelatihan/:id", usercontroller.UpdatePelatihan)
    r.DELETE("/api/pelatihan/:id", usercontroller.DeletePelatihan)

	// === ROUTES TESTIMONI (PUBLIC) ===
	r.GET("/api/testimonis", usercontroller.IndexTestimoni)
	r.GET("/api/testimonis/:id", usercontroller.ShowTestimoni)
	r.POST("/api/testimonis", usercontroller.CreateTestimoni)
	r.PUT("/api/testimonis/:id", usercontroller.UpdateTestimoni)
	r.DELETE("/api/testimonis/:id", usercontroller.DeleteTestimoni)

	// === ROUTES HERO BACKGROUND (PUBLIC) ===
	r.GET("/api/hero_backgrounds", usercontroller.IndexHeroBackground)
	r.GET("/api/hero_backgrounds/:id", usercontroller.ShowHeroBackground)
	r.POST("/api/hero_backgrounds", usercontroller.CreateHeroBackground)
	r.PUT("/api/hero_backgrounds/:id", usercontroller.UpdateHeroBackground)
	r.DELETE("/api/hero_backgrounds/:id", usercontroller.DeleteHeroBackground)

	r.Run()
}
