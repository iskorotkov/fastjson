package benchmarks

// Root response structure
type UserManagementResponse struct {
	APIVersion      string          `json:"api_version"`
	Timestamp       string          `json:"timestamp"`
	RequestID       string          `json:"request_id"`
	Environment     string          `json:"environment"`
	Region          string          `json:"region"`
	Pagination      Pagination      `json:"pagination"`
	Metadata        Metadata        `json:"metadata"`
	Users           []User          `json:"users"`
	SystemMetrics   SystemMetrics   `json:"system_metrics"`
	SecuritySummary SecuritySummary `json:"security_summary"`
}

// Pagination information
type Pagination struct {
	Page        int  `json:"page"`
	PerPage     int  `json:"per_page"`
	TotalPages  int  `json:"total_pages"`
	TotalCount  int  `json:"total_count"`
	HasNext     bool `json:"has_next"`
	HasPrevious bool `json:"has_previous"`
}

// Request metadata
type Metadata struct {
	ExecutionTimeMs int     `json:"execution_time_ms"`
	CacheHit        bool    `json:"cache_hit"`
	DatabaseQueries int     `json:"database_queries"`
	MemoryUsageMB   float64 `json:"memory_usage_mb"`
	CPUUsagePercent float64 `json:"cpu_usage_percent"`
}

// User represents a complete user profile
type User struct {
	ID               string      `json:"id"`
	Username         string      `json:"username"`
	Email            string      `json:"email"`
	FirstName        string      `json:"first_name"`
	LastName         string      `json:"last_name"`
	DisplayName      string      `json:"display_name"`
	AvatarURL        string      `json:"avatar_url"`
	Status           string      `json:"status"`
	EmailVerified    bool        `json:"email_verified"`
	PhoneVerified    bool        `json:"phone_verified"`
	TwoFactorEnabled bool        `json:"two_factor_enabled"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	LastLogin        string      `json:"last_login"`
	LoginCount       int         `json:"login_count"`
	Profile          Profile     `json:"profile"`
	Permissions      Permissions `json:"permissions"`
	Activity         Activity    `json:"activity"`
	Preferences      Preferences `json:"preferences"`
}

// Profile contains user profile information
type Profile struct {
	Bio            string `json:"bio"`
	Location       string `json:"location"`
	Timezone       string `json:"timezone"`
	Language       string `json:"language"`
	DateFormat     string `json:"date_format"`
	TimeFormat     string `json:"time_format"`
	Company        string `json:"company"`
	Department     string `json:"department"`
	Title          string `json:"title"`
	ManagerID      string `json:"manager_id"`
	HireDate       string `json:"hire_date"`
	SalaryBand     string `json:"salary_band"`
	EmploymentType string `json:"employment_type"`
}

// Permissions contains user access control information
type Permissions struct {
	Roles        []string     `json:"roles"`
	Groups       []string     `json:"groups"`
	AccessLevels AccessLevels `json:"access_levels"`
	FeatureFlags FeatureFlags `json:"feature_flags"`
}

// AccessLevels defines what the user can access
type AccessLevels struct {
	Repositories  []string `json:"repositories"`
	Environments  []string `json:"environments"`
	SensitiveData bool     `json:"sensitive_data"`
	AdminPanel    bool     `json:"admin_panel"`
	Billing       bool     `json:"billing"`
}

// FeatureFlags contains feature flag settings
type FeatureFlags struct {
	NewDashboard      bool `json:"new_dashboard"`
	ExperimentalAI    bool `json:"experimental_ai"`
	BetaMobileApp     bool `json:"beta_mobile_app"`
	AdvancedAnalytics bool `json:"advanced_analytics"`
}

// Activity contains user activity information
type Activity struct {
	Last30Days    ActivityStats  `json:"last_30_days"`
	RecentActions []RecentAction `json:"recent_actions"`
}

// ActivityStats contains activity metrics
type ActivityStats struct {
	Logins         int `json:"logins"`
	Commits        int `json:"commits"`
	PullRequests   int `json:"pull_requests"`
	CodeReviews    int `json:"code_reviews"`
	Deployments    int `json:"deployments"`
	SupportTickets int `json:"support_tickets"`
}

// RecentAction represents a recent user action
type RecentAction struct {
	Action    string `json:"action"`
	Resource  string `json:"resource"`
	Timestamp string `json:"timestamp"`
	IPAddress string `json:"ip_address"`
}

// Preferences contains user preference settings
type Preferences struct {
	Notifications NotificationSettings `json:"notifications"`
	UI            UISettings           `json:"ui"`
}

// NotificationSettings contains notification preferences
type NotificationSettings struct {
	Email  EmailNotifications  `json:"email"`
	Slack  SlackNotifications  `json:"slack"`
	Mobile MobileNotifications `json:"mobile"`
}

// EmailNotifications contains email notification settings
type EmailNotifications struct {
	SystemUpdates    bool `json:"system_updates"`
	SecurityAlerts   bool `json:"security_alerts"`
	TeamMentions     bool `json:"team_mentions"`
	DeploymentStatus bool `json:"deployment_status"`
	WeeklySummary    bool `json:"weekly_summary"`
}

// SlackNotifications contains Slack notification settings
type SlackNotifications struct {
	DirectMessages bool `json:"direct_messages"`
	TeamChannels   bool `json:"team_channels"`
	UrgentAlerts   bool `json:"urgent_alerts"`
}

// MobileNotifications contains mobile notification settings
type MobileNotifications struct {
	PushEnabled bool       `json:"push_enabled"`
	QuietHours  QuietHours `json:"quiet_hours"`
}

// QuietHours defines quiet hours for notifications
type QuietHours struct {
	Enabled bool   `json:"enabled"`
	Start   string `json:"start"`
	End     string `json:"end"`
}

// UISettings contains UI preference settings
type UISettings struct {
	Theme             string `json:"theme"`
	SidebarCollapsed  bool   `json:"sidebar_collapsed"`
	CompactMode       bool   `json:"compact_mode"`
	AnimationsEnabled bool   `json:"animations_enabled"`
}

// SystemMetrics contains overall system metrics
type SystemMetrics struct {
	ActiveUsersLast24h            int                `json:"active_users_last_24h"`
	TotalLoginsToday              int                `json:"total_logins_today"`
	FailedLoginAttempts           int                `json:"failed_login_attempts"`
	PasswordResetsRequested       int                `json:"password_resets_requested"`
	NewUserRegistrations          int                `json:"new_user_registrations"`
	AverageSessionDurationMinutes int                `json:"average_session_duration_minutes"`
	FeatureAdoptionRates          map[string]float64 `json:"feature_adoption_rates"`
}

// SecuritySummary contains security-related information
type SecuritySummary struct {
	SuspiciousActivities int              `json:"suspicious_activities"`
	BlockedIPs           []string         `json:"blocked_ips"`
	SecurityAlerts       []SecurityAlert  `json:"security_alerts"`
	ComplianceStatus     ComplianceStatus `json:"compliance_status"`
}

// SecurityAlert represents a security alert
type SecurityAlert struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	UserID    string `json:"user_id"`
	Timestamp string `json:"timestamp"`
	Severity  string `json:"severity"`
	Resolved  bool   `json:"resolved"`
}

// ComplianceStatus contains compliance information
type ComplianceStatus struct {
	GDPRCompliant bool   `json:"gdpr_compliant"`
	SOXCompliant  bool   `json:"sox_compliant"`
	LastAudit     string `json:"last_audit"`
	NextAudit     string `json:"next_audit"`
}
