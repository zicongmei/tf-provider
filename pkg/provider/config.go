package provider

import "github.com/hashicorp/terraform-plugin-sdk/v2/diag"

type Config struct {
	DBFile1 string
	DBFile2 string
}

type DBClient struct {
	Path1 string
	Path2 string
}

func (c *Config) Client() (*DBClient, diag.Diagnostics) {
	return &DBClient{
		Path1: c.DBFile1,
		Path2: c.DBFile2,
	}, nil
}
