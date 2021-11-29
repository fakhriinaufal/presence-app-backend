package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	_middleware "presence-app-backend/app/middlewares"
	"presence-app-backend/app/routes"
	_departmentUsecase "presence-app-backend/business/departments"
	_presenceUsecase "presence-app-backend/business/presences"
	_scheduleUsecase "presence-app-backend/business/schedules"
	_userUsecase "presence-app-backend/business/users"
	_departmentController "presence-app-backend/controllers/departments"
	_presenceController "presence-app-backend/controllers/presences"
	_scheduleController "presence-app-backend/controllers/schedules"
	_userController "presence-app-backend/controllers/users"
	_departmentRepo "presence-app-backend/drivers/databases/departments"
	"presence-app-backend/drivers/databases/mysql"
	_presenceRepo "presence-app-backend/drivers/databases/presences"
	_scheduleRepo "presence-app-backend/drivers/databases/schedules"
	_userRepo "presence-app-backend/drivers/databases/users"
	"time"
)

func init() {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func DbMigrate(db *gorm.DB) {
	db.AutoMigrate(&_departmentRepo.Department{})
	db.AutoMigrate(&_userRepo.User{})
	db.AutoMigrate(&_scheduleRepo.Schedule{})
	db.AutoMigrate(&_presenceRepo.Presence{})
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: viper.GetString("DB_USERNAME"),
		DB_Password: viper.GetString("DB_PASSWORD"),
		DB_Host:     viper.GetString("DB_HOST"),
		DB_Port:     viper.GetString("DB_PORT"),
		DB_Database: viper.GetString("DB_NAME"),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("JWT_SECRET"),
		ExpiresDuration: viper.GetInt("JWT_EXPIRED"),
	}

	conn := configDB.InitDB()
	DbMigrate(conn)

	e := echo.New()
	timeoutContext := time.Duration(2)

	// department route
	departmentRepository := _departmentRepo.NewMysqlDepartmentRepository(conn)
	departmentUsecase := _departmentUsecase.NewDepartmentUsecase(departmentRepository, timeoutContext)
	departmentController := _departmentController.NewDepartmentController(departmentUsecase)

	// user route
	userRepository := _userRepo.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, departmentRepository, timeoutContext, configJWT)
	userController := _userController.NewUserController(userUsecase)

	// schedule route
	scheduleRepository := _scheduleRepo.NewMysqlScheduleRepository(conn)
	scheduleUsecase := _scheduleUsecase.NewScheduleUsecase(scheduleRepository, departmentRepository, timeoutContext)
	scheduleController := _scheduleController.NewScheduleController(scheduleUsecase)

	// presence route
	presenceRepository := _presenceRepo.NewMysqlPresenceRepository(conn)
	presenceUsecase := _presenceUsecase.NewPresenceUsecase(presenceRepository, userRepository, scheduleUsecase, timeoutContext)
	presenceController := _presenceController.NewPresenceController(presenceUsecase)

	routeInit := routes.ControllerList{
		JwtConfig:            configJWT.Init(),
		DepartmentController: *departmentController,
		UserController:       *userController,
		ScheduleController:   *scheduleController,
		PresenceController:   *presenceController,
	}

	routeInit.RouteRegister(e)
	log.Fatal(e.Start(":" + viper.GetString("APP_PORT")))

}
