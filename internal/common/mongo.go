package common

import (
	"github.com/SeeDAO-OpenSource/sgn/pkg/app"
	"github.com/SeeDAO-OpenSource/sgn/pkg/db/mongodb"
	"github.com/SeeDAO-OpenSource/sgn/pkg/services"
	"github.com/SeeDAO-OpenSource/sgn/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMongoClient(ac *app.AppBuilder) {
	ac.ConfigureServices(func() error {
		var mongoOptions = &mongodb.MongoOptions{
			URL: "mongodb://localhost:27017",
		}
		utils.ViperBind("Mongo", mongoOptions)
		services.AddValue(mongoOptions)
		services.AddTransient(func(c *services.Container) *mongo.Client {
			client, err := mongodb.GetClient(mongoOptions)
			if err != nil {
				return nil
			}
			return client
		})
		return nil
	})
}
