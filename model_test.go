package main

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestCompareTime(t *testing.T) {
	experiment := &TimePrecisionExperience{
		ExpTime: time.Now(),
	}

	err := gorm.G[TimePrecisionExperience](testDB).Create(context.Background(), experiment)
	require.NoError(t, err)

	savedExperiment, err := gorm.G[TimePrecisionExperience](testDB).Where("id = ?", experiment.ID).First(context.Background())
	require.NoError(t, err)

	t.Logf("local  experiment time: %s\n", experiment.ExpTime)
	t.Logf("stored experiment time: %s\n", savedExperiment.ExpTime)

	if diff := cmp.Diff(experiment, savedExperiment); diff != "" {
		t.Fatalf("experiment value is mismatch (-local, +db)\n%s", diff)
	}
}
