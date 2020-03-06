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
package service

import (
	"fmt"

	"github.com/superhero-offline-messages/cmd/api/model"
	"github.com/superhero-offline-messages/cmd/api/service/mapper"
)

// GetCachedSuggestion fetches suggestion from cache and maps it into result.
func (srv *Service) GetMessages(superheroID string) ([]model.Message, error) {
	messages, err := srv.Cache.GetMessages(fmt.Sprintf(srv.Cache.MessagesKeyFormat, superheroID))
	if err != nil {
		return nil, err
	}

	result := mapper.MapFromCacheModelToAPI(messages)

	return result, nil
}
