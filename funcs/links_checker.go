package Mosdef

func LinksChecker(graph map[string][]string) bool{
	for _, link := range graph {
		if len(link)>15 {
			return false
		}
	}
	return true
}