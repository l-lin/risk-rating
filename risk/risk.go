package risk

import "fmt"

// Risk is the representation of the OWASP risk rating model
// See https://www.owasp.org/index.php/OWASP_Risk_Rating_Methodology
type Risk struct {
	SkillLevel, Motive, Opportunity, Size                                            *Value
	EaseOfDiscovery, EaseOfExploit, Awareness, IntrusionDetection                    *Value
	LossOfConfidentiality, LossOfIntegrity, LossOfAvailability, LossOfAccountability *Value
	FinancialDamage, ReputationDamage, NonCompliance, PrivacyViolation               *Value
}

// ComputeOverallLikelihood is the average value of the Likelihood
func (r *Risk) ComputeOverallLikelihood() float64 {
	return computeAverage([]*Value{r.SkillLevel, r.Motive, r.Opportunity, r.Size, r.EaseOfDiscovery, r.EaseOfExploit, r.Awareness, r.IntrusionDetection})
}

// ComputeOverallImpact is the average value of the impact
func (r *Risk) ComputeOverallImpact() float64 {
	return computeAverage([]*Value{r.LossOfConfidentiality, r.LossOfIntegrity, r.LossOfAvailability, r.LossOfAccountability, r.FinancialDamage, r.ReputationDamage, r.NonCompliance, r.PrivacyViolation})
}

// DisplayOverallLikelihood is the average value of the Likelihood with the correct background color
func (r *Risk) DisplayOverallLikelihood() string {
	return displayAverage(r.ComputeOverallLikelihood())
}

// DisplayOverallImpact is the average value of the impact with the correct background color
func (r *Risk) DisplayOverallImpact() string {
	return displayAverage(r.ComputeOverallImpact())
}

// DisplaySeverity is the overall risk severity with the correct background color
func (r *Risk) DisplaySeverity() string {
	rateLikelihood := &Rate{r.ComputeOverallLikelihood()}
	rateImpact := &Rate{r.ComputeOverallImpact()}
	s := NewSeverity(rateLikelihood, rateImpact)
	return s.Display()
}

func computeAverage(values []*Value) float64 {
	total := 0
	for _, v := range values {
		total += v.Weight
	}
	return float64(total) / float64(len(values))
}

func displayAverage(average float64) string {
	rate := &Rate{average}
	return fmt.Sprintf("{set:cellbgcolor:%s}%v (%s)", rate.Color(), average, rate.Level())
}
