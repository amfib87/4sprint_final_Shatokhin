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
	// TODO: реализовать функцию
	sl := strings.Split(data, ",")
	var tm time.Duration

	if len(sl) != 2 {
		return 0, tm, fmt.Errorf("lenght of slice != 2")
	}

	steps, err := strconv.Atoi(sl[0])
	if err != nil {
		return 0, tm, err
	}

	if steps <= 0 {
		return 0, tm, fmt.Errorf("numbers of steps <= 0")
	}

	tim, err := time.ParseDuration(sl[1])
	if err != nil {
		return 0, tm, err
	}

	if tim <= 0 {
		return 0, tm, fmt.Errorf("wrong duration <= 0")
	}

	return steps, tim, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, tm, err := parsePackage(data)
	if err != nil {
		log.Println(err.Error())
		return ""
	}

	if steps <= 0 {
		return ""
	}

	dist := float64(steps) * stepLength / float64(mInKm)

	calors, err := spentcalories.WalkingSpentCalories(steps, weight, height, tm)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, calors)
}
