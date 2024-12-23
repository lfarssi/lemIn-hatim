package Mosdef

func LinksChecker(graph map[string][]string) bool{
	if len(graph) > 30 {
		return false
	}
	for _, link := range graph {
		if len(link) > 15 {
			return false
		}
	}
	return true
}