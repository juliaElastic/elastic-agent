// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildfleetcfg/buildfleetcfg.go - DO NOT EDIT.

package application

import "github.com/elastic/beats/v7/x-pack/agent/pkg/packer"

// DefaultAgentFleetConfig is the content of the default configuration when we enroll a beat, the agent.yml
// will be replaced with this variables.
var DefaultAgentFleetConfig []byte

func init() {
	// Packed File
	// _meta/agent.fleet.yml
	unpacked := packer.MustUnpack("eJyMVV+Po74Vfe/H2A/wK3+G1VLp9xBnBmMW2IbM2NgvFbYzhsROUBMgUPW7VyZMZqYdVfsQRQJ8zz3nnnv8r2//MLtL9ddK7Y6XP171bnf5YzT629++ockJ//78ez/0JDU3eKRkUL975v7bnPboKdG0LCYE9YRg3nPDWubjkZGNkiQ4sBKpwlw1KzdnFOtebsGekgfFIDYVCTSCbi/8jRIediTUHfOzn2gERpjwguLixLYgrUhSc/Ok8MuhQ9F7DVmCAy0Lbb+rykwxo89sCxw+gj33AlMR6QrzoiSsNYpzLeOi5UZO9ntabixObd9zEzoozl0Rg14cC83WoNltAeQQa7l++MlJ1FEiNSe4k4+nn2gNLtwr9K9m1bAbt2at2gP3ZcdhWLP5m9UexXkvy2TPtqBnDXAq+KJomTgVYTX1iyldg5GVkVuViRYj0BxGk4R6j+C1rbwXJfxilCR3ZJloBKOOrUFHidtyIxT36I0XTFpuorEieFo3K4VGoFMjVOrLlpuil3522T0qLz22rVid1L0njLfb55NCpnZkDKZfzY/+U28Gn2mZOxXJOur96BkMfW6uPSXFlHq3macjWniCWkKlJPyx9FS0wuC9hOG4286zcCjRXUWCI4JFL73gzL3oYD0iYLSvvOjIyszq6tAyObKyeGVQj4zkDveTwOqcxJeBlsUpNUUtYRSm3nsdce+jaDnBvSw3iplw/Ig1Y8egpt6lZt7LjFURV3MfO7+a1ZTF2a0GZC2HL7Yvp1rPXqpRDFxqri0dwT9ZeVCyzK1HDPeRkgaPFWHBV9w+6rnMpmUN6Li/efNqK0YwCIMnRgIHweg4/8ey5XBQMk4CFBc9glb3jRIxbvjsj4vebZcdgZ/mpKSHG+Fph20fLMeB+cBU5Kq/1vD9rIByQOqk0DpTFQkmCaMzj8L5ewQLLbx8rMjs09nHHIZ7WhYt9x5mL97PQHx41xlPqQ9G7smREndKG5DI9cy55nHeo9jy3rzp3osGjIzcZ7Wcc5TwgaaeNhXJawl1z49ZZ/X87z6/5qg74Rc1h4Pd0UGYcM/KfPq1cOWzL92ar8HA/cRBMHC55dnc5it8UEsvsPO+n128eqlI8CpgOMrHk8qeV0P2uFq4uPUO6ou0vdv3a9BJ4jY2C2dfxvlAS9nN/o+Tnnp4Et5tfzmJHkrf5tDmO3rMbM1h8c7dr4tWe2FwPWfmCDpWik+10gb4FcTdTcdcCz9TNiPfso7Oz3IjyfVsd/RWCx8QnD0356mE+sJI6MrVTSurozhuFPeBFiZyKhJ299w/5i2Dt9ynW2Ao0WebW7b3RcPXez2r1wSmtWpHVhbj7nnJS4g7SpKzrcEgdliZWI2t5ywXp4J4RHGi572wuDav/byl5qrTBjwxwmpJrs6tJ2Y9plM75/e631FsM/VlyYu5VsvK/JV6ocuPmw9ew9MHL9j8nnu1z/8H6/HmRzSC7czHZp6PO7m251ZXBF1NSWDvyPtuU/OihMGOMNp6q+fzPaUvbAtqdixmTPuMmajlMR6Xfb73UHp5L0kwZ9eCfWDkWu+2YBIQ7yvCWmbvxDkfDqoiD4qXmaIGO9LDmv9/PoMYrG6FvZOCz/no1rso/BLLercirsu3915a3oCLGEHDysJnxPrxTfOsQ0+FZiZyebyZvZlNqwX3A8bzSeXEuWE/4QcBw85mZEWis+1PQjzaHbb40qu1vX9E845v8+rd29GE4Oca6ZID2NNOtZ553jWjJDigx+w+JzFZ367O6JEO6Xp1zfbgTfsvvWffsbIeuBdofixaSobv93fqzz+//fsv/wkAAP//30VSPA==")
	raw, ok := unpacked["_meta/agent.fleet.yml"]
	if !ok {
		// ensure we have something loaded.
		panic("agent.fleet.yml is not included in the binary")
	}
	DefaultAgentFleetConfig = raw
}