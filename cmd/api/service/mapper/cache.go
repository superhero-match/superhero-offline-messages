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
package mapper

import (
	"github.com/superhero-offline-messages/cmd/api/model"
	cm "github.com/superhero-offline-messages/internal/cache/model"
)

// MapFromCacheModelToAPI maps Cache Message model to API Message model.
func MapFromCacheModelToAPI(messages []*cm.Message) []model.Message {
	result := make([]model.Message, 0)

	for _, msg := range messages {
		result = append(result, model.Message{
			SenderID:  msg.SenderID,
			ReceiverID: msg.ReceiverID,
			Message:    msg.Message,
			CreatedAt:  msg.CreatedAt,
		})
	}

	return result
}
