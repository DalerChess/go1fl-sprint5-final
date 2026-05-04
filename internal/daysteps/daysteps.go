package daysteps

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) error {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return errors.New("invalid data format")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return err
	}

	if steps <= 0 {
		return errors.New("invalid steps: must be positive")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return err
	}

	if duration <= 0 {
		return errors.New("invalid duration: must be positive")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {

	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)

	if err != nil {
		return "", err
	}

	result :=
		"Количество шагов: " + strconv.Itoa(ds.Steps) + ".\n" +
			"Дистанция составила " + strconv.FormatFloat(distance, 'f', 2, 64) + " км.\n" +
			"Вы сожгли " + strconv.FormatFloat(calories, 'f', 2, 64) + " ккал.\n"

	return result, nil
}
