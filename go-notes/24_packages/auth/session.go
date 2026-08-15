package auth

// private function since started with small letter
func extractSession() string{
	return "loggedin"
}

 
// public function since started with Capital letter
func GetSession() string{
	return extractSession()
} 