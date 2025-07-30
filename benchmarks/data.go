package benchmarks

var Data = []byte(`{
  "api_version": "v2.1.0",
  "timestamp": "2025-07-18T14:30:45.123Z",
  "request_id": "req_7d8f9e2a1b4c5f6g7h8i9j0k",
  "environment": "production",
  "region": "us-east-1",
  "pagination": {
    "page": 1,
    "per_page": 50,
    "total_pages": 3,
    "total_count": 143,
    "has_next": true,
    "has_previous": false
  },
  "metadata": {
    "execution_time_ms": 234,
    "cache_hit": false,
    "database_queries": 7,
    "memory_usage_mb": 45.2,
    "cpu_usage_percent": 12.8
  },
  "users": [
    {
      "id": "usr_9f8e7d6c5b4a3f2e1d0c9b8a",
      "username": "alexandra.chen",
      "email": "alexandra.chen@company.com",
      "first_name": "Alexandra",
      "last_name": "Chen",
      "display_name": "Alex Chen",
      "avatar_url": "https://cdn.company.com/avatars/alexandra_chen.jpg",
      "status": "active",
      "email_verified": true,
      "phone_verified": false,
      "two_factor_enabled": true,
      "created_at": "2023-02-15T09:30:00Z",
      "updated_at": "2025-07-17T16:45:30Z",
      "last_login": "2025-07-18T08:15:22Z",
      "login_count": 847,
      "profile": {
        "bio": "Senior Software Engineer with 8+ years of experience in distributed systems and cloud architecture. Passionate about building scalable solutions.",
        "location": "San Francisco, CA",
        "timezone": "America/Los_Angeles",
        "language": "en-US",
        "date_format": "MM/DD/YYYY",
        "time_format": "12h",
        "company": "TechCorp Inc.",
        "department": "Engineering",
        "title": "Senior Software Engineer",
        "manager_id": "usr_1a2b3c4d5e6f7g8h9i0j1k2l",
        "hire_date": "2023-02-15",
        "salary_band": "L5",
        "employment_type": "full_time"
      },
      "permissions": {
        "roles": ["engineer", "code_reviewer", "deployment_approver"],
        "groups": ["frontend_team", "platform_team", "on_call_rotation"],
        "access_levels": {
          "repositories": ["read", "write", "admin"],
          "environments": ["development", "staging"],
          "sensitive_data": false,
          "admin_panel": false,
          "billing": false
        },
        "feature_flags": {
          "new_dashboard": true,
          "experimental_ai": true,
          "beta_mobile_app": false,
          "advanced_analytics": true
        }
      },
      "activity": {
        "last_30_days": {
          "logins": 28,
          "commits": 156,
          "pull_requests": 23,
          "code_reviews": 45,
          "deployments": 12,
          "support_tickets": 3
        },
        "recent_actions": [
          {
            "action": "code_review_approved",
            "resource": "PR #1247: Implement user authentication cache",
            "timestamp": "2025-07-18T13:45:00Z",
            "ip_address": "192.168.1.100"
          },
          {
            "action": "deployment_triggered",
            "resource": "staging environment",
            "timestamp": "2025-07-18T12:30:15Z",
            "ip_address": "192.168.1.100"
          },
          {
            "action": "repository_cloned",
            "resource": "user-service-v2",
            "timestamp": "2025-07-18T10:20:45Z",
            "ip_address": "192.168.1.100"
          }
        ]
      },
      "preferences": {
        "notifications": {
          "email": {
            "system_updates": true,
            "security_alerts": true,
            "team_mentions": true,
            "deployment_status": true,
            "weekly_summary": false
          },
          "slack": {
            "direct_messages": true,
            "team_channels": true,
            "urgent_alerts": true
          },
          "mobile": {
            "push_enabled": false,
            "quiet_hours": {
              "enabled": true,
              "start": "22:00",
              "end": "08:00"
            }
          }
        },
        "ui": {
          "theme": "dark",
          "sidebar_collapsed": false,
          "compact_mode": true,
          "animations_enabled": true
        }
      }
    },
    {
      "id": "usr_8e7d6c5b4a3f2e1d0c9b8a7f",
      "username": "michael.rodriguez",
      "email": "michael.rodriguez@company.com",
      "first_name": "Michael",
      "last_name": "Rodriguez",
      "display_name": "Mike Rodriguez",
      "avatar_url": "https://cdn.company.com/avatars/michael_rodriguez.jpg",
      "status": "active",
      "email_verified": true,
      "phone_verified": true,
      "two_factor_enabled": true,
      "created_at": "2022-11-08T14:20:00Z",
      "updated_at": "2025-07-16T11:30:45Z",
      "last_login": "2025-07-17T19:45:12Z",
      "login_count": 1205,
      "profile": {
        "bio": "Product Manager focused on user experience and data-driven decisions. Building products that users love.",
        "location": "Austin, TX",
        "timezone": "America/Chicago",
        "language": "en-US",
        "date_format": "DD/MM/YYYY",
        "time_format": "24h",
        "company": "TechCorp Inc.",
        "department": "Product",
        "title": "Senior Product Manager",
        "manager_id": "usr_2b3c4d5e6f7g8h9i0j1k2l3m",
        "hire_date": "2022-11-08",
        "salary_band": "L6",
        "employment_type": "full_time"
      },
      "permissions": {
        "roles": ["product_manager", "analytics_viewer", "user_researcher"],
        "groups": ["product_team", "growth_team", "executive_dashboard"],
        "access_levels": {
          "repositories": ["read"],
          "environments": ["development", "staging", "production"],
          "sensitive_data": true,
          "admin_panel": true,
          "billing": false
        },
        "feature_flags": {
          "new_dashboard": true,
          "experimental_ai": false,
          "beta_mobile_app": true,
          "advanced_analytics": true
        }
      },
      "activity": {
        "last_30_days": {
          "logins": 25,
          "commits": 0,
          "pull_requests": 0,
          "code_reviews": 8,
          "deployments": 0,
          "support_tickets": 12
        },
        "recent_actions": [
          {
            "action": "analytics_report_generated",
            "resource": "User Engagement Q3 2025",
            "timestamp": "2025-07-17T18:30:00Z",
            "ip_address": "10.0.1.50"
          },
          {
            "action": "feature_flag_updated",
            "resource": "beta_mobile_app enabled for 10% users",
            "timestamp": "2025-07-17T15:45:30Z",
            "ip_address": "10.0.1.50"
          },
          {
            "action": "user_feedback_reviewed",
            "resource": "Mobile app feedback batch #447",
            "timestamp": "2025-07-17T14:20:15Z",
            "ip_address": "10.0.1.50"
          }
        ]
      },
      "preferences": {
        "notifications": {
          "email": {
            "system_updates": false,
            "security_alerts": true,
            "team_mentions": true,
            "deployment_status": false,
            "weekly_summary": true
          },
          "slack": {
            "direct_messages": true,
            "team_channels": false,
            "urgent_alerts": true
          },
          "mobile": {
            "push_enabled": true,
            "quiet_hours": {
              "enabled": false,
              "start": "23:00",
              "end": "07:00"
            }
          }
        },
        "ui": {
          "theme": "light",
          "sidebar_collapsed": true,
          "compact_mode": false,
          "animations_enabled": false
        }
      }
    },
    {
      "id": "usr_7d6c5b4a3f2e1d0c9b8a7f6e",
      "username": "sarah.johnson",
      "email": "sarah.johnson@company.com",
      "first_name": "Sarah",
      "last_name": "Johnson",
      "display_name": "Sarah Johnson",
      "avatar_url": "https://cdn.company.com/avatars/sarah_johnson.jpg",
      "status": "active",
      "email_verified": true,
      "phone_verified": true,
      "two_factor_enabled": true,
      "created_at": "2021-05-22T10:15:00Z",
      "updated_at": "2025-07-18T09:20:30Z",
      "last_login": "2025-07-18T09:20:30Z",
      "login_count": 2156,
      "profile": {
        "bio": "Engineering Manager with extensive experience in team leadership and system architecture. Focused on building high-performing engineering teams.",
        "location": "New York, NY",
        "timezone": "America/New_York",
        "language": "en-US",
        "date_format": "YYYY-MM-DD",
        "time_format": "24h",
        "company": "TechCorp Inc.",
        "department": "Engineering",
        "title": "Engineering Manager",
        "manager_id": "usr_3c4d5e6f7g8h9i0j1k2l3m4n",
        "hire_date": "2021-05-22",
        "salary_band": "M1",
        "employment_type": "full_time"
      },
      "permissions": {
        "roles": ["engineering_manager", "hiring_manager", "performance_reviewer"],
        "groups": ["management_team", "platform_team", "hiring_committee"],
        "access_levels": {
          "repositories": ["read", "write", "admin"],
          "environments": ["development", "staging", "production"],
          "sensitive_data": true,
          "admin_panel": true,
          "billing": true
        },
        "feature_flags": {
          "new_dashboard": true,
          "experimental_ai": true,
          "beta_mobile_app": true,
          "advanced_analytics": true
        }
      },
      "activity": {
        "last_30_days": {
          "logins": 30,
          "commits": 45,
          "pull_requests": 8,
          "code_reviews": 89,
          "deployments": 25,
          "support_tickets": 15
        },
        "recent_actions": [
          {
            "action": "team_meeting_scheduled",
            "resource": "Sprint Planning - Sprint 47",
            "timestamp": "2025-07-18T09:00:00Z",
            "ip_address": "172.16.0.25"
          },
          {
            "action": "performance_review_submitted",
            "resource": "Q2 2025 - Alexandra Chen",
            "timestamp": "2025-07-17T16:30:45Z",
            "ip_address": "172.16.0.25"
          },
          {
            "action": "infrastructure_approval",
            "resource": "AWS RDS scaling for user-db-prod",
            "timestamp": "2025-07-17T14:15:20Z",
            "ip_address": "172.16.0.25"
          }
        ]
      },
      "preferences": {
        "notifications": {
          "email": {
            "system_updates": true,
            "security_alerts": true,
            "team_mentions": true,
            "deployment_status": true,
            "weekly_summary": true
          },
          "slack": {
            "direct_messages": true,
            "team_channels": true,
            "urgent_alerts": true
          },
          "mobile": {
            "push_enabled": true,
            "quiet_hours": {
              "enabled": true,
              "start": "21:00",
              "end": "07:00"
            }
          }
        },
        "ui": {
          "theme": "auto",
          "sidebar_collapsed": false,
          "compact_mode": true,
          "animations_enabled": true
        }
      }
    }
  ],
  "system_metrics": {
    "active_users_last_24h": 1247,
    "total_logins_today": 3456,
    "failed_login_attempts": 23,
    "password_resets_requested": 7,
    "new_user_registrations": 12,
    "average_session_duration_minutes": 127,
    "feature_adoption_rates": {
      "new_dashboard": 0.78,
      "experimental_ai": 0.34,
      "beta_mobile_app": 0.12,
      "advanced_analytics": 0.89
    }
  },
  "security_summary": {
    "suspicious_activities": 0,
    "blocked_ips": ["203.0.113.45", "198.51.100.78"],
    "security_alerts": [
      {
        "id": "alert_789",
        "type": "multiple_failed_logins",
        "user_id": "usr_4d5e6f7g8h9i0j1k2l3m4n5o",
        "timestamp": "2025-07-18T06:30:00Z",
        "severity": "medium",
        "resolved": true
      }
    ],
    "compliance_status": {
      "gdpr_compliant": true,
      "sox_compliant": true,
      "last_audit": "2025-06-15T00:00:00Z",
      "next_audit": "2025-09-15T00:00:00Z"
    }
  }
}`)

