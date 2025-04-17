package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

const (
	run  = "БЕГ"
	walk = "ХОДЬБА"
)

func parseTraining(data string) (int, string, time.Duration, error) {
	// TODO: реализовать функцию
	sl := strings.Split(data, ",")
	var kind string

	if len(sl) != 3 {
		return 0, " ", 0, fmt.Errorf("lenght of slice != 3")
	}

	steps, err := strconv.Atoi(sl[0])
	if err != nil {
		return 0, " ", 0, err
	}
	if steps <= 0 {
		return 0, " ", 0, fmt.Errorf("wrong number of steps <= 0")
	}

	kind = sl[1]

	dur, err := time.ParseDuration(sl[2])
	if err != nil {
		return 0, " ", 0, err
	}

	if dur <= 0 {
		return 0, " ", 0, fmt.Errorf("wrong duration <= 0")
	}

	return steps, kind, dur, nil
}

func distance(steps int, height float64) float64 {
	// TODO: реализовать функцию
	lenghtStep := height * stepLengthCoefficient
	dist := float64(steps) * lenghtStep
	return dist / float64(mInKm)

}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	// TODO: реализовать функцию

	if duration <= 0 {
		return 0
	}
	dist := distance(steps, height)

	return dist / duration.Hours()
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	// TODO: реализовать функцию
	var calors float64
	steps, kind, tm, err := parseTraining(data)
	if err != nil {
		return "", err
	}

	kindUp := strings.ToUpper(kind)

	switch kindUp {
	case run:
		cal, err := RunningSpentCalories(steps, weight, height, tm)
		if err != nil {
			return "", err
		}
		calors = cal

	case walk:
		cal, err := WalkingSpentCalories(steps, weight, height, tm)
		if err != nil {
			return "", err
		}
		calors = cal

	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	dist := distance(steps, height)
	speed := meanSpeed(steps, height, tm)

	text := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
		kind, tm.Hours(), dist, speed, calors)
	return text, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("number of steps <= 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("wrong weight <= 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("wrong height <= 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("wrong duration <= 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	calor := (duration.Minutes() * weight * meanSpeed) / float64(minInH)
	return calor, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	// TODO: реализовать функцию
	if steps <= 0 {
		return 0, fmt.Errorf("number of steps <= 0")
	}

	if weight <= 0 {
		return 0, fmt.Errorf("wrong weight <= 0")
	}

	if height <= 0 {
		return 0, fmt.Errorf("wrong height <= 0")
	}

	if duration <= 0 {
		return 0, fmt.Errorf("wrong duration <= 0")
	}

	meanSpeed := meanSpeed(steps, height, duration)

	calor := (duration.Minutes() * weight * meanSpeed) / float64(minInH) * walkingCaloriesCoefficient
	return calor, nil
}
