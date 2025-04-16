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
	steps := 0

	if len(sl) != 2 {
		return 0, tm, fmt.Errorf("Длина слайса != 2")
	}

	for ind, val := range sl {
		if ind == 0 {

			st, err := strconv.Atoi(val)
			if err != nil {
				return 0, tm, err
			}

			if st <= 0 {
				return 0, tm, fmt.Errorf("Кол-во шагов меньше или равно 0")
			}

			steps = st

		} else if ind == 1 {
			tim, err := time.ParseDuration(val)
			if err != nil {
				return 0, tm, err
			}

			if tim <= 0 {
				return 0, tm, fmt.Errorf("неверная продолжительность <= 0")
			}
			tm = tim
		}
	}

	return steps, tm, nil
}

func DayActionInfo(data string, weight, height float64) string {
	// TODO: реализовать функцию

	steps, tm, err := parsePackage(data)
	if err != nil {
		//fmt.Println(err.Error())
		log.Println(err.Error())
		return ""
	}

	if steps <= 0 {
		return ""
	}

	dist := float64(steps) * stepLength
	dist /= float64(mInKm)

	calors, err := spentcalories.WalkingSpentCalories(steps, weight, height, tm)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, dist, calors)
}