var Response = UserManagementResponse{
	APIVersion:  "v2.1.0",
	Timestamp:   "2025-07-18T14:30:45.123Z",
	RequestID:   "req_7d8f9e2a1b4c5f6g7h8i9j0k",
	Environment: "production",
	Region:      "us-east-1",
	Pagination: Pagination{
		Page:        1,
		PerPage:     50,
		TotalPages:  3,
		TotalCount:  143,
		HasNext:     true,
		HasPrevious: false,
	},
	Metadata: Metadata{
		ExecutionTimeMs: 234,
		CacheHit:        false,
		DatabaseQueries: 7,
		MemoryUsageMB:   45.2,
		CPUUsagePercent: 12.8,
	},
	Users: []User{
		{
			ID:               "usr_9f8e7d6c5b4a3f2e1d0c9b8a",
			Username:         "alexandra.chen",
			Email:            "alexandra.chen@company.com",
			FirstName:        "Alexandra",
			LastName:         "Chen",
			DisplayName:      "Alex Chen",
			AvatarURL:        "https://cdn.company.com/avatars/alexandra_chen.jpg",
			Status:           "active",
			EmailVerified:    true,
			PhoneVerified:    false,
			TwoFactorEnabled: true,
			CreatedAt:        "2023-02-15T09:30:00Z",
			UpdatedAt:        "2025-07-17T16:45:30Z",
			LastLogin:        "2025-07-18T08:15:22Z",
			LoginCount:       847,
			Profile: Profile{
				Bio:            "Senior Software Engineer with 8+ years of experience in distributed systems and cloud architecture. Passionate about building scalable solutions.",
				Location:       "San Francisco, CA",
				Timezone:       "America/Los_Angeles",
				Language:       "en-US",
				DateFormat:     "MM/DD/YYYY",
				TimeFormat:     "12h",
				Company:        "TechCorp Inc.",
				Department:     "Engineering",
				Title:          "Senior Software Engineer",
				ManagerID:      "usr_1a2b3c4d5e6f7g8h9i0j1k2l",
				HireDate:       "2023-02-15",
				SalaryBand:     "L5",
				EmploymentType: "full_time",
			},
			Permissions: Permissions{
				Roles:  []string{"engineer", "code_reviewer", "deployment_approver"},
				Groups: []string{"frontend_team", "platform_team", "on_call_rotation"},
				AccessLevels: AccessLevels{
					Repositories:  []string{"read", "write", "admin"},
					Environments:  []string{"development", "staging"},
					SensitiveData: false,
					AdminPanel:    false,
					Billing:       false,
				},
				FeatureFlags: FeatureFlags{
					NewDashboard:      true,
					ExperimentalAI:    true,
					BetaMobileApp:     false,
					AdvancedAnalytics: true,
				},
			},
			Activity: Activity{
				Last30Days: ActivityStats{
					Logins:         28,
					Commits:        156,
					PullRequests:   23,
					CodeReviews:    45,
					Deployments:    12,
					SupportTickets: 3,
				},
				RecentActions: []RecentAction{
					{
						Action:    "code_review_approved",
						Resource:  "PR #1247: Implement user authentication cache",
						Timestamp: "2025-07-18T13:45:00Z",
						IPAddress: "192.168.1.100",
					},
					{
						Action:    "deployment_triggered",
						Resource:  "staging environment",
						Timestamp: "2025-07-18T12:30:15Z",
						IPAddress: "192.168.1.100",
					},
					{
						Action:    "repository_cloned",
						Resource:  "user-service-v2",
						Timestamp: "2025-07-18T10:20:45Z",
						IPAddress: "192.168.1.100",
					},
				},
			},
			Preferences: Preferences{
				Notifications: NotificationSettings{
					Email: EmailNotifications{
						SystemUpdates:    true,
						SecurityAlerts:   true,
						TeamMentions:     true,
						DeploymentStatus: true,
						WeeklySummary:    false,
					},
					Slack: SlackNotifications{
						DirectMessages: true,
						TeamChannels:   true,
						UrgentAlerts:   true,
					},
					Mobile: MobileNotifications{
						PushEnabled: false,
						QuietHours: QuietHours{
							Enabled: true,
							Start:   "22:00",
							End:     "08:00",
						},
					},
				},
				UI: UISettings{
					Theme:             "dark",
					SidebarCollapsed:  false,
					CompactMode:       true,
					AnimationsEnabled: true,
				},
			},
		},
		{
			ID:               "usr_8e7d6c5b4a3f2e1d0c9b8a7f",
			Username:         "michael.rodriguez",
			Email:            "michael.rodriguez@company.com",
			FirstName:        "Michael",
			LastName:         "Rodriguez",
			DisplayName:      "Mike Rodriguez",
			AvatarURL:        "https://cdn.company.com/avatars/michael_rodriguez.jpg",
			Status:           "active",
			EmailVerified:    true,
			PhoneVerified:    true,
			TwoFactorEnabled: true,
			CreatedAt:        "2022-11-08T14:20:00Z",
			UpdatedAt:        "2025-07-16T11:30:45Z",
			LastLogin:        "2025-07-17T19:45:12Z",
			LoginCount:       1205,
			Profile: Profile{
				Bio:            "Product Manager focused on user experience and data-driven decisions. Building products that users love.",
				Location:       "Austin, TX",
				Timezone:       "America/Chicago",
				Language:       "en-US",
				DateFormat:     "DD/MM/YYYY",
				TimeFormat:     "24h",
				Company:        "TechCorp Inc.",
				Department:     "Product",
				Title:          "Senior Product Manager",
				ManagerID:      "usr_2b3c4d5e6f7g8h9i0j1k2l3m",
				HireDate:       "2022-11-08",
				SalaryBand:     "L6",
				EmploymentType: "full_time",
			},
			Permissions: Permissions{
				Roles:  []string{"product_manager", "analytics_viewer", "user_researcher"},
				Groups: []string{"product_team", "growth_team", "executive_dashboard"},
				AccessLevels: AccessLevels{
					Repositories:  []string{"read"},
					Environments:  []string{"development", "staging", "production"},
					SensitiveData: true,
					AdminPanel:    true,
					Billing:       false,
				},
				FeatureFlags: FeatureFlags{
					NewDashboard:      true,
					ExperimentalAI:    false,
					BetaMobileApp:     true,
					AdvancedAnalytics: true,
				},
			},
			Activity: Activity{
				Last30Days: ActivityStats{
					Logins:         25,
					Commits:        0,
					PullRequests:   0,
					CodeReviews:    8,
					Deployments:    0,
					SupportTickets: 12,
				},
				RecentActions: []RecentAction{
					{
						Action:    "analytics_report_generated",
						Resource:  "User Engagement Q3 2025",
						Timestamp: "2025-07-17T18:30:00Z",
						IPAddress: "10.0.1.50",
					},
					{
						Action:    "feature_flag_updated",
						Resource:  "beta_mobile_app enabled for 10% users",
						Timestamp: "2025-07-17T15:45:30Z",
						IPAddress: "10.0.1.50",
					},
					{
						Action:    "user_feedback_reviewed",
						Resource:  "Mobile app feedback batch #447",
						Timestamp: "2025-07-17T14:20:15Z",
						IPAddress: "10.0.1.50",
					},
				},
			},
			Preferences: Preferences{
				Notifications: NotificationSettings{
					Email: EmailNotifications{
						SystemUpdates:    false,
						SecurityAlerts:   true,
						TeamMentions:     true,
						DeploymentStatus: false,
						WeeklySummary:    true,
					},
					Slack: SlackNotifications{
						DirectMessages: true,
						TeamChannels:   false,
						UrgentAlerts:   true,
					},
					Mobile: MobileNotifications{
						PushEnabled: true,
						QuietHours: QuietHours{
							Enabled: false,
							Start:   "23:00",
							End:     "07:00",
						},
					},
				},
				UI: UISettings{
					Theme:             "light",
					SidebarCollapsed:  true,
					CompactMode:       false,
					AnimationsEnabled: false,
				},
			},
		},
		{
			ID:               "usr_7d6c5b4a3f2e1d0c9b8a7f6e",
			Username:         "sarah.johnson",
			Email:            "sarah.johnson@company.com",
			FirstName:        "Sarah",
			LastName:         "Johnson",
			DisplayName:      "Sarah Johnson",
			AvatarURL:        "https://cdn.company.com/avatars/sarah_johnson.jpg",
			Status:           "active",
			EmailVerified:    true,
			PhoneVerified:    true,
			TwoFactorEnabled: true,
			CreatedAt:        "2021-05-22T10:15:00Z",
			UpdatedAt:        "2025-07-18T09:20:30Z",
			LastLogin:        "2025-07-18T09:20:30Z",
			LoginCount:       2156,
			Profile: Profile{
				Bio:            "Engineering Manager with extensive experience in team leadership and system architecture. Focused on building high-performing engineering teams.",
				Location:       "New York, NY",
				Timezone:       "America/New_York",
				Language:       "en-US",
				DateFormat:     "YYYY-MM-DD",
				TimeFormat:     "24h",
				Company:        "TechCorp Inc.",
				Department:     "Engineering",
				Title:          "Engineering Manager",
				ManagerID:      "usr_3c4d5e6f7g8h9i0j1k2l3m4n",
				HireDate:       "2021-05-22",
				SalaryBand:     "M1",
				EmploymentType: "full_time",
			},
			Permissions: Permissions{
				Roles:  []string{"engineering_manager", "hiring_manager", "performance_reviewer"},
				Groups: []string{"management_team", "platform_team", "hiring_committee"},
				AccessLevels: AccessLevels{
					Repositories:  []string{"read", "write", "admin"},
					Environments:  []string{"development", "staging", "production"},
					SensitiveData: true,
					AdminPanel:    true,
					Billing:       true,
				},
				FeatureFlags: FeatureFlags{
					NewDashboard:      true,
					ExperimentalAI:    true,
					BetaMobileApp:     true,
					AdvancedAnalytics: true,
				},
			},
			Activity: Activity{
				Last30Days: ActivityStats{
					Logins:         30,
					Commits:        45,
					PullRequests:   8,
					CodeReviews:    89,
					Deployments:    25,
					SupportTickets: 15,
				},
				RecentActions: []RecentAction{
					{
						Action:    "team_meeting_scheduled",
						Resource:  "Sprint Planning - Sprint 47",
						Timestamp: "2025-07-18T09:00:00Z",
						IPAddress: "172.16.0.25",
					},
					{
						Action:    "performance_review_submitted",
						Resource:  "Q2 2025 - Alexandra Chen",
						Timestamp: "2025-07-17T16:30:45Z",
						IPAddress: "172.16.0.25",
					},
					{
						Action:    "infrastructure_approval",
						Resource:  "AWS RDS scaling for user-db-prod",
						Timestamp: "2025-07-17T14:15:20Z",
						IPAddress: "172.16.0.25",
					},
				},
			},
			Preferences: Preferences{
				Notifications: NotificationSettings{
					Email: EmailNotifications{
						SystemUpdates:    true,
						SecurityAlerts:   true,
						TeamMentions:     true,
						DeploymentStatus: true,
						WeeklySummary:    true,
					},
					Slack: SlackNotifications{
						DirectMessages: true,
						TeamChannels:   true,
						UrgentAlerts:   true,
					},
					Mobile: MobileNotifications{
						PushEnabled: true,
						QuietHours: QuietHours{
							Enabled: true,
							Start:   "21:00",
							End:     "07:00",
						},
					},
				},
				UI: UISettings{
					Theme:             "auto",
					SidebarCollapsed:  false,
					CompactMode:       true,
					AnimationsEnabled: true,
				},
			},
		},
	},
	SystemMetrics: SystemMetrics{
		ActiveUsersLast24h:            1247,
		TotalLoginsToday:              3456,
		FailedLoginAttempts:           23,
		PasswordResetsRequested:       7,
		NewUserRegistrations:          12,
		AverageSessionDurationMinutes: 127,
		FeatureAdoptionRates: map[string]float64{
			"new_dashboard":      0.78,
			"experimental_ai":    0.34,
			"beta_mobile_app":    0.12,
			"advanced_analytics": 0.89,
		},
	},
	SecuritySummary: SecuritySummary{
		SuspiciousActivities: 0,
		BlockedIPs:           []string{"203.0.113.45", "198.51.100.78"},
		SecurityAlerts: []SecurityAlert{
			{
				ID:        "alert_789",
				Type:      "multiple_failed_logins",
				UserID:    "usr_4d5e6f7g8h9i0j1k2l3m4n5o",
				Timestamp: "2025-07-18T06:30:00Z",
				Severity:  "medium",
				Resolved:  true,
			},
		},
		ComplianceStatus: ComplianceStatus{
			GDPRCompliant: true,
			SOXCompliant:  true,
			LastAudit:     "2025-06-15T00:00:00Z",
			NextAudit:     "2025-09-15T00:00:00Z",
		},
	},
}
