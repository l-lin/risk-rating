package cmd

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/l-lin/risk-rating/risk"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new risk rating in ascii doctor table",
	Args:  cobra.ExactArgs(0),
	Run:   executeNew,
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func executeNew(cmd *cobra.Command, args []string) {
	r := askInfo()
	content := generate(r)
	fmt.Println()
	fmt.Println("Asciidoctor table generated:")
	fmt.Println()
	fmt.Println(content)
	fmt.Println()
	clipboard.WriteAll(content)
	fmt.Println("Content copied in clipboard")
}

func generate(r *risk.Risk) string {
	tmpl, err := template.New("risk-rating").Parse(`|===
|Threat agent factors|Skill level|Motive|Opportunity|Size
||{{.SkillLevel}}|{{.Motive}}|{{.Opportunity}}|{{.Size}}
|Vulnerability factors|Ease of discovery|Ease of exploit|Awareness|Intrusion detection
||{{.EaseOfDiscovery}}|{{.EaseOfExploit}}|{{.Awareness}}|{{.IntrusionDetection}}
|Overall likelihood|{{.DisplayOverallLikelihood}}|{set:cellbgcolor!}||

|{set:cellbgcolor!}||||

|Technical impact|Loss of confidentiality|Loss of integrity|Loss of availability|Loss of accountability
||{{.LossOfConfidentiality}}|{{.LossOfIntegrity}}|{{.LossOfAvailability}}|{{.LossOfAccountability}}
|Business Impact|Financial damage|Reputation damage|Non-compliance|Privacy violation
||{{.FinancialDamage}}|{{.ReputationDamage}}|{{.NonCompliance}}|{{.PrivacyViolation}}
|Overall impact|{{.DisplayOverallImpact}}|{set:cellbgcolor!}||

|{set:cellbgcolor!}||||
|Overrall Risk Severity|{{.DisplaySeverity}}|{set:cellbgcolor!}||
|===`)
	if err != nil {
		log.Fatalln("Could not generate risk rating", err)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, r)
	if err != nil {
		log.Fatalln("Error when execute templating", err)
	}
	return strings.Trim(buf.String(), " ")
}

