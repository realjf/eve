package nlp

const SENSES_DUP_ANALYSIS = 1

type Senses struct {
	duplicate bool
	semdb     *SemanticDB
}

func NewSenses(wsdFile string) *Senses {
	this := Senses{
		semdb: NewSemanticDB(wsdFile),
	}

	cfg := NewConfigFile(true, "")
	cfg.AddSection("DuplicateAnalysis", SENSES_DUP_ANALYSIS)

	if !cfg.Open(wsdFile) {
		LOG.Panic("Error opening file " + wsdFile)
	}

	line := ""
	for cfg.GetContentLine(&line) {
		items := Split(line, " ")
		switch cfg.GetSection() {
		case SENSES_DUP_ANALYSIS:
			{
				key := items[0]
				if key == "yes" {
					this.duplicate = true
				}
				break
			}
		default:
			break
		}
	}

	LOG.Trace("Analyzer succesfully created")

	return &this
}

func (this *Senses) Analyze(sentence *Sentence) {
	// @todo
}
