package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type CalculatorService struct {
	window     []int
	mutex      sync.Mutex
	httpClient *http.Client
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{
		window: []int{},
		httpClient: &http.Client{
			Timeout: RequestTimeout,
		},
	}
}

func (s *CalculatorService) HandleNumberRequest(c *fiber.Ctx) error {
	numberID := NumberType(c.Params("numberid"))
	if !numberID.IsValid() {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "Invalid number type. Use 'p' for prime, 'f' for Fibonacci, 'e' for even, or 'r' for random"})
	}

	numbers, err := s.fetchNumbers(numberID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": fmt.Sprintf("Failed to fetch numbers: %v", err)})
	}

	s.mutex.Lock()
	prevState := make([]int, len(s.window))
	copy(prevState, s.window)

	for _, num := range numbers {
		if !contains(s.window, num) {
			s.addToWindow(num)
		}
	}

	currState := make([]int, len(s.window))
	copy(currState, s.window)
	s.mutex.Unlock()

	var avg float64
	if len(currState) > 0 {
		sum := 0
		for _, num := range currState {
			sum += num
		}
		avg = float64(sum) / float64(len(currState))
		avg = math.Round(avg*100) / 100
	}

	response := NumberResponse{
		WindowPrevState: prevState,
		WindowCurrState: currState,
		Numbers:         numbers,
		Avg:             avg,
	}

	return c.JSON(response)
}

func (s *CalculatorService) fetchNumbers(numberType NumberType) ([]int, error) {
	url := GetAPIEndpoint(numberType)
	
	if url == "" {
		return nil, fmt.Errorf("invalid number type: %s", numberType)
	}
	
	fmt.Printf("Fetching numbers from: %s\n", url)
	
	ctx, cancel := context.WithTimeout(context.Background(), RequestTimeout)
	defer cancel()
	
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	
	req.Header.Set("Authorization", AuthToken)
	fmt.Printf("Using Authorization header: %s\n", AuthToken[:40]+"...")
	
	for key, values := range req.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	
	resp, err := s.httpClient.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	
	fmt.Printf("Response status: %d\n", resp.StatusCode)
	
	fmt.Println("Response headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		fmt.Printf("Response body: %s\n", string(bodyBytes))
		return nil, fmt.Errorf("API server returned status: %d", resp.StatusCode)
	}
	
	var response struct {
		Numbers []int `json:"numbers"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return nil, err
	}
	
	fmt.Printf("Received numbers: %v\n", response.Numbers)
	
	// Check if we received an empty array
	if len(response.Numbers) == 0 {
		return nil, fmt.Errorf("API returned empty array of numbers")
	}
	
	return response.Numbers, nil
}

func (s *CalculatorService) addToWindow(num int) {
	if len(s.window) >= WindowSize {
		s.window = append(s.window[1:], num)
	} else {
		s.window = append(s.window, num)
	}
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
} 