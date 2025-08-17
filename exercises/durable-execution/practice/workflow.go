package translation

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SayHelloGoodbye(ctx workflow.Context, input TranslationWorkflowInput) (TranslationWorkflowOutput, error) {
	// TODO define the Workflow logger here
	logger := workflow.GetLogger(ctx)
	// TODO Log, at the Info level, when the Workflow function is invoked
	//      and be sure to include the name passed as input
	logger.Info("SayHelloGoodbye Workflow started", "name", input.Name)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 45,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// TODO Log, at the Debug level, a message about the Activity to be executed,
	//      be sure to include the language code passed as input
	helloInput := TranslationActivityInput{
		Term:         "Hello",
		LanguageCode: input.LanguageCode,
	}
	logger.Debug("Executing TranslateTerm Activity helloInput", "LanguageCode", helloInput.LanguageCode)
	var helloResult TranslationActivityOutput
	err := workflow.ExecuteActivity(ctx, TranslateTerm, helloInput).Get(ctx, &helloResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	helloMessage := fmt.Sprintf("%s, %s", helloResult.Translation, input.Name)

	// TODO: (Part C): log a message at the Debug level and then start a Timer for 10 seconds
	logger.Debug("Sleeping between translation calls")
	workflow.Sleep(ctx, time.Second*10)

	goodbyeInput := TranslationActivityInput{
		Term:         "Goodbye",
		LanguageCode: input.LanguageCode,
	}
	logger.Debug("Executing TranslateTerm Activity goodbyeInput", "LanguageCode", goodbyeInput.LanguageCode)
	var goodbyeResult TranslationActivityOutput
	// TODO Log, at the Debug level, a message about the Activity to be executed,
	//      be sure to include the language code passed as input
	err = workflow.ExecuteActivity(ctx, TranslateTerm, goodbyeInput).Get(ctx, &goodbyeResult)
	if err != nil {
		return TranslationWorkflowOutput{}, err
	}
	goodbyeMessage := fmt.Sprintf("%s, %s", goodbyeResult.Translation, input.Name)

	output := TranslationWorkflowOutput{
		HelloMessage:   helloMessage,
		GoodbyeMessage: goodbyeMessage,
	}

	return output, nil
}
