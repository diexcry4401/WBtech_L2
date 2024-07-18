package pattern

import "fmt"

// Service — интерфейс, который определяет метод выполнения операции и метод установки следующего обработчика в цепочке.
type Service interface {
	Execute(d *Data)
	SetNext(service Service)
}

// Data — структура, содержащая данные и флаги, указывающие на стадии обработки.
type Data struct {
	GetSource    bool
	UpdateSource bool
}

// Device — структура, представляющая источник данных.
type Device struct {
	Name string
	Next Service
}

// Execute — метод получения данных из устройства.
func (dev *Device) Execute(d *Data) {
	if d.GetSource {
		fmt.Printf("Data from device [%s] already get.\n", dev.Name)
		dev.Next.Execute(d)
	} else {
		fmt.Printf("Get data from device [%s]\n", dev.Name)
		d.GetSource = true
		dev.Next.Execute(d)
	}
}

// SetNext — метод установки следующего обработчика в цепочке.
func (dev *Device) SetNext(service Service) {
	dev.Next = service
}

// UpdateDataService — структура, представляющая службу обновления данных.
type UpdateDataService struct {
	Name string
	Next Service
}

// Execute — метод обновления данных.
func (upd *UpdateDataService) Execute(d *Data) {
	if d.UpdateSource {
		fmt.Printf("Data from device [%s] already update.\n", upd.Name)
		upd.Next.Execute(d)
	} else {
		fmt.Printf("Update data from device [%s]\n", upd.Name)
		d.UpdateSource = true
		upd.Next.Execute(d)
	}
}

// SetNext — метод установки следующего обработчика в цепочке.
func (upd *UpdateDataService) SetNext(service Service) {
	upd.Next = service
}

// SaveDataService — структура, представляющая службу сохранения данных.
type SaveDataService struct {
	Next Service
}

// Execute — метод сохранения данных.
func (save *SaveDataService) Execute(d *Data) {
	if !d.UpdateSource {
		fmt.Println("Data not update")
	} else {
		fmt.Println("Data save")
	}
}

// SetNext — метод установки следующего обработчика в цепочке (не используется, т.к. это последний обработчик).
func (save *SaveDataService) SetNext(service Service) {
	save.Next = service
}

// NewDevice — конструктор для создания нового устройства.
func NewDevice(name string) *Device {
	return &Device{
		Name: name,
	}
}

// NewUpdateSvc — конструктор для создания новой службы обновления данных.
func NewUpdateSvc(name string) *UpdateDataService {
	return &UpdateDataService{
		Name: name,
	}
}

// NewSaveDataService — конструктор для создания новой службы сохранения данных.
func NewSaveDataService() *SaveDataService {
	return &SaveDataService{}
}

// NewData — конструктор для создания новой структуры данных.
func NewData() *Data {
	return &Data{}
}

// func main() {
// 	// Создаем обработчиков.
// 	device := NewDevice("Device1")
// 	updateService := NewUpdateSvc("UpdateService1")
// 	saveService := NewSaveDataService()

// 	// Строим цепочку обязанностей.
// 	device.SetNext(updateService)
// 	updateService.SetNext(saveService)

// 	// Создаем данные.
// 	data := NewData()

// 	// Передаем данные первому обработчику в цепочке.
// 	device.Execute(data)
// }
