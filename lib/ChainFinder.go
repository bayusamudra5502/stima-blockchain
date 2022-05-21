package lib

func FindFirstDiff(chain1, chain2 Chain) int64 {
	if chain1.length != chain2.length {
		panic("Ukuran chain harus sama")
	}

	l := int64(0)
	r := chain1.length

	for l < r {
		center := (l + r) / 2
		diff := CheckDifference(chain1, chain2, center)

		if diff == 0 {
			return center
		} else if diff > 0 {
			l = center + 1
		} else {
			r = center
		}
	}

	return -1
}

func CheckDifference(Chain1, Chain2 Chain, idx int64) int64 {
	it1 := Chain1.GetChainItem(idx)
	it2 := Chain2.GetChainItem(idx)

	if it1.GetBase64Hash() != it2.GetBase64Hash() {
		if idx > 0 {
			last1 := Chain1.GetChainItem(idx - 1)
			last2 := Chain2.GetChainItem(idx - 1)

			if last1.GetBase64Hash() == last2.GetBase64Hash() {
				return 0
			} else {
				return -1
			}
		} else {
			return 0
		}
	} else {
		return 1
	}
}

func FindFirstDiffBF(chain1, chain2 Chain) int64 {
	if chain1.length != chain2.length {
		panic("Ukuran chain harus sama")
	}

	for i := int64(0); i < chain1.length; i++ {
		if CheckDifference(chain1, chain2, i) == 0 {
			return i
		}
	}

	return -1
}
