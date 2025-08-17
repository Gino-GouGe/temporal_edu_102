package translation

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.temporal.io/sdk/testsuite"
)

func TestSuccessfulCompleteFrenchTranslation(t *testing.T) {
	s := testsuite.WorkflowTestSuite{}

	env := s.NewTestWorkflowEnvironment()
	env.RegisterActivity(TranslateTerm)

	workflowInput := TranslationWorkflowInput{
		Name:         "Pierre",
		LanguageCode: "fr",
	}

	env.ExecuteWorkflow(SayHelloGoodbye, workflowInput)

	env.OnActivity(TranslateTerm, mock.Anything, TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: workflowInput.LanguageCode,
	}).Return(TranslationActivityInput{
		Term:         "Pierre",
		LanguageCode: workflowInput.LanguageCode,
	}, nil)
	env.OnActivity(TranslateTerm, mock.Anything, TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: workflowInput.LanguageCode,
	}).Return(TranslationActivityInput{
		Term:         "Pierre",
		LanguageCode: workflowInput.LanguageCode,
	}, nil)

	// TODO: Assert that Workflow Execution completed
	assert.True(t, env.IsWorkflowCompleted())
	assert.NoError(t, env.GetWorkflowError())

	var result TranslationWorkflowOutput
	err := env.GetWorkflowResult(&result)

	helloMessage := fmt.Sprintf("Bonjour, %s", workflowInput.Name)
	goodbyeMessage := fmt.Sprintf("Au revoir, %s", workflowInput.Name)
	// TODO: Assert that the HelloMessage field in the
	//       result is: Bonjour, Pierre
	assert.NoError(t, err)
	assert.Equal(t, helloMessage, result.HelloMessage)

	// TODO: Assert that the GoodbyeMessage field in the
	//       result is: Au revoir, Pierre
	assert.Equal(t, goodbyeMessage, result.GoodbyeMessage)
}
