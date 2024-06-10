/* Specify the name of the generated file's package. */
package copygen

import (
	"github.com/cdevents/gerrit-translator/pkg/gerrit"
)

/* Copygen defines the functions that will be generated. */
type Copygen interface {
	// tag .* cdevents
	// deepcopy .*
	SrcToDest(*gerrit.Source) *gerrit.Destination
	//ToRepositoryCreatedEvent(*gerrit.ProjectCreated) *sdk.RepositoryCreatedEvent
}

// func Itoa(i int) string {
// 	return strconv.Itoa(i)
// }
