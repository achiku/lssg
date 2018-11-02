package main

import (
	"io"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/hcl"
	"github.com/pkg/errors"
)

func parseSG(f io.Reader) error {
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return errors.Wrap(err, "failed to read")
	}
	as, err := hcl.ParseBytes(buf)
	if err != nil {
		return err
	}
	spew.Dump(as)
	return nil
}
