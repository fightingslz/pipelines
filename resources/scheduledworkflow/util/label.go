// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package util

import (
	"fmt"
	"github.com/argoproj/argo/workflow/common"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/selection"
	"strconv"
)

func GetRequirementForCompletedWorkflowOrFatal(completed bool) *labels.Requirement {
	operator := selection.NotEquals
	if completed == true {
		operator = selection.Equals
	}
	req, err := labels.NewRequirement(common.LabelKeyCompleted, operator,
		[]string{"true"})
	if err != nil {
		log.Fatalf("Error while creating requirement: %s", err)
	}
	return req
}

func GetRequirementForScheduleNameOrFatal(scheduleName string) *labels.Requirement {
	req, err := labels.NewRequirement(LabelKeyWorkflowName, selection.Equals, []string{scheduleName})
	if err != nil {
		log.Fatalf("Error while creating requirement: %s", err)
	}
	return req
}

func GetRequirementForMinIndexOrFatal(minIndex int64) *labels.Requirement {
	req, err := labels.NewRequirement(LabelKeyWorkflowIndex, selection.GreaterThan,
		[]string{FormatInt64ForLabel(minIndex)})
	if err != nil {
		log.Fatalf("Error while creating requirement: %s", err)
	}
	return req
}

func FormatInt64ForLabel(epoch int64) string {
	return fmt.Sprintf("%d", epoch)
}

func RetrieveInt64FromLabel(epoch string) (int64, error) {
	return strconv.ParseInt(epoch, 10, 64)
}