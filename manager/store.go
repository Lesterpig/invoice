package manager

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func Generate(input, output string) error {
	var i Invoice

	data, err := ioutil.ReadFile(input)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &i)
	if err != nil {
		return err
	}

	return i.PDF(output)
}

func Next() (string, error) {
	var max uint64
	names, _ := filepath.Glob("*.yml")
	for _, name := range names {
		if len(name) < 6 {
			continue
		}

		value, err := strconv.ParseUint(name[:6], 10, 64)
		if err != nil {
			continue
		}

		if value > max {
			max = value
		}
	}

	i := Invoice{
		ID:        max + 1,
		Emitted:   time.Now(),
		Delivered: time.Now(),
		Services: []Service{
			{},
		},
		Currency:    '€',
		PaymentDays: 30,
	}

	data, err := yaml.Marshal(&i)
	if err != nil {
		return "", err
	}

	filename := fmt.Sprintf("%06d", i.ID) + ".yml"
	err = ioutil.WriteFile(filename, data, 0644)
	return filename, err
}
