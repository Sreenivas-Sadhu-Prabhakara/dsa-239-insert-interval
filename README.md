# Insert Interval

> NeetCode 150 · Intervals · Blind 75 — three phases over the sorted list: copy the before, swallow the overlapping into the newcomer, copy the after. Go, stdlib only, tested.

## Problem

Insert one interval into a sorted, non-overlapping list so it stays sorted and non-overlapping — merging as needed. The input's sortedness is the gift that distinguishes this from Merge Intervals (dsa-008), which must sort first.

## Approach

Because the list is sorted, the newcomer's interactions are *contiguous*: some prefix ends strictly before it (copied untouched), a middle run overlaps it (each widens the newcomer: min of starts, max of ends), and the suffix starts strictly after (copied untouched). Three sequential loops sharing one index — no sort, no second pass, and the overlap test `intervals[i][0] <= newEnd` paired with the before-test `intervals[i][1] < newStart` makes touching intervals merge (pinned by a test). Each phase's condition is exactly the negation boundary of the previous, which is why no interval is dropped or doubled.

**Complexity:** O(n) time, O(n) output.

## Run

    go test ./...   # verify
    go run .        # demo

---

Part of a 275-repo DSA series in Go — Striver's SDE Sheet (dsa-001–191) plus the NeetCode 150 / Blind 75 additions (dsa-192–275), stdlib only, every solution tested. Browse: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
