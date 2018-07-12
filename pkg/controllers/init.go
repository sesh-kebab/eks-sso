package controllers

import (
	"encoding/gob"

	"github.com/quintilesims/eks-sso/pkg/models"
)

func init() {
	gob.Register(&models.IAMCredentials{})
}
