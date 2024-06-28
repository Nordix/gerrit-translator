package gerrit

import (
	"testing"
)

func TestRepositoryCreatedCDEvent(t *testing.T) {
	projectCreated := ProjectCreated{
		ProjectName:  "TestRepo12",
		HeadName:     "refs/repo",
		CommonFields: CommonFields{Type: "project-created", EventCreatedOn: 234234234, RepoURL: "http://gerrit.est.tech"},
	}
	cdEvent, err := projectCreated.RepositoryCreatedCDEvent()
	Log().Info("Translated project-created gerrit event into dev.cdevents.repository.created: ", cdEvent)
	if err != nil {
		t.Errorf("Expected RepositoryCreated CDEvent to be successful.")
	}
}
