package cmd

import (
	"strings"
	"testing"

	"github.com/l-lin/risk-rating/risk"
)

const empty = ""

func Test_generate(t *testing.T) {
	type args struct {
		r *risk.Risk
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Nominal case",
			args{&risk.Risk{
				SkillLevel:            risk.NewValue(5, empty),
				Motive:                risk.NewValue(4, empty),
				Opportunity:           risk.NewValue(4, empty),
				Size:                  risk.NewValue(2, empty),
				EaseOfDiscovery:       risk.NewValue(3, empty),
				EaseOfExploit:         risk.NewValue(5, empty),
				Awareness:             risk.NewValue(6, empty),
				IntrusionDetection:    risk.NewValue(1, empty),
				LossOfConfidentiality: risk.NewValue(7, empty),
				LossOfIntegrity:       risk.NewValue(1, empty),
				LossOfAvailability:    risk.NewValue(1, empty),
				LossOfAccountability:  risk.NewValue(1, empty),
				FinancialDamage:       risk.NewValue(1, empty),
				ReputationDamage:      risk.NewValue(4, empty),
				NonCompliance:         risk.NewValue(5, empty),
				PrivacyViolation:      risk.NewValue(7, empty),
			}},
			strings.Trim(`|===
|Threat agent factors|Skill level|Motive|Opportunity|Size
||5|4|4|2
|Vulnerability factors|Ease of discovery|Ease of exploit|Awareness|Intrusion detection
||3|5|6|1
|Overall likelihood|{set:cellbgcolor:yellow}3.75 (MEDIUM)|{set:cellbgcolor!}||

|{set:cellbgcolor!}||||

|Technical impact|Loss of confidentiality|Loss of integrity|Loss of availability|Loss of accountability
||7|1|1|1
|Business Impact|Financial damage|Reputation damage|Non-compliance|Privacy violation
||1|4|5|7
|Overall impact|{set:cellbgcolor:yellow}3.375 (MEDIUM)|{set:cellbgcolor!}||

|{set:cellbgcolor!}||||
|Overrall Risk Severity|{set:cellbgcolor:orange} MEDIUM|{set:cellbgcolor!}||
|===`, " "),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generate(tt.args.r); got != tt.want {
				t.Errorf("generate() = %v, want %v", got, tt.want)
			}
		})
	}
}
