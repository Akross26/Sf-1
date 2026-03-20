package daysteps

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	sl := strings.Split(data, ",")

	if len(sl) != 2 {
		return 0, 0, fmt.Errorf("Неверные вводные данные")
	}

	step, err := strconv.Atoi(sl[0])
	if err != nil {
		return 0, 0, err
	}

	if step <= 0 {
		return 0, 0, fmt.Errorf("Неверное кол-во шагов")
	}

	t, err := time.ParseDuration(sl[1])
	if err != nil {
		return 0, 0, err
	}

	if t <= 0 {
		return 0, 0, fmt.Errorf("Неверная длительность")
	}

	return step, t, nil
}
func DayActionInfo(data string, weight, height float64) string {
	step, t, err := parsePackage(data)
	if err != nil {
		log.Println(err)
		return ""
	}

	if step <= 0 {
		return ""
	}

	dist := float64(step) * stepLength
	km := dist / mInKm
	calories, err := spentcalories.WalkingSpentCalories(step, weight, height, t)

	if err != nil {
		log.Println(err)
		return ""
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", step, km, calories)

	return result
}
