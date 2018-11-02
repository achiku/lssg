package main

import (
	"strings"
	"testing"

	"github.com/hashicorp/hcl"
)

var simpleconif = `
directory "config" {
    source_dir = "/etc/eventstore"
    dest_prefix = "escluster/config"
    exclude = []
    pre_backup_script = "before_backup.sh"
    post_backup_script = "after_backup.sh"
    pre_restore_script = "before_restore.sh"
    post_restore_script = "after_restore.sh"
}
`

var simplesg = `
resource "aws_security_group" "bastion-1" {
  description = "bastion-1 security group"
  name = "sg-bastion-1"
  egress = {
    cidr_blocks = ["0.0.0.0/0"]
    from_port = "0"
    protocol = "-1"
    to_port = "0"
  }
  ingress = {
    cidr_blocks = [
        "172.12.99.82/32",
    ]
    protocol = "tcp"
    from_port = "22"
    to_port = "22"
  }
  tags {
    Name = "sg-bastion-1"
  }
  vpc_id = "xxxxxxxxx"
}

resource "aws_security_group" "bastion-2" {
  description = "bastion-2 security group"
  name = "sg-bastion-2"
  egress = {
    cidr_blocks = ["0.0.0.0/0"]
    from_port = "0"
    protocol = "-1"
    to_port = "0"
  }
  ingress = {
    cidr_blocks = [
        "172.12.99.11/32",
    ]
    protocol = "tcp"
    from_port = "22"
    to_port = "22"
  }
  tags {
    Name = "sg-bastion-2"
  }
  vpc_id = "xxxxxxxxx"
}
`

func TestDecodeSimpleSG(t *testing.T) {
	var f SecurityGroupResource
	if err := hcl.Decode(&f, simplesg); err != nil {
		t.Fatal(err)
	}
	for k, v := range f.Resources {
		for _, sg := range v {
			t.Logf("id=%s: %+v", k, sg)
		}
	}
}

func TestParseSG(t *testing.T) {
	if err := parseSG(strings.NewReader(simplesg)); err != nil {
		t.Fatal(err)
	}
}
