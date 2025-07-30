package avro

import (
	"encoding/json"
	"os"
	"runtime"
	"runtime/debug"
	"testing"

	"github.com/hamba/avro/v2"
	"github.com/iskorotkov/fastjson/benchmarks"
)

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(1)
	debug.SetGCPercent(-1)
	debug.SetMemoryLimit(25 * (1 << 20))

	os.Exit(m.Run())
}

var schema = avro.MustParse(`{
	"type": "record",
	"name": "UserManagementResponse",
	"namespace": "benchmarks",
	"fields": [
		{"name": "api_version", "type": "string"},
		{"name": "timestamp", "type": "string"},
		{"name": "request_id", "type": "string"},
		{"name": "environment", "type": "string"},
		{"name": "region", "type": "string"},
		{"name": "pagination", "type": {
			"type": "record",
			"name": "Pagination",
			"fields": [
				{"name": "page", "type": "int"},
				{"name": "per_page", "type": "int"},
				{"name": "total_pages", "type": "int"},
				{"name": "total_count", "type": "int"},
				{"name": "has_next", "type": "boolean"},
				{"name": "has_previous", "type": "boolean"}
			]
		}},
		{"name": "metadata", "type": {
			"type": "record",
			"name": "Metadata",
			"fields": [
				{"name": "execution_time_ms", "type": "int"},
				{"name": "cache_hit", "type": "boolean"},
				{"name": "database_queries", "type": "int"},
				{"name": "memory_usage_mb", "type": "double"},
				{"name": "cpu_usage_percent", "type": "double"}
			]
		}},
		{"name": "users", "type": {
			"type": "array",
			"items": {
				"type": "record",
				"name": "User",
				"fields": [
					{"name": "id", "type": "string"},
					{"name": "username", "type": "string"},
					{"name": "email", "type": "string"},
					{"name": "first_name", "type": "string"},
					{"name": "last_name", "type": "string"},
					{"name": "display_name", "type": "string"},
					{"name": "avatar_url", "type": "string"},
					{"name": "status", "type": "string"},
					{"name": "email_verified", "type": "boolean"},
					{"name": "phone_verified", "type": "boolean"},
					{"name": "two_factor_enabled", "type": "boolean"},
					{"name": "created_at", "type": "string"},
					{"name": "updated_at", "type": "string"},
					{"name": "last_login", "type": "string"},
					{"name": "login_count", "type": "int"},
					{"name": "profile", "type": {
						"type": "record",
						"name": "Profile",
						"fields": [
							{"name": "bio", "type": "string"},
							{"name": "location", "type": "string"},
							{"name": "timezone", "type": "string"},
							{"name": "language", "type": "string"},
							{"name": "date_format", "type": "string"},
							{"name": "time_format", "type": "string"},
							{"name": "company", "type": "string"},
							{"name": "department", "type": "string"},
							{"name": "title", "type": "string"},
							{"name": "manager_id", "type": "string"},
							{"name": "hire_date", "type": "string"},
							{"name": "salary_band", "type": "string"},
							{"name": "employment_type", "type": "string"}
						]
					}},
					{"name": "permissions", "type": {
						"type": "record",
						"name": "Permissions",
						"fields": [
							{"name": "roles", "type": {"type": "array", "items": "string"}},
							{"name": "groups", "type": {"type": "array", "items": "string"}},
							{"name": "access_levels", "type": {
								"type": "record",
								"name": "AccessLevels",
								"fields": [
									{"name": "repositories", "type": {"type": "array", "items": "string"}},
									{"name": "environments", "type": {"type": "array", "items": "string"}},
									{"name": "sensitive_data", "type": "boolean"},
									{"name": "admin_panel", "type": "boolean"},
									{"name": "billing", "type": "boolean"}
								]
							}},
							{"name": "feature_flags", "type": {
								"type": "record",
								"name": "FeatureFlags",
								"fields": [
									{"name": "new_dashboard", "type": "boolean"},
									{"name": "experimental_ai", "type": "boolean"},
									{"name": "beta_mobile_app", "type": "boolean"},
									{"name": "advanced_analytics", "type": "boolean"}
								]
							}}
						]
					}},
					{"name": "activity", "type": {
						"type": "record",
						"name": "Activity",
						"fields": [
							{"name": "last_30_days", "type": {
								"type": "record",
								"name": "ActivityStats",
								"fields": [
									{"name": "logins", "type": "int"},
									{"name": "commits", "type": "int"},
									{"name": "pull_requests", "type": "int"},
									{"name": "code_reviews", "type": "int"},
									{"name": "deployments", "type": "int"},
									{"name": "support_tickets", "type": "int"}
								]
							}},
							{"name": "recent_actions", "type": {
								"type": "array",
								"items": {
									"type": "record",
									"name": "RecentAction",
									"fields": [
										{"name": "action", "type": "string"},
										{"name": "resource", "type": "string"},
										{"name": "timestamp", "type": "string"},
										{"name": "ip_address", "type": "string"}
									]
								}
							}}
						]
					}},
					{"name": "preferences", "type": {
						"type": "record",
						"name": "Preferences",
						"fields": [
							{"name": "notifications", "type": {
								"type": "record",
								"name": "NotificationSettings",
								"fields": [
									{"name": "email", "type": {
										"type": "record",
										"name": "EmailNotifications",
										"fields": [
											{"name": "system_updates", "type": "boolean"},
											{"name": "security_alerts", "type": "boolean"},
											{"name": "team_mentions", "type": "boolean"},
											{"name": "deployment_status", "type": "boolean"},
											{"name": "weekly_summary", "type": "boolean"}
										]
									}},
									{"name": "slack", "type": {
										"type": "record",
										"name": "SlackNotifications",
										"fields": [
											{"name": "direct_messages", "type": "boolean"},
											{"name": "team_channels", "type": "boolean"},
											{"name": "urgent_alerts", "type": "boolean"}
										]
									}},
									{"name": "mobile", "type": {
										"type": "record",
										"name": "MobileNotifications",
										"fields": [
											{"name": "push_enabled", "type": "boolean"},
											{"name": "quiet_hours", "type": {
												"type": "record",
												"name": "QuietHours",
												"fields": [
													{"name": "enabled", "type": "boolean"},
													{"name": "start", "type": "string"},
													{"name": "end", "type": "string"}
												]
											}}
										]
									}}
								]
							}},
							{"name": "ui", "type": {
								"type": "record",
								"name": "UISettings",
								"fields": [
									{"name": "theme", "type": "string"},
									{"name": "sidebar_collapsed", "type": "boolean"},
									{"name": "compact_mode", "type": "boolean"},
									{"name": "animations_enabled", "type": "boolean"}
								]
							}}
						]
					}}
				]
			}
		}},
		{"name": "system_metrics", "type": {
			"type": "record",
			"name": "SystemMetrics",
			"fields": [
				{"name": "active_users_last_24h", "type": "int"},
				{"name": "total_logins_today", "type": "int"},
				{"name": "failed_login_attempts", "type": "int"},
				{"name": "password_resets_requested", "type": "int"},
				{"name": "new_user_registrations", "type": "int"},
				{"name": "average_session_duration_minutes", "type": "int"},
				{"name": "feature_adoption_rates", "type": {"type": "map", "values": "double"}}
			]
		}},
		{"name": "security_summary", "type": {
			"type": "record",
			"name": "SecuritySummary",
			"fields": [
				{"name": "suspicious_activities", "type": "int"},
				{"name": "blocked_ips", "type": {"type": "array", "items": "string"}},
				{"name": "security_alerts", "type": {
					"type": "array",
					"items": {
						"type": "record",
						"name": "SecurityAlert",
						"fields": [
							{"name": "id", "type": "string"},
							{"name": "type", "type": "string"},
							{"name": "user_id", "type": "string"},
							{"name": "timestamp", "type": "string"},
							{"name": "severity", "type": "string"},
							{"name": "resolved", "type": "boolean"}
						]
					}
				}},
				{"name": "compliance_status", "type": {
					"type": "record",
					"name": "ComplianceStatus",
					"fields": [
						{"name": "gdpr_compliant", "type": "boolean"},
						{"name": "sox_compliant", "type": "boolean"},
						{"name": "last_audit", "type": "string"},
						{"name": "next_audit", "type": "string"}
					]
				}}
			]
		}}
	]
}`)

var avroData []byte

func init() {
	// Convert JSON to Go struct with Avro tags
	var jsonResp UserManagementResponse
	if err := json.Unmarshal(benchmarks.Data, &jsonResp); err != nil {
		panic(err)
	}

	// Marshal to Avro binary format
	var err error
	avroData, err = avro.Marshal(schema, jsonResp)
	if err != nil {
		panic(err)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	b.Run("hamba/avro/v2", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		b.SetBytes(int64(len(avroData)))

		for b.Loop() {
			var resp UserManagementResponse
			if err := avro.Unmarshal(schema, avroData, &resp); err != nil {
				b.Fatalf("unexpected error: %v", err)
			}
		}
	})
}
