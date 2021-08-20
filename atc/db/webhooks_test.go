package db_test

import (
	"encoding/json"

	"github.com/concourse/concourse/atc"
	"github.com/concourse/concourse/atc/db"
	"github.com/concourse/concourse/atc/db/dbtest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Webhooks", func() {
	It("creates a check for each resource matching the payload and type", func() {
		scenario1 := dbtest.Setup(
			builder.WithExistingTeam(defaultTeam),
			builder.WithNamedPipeline("p1", atc.Config{
				Resources: atc.ResourceConfigs{
					{
						Name: "match-type-and-payload",
						Type: dbtest.BaseResourceType,
						Webhooks: []atc.WebhookConfig{
							{
								Type:   "github",
								Filter: json.RawMessage(`{"foo": "bar"}`),
							},
						},
					},
					{
						Name: "mismatched-type",
						Type: dbtest.BaseResourceType,
						Webhooks: []atc.WebhookConfig{
							{
								Type:   "not-github",
								Filter: json.RawMessage(`{"foo": "bar"}`),
							},
						},
					},
					{
						Name: "mismatched-payload",
						Type: dbtest.BaseResourceType,
						Webhooks: []atc.WebhookConfig{
							{
								Type:   "github",
								Filter: json.RawMessage(`{"foo": "qux"}`),
							},
						},
					},
				},
			}),
		)

		scenario2 := dbtest.Setup(
			builder.WithExistingTeam(defaultTeam),
			builder.WithNamedPipeline("p2", atc.Config{
				Resources: atc.ResourceConfigs{
					{
						Name: "match-type-and-payload",
						Type: dbtest.BaseResourceType,
						Webhooks: []atc.WebhookConfig{
							{
								Type:   "github",
								Filter: json.RawMessage(`{"foo": "bar", "array": [{"hello": "world"}]}`),
							},
						},
					},
				},
			}),
		)

		err := webhooks.CreateWebhook(defaultTeam.ID(), "some-webhook", "github", "abc123")
		Expect(err).ToNot(HaveOccurred())

		payload := json.RawMessage(`{"foo": "bar", "array": [{"hello": "you", "num": 1}, {"hello": "world", "num": 2}]}`)
		numChecksCreated, err := webhooks.CheckResourcesMatchingWebhookPayload(logger, defaultTeam.ID(), "some-webhook", payload, "abc123")
		Expect(err).ToNot(HaveOccurred())
		Expect(numChecksCreated).To(Equal(2))

		Expect(scenario1.Resource("match-type-and-payload").BuildSummary()).ToNot(BeNil())
		Expect(scenario2.Resource("match-type-and-payload").BuildSummary()).ToNot(BeNil())

		Expect(scenario1.Resource("mismatched-type").BuildSummary()).To(BeNil())
		Expect(scenario1.Resource("mismatched-payload").BuildSummary()).To(BeNil())
	})

	It("errors if the webhook does not exist", func() {
		_, err := webhooks.CheckResourcesMatchingWebhookPayload(logger, defaultTeam.ID(), "invalid-webhook", nil, "abc123")
		Expect(err).To(MatchError(db.ErrMissingWebhook))
	})

	It("errors if the token is invalid", func() {
		err := webhooks.CreateWebhook(defaultTeam.ID(), "some-webhook", "github", "abc123")
		Expect(err).ToNot(HaveOccurred())

		_, err = webhooks.CheckResourcesMatchingWebhookPayload(logger, defaultTeam.ID(), "some-webhook", nil, "bad-token")
		Expect(err).To(MatchError(db.ErrInvalidWebhookToken))
	})
})
