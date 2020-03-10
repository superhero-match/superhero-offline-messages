/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package cache

import (
	"github.com/go-redis/redis"
	"github.com/superhero-match/superhero-offline-messages/internal/cache/model"
)

// GetMessages fetches suggestions from cache.
func (c *Cache) GetMessages(key string) ([]*model.Message, error) {
	res, err := c.Redis.SMembers(key).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, nil
	}

	messages := make([]*model.Message, 0)

	for _, msg := range res {
		var message model.Message

		if err := message.UnmarshalBinary([]byte(msg)); err != nil {
			return nil, err
		}

		messages = append(messages, &message)
	}


	return messages, nil
}