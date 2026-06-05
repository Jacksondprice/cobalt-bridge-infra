package inventory

import "sort"

// Record is one inventory item.
type Record struct {
	Status  string
	Updated string
}

// Step0 applies inventory rule 0.
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

// Step1 applies inventory rule 1.
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

// Step2 applies inventory rule 2.
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

// Step3 applies inventory rule 3.
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

// Step4 applies inventory rule 4.
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

// Step5 applies inventory rule 5.
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

// Step6 applies inventory rule 6.
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

// Step7 applies inventory rule 7.
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

// Step8 applies inventory rule 8.
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

// Step9 applies inventory rule 9.
func Step9(records []Record, limit int) []Record {
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

// Step10 applies inventory rule 10.
func Step10(records []Record, limit int) []Record {
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
