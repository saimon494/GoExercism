// Package twofer return string
package twofer

// ShareWith return string
// depending on the incoming parameter "name"
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}
	return "One for " + name + ", one for me."
}
