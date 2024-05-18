package cacher

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"
)

func cacheUpdate() {
	/*
		// здесь настраиваем названия таблиц, которые будут сохранятся в кэше
		cacheTables := []string{"users"}

		// сначала загружаем весь кэш в БД
		for table := range cacheTables {
			// для этого получаем весь кэш
			cacheMap, err := cache.ReadCache(cacheTables[table])
			if err != nil {
				return
			}
			// затем загружаем в БД
			db.PullData(cacheTables[table], cacheMap)
		}
	*/
}

func Init(interval string) {
	intervalStr := fmt.Sprintf("0 */%s * * * *", interval)

	c := cron.New()
	_, err := c.AddFunc(intervalStr, cacheUpdate)
	if err != nil {
		log.Fatalf("ошибка при инициализации кэша: %v", err.Error())
		return
	}

	c.Start()
}
