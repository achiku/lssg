package main

import (
	"io"
	"io/ioutil"

	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

// SecurityGroup aws security group
type SecurityGroup struct {
	Name        string `hcl:"name"`
	Description string `hcl:"description"`
	Egress      struct {
		CIDR     []string `hcl:"cidr_blocks"`
		FromPort string   `hcl:"from_port"`
		ToPort   string   `hcl:"to_port"`
		Protocol string   `hcl:"protocol"`
	} `hcl:"egress"`
	Inress struct {
		CIDR     []string `hcl:"cidr_blocks"`
		FromPort string   `hcl:"from_port"`
		ToPort   string   `hcl:"to_port"`
		Protocol string   `hcl:"protocol"`
	} `hcl:"ingress"`
	VpcID string `hcl:"vpc_id"`
}

// SecurityGroupResource resource
type SecurityGroupResource struct {
	Resources map[string]map[string]*SecurityGroup `hcl:"resource,expand"`
}

// ParseSG parse security groups
func ParseSG(f io.Reader) (*SecurityGroupResource, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read")
	}
	var sgs SecurityGroupResource
	if err := hcl.Decode(&sgs, string(buf)); err != nil {
		return nil, err
	}
	return &sgs, nil
}
