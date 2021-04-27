package nlp

const VERTEX_NOT_FOUND = -1
const RE_WNP = "^[NARV]"

const (
	UKB_RELATION_FILE = 1 + iota
	UKB_REX_WNPOS
	UKB_PR_PARAMS
)

type CSRKB struct {
	maxIterations int
	threshold     float64
	damping       float64
	vertexIndex   map[string]int
	outCoef       []float64
	firstEdge     []int
	numEdges      []int
	edges         []int
	numVertices   int
}

type IntPair struct {
	first  int
	second int
}

type IntPairsArray []IntPair

func (a IntPairsArray) Less(i, j int) bool {
	return a[i].first < a[j].first || (a[i].first == a[j].first && a[i].second < a[j].second)
}

func (a IntPairsArray) Len() int { return len(a) }

func (a IntPairsArray) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
