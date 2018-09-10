package delay109

import (
	"regexp"
	"strconv"
	"time"
)

func parseDelayDuration(htmlBytes []byte) (time.Duration, error) {
	r := regexp.MustCompile(`<td class="left">　：　(\d+)分程度</td>`)
	found := r.FindSubmatch(htmlBytes)
	if len(found) < 2 {
		return 0, nil
	}

	duration, err := strconv.Atoi(string(found[1]))
	if err != nil {
		return 0, err
	}
	return time.Duration(duration) * time.Minute, nil
}
