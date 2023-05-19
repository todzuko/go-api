package cache

import "github.com/todzuko/go-api/models"

type QuestsCache interface {
	Set(key string, value *models.Quest)
	Get(key string) *models.Quest
}
