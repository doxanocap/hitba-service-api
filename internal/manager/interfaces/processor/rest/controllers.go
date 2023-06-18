package rest

import "app/internal/manager/interfaces/processor/rest/controllers"

type IControllersManager interface {
	Storage() controllers.IStorageManager
}
