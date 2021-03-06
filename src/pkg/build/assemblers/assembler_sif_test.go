// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the URIs of this project regarding your
// rights to use or distribute this software.

package assemblers_test

import (
	"os"
	"testing"

	"github.com/singularityware/singularity/src/pkg/build/assemblers"
	"github.com/singularityware/singularity/src/pkg/build/sources"
	"github.com/singularityware/singularity/src/pkg/build/types"
	"github.com/singularityware/singularity/src/pkg/test"
)

const (
	assemblerDockerURI  = "docker://alpine"
	assemblerDockerDest = "/tmp/docker_alpine_assemble_test.sif"
	assemblerShubURI    = "shub://ikaneshiro/singularityhub:latest"
	assemblerShubDest   = "/tmp/shub_alpine_assemble_test.sif"
)

// TestAssembler sees if we can build a SIF image from a docke based kitchen to /tmp
func TestSIFAssemblerDocker(t *testing.T) {
	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	def, err := types.NewDefinitionFromURI(assemblerDockerURI)
	if err != nil {
		t.Fatalf("unable to parse URI %s: %v\n", assemblerDockerURI, err)
	}

	ocp := &sources.OCIConveyorPacker{}

	if err := ocp.Get(def); err != nil {
		t.Fatalf("failed to Get from %s: %v\n", assemblerDockerURI, err)
	}

	b, err := ocp.Pack()
	if err != nil {
		t.Fatalf("failed to Pack from %s: %v\n", assemblerDockerURI, err)
	}

	a := &assemblers.SIFAssembler{}

	err = a.Assemble(b, assemblerDockerDest)
	if err != nil {
		t.Fatalf("failed to assemble from %s: %v\n", assemblerDockerURI, err)
	}

	defer os.Remove(assemblerDockerDest)
}
func TestSIFAssemblerShub(t *testing.T) {
	test.DropPrivilege(t)
	defer test.ResetPrivilege(t)

	def, err := types.NewDefinitionFromURI(assemblerShubURI)
	if err != nil {
		t.Fatalf("unable to parse URI %s: %v\n", assemblerShubURI, err)
	}

	scp := &sources.ShubConveyorPacker{}

	if err := scp.Get(def); err != nil {
		t.Fatalf("failed to Get from %s: %v\n", assemblerShubURI, err)
	}

	b, err := scp.Pack()
	if err != nil {
		t.Fatalf("failed to Pack from %s: %v\n", assemblerShubURI, err)
	}

	a := &assemblers.SIFAssembler{}

	err = a.Assemble(b, assemblerShubDest)
	if err != nil {
		t.Fatalf("failed to assemble from %s: %v\n", assemblerShubURI, err)
	}

	defer os.Remove(assemblerShubDest)
}
