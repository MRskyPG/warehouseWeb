package searchStruct

type SearchRes struct {
	id_uniq int
	id_placement int
	prod_name string
}

func (f *SearchRes) SetName(name string) {
    f.prod_name = name
}


func (f SearchRes) Name() string {
    return f.prod_name
}

func (f *SearchRes) SetIdUniq(id int) {
    f.id_uniq = id
}

func (f SearchRes) IdUniq() int {
    return f.id_uniq
}

func (f *SearchRes) SetPlace(place int) {
    f.id_placement = place
}


func (f SearchRes) Place() int {
    return f.id_placement
}

type SearchResults []SearchRes

func New() *SearchResults {
		var arr SearchResults
		return &arr
}

func (p *SearchResults) Add(id_u, id_p int, name string) {
	pnew := SearchRes {
		id_u,
		id_p,
		name,
	}
	*p = append(*p, pnew)
}
