package logging

import (
	"github.com/sirupsen/logrus"
)

func Init() {

	//выбираем форматтер -- Json
	logrus.SetFormatter(&logrus.JSONFormatter{})

	//под капотом определяет, из какого конкретно файла и какой строки
	//кода был вызван логгер, и добавляет их в JSON:
	logrus.SetReportCaller(true)

	//ставим минимальный порог логов, которые мы видим (Debug самый минимальный)
	//на продакшене обычно начинают с ErrorLevel
	logrus.SetLevel(logrus.DebugLevel)
}
