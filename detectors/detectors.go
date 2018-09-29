package detectors

type Detector interface {
	Detect(string) (string, error)
}
