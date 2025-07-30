package msgpack

type UserManagementResponse struct {
	APIVersion      string          `msgpack:"api_version" json:"api_version"`
	Timestamp       string          `msgpack:"timestamp" json:"timestamp"`
	RequestID       string          `msgpack:"request_id" json:"request_id"`
	Environment     string          `msgpack:"environment" json:"environment"`
	Region          string          `msgpack:"region" json:"region"`
	Pagination      Pagination      `msgpack:"pagination" json:"pagination"`
	Metadata        Metadata        `msgpack:"metadata" json:"metadata"`
	Users           []User          `msgpack:"users" json:"users"`
	SystemMetrics   SystemMetrics   `msgpack:"system_metrics" json:"system_metrics"`
	SecuritySummary SecuritySummary `msgpack:"security_summary" json:"security_summary"`
}

type Pagination struct {
	Page        int  `msgpack:"page" json:"page"`
	PerPage     int  `msgpack:"per_page" json:"per_page"`
	TotalPages  int  `msgpack:"total_pages" json:"total_pages"`
	TotalCount  int  `msgpack:"total_count" json:"total_count"`
	HasNext     bool `msgpack:"has_next" json:"has_next"`
	HasPrevious bool `msgpack:"has_previous" json:"has_previous"`
}

type Metadata struct {
	ExecutionTimeMs int     `msgpack:"execution_time_ms" json:"execution_time_ms"`
	CacheHit        bool    `msgpack:"cache_hit" json:"cache_hit"`
	DatabaseQueries int     `msgpack:"database_queries" json:"database_queries"`
	MemoryUsageMB   float64 `msgpack:"memory_usage_mb" json:"memory_usage_mb"`
	CPUUsagePercent float64 `msgpack:"cpu_usage_percent" json:"cpu_usage_percent"`
}

type User struct {
	ID               string      `msgpack:"id" json:"id"`
	Username         string      `msgpack:"username" json:"username"`
	Email            string      `msgpack:"email" json:"email"`
	FirstName        string      `msgpack:"first_name" json:"first_name"`
	LastName         string      `msgpack:"last_name" json:"last_name"`
	DisplayName      string      `msgpack:"display_name" json:"display_name"`
	AvatarURL        string      `msgpack:"avatar_url" json:"avatar_url"`
	Status           string      `msgpack:"status" json:"status"`
	EmailVerified    bool        `msgpack:"email_verified" json:"email_verified"`
	PhoneVerified    bool        `msgpack:"phone_verified" json:"phone_verified"`
	TwoFactorEnabled bool        `msgpack:"two_factor_enabled" json:"two_factor_enabled"`
	CreatedAt        string      `msgpack:"created_at" json:"created_at"`
	UpdatedAt        string      `msgpack:"updated_at" json:"updated_at"`
	LastLogin        string      `msgpack:"last_login" json:"last_login"`
	LoginCount       int         `msgpack:"login_count" json:"login_count"`
	Profile          Profile     `msgpack:"profile" json:"profile"`
	Permissions      Permissions `msgpack:"permissions" json:"permissions"`
	Activity         Activity    `msgpack:"activity" json:"activity"`
	Preferences      Preferences `msgpack:"preferences" json:"preferences"`
}

type Profile struct {
	Bio            string `msgpack:"bio" json:"bio"`
	Location       string `msgpack:"location" json:"location"`
	Timezone       string `msgpack:"timezone" json:"timezone"`
	Language       string `msgpack:"language" json:"language"`
	DateFormat     string `msgpack:"date_format" json:"date_format"`
	TimeFormat     string `msgpack:"time_format" json:"time_format"`
	Company        string `msgpack:"company" json:"company"`
	Department     string `msgpack:"department" json:"department"`
	Title          string `msgpack:"title" json:"title"`
	ManagerID      string `msgpack:"manager_id" json:"manager_id"`
	HireDate       string `msgpack:"hire_date" json:"hire_date"`
	SalaryBand     string `msgpack:"salary_band" json:"salary_band"`
	EmploymentType string `msgpack:"employment_type" json:"employment_type"`
}

type Permissions struct {
	Roles        []string     `msgpack:"roles" json:"roles"`
	Groups       []string     `msgpack:"groups" json:"groups"`
	AccessLevels AccessLevels `msgpack:"access_levels" json:"access_levels"`
	FeatureFlags FeatureFlags `msgpack:"feature_flags" json:"feature_flags"`
}

type AccessLevels struct {
	Repositories  []string `msgpack:"repositories" json:"repositories"`
	Environments  []string `msgpack:"environments" json:"environments"`
	SensitiveData bool     `msgpack:"sensitive_data" json:"sensitive_data"`
	AdminPanel    bool     `msgpack:"admin_panel" json:"admin_panel"`
	Billing       bool     `msgpack:"billing" json:"billing"`
}

type FeatureFlags struct {
	NewDashboard      bool `msgpack:"new_dashboard" json:"new_dashboard"`
	ExperimentalAI    bool `msgpack:"experimental_ai" json:"experimental_ai"`
	BetaMobileApp     bool `msgpack:"beta_mobile_app" json:"beta_mobile_app"`
	AdvancedAnalytics bool `msgpack:"advanced_analytics" json:"advanced_analytics"`
}

