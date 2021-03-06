package commands

import (
	"fmt"

	"github.com/cloudfoundry/bosh-bootloader/storage"
)

const (
	EnvIDCommand            = "env-id"
	SSHKeyCommand           = "ssh-key"
	DirectorUsernameCommand = "director-username"
	DirectorPasswordCommand = "director-password"
	DirectorAddressCommand  = "director-address"
	DirectorCACertCommand   = "director-ca-cert"
	BOSHCACertCommand       = "bosh-ca-cert"

	EnvIDPropertyName            = "environment id"
	SSHKeyPropertyName           = "ssh key"
	DirectorUsernamePropertyName = "director username"
	DirectorPasswordPropertyName = "director password"
	DirectorAddressPropertyName  = "director address"
	DirectorCACertPropertyName   = "director ca cert"
	BOSHCACertPropertyName       = "bosh ca cert"
)

type StateQuery struct {
	logger         logger
	stateValidator stateValidator
	propertyName   string
	getProperty    getPropertyFunc
}

type getPropertyFunc func(storage.State) string

func NewStateQuery(logger logger, stateValidator stateValidator, propertyName string, getProperty getPropertyFunc) StateQuery {
	return StateQuery{
		logger:         logger,
		stateValidator: stateValidator,
		propertyName:   propertyName,
		getProperty:    getProperty,
	}
}

func (s StateQuery) Execute(subcommandFlags []string, state storage.State) error {
	err := s.stateValidator.Validate()
	if err != nil {
		return err
	}

	propertyValue := s.getProperty(state)
	if propertyValue == "" {
		return fmt.Errorf("Could not retrieve %s, please make sure you are targeting the proper state dir.", s.propertyName)
	}

	s.logger.Println(propertyValue)
	return nil
}
