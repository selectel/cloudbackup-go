package cloudbackup

import (
	"context"
	"errors"
	"testing"

	"github.com/selectel/cloudbackup-go/pkg/httptest"
	"github.com/stretchr/testify/require"
)

func TestServiceClient_Checkpoints(t *testing.T) {
	t.Run("SuccessWithQuery", func(t *testing.T) {
		// Prepare
		body := `{
			"checkpoints": [{
				"id": "checkpoint-id-1",
				"plan_id": "plan-id-1",
				"created_at": "2023-02-01T00:00:00Z",
				"status": "completed",
				"checkpoint_items": [{
					"id": "item-id-1",
					"backup_id": "backup-id-1",
					"chain_id": "chain-id-1",
					"checkpoint_id": "checkpoint-id-1",
					"created_at": "2023-02-01T00:01:00Z",
					"backup_created_at": "2023-02-01T00:01:00Z",
					"is_incremental": false,
					"status": "available",
					"resource": {
						"id": "resource-id-1",
						"name": "resource-name-1",
						"type": "volume"
					}
				}]
			}],
			"total": 1
		}`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Checkpoints(context.Background(), &CheckpointsQuery{PlanName: "test-plan", VolumeName: "test-volume"})

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &CheckpointsResponse{
			Checkpoints: []*Checkpoint{
				{
					ID:        "checkpoint-id-1",
					PlanID:    "plan-id-1",
					CreatedAt: "2023-02-01T00:00:00Z",
					Status:    "completed",
					CheckpointItems: []CheckpointItem{
						{
							ID:              "item-id-1",
							BackupID:        "backup-id-1",
							ChainID:         "chain-id-1",
							CheckpointID:    "checkpoint-id-1",
							CreatedAt:       "2023-02-01T00:01:00Z",
							BackupCreatedAt: "2023-02-01T00:01:00Z",
							IsIncremental:   false,
							Status:          "available",
							Resource: CheckpointResource{
								ID:   "resource-id-1",
								Name: "resource-name-1",
								Type: "volume",
							},
						},
					},
				},
			},
			Total: 1,
		}
		require.Equal(t, want, res)
	})

	t.Run("SuccessWithoutQuery", func(t *testing.T) {
		// Prepare
		body := `{
			"checkpoints": [{
				"id": "checkpoint-id-1",
				"plan_id": "plan-id-1",
				"created_at": "2023-02-01T00:00:00Z",
				"status": "completed",
				"checkpoint_items": []
			}],
			"total": 1
		}`
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Checkpoints(context.Background(), nil)

		// Analyse
		require.NoError(t, err)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
		want := &CheckpointsResponse{
			Checkpoints: []*Checkpoint{
				{
					ID:              "checkpoint-id-1",
					PlanID:          "plan-id-1",
					CreatedAt:       "2023-02-01T00:00:00Z",
					Status:          "completed",
					CheckpointItems: []CheckpointItem{},
				},
			},
			Total: 1,
		}
		require.Equal(t, want, res)
	})

	t.Run("InvalidJSON", func(t *testing.T) {
		// Prepare
		body := invalidJSONBody
		fakeResp := httptest.NewFakeResponse(200, body) //nolint:bodyclose
		fakeTransport := httptest.NewFakeTransport(fakeResp, nil)
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Checkpoints(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.NotNil(t, respRes)
		require.Equal(t, 200, respRes.StatusCode)
	})

	t.Run("HTTPError", func(t *testing.T) {
		// Prepare
		body := httpErrorBody
		fakeResp := httptest.NewFakeResponse(404, body) //nolint:bodyclose
		client := newFakeClient("http://fake", httptest.NewFakeTransport(fakeResp, nil))

		// Execute
		res, respRes, err := client.Checkpoints(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.NotNil(t, respRes)
		require.NotNil(t, respRes.Err)
		require.Nil(t, res)
		require.EqualError(t, respRes.Err, httpErrorMessage)
	})

	t.Run("DoRequestError", func(t *testing.T) {
		// Prepare
		fakeTransport := httptest.NewFakeTransport(nil, errors.New("network failure"))
		client := newFakeClient("http://fake", fakeTransport)

		// Execute
		res, respRes, err := client.Checkpoints(context.Background(), nil)

		// Analyse
		require.Error(t, err)
		require.Nil(t, res)
		require.Nil(t, respRes)
	})
}
