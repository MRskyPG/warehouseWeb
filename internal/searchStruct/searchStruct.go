package searchStruct

type SearchRes struct {
	id_uniq int
	id_placement int
	prod_name string
	client_name string
	client_surname string
	date string
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

func (f SearchRes) Place() int {
	return f.id_placement
}
func (f *SearchRes) SetPlace(place int) {
	f.id_placement = place
}


func (f SearchRes) ClName() string {
    return f.client_name
}

func (f *SearchRes) SetClName(name string) {
    f.client_name = name
}


func (f SearchRes) Surname() string {
    return f.client_surname
}

func (f *SearchRes) SetSurname(surname string) {
    f.client_surname = surname
}


func (f SearchRes) Date() string {
    return f.date
}
func (f *SearchRes) SetDate(date string) {
    f.date = date
}



type SearchResults []SearchRes

func New() *SearchResults {
		var arr SearchResults
		return &arr
}

func (p *SearchResults) Add(id_u, id_p int, name string, cl_name string, cl_surname string, date string) {
	pnew := SearchRes {
		id_u,
		id_p,
		name,
		cl_name,
		cl_surname,
		date,
	}
	*p = append(*p, pnew)
}
