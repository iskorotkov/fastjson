package avro

// Root response structure with Avro tags
type UserManagementResponse struct {
	APIVersion      string          `avro:"api_version" json:"api_version"`
	Timestamp       string          `avro:"timestamp" json:"timestamp"`
	RequestID       string          `avro:"request_id" json:"request_id"`
	Environment     string          `avro:"environment" json:"environment"`
	Region          string          `avro:"region" json:"region"`
	Pagination      Pagination      `avro:"pagination" json:"pagination"`
	Metadata        Metadata        `avro:"metadata" json:"metadata"`
	Users           []User          `avro:"users" json:"users"`
	SystemMetrics   SystemMetrics   `avro:"system_metrics" json:"system_metrics"`
	SecuritySummary SecuritySummary `avro:"security_summary" json:"security_summary"`
}

// Pagination information
type Pagination struct {
	Page        int  `avro:"page" json:"page"`
	PerPage     int  `avro:"per_page" json:"per_page"`
	TotalPages  int  `avro:"total_pages" json:"total_pages"`
	TotalCount  int  `avro:"total_count" json:"total_count"`
	HasNext     bool `avro:"has_next" json:"has_next"`
	HasPrevious bool `avro:"has_previous" json:"has_previous"`
}

// Request metadata
type Metadata struct {
	ExecutionTimeMs int     `avro:"execution_time_ms" json:"execution_time_ms"`
	CacheHit        bool    `avro:"cache_hit" json:"cache_hit"`
	DatabaseQueries int     `avro:"database_queries" json:"database_queries"`
	MemoryUsageMB   float64 `avro:"memory_usage_mb" json:"memory_usage_mb"`
	CPUUsagePercent float64 `avro:"cpu_usage_percent" json:"cpu_usage_percent"`
}

// User represents a complete user profile
type User struct {
	ID               string      `avro:"id" json:"id"`
	Username         string      `avro:"username" json:"username"`
	Email            string      `avro:"email" json:"email"`
	FirstName        string      `avro:"first_name" json:"first_name"`
	LastName         string      `avro:"last_name" json:"last_name"`
	DisplayName      string      `avro:"display_name" json:"display_name"`
	AvatarURL        string      `avro:"avatar_url" json:"avatar_url"`
	Status           string      `avro:"status" json:"status"`
	EmailVerified    bool        `avro:"email_verified" json:"email_verified"`
	PhoneVerified    bool        `avro:"phone_verified" json:"phone_verified"`
	TwoFactorEnabled bool        `avro:"two_factor_enabled" json:"two_factor_enabled"`
	CreatedAt        string      `avro:"created_at" json:"created_at"`
	UpdatedAt        string      `avro:"updated_at" json:"updated_at"`
	LastLogin        string      `avro:"last_login" json:"last_login"`
	LoginCount       int         `avro:"login_count" json:"login_count"`
	Profile          Profile     `avro:"profile" json:"profile"`
	Permissions      Permissions `avro:"permissions" json:"permissions"`
	Activity         Activity    `avro:"activity" json:"activity"`
	Preferences      Preferences `avro:"preferences" json:"preferences"`
}

// Profile contains user profile information
type Profile struct {
	Bio            string `avro:"bio" json:"bio"`
	Location       string `avro:"location" json:"location"`
	Timezone       string `avro:"timezone" json:"timezone"`
	Language       string `avro:"language" json:"language"`
	DateFormat     string `avro:"date_format" json:"date_format"`
	TimeFormat     string `avro:"time_format" json:"time_format"`
	Company        string `avro:"company" json:"company"`
	Department     string `avro:"department" json:"department"`
	Title          string `avro:"title" json:"title"`
	ManagerID      string `avro:"manager_id" json:"manager_id"`
	HireDate       string `avro:"hire_date" json:"hire_date"`
	SalaryBand     string `avro:"salary_band" json:"salary_band"`
	EmploymentType string `avro:"employment_type" json:"employment_type"`
}

// Permissions contains user access control information
type Permissions struct {
	Roles        []string     `avro:"roles" json:"roles"`
	Groups       []string     `avro:"groups" json:"groups"`
	AccessLevels AccessLevels `avro:"access_levels" json:"access_levels"`
	FeatureFlags FeatureFlags `avro:"feature_flags" json:"feature_flags"`
}

// AccessLevels defines what the user can access
type AccessLevels struct {
	Repositories  []string `avro:"repositories" json:"repositories"`
	Environments  []string `avro:"environments" json:"environments"`
	SensitiveData bool     `avro:"sensitive_data" json:"sensitive_data"`
	AdminPanel    bool     `avro:"admin_panel" json:"admin_panel"`
	Billing       bool     `avro:"billing" json:"billing"`
}

// FeatureFlags contains feature flag settings
type FeatureFlags struct {
	NewDashboard      bool `avro:"new_dashboard" json:"new_dashboard"`
	ExperimentalAI    bool `avro:"experimental_ai" json:"experimental_ai"`
	BetaMobileApp     bool `avro:"beta_mobile_app" json:"beta_mobile_app"`
	AdvancedAnalytics bool `avro:"advanced_analytics" json:"advanced_analytics"`
}

