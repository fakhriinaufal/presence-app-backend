package main

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"presence-app-backend/app/routes"
	_departmentUsecase "presence-app-backend/business/departments"
	_scheduleUsecase "presence-app-backend/business/schedules"
	_userUsecase "presence-app-backend/business/users"
	_departmentController "presence-app-backend/controllers/departments"
	_scheduleController "presence-app-backend/controllers/schedules"
	_userController "presence-app-backend/controllers/users"
	_departmentRepo "presence-app-backend/drivers/databases/departments"
	"presence-app-backend/drivers/databases/mysql"
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
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: viper.GetString("DB_USERNAME"),
		DB_Password: viper.GetString("DB_PASSWORD"),
		DB_Host:     viper.GetString("DB_HOST"),
		DB_Port:     viper.GetString("DB_PORT"),
		DB_Database: viper.GetString("DB_NAME"),
	}

	conn := configDB.InitDB()
	DbMigrate(conn)

	e := echo.New()
	timeoutContext := time.Duration(2)

	departmentRepository := _departmentRepo.NewMysqlDepartmentRepository(conn)
	departmentUsecase := _departmentUsecase.NewDepartmentUsecase(departmentRepository, timeoutContext)
	departmentController := _departmentController.NewDepartmentController(departmentUsecase)

	// _userRepo route
	userRepository := _userRepo.NewMysqlUserRepository(conn)
	userUsecase := _userUsecase.NewUserUsecase(userRepository, departmentRepository,timeoutContext)
	userController := _userController.NewUserController(userUsecase)


	// schedule route
	scheduleRepository := _scheduleRepo.NewMysqlScheduleRepository(conn)
	scheduleUsecase := _scheduleUsecase.NewScheduleUsecase(scheduleRepository, departmentRepository, timeoutContext)
	scheduleController := _scheduleController.NewScheduleController(scheduleUsecase)

	routeInit := routes.ControllerList{
		DepartmentController: *departmentController,
		UserController: *userController,
		ScheduleController: *scheduleController,
	}

	routeInit.RouteRegister(e)
	log.Fatal(e.Start(":" + viper.GetString("APP_PORT")))

}
