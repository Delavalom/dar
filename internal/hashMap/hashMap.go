package hashMap

type Tree struct {
	HasSubTree  bool             `json:"hasSubTree"`
	SubTree     map[string]*Tree `json:"subTree"`     // nil if HasSubTree is false
	FileName    string           `json:"fileName"`    // empty if HasSubTree is true
	FileMode    string           `json:"fileMode"`    // empty if HasSubTree is true
	FileContent string           `json:"fileContent"` // empty if HasSubTree is true
}

type HashMap map[string]*Tree

func New() HashMap {
	return make(HashMap)
}
