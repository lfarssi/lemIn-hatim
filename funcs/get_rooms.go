package Mosdef

import (
	"strconv"
	"strings"
)

func GetRooms(lines []string) (string, string, int, map[string][]string, string) {
	antsnumber := 0
	startroom := false
	endroom := false
	start := ""
	end := ""
	rooms := make(map[string]bool)
	coords := make(map[string]bool)
	links := make(map[string][]string)
	linkscomb := make(map[string]bool)
	for i, line := range lines {
		var next []string
		room := strings.Split(line, " ")
		tunnel := strings.Split(line, "-")
		if i+1 != len(lines) {
			next = strings.Split(lines[i+1], " ")
		}
		if line == "" || (len(room) != 3 && len(tunnel) != 2 && line[0] != '#' && i != 0) {
			return "", "", 0, nil, "ERROR: invalid data format. invalid syntax"
		}
		if i == 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				return "", "", 0, nil, "ERROR: invalid data format. invalid ants number"
			}
			if number <= 0 || number > 10000 {
				return "", "", 0, nil, "ERROR: invalid data format. invalid ants number"
			}
			antsnumber = number
		}
		if len(room) == 3 && line[0] != '#' {
			if line[0] == 'L' || line[0] == 'l' {
				return "", "", 0, nil, "ERROR: invalid data format. invalid room name"

			}
			if _, err := strconv.Atoi(room[1]); err != nil {
				return "", "", 0, nil, "ERROR: invalid data format. invalid coordinates"
			}
			if _, err := strconv.Atoi(room[2]); err != nil {
				return "", "", 0, nil, "ERROR: invalid data format. invalid coordinates"
			}
			if _, exist := rooms[room[0]]; exist {
				return "", "", 0, nil, "ERROR: invalid data format. room already exists"
			}
			rooms[room[0]] = true
			if _, exist := coords[room[1]+"+"+room[2]]; exist {
				return "", "", 0, nil, "ERROR: invalid data format. repeated coordinates\n" + line + "\n" + strconv.Itoa(i)

			}
			coords[room[1]+"+"+room[2]] = true
		}
		if line == "##start" && !startroom {
			if i+1 != len(lines) && len(next) == 3 {
				startroom = true
				start = next[0]
			} else {
				return "", "", 0, nil, "ERROR: invalid data format. there is no start room"
			}
		} else if line == "##start" && startroom {
			return "", "", 0, nil, "ERROR: invalid data format. there is too many start rooms"
		}
		if line == "##end" && !endroom {
			if i+1 != len(lines) && len(next) == 3 {
				endroom = true
				end = next[0]
			} else {
				return "", "", 0, nil, "ERROR: invalid data format. there is no end room"
			}
		} else if line == "##end" && endroom {
			return "", "", 0, nil, "ERROR: invalid data format. there is too many end rooms"
		}

		if len(tunnel) == 2 && line[0] != '#' {
			if tunnel[0] == tunnel[1] {
				return "", "", 0, nil, "ERROR: invalid data format. tunnel to the same room"
			}
			if _, exist := rooms[tunnel[0]]; !exist {
				return "", "", 0, nil, "ERROR: invalid data format. unexisting room"
			}
			if _, exist := rooms[tunnel[1]]; !exist {
				return "", "", 0, nil, "ERROR: invalid data format. unexisting room"
			}
			if _, exist := linkscomb[tunnel[1]+"-"+tunnel[0]]; exist {
				return "", "", 0, nil, "ERROR: invalid data format. repeat links"
			}
			if _, exist := linkscomb[tunnel[0]+"-"+tunnel[1]]; exist {
				return "", "", 0, nil, "ERROR: invalid data format. repeat links"
			}
			linkscomb[line] = true
			links[tunnel[0]] = append(links[tunnel[0]], tunnel[1])
			links[tunnel[1]] = append(links[tunnel[1]], tunnel[0])
		}
	}
	if len(links) != len(rooms) {
		return "", "", 0, nil, "ERROR: invalid data format. rooms without links rooms"
	}

	return start, end, antsnumber, links, ""
}
