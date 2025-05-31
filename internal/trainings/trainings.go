package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

const (
	Running = "Бег"
	Walking = "Ходьба"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse метод парсит строку с данными формата "3456,Ходьба,3h00m" и записывает данные в соответствующие поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")

	if len(parts) != 3 {
		return fmt.Errorf("expected 3 params (count, duration), got %d", len(parts))
	}

	steps, err := strconv.Atoi(parts[0])

	if err != nil {
		return fmt.Errorf("invalid count format: %w", err)
	}

	if steps <= 0 {
		return fmt.Errorf("steps must be positive")
	}

	duration, err := time.ParseDuration(parts[2])

	if err != nil {
		return fmt.Errorf("invalid duration format: %w", err)
	}

	if duration <= 0 {
		return fmt.Errorf("duration must be positive")
	}

	t.Steps = steps
	t.TrainingType = parts[1]
	t.Duration = duration
	return nil
}

// ActionInfo метод формирует и возвращает строку с данными о тренировке, исходя из того, какой тип тренировки был передан.
func (t Training) ActionInfo() (string, error) {
	var calories float64
	var err error

	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	switch t.TrainingType {
	case Running:
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case Walking:
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	default:
		err = fmt.Errorf("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	pattern := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n", t.TrainingType, t.Duration.Hours(), dist, speed, calories)

	return pattern, nil
}
