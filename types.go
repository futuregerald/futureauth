package futureauth

// SignupData is what's sent to the signup endpoint and used to create the User model
type SignupData struct {
	ID           string   `json:"id,omitempty"`
	Email        string   `json:"email" validate:"required,email,max=30,min=6"`
	Tenant       string   `json:"tenantID"`
	Password     string   `json:"password" validate:"required,max=30,min=6"`
	AppMetaData  string   `json:"appMetaData"`
	UserMetaData string   `json:"userMetaData"`
	Confirmed    bool     `json:"confirmed"`
	IsAdmin      bool     `json:"isAdmin"`
	Disabled     bool     `json:"disabled"`
	Roles        []string `json:"roles"`
}

type User struct {
	Email        string   `json:"name"`
	Tenant       string   `json:"tenantID"`
	Password     string   `json:"password"`
	AppMetaData  string   `json:"appMetaData"`
	UserMetaData string   `json:"userMetaData"`
	Confirmed    bool     `json:"confirmed"`
	IsAdmin      bool     `json:"isAdmin"`
	Disabled     bool     `json:"disabled"`
	Roles        []string `json:"roles"`
}

// PasswordConfig is used to generate the argon2 password hash
type PasswordConfig struct {
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}
