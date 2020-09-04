package isodos

// Client store the Isodos
// credentials and parameters
type Client struct {
	S3Key     string
	S3Secret  string
	Project   string
	IsodosURL string
}

// Init initialize the Isodos
// Client and return it
func Init(S3Key, S3Secret, project string) (client *Client) {
	client = new(Client)

	client.S3Key = S3Key
	client.S3Secret = S3Secret
	client.Project = project
	client.IsodosURL = "https://isodos.archive.org"

	return client
}
