package logic

func GetPlayNextCardAction(id string) Action {
	return GetAction(Action_play_next_card, id)
}

func GetDisposeNextCardAction(id string) Action {
	return GetAction(Action_dispose_next_card, id)
}