type Activity struct {
	Last30Days    ActivityStats  `msgpack:"last_30_days" json:"last_30_days"`
	RecentActions []RecentAction `msgpack:"recent_actions" json:"recent_actions"`
}

type ActivityStats struct {
	Logins         int `msgpack:"logins" json:"logins"`
	Commits        int `msgpack:"commits" json:"commits"`
	PullRequests   int `msgpack:"pull_requests" json:"pull_requests"`
	CodeReviews    int `msgpack:"code_reviews" json:"code_reviews"`
	Deployments    int `msgpack:"deployments" json:"deployments"`
	SupportTickets int `msgpack:"support_tickets" json:"support_tickets"`
}

type RecentAction struct {
	Action    string `msgpack:"action" json:"action"`
	Resource  string `msgpack:"resource" json:"resource"`
	Timestamp string `msgpack:"timestamp" json:"timestamp"`
	IPAddress string `msgpack:"ip_address" json:"ip_address"`
}

type Preferences struct {
	Notifications NotificationSettings `msgpack:"notifications" json:"notifications"`
	UI            UISettings           `msgpack:"ui" json:"ui"`
}

type NotificationSettings struct {
	Email  EmailNotifications  `msgpack:"email" json:"email"`
	Slack  SlackNotifications  `msgpack:"slack" json:"slack"`
	Mobile MobileNotifications `msgpack:"mobile" json:"mobile"`
}

type EmailNotifications struct {
	SystemUpdates    bool `msgpack:"system_updates" json:"system_updates"`
	SecurityAlerts   bool `msgpack:"security_alerts" json:"security_alerts"`
	TeamMentions     bool `msgpack:"team_mentions" json:"team_mentions"`
	DeploymentStatus bool `msgpack:"deployment_status" json:"deployment_status"`
	WeeklySummary    bool `msgpack:"weekly_summary" json:"weekly_summary"`
}

type SlackNotifications struct {
	DirectMessages bool `msgpack:"direct_messages" json:"direct_messages"`
	TeamChannels   bool `msgpack:"team_channels" json:"team_channels"`
	UrgentAlerts   bool `msgpack:"urgent_alerts" json:"urgent_alerts"`
}

type MobileNotifications struct {
	PushEnabled bool       `msgpack:"push_enabled" json:"push_enabled"`
	QuietHours  QuietHours `msgpack:"quiet_hours" json:"quiet_hours"`
}

type QuietHours struct {
	Enabled bool   `msgpack:"enabled" json:"enabled"`
	Start   string `msgpack:"start" json:"start"`
	End     string `msgpack:"end" json:"end"`
}

type UISettings struct {
	Theme             string `msgpack:"theme" json:"theme"`
	SidebarCollapsed  bool   `msgpack:"sidebar_collapsed" json:"sidebar_collapsed"`
	CompactMode       bool   `msgpack:"compact_mode" json:"compact_mode"`
	AnimationsEnabled bool   `msgpack:"animations_enabled" json:"animations_enabled"`
}

type SystemMetrics struct {
	ActiveUsersLast24h            int                `msgpack:"active_users_last_24h" json:"active_users_last_24h"`
	TotalLoginsToday              int                `msgpack:"total_logins_today" json:"total_logins_today"`
	FailedLoginAttempts           int                `msgpack:"failed_login_attempts" json:"failed_login_attempts"`
	PasswordResetsRequested       int                `msgpack:"password_resets_requested" json:"password_resets_requested"`
	NewUserRegistrations          int                `msgpack:"new_user_registrations" json:"new_user_registrations"`
	AverageSessionDurationMinutes int                `msgpack:"average_session_duration_minutes" json:"average_session_duration_minutes"`
	FeatureAdoptionRates          map[string]float64 `msgpack:"feature_adoption_rates" json:"feature_adoption_rates"`
}

type SecuritySummary struct {
	SuspiciousActivities int              `msgpack:"suspicious_activities" json:"suspicious_activities"`
	BlockedIPs           []string         `msgpack:"blocked_ips" json:"blocked_ips"`
	SecurityAlerts       []SecurityAlert  `msgpack:"security_alerts" json:"security_alerts"`
	ComplianceStatus     ComplianceStatus `msgpack:"compliance_status" json:"compliance_status"`
}

type SecurityAlert struct {
	ID        string `msgpack:"id" json:"id"`
	Type      string `msgpack:"type" json:"type"`
	UserID    string `msgpack:"user_id" json:"user_id"`
	Timestamp string `msgpack:"timestamp" json:"timestamp"`
	Severity  string `msgpack:"severity" json:"severity"`
	Resolved  bool   `msgpack:"resolved" json:"resolved"`
}

type ComplianceStatus struct {
	GDPRCompliant bool   `msgpack:"gdpr_compliant" json:"gdpr_compliant"`
	SOXCompliant  bool   `msgpack:"sox_compliant" json:"sox_compliant"`
	LastAudit     string `msgpack:"last_audit" json:"last_audit"`
	NextAudit     string `msgpack:"next_audit" json:"next_audit"`
}
