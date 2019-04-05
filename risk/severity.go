package risk

type Severity interface {
	Display() string
}

func NewSeverity(rateLikelihood, rateImpact *Rate) Severity {
	if rateLikelihood.IsLow() {
		if rateImpact.IsLow() {
			return &NoteSeverity{}
		} else if rateImpact.IsMedium() {
			return &LowSeverity{}
		}
		return &MediumSeverity{}
	} else if rateLikelihood.IsMedium() {
		if rateImpact.IsLow() {
			return &LowSeverity{}
		} else if rateImpact.IsMedium() {
			return &MediumSeverity{}
		}
		return &HighSeverity{}
	}
	if rateImpact.IsLow() {
		return &MediumSeverity{}
	} else if rateImpact.IsMedium() {
		return &HighSeverity{}
	}
	return &CriticalSeverity{}
}

type NoteSeverity struct {
}

func (s *NoteSeverity) Display() string {
	return "{set:cellbgcolor:lightgreen} NOTE"
}

type LowSeverity struct {
}

func (s *LowSeverity) Display() string {
	return "{set:cellbgcolor:yellow} LOW"
}

type MediumSeverity struct {
}

func (s *MediumSeverity) Display() string {
	return "{set:cellbgcolor:orange} MEDIUM"
}

type HighSeverity struct {
}

func (s *HighSeverity) Display() string {
	return "{set:cellbgcolor:red} HIGH"
}

type CriticalSeverity struct {
}

func (s *CriticalSeverity) Display() string {
	return "{set:cellbgcolor:pink} CRITICAL"
}
