package gerrit

import (
	"testing"

	mimic "github.com/cdevents/mimic"
)

func TestRepositoryCreatedCDEvent(t *testing.T) {
	projectCreated := ProjectCreated{
		ProjectName:  "TestRepo12",
		HeadName:     "refs/repo",
		CommonFields: CommonFields{Type: "project-created", EventCreatedOn: 234234234, RepoURL: "http://gerrit.est.tech"},
	}
	cdEvent, err := projectCreated.RepositoryCreatedCDEvent()
	if err != nil {
		t.Errorf("Expected RepositoryCreated CDEvent to be successful.")
	}
	Log().Info("Translated project-created gerrit event into dev.cdevents.repository.created: ", cdEvent)
}

type Custom struct {
	CopyThisField      int `mimic:"Field1,FiealdA"`
	CopyThisOtherField int `mimic:"Child.Field2"`
}
type NoTagDst struct {
	Field1 int
	Child  NoTagChild
}

type NoTagChild struct {
	Field2 int
}

func TestWithFieldsOption(t *testing.T) {

	tagVisitor := mimic.NewTagVisitor(mimic.TagVisitorOptions.WithFlattenKeys(true))
	fieldVisitor := mimic.NewFieldVisitor(mimic.FieldVisitorOptions.WithFlattenKeys(true))
	mimic.NewCopier(
		mimic.WithSrcVisitor(tagVisitor),
		mimic.WithDstVisitor(fieldVisitor),
		mimic.WithInstance(),
	)

	// dst, err := mimic.Copy[*NoTagDst](&Custom{
	// 	CopyThisField:      1337,
	// 	CopyThisOtherField: 1010,
	// })

	dst := &NoTagDst{}
	src := &Custom{
		CopyThisField:      1337,
		CopyThisOtherField: 1010,
	}
	//err := copier.Copy(src, dst)
	err := mimic.CopyTo(src, dst)
	//dst, err := mimic.Copy[NoTagDst](src)

	if err != nil {
		Log().Error("Copied to  NoTagDst: ", err)
		t.Errorf("Expected copy to NoTagDst to be successful.")
	} else {
		Log().Info("Copied to  NoTagDst Child: ", dst.Child.Field2)
		Log().Info("Copied to  NoTagDst: ", dst.Field1)
	}
}
