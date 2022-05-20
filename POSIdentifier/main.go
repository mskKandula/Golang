package main

import (
	"fmt"
	"log"

	"github.com/jdkato/prose/v2"
)

func main() {
	// Create a new document with the default configuration:
	doc, err := prose.NewDocument(`Akbar first attacked Malwa, a state of strategic and economic importance commanding the route through the Vindhya Range to the plateau region of the Deccan (peninsular India) and containing rich agricultural land; it fell to him in 1561.
    Toward the zealously independent Hindu Rajputs (warrior ruling class) inhabiting the rugged hilly Rajputana region, Akbar adopted a policy of conciliation and conquest. Successive Muslim rulers had found the Rajputs dangerous, however weakened by disunity. But in 1562, when Raja Bihari Mal of Amber (now Jaipur), threatened by a succession dispute, offered Akbar his daughter in marriage, Akbar accepted the offer. The Raja acknowledged Akbar’s suzerainty, and his sons prospered in Akbar’s service. Akbar followed the same feudal policy toward the other Rajput chiefs. They were allowed to hold their ancestral territories, provided that they acknowledged Akbar as emperor, paid tribute, supplied troops when required, and concluded a marriage alliance with him. The emperor’s service was also opened to them and their sons, which offered financial rewards as well as honour.`)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate over the doc's tokens:
	for _, tok := range doc.Tokens() {
		if tok.Tag == "NNP" || tok.Tag == "NN" || tok.Tag == "NNS" {
			fmt.Println(tok.Text)
		}
	}

}
