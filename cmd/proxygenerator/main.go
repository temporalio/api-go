// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type config struct {
	licenseFile    string
	license        string
	descriptorPath string
	verifyOnly     bool
}

func main() {
	var cfg config
	flag.StringVar(&cfg.licenseFile, "licenseFile", "../../LICENSE", "license file")
	flag.StringVar(&cfg.descriptorPath, "descriptorPath", "../../descriptor_set.pb", "path to the proto descriptor set")
	flag.BoolVar(&cfg.verifyOnly, "verifyOnly", false,
		"don't write to the filesystem, just verify output has not changed")
	flag.Parse()

	data, err := os.ReadFile(cfg.licenseFile)
	if err != nil {
		log.Fatalf("error reading license file, err=%v", err.Error())
	}

	cfg.license, err = commentOutLines(string(data))
	if err != nil {
		log.Fatalf("error re-writing license, err=%v", err.Error())
	}

	serviceErr := generateService(cfg)
	if serviceErr != nil {
		log.Print(serviceErr)
	}

	interceptorErr := generateInterceptor(cfg)
	if interceptorErr != nil {
		log.Print(interceptorErr)
	}

	if serviceErr != nil || interceptorErr != nil {
		os.Exit(1)
	}
}

func commentOutLines(str string) (string, error) {
	var lines []string
	scanner := bufio.NewScanner(strings.NewReader(str))
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			lines = append(lines, "//\n")
		} else {
			lines = append(lines, fmt.Sprintf("// %s\n", line))
		}
	}
	lines = append(lines, "\n")

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return strings.Join(lines, ""), nil
}
