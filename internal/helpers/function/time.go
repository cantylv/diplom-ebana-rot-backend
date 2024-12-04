package function

import "time"

// FormatTime возвращает отформатированное время | `Mon Jan 2 15:04:05 MST 2006` настраиваемый шаблон
func FormatTime(t time.Time) string {
	return t.Format("02.01.2006 15:04:05 UTC-07")
}
