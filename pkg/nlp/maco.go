package nlp

type MacoOptions struct {
}

func NewMacoOptions(lang string) *MacoOptions {
	return &MacoOptions{}
}

func (this *MacoOptions) SetDataFiles(usr, pun, dic, aff, comp, loc, nps, qty, prb string) {

}