func askInfo() *risk.Risk {
	r := &risk.Risk{}
	fmt.Println("Threat Agent Factors")
	r.SkillLevel = ask("Skill level - How technically skilled is this group of threat agents? ", []*risk.Value{
		risk.NewValue(1, "no technical skill"),
		risk.NewValue(3, "some technical skills"),
		risk.NewValue(5, "advanced computer user"),
		risk.NewValue(6, "network and programming skills"),
		risk.NewValue(9, "security penetration skills"),
	})
	r.Motive = ask("Motive - How motivated is this group of threat agents to find and exploit this vulnerability? ", []*risk.Value{
		risk.NewValue(1, "low or no reward"),
		risk.NewValue(4, "possible reward"),
		risk.NewValue(9, "high reward"),
	})
	r.Opportunity = ask("Opportunity - What resources and opportunities are required for this group of threat agents to find and exploit this vulnerability? ", []*risk.Value{
		risk.NewValue(0, "full access or expensive resources required"),
		risk.NewValue(4, "special access or resources required"),
		risk.NewValue(7, "some access or resources required"),
		risk.NewValue(9, "no access or resources required"),
	})
	r.Size = ask("Size - How large is this group of threat agents? ", []*risk.Value{
		risk.NewValue(2, "developers"),
		risk.NewValue(2, "system administrators"),
		risk.NewValue(4, "intranet users"),
		risk.NewValue(5, "partners"),
		risk.NewValue(6, "authenticated users"),
		risk.NewValue(9, "anonymous internet users"),
	})
	fmt.Println("Vulnerability Factors")
	r.EaseOfDiscovery = ask("Ease of discovery - How easy is it for this group of threat agents to discover this vulnerability? ", []*risk.Value{
		risk.NewValue(1, "practically impossible"),
		risk.NewValue(3, "difficult"),
		risk.NewValue(7, "easy"),
		risk.NewValue(9, "automated tools available"),
	})
	r.EaseOfExploit = ask("Ease of exploit - How easy is it for this group of threat agents to actually exploit this vulnerability? ", []*risk.Value{
		risk.NewValue(1, "theoretical"),
		risk.NewValue(3, "difficult"),
		risk.NewValue(5, "easy"),
		risk.NewValue(9, "automated tools available"),
	})
	r.Awareness = ask("Awareness - How well known is this vulnerability to this group of threat agents? ", []*risk.Value{
		risk.NewValue(1, "unknown"),
		risk.NewValue(4, "hidden"),
		risk.NewValue(6, "obvious"),
		risk.NewValue(9, "public knowledge"),
	})
	r.IntrusionDetection = ask("Intrusion detection - How likely is an exploit to be detected? ", []*risk.Value{
		risk.NewValue(1, "active detection in application"),
		risk.NewValue(3, "logged and reviewed"),
		risk.NewValue(8, "logged without review"),
		risk.NewValue(9, "not logged"),
	})
	fmt.Println("Technical Impact Factors")
	r.LossOfConfidentiality = ask("Loss of confidentiality - How much data could be disclosed and how sensitive is it? ", []*risk.Value{
		risk.NewValue(2, "minimal non-sensitive data disclosed"),
		risk.NewValue(6, "minimal critical data disclosed"),
		risk.NewValue(6, "extensive non-sensitive data disclosed"),
		risk.NewValue(7, "extensive critical data disclosed"),
		risk.NewValue(9, "all data disclosed"),
	})
	r.LossOfIntegrity = ask("Loss of integrity - How much data could be corrupted and how damaged is it? ", []*risk.Value{
		risk.NewValue(1, "minimal slightly corrupt data"),
		risk.NewValue(3, "minimal seriously corrupt data"),
		risk.NewValue(5, "extensive slightly corrupt data"),
		risk.NewValue(7, "extensive seriously corrupt data"),
		risk.NewValue(9, "all data totally corrupt"),
	})
	r.LossOfAvailability = ask("Loss of availability - How much service could be lost and how vital is it? ", []*risk.Value{
		risk.NewValue(1, "minimal secondary services interrupted"),
		risk.NewValue(5, "minimal primary services interrupted"),
		risk.NewValue(5, "extensive secondary services interrupted"),
		risk.NewValue(7, "extensive primary services interrupted"),
		risk.NewValue(9, "all services completely lost"),
	})
	r.LossOfAccountability = ask("Loss of accountability - Are the threat agents' actions traceable to an individual? ", []*risk.Value{
		risk.NewValue(1, "fully traceable"),
		risk.NewValue(7, "possibly traceable"),
		risk.NewValue(9, "completely anonymous"),
	})
	fmt.Println("Business Impact Factors")
	r.FinancialDamage = ask("Financial damage - How much financial damage will result from an exploit? ", []*risk.Value{
		risk.NewValue(1, "less than the cost to fix the vulnerability"),
		risk.NewValue(3, "minor effect on annual profit"),
		risk.NewValue(7, "significant effect on annual profit"),
		risk.NewValue(9, "bankruptcy"),
	})
	r.ReputationDamage = ask("Reputation damage - Would an exploit result in reputation damage that would harm the business? ", []*risk.Value{
		risk.NewValue(1, "minimal damage"),
		risk.NewValue(4, "loss of major accounts"),
		risk.NewValue(5, "loss of goodwill"),
		risk.NewValue(9, "brand damage"),
	})
	r.NonCompliance = ask("Non compliance - How much exposure does non-compliance introduce? ", []*risk.Value{
		risk.NewValue(2, "minor violation"),
		risk.NewValue(5, "clear violation"),
		risk.NewValue(7, "high profile violation"),
	})
	r.PrivacyViolation = ask("Privacy violation - How much personally identifiable information could be disclosed? ", []*risk.Value{
		risk.NewValue(3, "one individual"),
		risk.NewValue(5, "hundreds of people"),
		risk.NewValue(7, "thousands of people"),
		risk.NewValue(9, "millions of people"),
	})
	return r
}

func ask(label string, items []*risk.Value) *risk.Value {
	prompt := promptui.Select{
		Label:             label,
		Items:             items,
		StartInSearchMode: true,
		Searcher:          buildSearcher(items),
	}

	_, result, err := prompt.Run()

	if err != nil {
		log.Fatalln(err)
	}
	return risk.Parse(result)
}

func buildSearcher(items []*risk.Value) func(input string, index int) bool {
	return func(input string, index int) bool {
		item := items[index]
		return strings.Contains(item.String(), input)
	}
}
