package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"openai_connect/functions"
	"os"

	"github.com/openai/openai-go/v2"
	"github.com/openai/openai-go/v2/option"
)

func OpenAICompletions() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(
		option.WithAPIKey(apiKey),
	)

	ctx := context.Background()

	question := "Como está o clima em São Paulo?"

	print("> ")
	println(question)

	params := openai.ChatCompletionNewParams{
		Messages: openai.F([]openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(question),
		}),
		Tools: openai.F([]openai.ChatCompletionToolParam{
			{
				Type: openai.F(openai.ChatCompletionToolTypeFunction),
				Function: openai.F(openai.FunctionDefinitionParam{
					Name:        openai.String("get_weather"),
					Description: openai.String("Get weather at the given location"),
					Parameters: openai.F(openai.FunctionParameters{
						"type": "object",
						"properties": map[string]interface{}{
							"location": map[string]string{
								"type": "string",
							},
						},
						"required": []string{"location"},
					}),
				}),
			},
		}),
		Seed:  openai.Int(0),
		Model: openai.F(openai.ChatModelGPT4o),
	}

	// Make initial chat completion request
	completion, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}

	toolCalls := completion.Choices[0].Message.ToolCalls

	// Abort early if there are no tool calls
	if len(toolCalls) == 0 {
		return completion.Choices[0].Message.Content, nil
	}

	// If there is a was a function call, continue the conversation
	params.Messages.Value = append(params.Messages.Value, completion.Choices[0].Message)
	if len(toolCalls) > 0 {
		for _, toolCall := range toolCalls {
			if toolCall.Function.Name == "get_weather" {
				// Extract the location from the function call arguments
				var args map[string]interface{}
				if err := json.Unmarshal([]byte(toolCall.Function.Arguments), &args); err != nil {
					panic(err)
				}
				location := args["location"].(string)

				// Simulate getting weather data
				weatherData := functions.GetWeather(location)

				// Print the weather data
				fmt.Printf("Weather in %s: %s\n", location, weatherData)

				params.Messages.Value = append(params.Messages.Value, openai.ToolMessage(toolCall.ID, weatherData))
			}
		}

	}

	completion, err = client.Chat.Completions.New(ctx, params)
	if err != nil {
		panic(err)
	}

	return completion.Choices[0].Message.Content, nil
}
