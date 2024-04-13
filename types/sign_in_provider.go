package types

type SignInProvider string

const (
	EmailProvider     SignInProvider = "email"
	PhoneProvider     SignInProvider = "phone"
	GoogleProvider    SignInProvider = "google"
	FacebookProvider  SignInProvider = "facebook"
	AppleProvider     SignInProvider = "apple"
	MicrosoftProvider SignInProvider = "microsoft"
)
