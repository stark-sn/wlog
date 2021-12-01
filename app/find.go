package app

import (
	"io/fs"
	"log"
	"os/user"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"s-stark.net/code/wlog/persistence"
)

func getAllActivityTitles(text string) []string {
	now := time.Now()
	user, err := user.Current()

	if err != nil {
		log.Fatalf("Failed to get current user, %v", err)
	}

	activities := make(map[string]bool)

	filepath.WalkDir(path.Join(user.HomeDir, dir), func(path string, d fs.DirEntry, err error) error {
		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		week, err := persistence.Read(path)

		if err != nil {
			return err
		}

		for _, day := range week.Days {
			for _, act := range day.GetActivities(now) {
				if strings.HasPrefix(act.Title, text) {
					activities[act.Title] = true
				}
			}
		}

		return nil
	})

	var titles []string
	for title, _ := range activities {
		titles = append(titles, title)
	}
	sort.Strings(titles)

	return titles
}
