package gerrit

import (
	"testing"

	mimic "github.com/cdevents/mimic"
)

// func TestRepositoryCreatedCDEvent(t *testing.T) {
// 	projectCreated := ProjectCreated{
// 		ProjectName:  "TestRepo12",
// 		HeadName:     "refs/repo",
// 		CommonFields: CommonFields{Type: "project-created", EventCreatedOn: 234234234, RepoURL: "http://gerrit.est.tech"},
// 	}
// 	cdEvent, err := projectCreated.RepositoryCreatedCDEvent()
// 	Log().Info("Translated project-created gerrit event into dev.cdevents.repository.created: ", cdEvent)
// 	if err != nil {
// 		t.Errorf("Expected RepositoryCreated CDEvent to be successful.")
// 	}
// }

type Custom struct {
	CopyThisField      int `mimic:Field1`
	CopyThisOtherField int `mimic:Child.Field2`
}
type NoTagDst struct {
	Field1 int
	Child  NoTagChild
}

type NoTagChild struct {
	Field2 int
}

func TestWithFieldsOption(t *testing.T) {

	tagVisitor := mimic.NewTagVisitor(mimic.NewTags().AddTag("mimic"), mimic.TagVisitorOptions.WithFlattenKeys(true))
	fieldVisitor := mimic.NewFieldVisitor(mimic.FieldVisitorOptions.WithFlattenKeys(true))
	copier := mimic.NewCopier(
		mimic.WithSrcVisitor(tagVisitor),
		mimic.WithDstVisitor(fieldVisitor),
		mimic.WithInstance(),
	)

	dst := &NoTagDst{}
	src := &Custom{
		CopyThisField:      1337,
		CopyThisOtherField: 1010,
	}
	err := copier.Copy(src, dst)
	//err := mimic.CopyTo(src, dst)

	if err != nil {
		Log().Error("Copied to  NoTagDst: ", err)
		t.Errorf("Expected copy to NoTagDst to be successful.")
	} else {
		Log().Info("Copied to  NoTagDst: ", dst.Field1)
	}
}
