package app

import "time"

const (
	WindowSize = 10

	RequestTimeout = 2 * time.Second

	AuthToken = "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJNYXBDbGFpbXMiOnsiZXhwIjoxNzQzMTU1MDM1LCJpYXQiOjE3NDMxNTQ3MzUsImlzcyI6IkFmZm9yZG1lZCIsImp0aSI6ImY5YmY5YWZjLTBkYmEtNDNjZS04MTViLTJjNTlhZDdiOGNjYiIsInN1YiI6ImtyaXNoc3B5azEyMzBAZ21haWwuY29tIn0sImNvbXBhbnlOYW1lIjoiZ29NYXJ0IiwiY2xpZW50SUQiOiJmOWJmOWFmYy0wZGJhLTQzY2UtODE1Yi0yYzU5YWQ3YjhjY2IiLCJjbGllbnRTZWNyZXQiOiJPVmp1RHV5R2JwQWZoSGJ4Iiwib3duZXJOYW1lIjoiUmFodWwiLCJvd25lckVtYWlsIjoia3Jpc2hzcHlrMTIzMEBnbWFpbC5jb20iLCJyb2xsTm8iOiI3MTM1MjJJVDAzMSJ9.mDznIxHTSdSGUktKLmtY1FgzShQhmyIBzqA84tuqtxk"
)

func GetAPIEndpoint(numberType NumberType) string {
	switch numberType {
	case Prime:
		return "http://20.244.56.144/test/primes"
	case Fibonacci:
		return "http://20.244.56.144/test/fibo"
	case Even:
		return "http://20.244.56.144/test/even"
	case Random:
		return "http://20.244.56.144/test/rand"
	default:
		return ""
	}
} 