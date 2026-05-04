package trainings

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 3 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("invalid steps: must be positive")
	}

	duration, err := time.ParseDuration(parts[2])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("invalid duration: must be positive")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps, t.Weight, t.Height, t.Duration,
		)
	default:
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	result := "Тип тренировки: " + t.TrainingType + "\n" +
		"Длительность: " + strconv.FormatFloat(t.Duration.Hours(), 'f', 2, 64) + " ч.\n" +
		"Дистанция: " + strconv.FormatFloat(distance, 'f', 2, 64) + " км.\n" +
		"Скорость: " + strconv.FormatFloat(speed, 'f', 2, 64) + " км/ч\n" +
		"Сожгли калорий: " + strconv.FormatFloat(calories, 'f', 2, 64) + "\n"

	return result, nil
}