// Activity contains user activity information
type Activity struct {
	Last30Days    ActivityStats  `avro:"last_30_days" json:"last_30_days"`
	RecentActions []RecentAction `avro:"recent_actions" json:"recent_actions"`
}

// ActivityStats contains activity metrics
type ActivityStats struct {
	Logins         int `avro:"logins" json:"logins"`
	Commits        int `avro:"commits" json:"commits"`
	PullRequests   int `avro:"pull_requests" json:"pull_requests"`
	CodeReviews    int `avro:"code_reviews" json:"code_reviews"`
	Deployments    int `avro:"deployments" json:"deployments"`
	SupportTickets int `avro:"support_tickets" json:"support_tickets"`
}

// RecentAction represents a recent user action
type RecentAction struct {
	Action    string `avro:"action" json:"action"`
	Resource  string `avro:"resource" json:"resource"`
	Timestamp string `avro:"timestamp" json:"timestamp"`
	IPAddress string `avro:"ip_address" json:"ip_address"`
}

// Preferences contains user preference settings
type Preferences struct {
	Notifications NotificationSettings `avro:"notifications" json:"notifications"`
	UI            UISettings           `avro:"ui" json:"ui"`
}

// NotificationSettings contains notification preferences
type NotificationSettings struct {
	Email  EmailNotifications  `avro:"email" json:"email"`
	Slack  SlackNotifications  `avro:"slack" json:"slack"`
	Mobile MobileNotifications `avro:"mobile" json:"mobile"`
}

// EmailNotifications contains email notification settings
type EmailNotifications struct {
	SystemUpdates    bool `avro:"system_updates" json:"system_updates"`
	SecurityAlerts   bool `avro:"security_alerts" json:"security_alerts"`
	TeamMentions     bool `avro:"team_mentions" json:"team_mentions"`
	DeploymentStatus bool `avro:"deployment_status" json:"deployment_status"`
	WeeklySummary    bool `avro:"weekly_summary" json:"weekly_summary"`
}

// SlackNotifications contains Slack notification settings
type SlackNotifications struct {
	DirectMessages bool `avro:"direct_messages" json:"direct_messages"`
	TeamChannels   bool `avro:"team_channels" json:"team_channels"`
	UrgentAlerts   bool `avro:"urgent_alerts" json:"urgent_alerts"`
}

// MobileNotifications contains mobile notification settings
type MobileNotifications struct {
	PushEnabled bool       `avro:"push_enabled" json:"push_enabled"`
	QuietHours  QuietHours `avro:"quiet_hours" json:"quiet_hours"`
}

// QuietHours defines quiet hours for notifications
type QuietHours struct {
	Enabled bool   `avro:"enabled" json:"enabled"`
	Start   string `avro:"start" json:"start"`
	End     string `avro:"end" json:"end"`
}

// UISettings contains UI preference settings
type UISettings struct {
	Theme             string `avro:"theme" json:"theme"`
	SidebarCollapsed  bool   `avro:"sidebar_collapsed" json:"sidebar_collapsed"`
	CompactMode       bool   `avro:"compact_mode" json:"compact_mode"`
	AnimationsEnabled bool   `avro:"animations_enabled" json:"animations_enabled"`
}

// SystemMetrics contains overall system metrics
type SystemMetrics struct {
	ActiveUsersLast24h            int                `avro:"active_users_last_24h" json:"active_users_last_24h"`
	TotalLoginsToday              int                `avro:"total_logins_today" json:"total_logins_today"`
	FailedLoginAttempts           int                `avro:"failed_login_attempts" json:"failed_login_attempts"`
	PasswordResetsRequested       int                `avro:"password_resets_requested" json:"password_resets_requested"`
	NewUserRegistrations          int                `avro:"new_user_registrations" json:"new_user_registrations"`
	AverageSessionDurationMinutes int                `avro:"average_session_duration_minutes" json:"average_session_duration_minutes"`
	FeatureAdoptionRates          map[string]float64 `avro:"feature_adoption_rates" json:"feature_adoption_rates"`
}

// SecuritySummary contains security-related information
type SecuritySummary struct {
	SuspiciousActivities int              `avro:"suspicious_activities" json:"suspicious_activities"`
	BlockedIPs           []string         `avro:"blocked_ips" json:"blocked_ips"`
	SecurityAlerts       []SecurityAlert  `avro:"security_alerts" json:"security_alerts"`
	ComplianceStatus     ComplianceStatus `avro:"compliance_status" json:"compliance_status"`
}

// SecurityAlert represents a security alert
type SecurityAlert struct {
	ID        string `avro:"id" json:"id"`
	Type      string `avro:"type" json:"type"`
	UserID    string `avro:"user_id" json:"user_id"`
	Timestamp string `avro:"timestamp" json:"timestamp"`
	Severity  string `avro:"severity" json:"severity"`
	Resolved  bool   `avro:"resolved" json:"resolved"`
}

// ComplianceStatus contains compliance information
type ComplianceStatus struct {
	GDPRCompliant bool   `avro:"gdpr_compliant" json:"gdpr_compliant"`
	SOXCompliant  bool   `avro:"sox_compliant" json:"sox_compliant"`
	LastAudit     string `avro:"last_audit" json:"last_audit"`
	NextAudit     string `avro:"next_audit" json:"next_audit"`
}