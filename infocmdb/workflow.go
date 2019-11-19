package infocmdb

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
)

// Input parameters to a workflow.
// These are usually passed encoded as a json string as first process argument.
type WorkflowParams struct {
	Apikey              string `json:"apikey"`
	TriggerType         string `json:"triggerType"`
	WorkflowItemId      int    `json:"workflow_item_id"`
	WorkflowInstanceId  int    `json:"workflow_instance_id"`
	CiId                int    `json:"ciid"`
	CiAttributeId       int    `json:"ciAttributeId"`
	CiRelationId        int    `json:"ciRelationId"`
	CiProjectId         int    `json:"ciProjectId"`
	FileImportHistoryId int    `json:"fileImportHistoryId"`
}

// User defined workflow function that can be passed to `workflow.Run`.
type WorkflowFunc func(params WorkflowParams, cmdb *Client) (err error)

type Workflow struct {
	config string
}

func NewWorkflow() Workflow {
	return Workflow{
		config: "infocmdb.yml",
	}
}

func (w *Workflow) SetConfig(config string) {
	w.config = config
}

// Executes a workflow.
//
// First a infoCMDB client instance is created and the workflow parameters are parsed.
// The workflow parameters are decoded from the first process argument if available.
// Absence of any process argument will not lead to a failure.
//
// If everything is successful the workflow will be executed with the prepared parameters and client.
//
// Any errors that are returned from the workflow function will be logged and lead to a execution failure.
// Additionally the workflow will be marked as failed when something is printed to Stderr during execution.
func (w Workflow) Run(workflowFunc WorkflowFunc) {
	cmdb := NewClient()
	cmdbClientErr := cmdb.LoadConfig(w.config)
	if cmdbClientErr != nil {
		log.Fatalf("Failed to Login: %v", cmdbClientErr)
	}

	params, parseErr := parseParams()
	if parseErr != nil {
		log.Fatal(parseErr)
	}

	workflowErr := workflowFunc(params, cmdb)
	if workflowErr != nil {
		log.Fatal(workflowErr)
	}
}

func parseParams() (params WorkflowParams, err error) {
	if len(os.Args) < 2 {
		log.Fatal("Missing json encoded WorkflowParams as first program argument")
	}

	jsonParam := os.Args[1]
	err = json.Unmarshal([]byte(jsonParam), &params)
	if err != nil {
		return
	}

	return
}

type PreconditionType string

const (
	TYPE_ATTRIBUTE = "attribute"
	TYPE_RELATION  = "relation"
)

type Preconditions []struct {
	Type PreconditionType
	Name string
}

func (w Workflow) TestPreconditions(t *testing.T, preconditions Preconditions) {
	cmdb := NewClient()
	cmdbClientErr := cmdb.LoadConfig(w.config)
	if cmdbClientErr != nil {
		log.Fatalf("Failed to Login: %v", cmdbClientErr)
	}

	for _, precondition := range preconditions {
		testName := fmt.Sprintf("%v \"%v\" exists", precondition.Type, precondition.Name)

		t.Run(testName, func(t *testing.T) {
			switch precondition.Type {
			case TYPE_ATTRIBUTE:
				if _, err := cmdb.GetAttributeIdByAttributeName(precondition.Name); err != nil {
					t.Error(err)
				}
			case TYPE_RELATION:
				if _, err := cmdb.GetCiRelationTypeIdByRelationTypeName(precondition.Name); err != nil {
					t.Error(err)
				}
			}
		})
	}
}