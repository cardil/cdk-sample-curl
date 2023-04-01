//go:generate npx cdktf-cli get

package main

import (
	"testing"

	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/stretchr/testify/assert"
)

// The tests below are example tests, you can find more information at
// https://cdk.tf/testing

var stack = NewMyStack(cdktf.Testing_App(nil), "stack")
var synth = cdktf.Testing_Synth(stack, jsii.Bool(true))

func TestCheckValidity(t *testing.T) {
	assertion := cdktf.Testing_ToBeValidTerraform(cdktf.Testing_FullSynth(stack))

	assert.True(t, *assertion)
}
