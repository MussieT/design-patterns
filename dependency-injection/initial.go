package dependencyinjection

import "fmt"

// GreetingService is dependent on GreetingFormatter.

// Define formatter interface
type GreetingFormatter interface {
	FormatGreeting(name string) string
}

// Create a concrete implementation of formatter interface
type DefaultGreetingFormatter struct {}

func (f *DefaultGreetingFormatter) FormatGreeting(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}

// Define the GreetingService that depends on the GreetingFormatter:
type GreetingService struct {
	formatter GreetingFormatter
}

func NewGreetingService(formatter GreetingFormatter) *GreetingService {
	return &GreetingService{
		formatter: formatter,
	}
}

func (s *GreetingService) Greet(name string) string {
	return s.formatter.FormatGreeting(name)
}

// type fr struct {}

func InjectAndCheck() {
	// create the concrete implementation of the Greetings formatter
	formatter := &DefaultGreetingFormatter{}

	// type fr struct {}

	greetingService := NewGreetingService(formatter)

	greeting := greetingService.Greet("John")

	fmt.Println(greeting)
}
