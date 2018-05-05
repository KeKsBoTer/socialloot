package lib

import (
	"math"
	"sort"
	"time"

	"github.com/KeKsBoTer/socialloot/models"
)

// SortByRank a list of post by their rank
// see rank(...) function for closer explanation
func SortByRank(posts []*models.PostMetaData) {
	for _, p := range posts {
		p.Rank = rank(p.Date, p.Votes)
	}
	sort.Slice(posts, func(i int, j int) bool {
		return posts[i].Rank > posts[j].Rank
	})
}

// Computes hot page rank for a post
// see https://medium.com/hacking-and-gonzo/how-reddit-ranking-algorithms-work-ef111e33d0d9
func rank(date time.Time, votes int) float64 {
	t := date.Unix() - 1514764800 // 1.1.2018
	x := float64(votes)
	var y, z float64
	switch {
	case votes > 0:
		y = 1
		z = x
	case votes == 0:
		y = 0
		z = 1
	case votes < 0:
		y = -1
		z = -x
	}
	return math.Log10(z+1)*y + float64(t)/45000.0
}
