package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Zireael13/capstone-archive/server/internal/graph"
	"github.com/Zireael13/capstone-archive/server/internal/graph/generated"
	"github.com/gin-gonic/gin"
	"github.com/matthewhartstonge/argon2"
	"gorm.io/gorm"
)

func GraphQLHandler(db *gorm.DB, argon *argon2.Config) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db, Argon: argon}},
		),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func PlaygroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL playground", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
