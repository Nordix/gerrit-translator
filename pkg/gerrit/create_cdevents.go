/*
Copyright (C) 2024 Nordix Foundation.
For a full list of individual contributors, please see the commit history.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
SPDX-License-Identifier: Apache-2.0
*/

package gerrit

import (
	mimic "github.com/cdevents/mimic"
	sdk "github.com/cdevents/sdk-go/pkg/api"
)

// map If I want to build a custom Field, instead of direct field-field map??
// cdevents:"context_source,url" - doesn't work - map multiple fileds to set??
// map fields to a common Object, Ex: Repository *Reference

// ------------

// copygen elimantes only setting the fields manually, other code can be still exist/refactored.
// no chance of introducing bulk event types at this stage of gerrit for translating into CDEvents

func (projectCreated *ProjectCreated) RepositoryCreatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepositoryCreatedEvent with Mimic %s\n", projectCreated.ProjectName)
	cdEvent, err := sdk.NewRepositoryCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent RepositoryCreatedEvent %s\n", err)
		return "", err
	}
	customVisitor := mimic.NewTagVisitor(mimic.NewTags().AddTag("cdevents"))
	copier := mimic.NewCopier(
		mimic.WithSrcVisitor(customVisitor),
		mimic.WithDstVisitor(customVisitor),
		// Overrides the default singleton copier which allows us to use the generic
		// Copy function
		mimic.WithInstance(),
	)
	err = copier.Copy(projectCreated, cdEvent)

	if err != nil {
		Log().Error("Mimic copier: Error copying data from ProjectCreated gerrit event into RepositoryCreated CDEvent %s\n", err)
		return "", err
	} else {
		Log().Info("Mimic copier : Success cdEvent.GetSource() %s\n", cdEvent.GetSource())
		Log().Info("Mimic copier : Success cdEvent.GetSubject() type %s\n", cdEvent.GetSubject().GetSubjectType())
	}

	//ToRepositoryCreatedEvent(cdEvent, projectCreated)
	//cdEvent.SetSource(projectCreated.RepoURL + "sdewsrf") // cdevents:"context_source"
	// cdEvent.SetSubjectName(projectCreated.ProjectName) // cdevents:"name"
	// cdEvent.SetSubjectId(projectCreated.HeadName) // cdevents:"subject_id"
	//cdEvent.SetSubjectUrl(projectCreated.RepoURL) // cdevents:"context_source,url" - doesn't work - How to map this along with context_source??

	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepositoryCreated CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (projectHeadUpdated *ProjectHeadUpdated) RepositoryModifiedCDEvent() (string, error) {
	Log().Info("Creating CDEvent RepositoryModifiedEvent")
	cdEvent, err := sdk.NewRepositoryModifiedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent RepositoryModified %s\n", err)
		return "", err
	}
	ToRepositoryModifiedEvent(cdEvent, projectHeadUpdated)
	//cdEvent.SetSource(projectHeadUpdated.RepoURL) // cdevents:"context_source"
	//cdEvent.SetSubjectName(projectHeadUpdated.ProjectName) // cdevents:"name"
	//cdEvent.SetSubjectId(projectHeadUpdated.NewHead) // cdevents:"subject_id"
	cdEvent.SetSubjectUrl(projectHeadUpdated.RepoURL) // cdevents:"url"- how to map this along with context_source??
	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating RepositoryModified CDEvent as Json string %s\n", err)
		return "", err
	}

	return cdEventStr, nil
}

func (refUpdated *RefUpdated) BranchCreatedCDEvent() (string, error) {
	Log().Info("Creating CDEvent BranchCreatedEvent")
	cdEvent, err := sdk.NewBranchCreatedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent BranchCreatedEvent %s\n", err)
		return "", err
	}
	//cdEvent.SetSource(refUpdated.RepoURL)                                          // cdevents:"context_source"
	//cdEvent.SetSubjectId(refUpdated.RefUpdate.NewRev)                              // cdevents:"subject_id"
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: refUpdated.RefUpdate.RefName}) //map field to Object - Repository *Reference `json:"repository,omitempty" cdevents:"repository"`"
	//cdEvent.SetSubjectSource(refUpdated.RefUpdate.Project)                         // Source == SubjectSource / cdevents:"subject_source"

	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating BranchCreated CDEvent as Json string %s\n", err)
		return "", err
	}
	return cdEventStr, nil
}

func (refUpdated *RefUpdated) BranchDeletedCDEvent() (string, error) {
	Log().Info("Creating CDEvent BranchDeletedEvent")
	cdEvent, err := sdk.NewBranchDeletedEvent()
	if err != nil {
		Log().Error("Error creating CDEvent BranchDeletedEvent %s\n", err)
		return "", err
	}
	//cdEvent.SetSource(refUpdated.RepoURL)                                          // cdevents:"context_source"
	cdEvent.SetSubjectId(refUpdated.RefUpdate.OldRev)                              // cdevents:"subject_id"
	cdEvent.SetSubjectRepository(&sdk.Reference{Id: refUpdated.RefUpdate.RefName}) // Repository *Reference `json:"repository,omitempty" cdevents:"repository"`"
	//cdEvent.SetSubjectSource(refUpdated.RefUpdate.Project)                         // Source == SubjectSource / cdevents:"subject_source"

	cdEventStr, err := sdk.AsJsonString(cdEvent)
	if err != nil {
		Log().Error("Error creating BranchDeleted CDEvent as Json string %s\n", err)
		return "", err
	}
	return cdEventStr, nil
}
