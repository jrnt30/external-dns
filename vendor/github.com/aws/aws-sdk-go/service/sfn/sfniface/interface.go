// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package sfniface provides an interface to enable mocking the AWS Step Functions service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package sfniface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sfn"
)

// SFNAPI provides an interface to enable mocking the
// sfn.SFN service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Step Functions.
//    func myFunc(svc sfniface.SFNAPI) bool {
//        // Make svc.CreateActivity request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := sfn.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSFNClient struct {
//        sfniface.SFNAPI
//    }
//    func (m *mockSFNClient) CreateActivity(input *sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSFNClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type SFNAPI interface {
	CreateActivityRequest(*sfn.CreateActivityInput) (*request.Request, *sfn.CreateActivityOutput)

	CreateActivity(*sfn.CreateActivityInput) (*sfn.CreateActivityOutput, error)

	CreateStateMachineRequest(*sfn.CreateStateMachineInput) (*request.Request, *sfn.CreateStateMachineOutput)

	CreateStateMachine(*sfn.CreateStateMachineInput) (*sfn.CreateStateMachineOutput, error)

	DeleteActivityRequest(*sfn.DeleteActivityInput) (*request.Request, *sfn.DeleteActivityOutput)

	DeleteActivity(*sfn.DeleteActivityInput) (*sfn.DeleteActivityOutput, error)

	DeleteStateMachineRequest(*sfn.DeleteStateMachineInput) (*request.Request, *sfn.DeleteStateMachineOutput)

	DeleteStateMachine(*sfn.DeleteStateMachineInput) (*sfn.DeleteStateMachineOutput, error)

	DescribeActivityRequest(*sfn.DescribeActivityInput) (*request.Request, *sfn.DescribeActivityOutput)

	DescribeActivity(*sfn.DescribeActivityInput) (*sfn.DescribeActivityOutput, error)

	DescribeExecutionRequest(*sfn.DescribeExecutionInput) (*request.Request, *sfn.DescribeExecutionOutput)

	DescribeExecution(*sfn.DescribeExecutionInput) (*sfn.DescribeExecutionOutput, error)

	DescribeStateMachineRequest(*sfn.DescribeStateMachineInput) (*request.Request, *sfn.DescribeStateMachineOutput)

	DescribeStateMachine(*sfn.DescribeStateMachineInput) (*sfn.DescribeStateMachineOutput, error)

	GetActivityTaskRequest(*sfn.GetActivityTaskInput) (*request.Request, *sfn.GetActivityTaskOutput)

	GetActivityTask(*sfn.GetActivityTaskInput) (*sfn.GetActivityTaskOutput, error)

	GetExecutionHistoryRequest(*sfn.GetExecutionHistoryInput) (*request.Request, *sfn.GetExecutionHistoryOutput)

	GetExecutionHistory(*sfn.GetExecutionHistoryInput) (*sfn.GetExecutionHistoryOutput, error)

	GetExecutionHistoryPages(*sfn.GetExecutionHistoryInput, func(*sfn.GetExecutionHistoryOutput, bool) bool) error

	ListActivitiesRequest(*sfn.ListActivitiesInput) (*request.Request, *sfn.ListActivitiesOutput)

	ListActivities(*sfn.ListActivitiesInput) (*sfn.ListActivitiesOutput, error)

	ListActivitiesPages(*sfn.ListActivitiesInput, func(*sfn.ListActivitiesOutput, bool) bool) error

	ListExecutionsRequest(*sfn.ListExecutionsInput) (*request.Request, *sfn.ListExecutionsOutput)

	ListExecutions(*sfn.ListExecutionsInput) (*sfn.ListExecutionsOutput, error)

	ListExecutionsPages(*sfn.ListExecutionsInput, func(*sfn.ListExecutionsOutput, bool) bool) error

	ListStateMachinesRequest(*sfn.ListStateMachinesInput) (*request.Request, *sfn.ListStateMachinesOutput)

	ListStateMachines(*sfn.ListStateMachinesInput) (*sfn.ListStateMachinesOutput, error)

	ListStateMachinesPages(*sfn.ListStateMachinesInput, func(*sfn.ListStateMachinesOutput, bool) bool) error

	SendTaskFailureRequest(*sfn.SendTaskFailureInput) (*request.Request, *sfn.SendTaskFailureOutput)

	SendTaskFailure(*sfn.SendTaskFailureInput) (*sfn.SendTaskFailureOutput, error)

	SendTaskHeartbeatRequest(*sfn.SendTaskHeartbeatInput) (*request.Request, *sfn.SendTaskHeartbeatOutput)

	SendTaskHeartbeat(*sfn.SendTaskHeartbeatInput) (*sfn.SendTaskHeartbeatOutput, error)

	SendTaskSuccessRequest(*sfn.SendTaskSuccessInput) (*request.Request, *sfn.SendTaskSuccessOutput)

	SendTaskSuccess(*sfn.SendTaskSuccessInput) (*sfn.SendTaskSuccessOutput, error)

	StartExecutionRequest(*sfn.StartExecutionInput) (*request.Request, *sfn.StartExecutionOutput)

	StartExecution(*sfn.StartExecutionInput) (*sfn.StartExecutionOutput, error)

	StopExecutionRequest(*sfn.StopExecutionInput) (*request.Request, *sfn.StopExecutionOutput)

	StopExecution(*sfn.StopExecutionInput) (*sfn.StopExecutionOutput, error)
}

var _ SFNAPI = (*sfn.SFN)(nil)