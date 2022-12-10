package css

import (
	"bytes"
	"regexp"
	"strings"
)

// We do a similar thing we've done with html
// -split into tokens
// -parse the tokens

// then we do additional steps
// - create the rules
// - append the rules to the stylesheet
// Each rule has a selector and properties

type Rule struct {
	Selector   string
	Properties map[string]string
}

type Stylesheet struct {
	Rules []*Rule
}

// The stylesheet content as string
func (s *Stylesheet) String() string {
	var buffer bytes.Buffer

	for _, rule := range s.Rules {
		buffer.WriteString(rule.Selector)
		buffer.WriteString(" {\n")
		for property, value := range rule.Properties {
			buffer.WriteString("  ")
			buffer.WriteString(property)
			buffer.WriteString(": ")
			buffer.WriteString(value)
			buffer.WriteString("\n")
		}
		buffer.WriteString("}\n")
	}

	return buffer.String()
}

func ParseCSS(css string) (*Stylesheet, error) {
	stylesheet := &Stylesheet{
		Rules: []*Rule{},
	}

	tokens := tokenize(css)
	stylesheet.Rules = parseRules(tokens)

	return stylesheet, nil
}

func tokenize(css string) []string {
	// split into tokens
	regex := regexp.MustCompile("([^{]*{|}[^}]*}|[^{};]*;|[^{}]*)")
	tokens := regex.FindAllString(css, -1)

	return tokens
}

func parseRules(tokens []string) []*Rule {
	rules := []*Rule{}

	for _, token := range tokens {
		token = strings.TrimSpace(token)

		// Check if the token is a rule
		if strings.HasSuffix(token, "{") {
			// Get the selector from the token
			selector := strings.TrimSuffix(token, "{")

			rules = append(rules, &Rule{
				Selector:   selector,
				Properties: map[string]string{},
			})
		} else if strings.Contains(token, ":") {
			// If token is a property and value
			// we use the same method we used in the html parser to get the key value from the attributes
			partsOfProperty := strings.SplitN(token, ":", 2)
			property := strings.TrimSpace(partsOfProperty[0])
			value := strings.TrimSpace(partsOfProperty[1])

			lastRule := rules[len(rules)-1]
			lastRule.Properties[property] = value
		}
	}

	return rules
}
