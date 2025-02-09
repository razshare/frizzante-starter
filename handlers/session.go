package handlers

var memory = map[string]map[string]any{}

func Session(id string) (
	get func(key string, defaultValue any) (value any),
	set func(key string, value any),
	unset func(key string),
	destroy func(),
) {
	get = func(key string, defaultValue any) (value any) {
		localMemory, localMemoryExists := memory[id]

		if !localMemoryExists {
			localMemory = map[string]any{}
			memory[id] = localMemory
		}

		sessionItem, ok := localMemory[key]
		if !ok {
			localMemory[key] = defaultValue
			return localMemory[key]
		}

		return sessionItem
	}

	set = func(key string, value any) {
		localMemory, localMemoryExists := memory[id]

		if !localMemoryExists {
			localMemory = map[string]any{}
			memory[id] = localMemory
		}

		localMemory[key] = value
	}

	unset = func(key string) {
		localMemory, localMemoryExists := memory[id]

		if !localMemoryExists {
			localMemory = map[string]any{}
			memory[id] = localMemory
		}

		delete(memory, key)
	}

	destroy = func() {
		localMemory, localMemoryExists := memory[id]

		if !localMemoryExists {
			localMemory = map[string]any{}
			memory[id] = localMemory
		}

		delete(localMemory, id)
	}
	return

}
