package az

type Cli struct {
}

func NewCli() *Cli {
	return &Cli{}
}

func (cli *Cli) GetAccounts() ([]string, error) {
	return []string{"Account 1", "Account 2"}, nil
}
