package ken

// Environment represents an enumeration type that defines different environments, such as UAT and PROD.
// It is used to specify the environment in which an API or application is running.
//
// The Environment type is declared as an int and provides a String() method that returns the name of
// the environment as a string.
// The constants UAT and PROD are defined as values of the Environment type.
// UAT is set to 0, and PROD is set to 1.
//
// The Environment type can be used to create instances of different API environments.
// The uatApiEnvironment
// and prodApiEnvironment types implement the IApiEnvironment interface and provide specific base URL and
// token URL values for each environment.
//
// The Environment type is also used in the NewApiEnvironment function, which creates a new instance of
// IApiEnvironment based on the provided environment value.
// The returned instance will have the base URL and
// token URL set according to the provided environment.
type Environment int

// String returns the string representation of the Environment type.
// It returns "UAT" if the Environment value is UAT, and "PROD" if the
// Environment value is PROD.
//
// This method is used to convert the Environment type to its string representation,
// which is used in various parts of the code where the environment needs to be
// displayed or compared with other values.
func (e Environment) String() string {
	return [...]string{"UAT", "PROD"}[e]
}

const (
	UAT Environment = iota
	PROD
)

type apiEnvironment struct {
	environment Environment
	siteId      string
	listId      string
}

// Environment returns the environment associated with the API environment object.
func (a *apiEnvironment) Environment() Environment {
	return a.environment
}

func (a *apiEnvironment) SiteId() string {
	return a.siteId
}

func (a *apiEnvironment) ListId() string {
	return a.listId
}

type uatApiEnvironment struct {
	apiEnvironment
}

type prodApiEnvironment struct {
	apiEnvironment
}

type IApiEnvironment interface {
	Environment() Environment
	SiteId() string
	ListId() string
}

// NewApiEnvironment creates a new instance of IApiEnvironment based on the provided
// environment. It returns an implementation of the IApiEnvironment interface.
// The environment parameter can be either UAT or PROD.
// The returned IApiEnvironment instance will have the base URL and token URL
// set according to the provided environment value.
func NewApiEnvironment(environment Environment) IApiEnvironment {
	var env IApiEnvironment
	switch environment {
	case UAT:
		env = &uatApiEnvironment{
			apiEnvironment{
				environment: environment,
				siteId:      kenUatSiteId,
				listId:      kenUatListId,
			},
		}
	case PROD:
		env = &prodApiEnvironment{
			apiEnvironment{
				environment: environment,
				siteId:      kenSiteId,
				listId:      kenListId,
			},
		}
	}
	return env
}
