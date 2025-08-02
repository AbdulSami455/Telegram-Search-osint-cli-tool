package scraper

// ConvertChannelsToInterface converts []Channel to []interface{}
func ConvertChannelsToInterface(channels []Channel) []interface{} {
	result := make([]interface{}, len(channels))
	for i, ch := range channels {
		result[i] = map[string]interface{}{
			"Title":       ch.Title,
			"URL":         ch.URL,
			"Description": ch.Description,
		}
	}
	return result
}

// ConvertGroupsToInterface converts []Group to []interface{}
func ConvertGroupsToInterface(groups []Group) []interface{} {
	result := make([]interface{}, len(groups))
	for i, group := range groups {
		result[i] = map[string]interface{}{
			"Title":       group.Title,
			"URL":         group.URL,
			"Description": group.Description,
		}
	}
	return result
}

// ConvertBotsToInterface converts []Bot to []interface{}
func ConvertBotsToInterface(bots []Bot) []interface{} {
	result := make([]interface{}, len(bots))
	for i, bot := range bots {
		result[i] = map[string]interface{}{
			"Title":       bot.Title,
			"URL":         bot.URL,
			"Description": bot.Description,
		}
	}
	return result
}

// ConvertMessagesToInterface converts []Message to []interface{}
func ConvertMessagesToInterface(messages []Message) []interface{} {
	result := make([]interface{}, len(messages))
	for i, msg := range messages {
		result[i] = map[string]interface{}{
			"Title":       msg.Title,
			"URL":         msg.URL,
			"Description": msg.Description,
		}
	}
	return result
}
