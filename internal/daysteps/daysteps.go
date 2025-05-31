package daysteps

import (
	"fmt"
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

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 2 {
		return fmt.Errorf("expected 2 params (count, duration), got %d", len(parts))
	}

	stepsCount, err := strconv.Atoi(parts[0])

	if err != nil {
		return fmt.Errorf("invalid count format: %w", err)
	}
	if stepsCount <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	duration, err := time.ParseDuration(parts[1])

	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}
	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	ds.Steps, ds.Duration = stepsCount, duration
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)

	if err != nil {
		return "", err
	}

	pattern := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n", ds.Steps, distance, calories)

	return pattern, err
}
