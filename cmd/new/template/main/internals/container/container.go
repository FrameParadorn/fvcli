package container

import (
	"fmt"
	"net/http"

	"git.robodev.co/imp/shared-utility/validator"
	"{{ service_name }}/internals/config"
	"{{ service_name }}/internals/controller"
	"{{ service_name }}/internals/infrastructure/database"
	"{{ service_name }}/internals/infrastructure/grpcServer"
	"{{ service_name }}/internals/infrastructure/httpServer"
	"{{ service_name }}/internals/infrastructure/jaeger"
	"{{ service_name }}/internals/infrastructure/logrus"
	"{{ service_name }}/internals/repository/postgres"
	"{{ service_name }}/internals/utils"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/dig"
)

type Container struct {
	container *dig.Container
}

func (c *Container) Configure() error {
	servicesConstructors := []interface{}{
		config.NewConfiguration,
		grpcServer.NewServer,
		database.NewServerBase,
		http.NewServeMux,
		httpServer.NewServer,
		runtime.NewServeMux,
		jaeger.NewJaeger,
		logrus.NewLog,
		controller.NewHealthZController,
		validator.NewCustomValidator,
		postgres.NewRepository,
		utils.NewUtils,
		utils.NewCustomValidator,
	}

	for _, service := range servicesConstructors {
		if err := c.container.Provide(service); err != nil {
			return err
		}
	}
	appConfig := config.NewConfiguration()
	jaeger.NewJaeger(appConfig)
	return nil
}

func (c *Container) Start() error {
	fmt.Println("Start Container")

	if err := c.container.Invoke(func(s *grpcServer.Server, h *httpServer.Server, v *validator.CustomValidator) {
		go func() {
			_ = h.Start()
		}()
		s.Start()

	}); err != nil {
		fmt.Printf("%s", err)

		return err
	}

	return nil
}

// MigrateDB ...
func (c *Container) MigrateDB() error {
	fmt.Println("Start Container DB")

	if err := c.container.Invoke(func(d *database.DB) {
		d.MigrateDB()
	}); err != nil {
		return err
	}

	return nil
}

func NewContainer() (*Container, error) {
	d := dig.New()

	container := &Container{
		container: d,
	}

	if err := container.Configure(); err != nil {
		return nil, err
	}

	return container, nil
}
