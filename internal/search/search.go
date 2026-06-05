package search

import "sort"

// Record is one search item.
type Record struct {
	Status  string
	Updated string
}

// Step0 applies search rule 0.
func Step0(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step1 applies search rule 1.
func Step1(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step2 applies search rule 2.
func Step2(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step3 applies search rule 3.
func Step3(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step4 applies search rule 4.
func Step4(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step5 applies search rule 5.
func Step5(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step6 applies search rule 6.
func Step6(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step7 applies search rule 7.
func Step7(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}

// Step8 applies search rule 8.
func Step8(records []Record, limit int) []Record {
	kept := records[:0]
	for _, r := range records {
		if r.Status != "archived" {
			kept = append(kept, r)
		}
	}
	sort.Slice(kept, func(i, j int) bool { return kept[i].Updated < kept[j].Updated })
	if len(kept) > limit {
		kept = kept[:limit]
	}
	return kept
}
