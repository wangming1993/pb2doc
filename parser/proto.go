package parser

func ReadNote(lines []string) (string, int) {
	total := len(lines)
	var (
		i       int
		step    int
		comment string
	)
	for {
		if i >= total {
			break
		}
		line := lines[i]
		// Must done otherwise dead foreach
		i++
		if IsSingleComment(line) {
			step += 1
			comment += line + "\n"
		} else if StartMultiComment(line) {
			for {
				step += 1
				comment += line + "\n"
				line = lines[i]

				if EndMultiComment(line) {
					step++
					comment += line + "\n"
					break
				}
				i++
			}
		} else {
			return comment, step
		}
	}
	return comment, step
}
