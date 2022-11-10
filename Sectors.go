package percbar

// sectors for Bar
type sectors []*sector

// Len satisfy Sorter interfase
func (sect sectors) Len() int {
	return len(sect)
}

// Less satisfy Sorter interfase
func (sect sectors) Less(i, j int) bool {
	return sect[i].count > sect[j].count

}

// Swap satisfy Sorter interfase
func (sect sectors) Swap(i, j int) {
	sect[i], sect[j] = sect[j], sect[i]
}
