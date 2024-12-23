package Mosdef

import (
	"log"
	"strconv"
	"strings"
)

func GetRooms(lines []string) (string, string, int, map[string][]string) {
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
		room := strings.Split(line, " ") 
		tunnel := strings.Split(line, "-")
		if line == "" || (len(room) != 3 && len(tunnel) != 2 && line[0] != '#' && i != 0)  {

			log.Fatal("ERROR: invalid data format. invalid syntax")
		}
		if i == 0 {
			number, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal("ERROR: invalid data format. invalid ants number")
			}
			if number <= 0 || number > 10000 {
				log.Fatal("ERROR: invalid data format. invalid number of ants")
			}
			antsnumber = number
		}
		if len(room) == 3 && line[0] != '#' {
			if line[0]=='L' || line[0]=='l'{
				log.Fatal("ERROR: invalid data format. invalid room name")
				
			}
			if _,err:= strconv.Atoi(room[1]);err!=nil  {
				log.Fatal("ERROR: invalid data format. invalid coordinates")
			}
			if _,err:= strconv.Atoi(room[2]);err!=nil {
				log.Fatal("ERROR: invalid data format. invalid coordinates")
			}
			if _, exist := rooms[room[0]]; exist {
				log.Fatal("ERROR: invalid data format. room already exists")
			}
			rooms[room[0]] =true
			if _, exist := coords[room[1]+"+"+room[2]]; exist {
				log.Fatal("ERROR: invalid data format. repeated coordinates\n" + line + "\n" + strconv.Itoa(i))
			}

			coords[room[1]+"+"+room[2]] =true
		}
		if line == "##start" && !startroom {
			if i+1 != len(lines) && len(strings.Split(lines[i+1], " ")) == 3 {
				startroom = true
				start = strings.Split(lines[i+1], " ")[0]
			} else {
				log.Fatal("ERROR: invalid data format.there is no start room")
			}
		} else if line == "##start" && startroom {
			log.Fatal("ERROR: invalid data format.there is too many start rooms")
		}
		if line == "##end" && !endroom {
			if i+1 != len(lines) && len(strings.Split(lines[i+1], " ")) == 3 {
				endroom = true
				end = strings.Split(lines[i+1], " ")[0]
			} else {
				log.Fatal("ERROR: invalid data format.there is no end room")
			}
		} else if line == "##end" && endroom {
			log.Fatal("ERROR: invalid data format.there is too many end rooms")
		}

		if len(tunnel) == 2 && line[0] != '#' {
			
			if _, exist := rooms[tunnel[0]]; !exist {
				log.Fatal("ERROR: invalid data format. unexisting room")
			}
			if _, exist := rooms[tunnel[1]]; !exist {
				log.Fatal("ERROR: invalid data format. unexisting room")
			}
			if _, exist := linkscomb[tunnel[1]+"-"+tunnel[0]]; exist {
				log.Fatal("ERROR: invalid data format. repeated link")
			}
			if _, exist := linkscomb[tunnel[0]+"-"+tunnel[1]]; exist {
				log.Fatal("ERROR: invalid data format. repeated link")
			}
			linkscomb[line] = true
			links[tunnel[0]] = append(links[tunnel[0]], tunnel[1])
			links[tunnel[1]] = append(links[tunnel[1]], tunnel[0])
		}
	}
	if len(links) != len(rooms) {
		log.Fatal("ERROR: invalid data format. room without links room")
	}

	return start, end, antsnumber, links
}
