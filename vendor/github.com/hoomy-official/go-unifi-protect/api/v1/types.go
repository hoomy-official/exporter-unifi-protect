package v1

import "time"

type Event struct {
	ID                string        `json:"id"`
	ModelKey          string        `json:"modelKey"`
	Type              string        `json:"type"`
	Start             int64         `json:"start"`
	End               int64         `json:"end"`
	Score             int           `json:"score"`
	SmartDetectTypes  []interface{} `json:"smartDetectTypes"`
	SmartDetectEvents []interface{} `json:"smartDetectEvents"`
	Camera            interface{}   `json:"camera"`
	Partition         interface{}   `json:"partition"`
	User              interface{}   `json:"user"`
	Metadata          map[string]struct {
		SensorID struct {
			Text string `json:"text"`
		} `json:"sensorId"`
		SensorName struct {
			Text string `json:"text"`
		} `json:"sensorName"`
		Type struct {
			Text string `json:"text"`
		} `json:"type"`
		MountType struct {
			Text string `json:"text"`
		} `json:"mountType"`
	} `json:"metadata"`
	Thumbnail string `json:"thumbnail"`
	Heatmap   string `json:"heatmap"`
	Timestamp int64  `json:"timestamp"`
}

type User struct {
	UniqueID           string `json:"unique_id"`
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	Alias              string `json:"alias"`
	FullName           string `json:"full_name"`
	Email              string `json:"email"`
	EmailStatus        string `json:"email_status"`
	EmailIsNull        bool   `json:"email_is_null"`
	Phone              string `json:"phone"`
	AvatarRelativePath string `json:"avatar_relative_path"`
	AvatarRpath2       string `json:"avatar_rpath2"`
	Status             string `json:"status"`
	EmployeeNumber     string `json:"employee_number"`
	CreateTime         int    `json:"create_time"`
	Extras             struct {
	} `json:"extras"`
	LoginTime         int    `json:"login_time"`
	Username          string `json:"username"`
	LocalAccountExist bool   `json:"local_account_exist"`
	PasswordRevision  int    `json:"password_revision"`
	OnlyUIAccount     bool   `json:"only_ui_account"`
	OnlyLocalAccount  bool   `json:"only_local_account"`
	SsoAccount        string `json:"sso_account"`
	SSOUID            string `json:"sso_uuid"`
	SSOUsername       string `json:"sso_username"`
	SSOPicture        string `json:"sso_picture"`
	UIDSSOId          string `json:"uid_sso_id"`
	UIDSSAccount      string `json:"uid_sso_account"`
	UIDAccountStatus  string `json:"uid_account_status"`
	Groups            []struct {
		UniqueID   string    `json:"unique_id"`
		Name       string    `json:"name"`
		UpID       string    `json:"up_id"`
		UpIDs      []string  `json:"up_ids"`
		SystemName string    `json:"system_name"`
		CreateTime time.Time `json:"create_time"`
	} `json:"groups"`
	Roles []struct {
		UniqueID   string    `json:"unique_id"`
		Name       string    `json:"name"`
		SystemRole bool      `json:"system_role"`
		SystemKey  string    `json:"system_key"`
		Level      int       `json:"level"`
		CreateTime time.Time `json:"create_time"`
		UpdateTime time.Time `json:"update_time"`
		IsPrivate  bool      `json:"is_private"`
	} `json:"roles"`
	Permissions struct {
		AccessManagement         []string `json:"access.management"`
		CalculusManagement       []string `json:"calculus.management"`
		ConnectManagement        []string `json:"connect.management"`
		DriveManagement          []string `json:"drive.management"`
		LedManagement            []string `json:"led.management"`
		NetworkManagement        []string `json:"network.management"`
		OlympusManagement        []string `json:"olympus.management"`
		ProtectManagement        []string `json:"protect.management"`
		SystemManagementLocation []string `json:"system.management.location"`
		SystemManagementUser     []string `json:"system.management.user"`
		TalkManagement           []string `json:"talk.management"`
	} `json:"permissions"`
	Scopes             []string    `json:"scopes"`
	CloudAccessGranted bool        `json:"cloud_access_granted"`
	UpdateTime         int         `json:"update_time"`
	Avatar             interface{} `json:"avatar"`
	NfcToken           string      `json:"nfc_token"`
	NfcDisplayID       string      `json:"nfc_display_id"`
	NfcCardType        string      `json:"nfc_card_type"`
	NfcCardStatus      string      `json:"nfc_card_status"`
	Role               string      `json:"role"`
	ID                 string      `json:"id"`
	IsOwner            bool        `json:"isOwner"`
	IsSuperAdmin       bool        `json:"isSuperAdmin"`
	IsMember           bool        `json:"isMember"`
	UcorePermission    struct {
		HasViewUserPermission         bool `json:"hasViewUserPermission"`
		HasEditUserPermission         bool `json:"hasEditUserPermission"`
		HasViewSettingsPermission     bool `json:"hasViewSettingsPermission"`
		HasUpdateChannelPermission    bool `json:"hasUpdateChannelPermission"`
		HasGeneralSettingsPermission  bool `json:"hasGeneralSettingsPermission"`
		HasUpdateAndInstallPermission bool `json:"hasUpdateAndInstallPermission"`
		HasAutoUpdatePermission       bool `json:"hasAutoUpdatePermission"`
		HasNotificationPermission     bool `json:"hasNotificationPermission"`
		HasRemoteAccessPermission     bool `json:"hasRemoteAccessPermission"`
		HasBackupPermission           bool `json:"hasBackupPermission"`
		HasRestartConsolePermission   bool `json:"hasRestartConsolePermission"`
		HasPoweroffConsolePermission  bool `json:"hasPoweroffConsolePermission"`
		HasResetConsolePermission     bool `json:"hasResetConsolePermission"`
		HasTransferOwnerPermission    bool `json:"hasTransferOwnerPermission"`
		HasSSHPermission              bool `json:"hasSSHPermission"`
		HasSupportFilePermission      bool `json:"hasSupportFilePermission"`
	} `json:"ucorePermission"`
	MaskedEmail string `json:"maskedEmail"`
	DeviceToken string `json:"deviceToken"`
	SsoAuth     struct {
		Name    string    `json:"name"`
		Value   string    `json:"value"`
		Expires time.Time `json:"expires"`
	} `json:"ssoAuth"`
}

type Sensor struct {
	Mac                   string      `json:"mac"`
	Host                  interface{} `json:"host"`
	ConnectionHost        string      `json:"connectionHost"`
	Type                  string      `json:"type"`
	Name                  string      `json:"name"`
	UpSince               int64       `json:"upSince"`
	Uptime                interface{} `json:"uptime"`
	LastSeen              int64       `json:"lastSeen"`
	ConnectedSince        int64       `json:"connectedSince"`
	State                 string      `json:"state"`
	LastDisconnect        interface{} `json:"lastDisconnect"`
	HardwareRevision      string      `json:"hardwareRevision"`
	FirmwareVersion       string      `json:"firmwareVersion"`
	LatestFirmwareVersion string      `json:"latestFirmwareVersion"`
	FirmwareBuild         interface{} `json:"firmwareBuild"`
	IsUpdating            bool        `json:"isUpdating"`
	IsDownloadingFW       bool        `json:"isDownloadingFW"`
	FwUpdateState         string      `json:"fwUpdateState"`
	IsAdopting            bool        `json:"isAdopting"`
	IsRestoring           bool        `json:"isRestoring"`
	IsAdopted             bool        `json:"isAdopted"`
	IsAdoptedByOther      bool        `json:"isAdoptedByOther"`
	IsProvisioned         bool        `json:"isProvisioned"`
	IsRebooting           bool        `json:"isRebooting"`
	IsSSHEnabled          bool        `json:"isSshEnabled"`
	CanAdopt              bool        `json:"canAdopt"`
	IsAttemptingToConnect bool        `json:"isAttemptingToConnect"`
	GUID                  interface{} `json:"guid"`
	AnonymousDeviceID     interface{} `json:"anonymousDeviceId"`
	IsMotionDetected      bool        `json:"isMotionDetected"`
	MountType             string      `json:"mountType"`
	LeakDetectedAt        interface{} `json:"leakDetectedAt"`
	TamperingDetectedAt   interface{} `json:"tamperingDetectedAt"`
	IsOpened              bool        `json:"isOpened"`
	OpenStatusChangedAt   int64       `json:"openStatusChangedAt"`
	AlarmTriggeredAt      interface{} `json:"alarmTriggeredAt"`
	MotionDetectedAt      int64       `json:"motionDetectedAt"`
	WiredConnectionState  struct {
		PhyRate interface{} `json:"phyRate"`
	} `json:"wiredConnectionState"`
	Stats struct {
		Light struct {
			Value  int    `json:"value"`
			Status string `json:"status"`
		} `json:"light"`
		Humidity struct {
			Value  int    `json:"value"`
			Status string `json:"status"`
		} `json:"humidity"`
		Temperature struct {
			Value  float64 `json:"value"`
			Status string  `json:"status"`
		} `json:"temperature"`
	} `json:"stats"`
	BluetoothConnectionState struct {
		SignalQuality  int `json:"signalQuality"`
		SignalStrength int `json:"signalStrength"`
	} `json:"bluetoothConnectionState"`
	BatteryStatus struct {
		Percentage int  `json:"percentage"`
		IsLow      bool `json:"isLow"`
	} `json:"batteryStatus"`
	AlarmSettings struct {
		IsEnabled bool `json:"isEnabled"`
	} `json:"alarmSettings"`
	LightSettings struct {
		IsEnabled     bool        `json:"isEnabled"`
		LowThreshold  interface{} `json:"lowThreshold"`
		HighThreshold interface{} `json:"highThreshold"`
		Margin        int         `json:"margin"`
	} `json:"lightSettings"`
	MotionSettings struct {
		IsEnabled   bool `json:"isEnabled"`
		Sensitivity int  `json:"sensitivity"`
	} `json:"motionSettings"`
	TemperatureSettings struct {
		IsEnabled     bool        `json:"isEnabled"`
		LowThreshold  interface{} `json:"lowThreshold"`
		HighThreshold interface{} `json:"highThreshold"`
		Margin        float64     `json:"margin"`
	} `json:"temperatureSettings"`
	HumiditySettings struct {
		IsEnabled     bool        `json:"isEnabled"`
		LowThreshold  interface{} `json:"lowThreshold"`
		HighThreshold interface{} `json:"highThreshold"`
		Margin        int         `json:"margin"`
	} `json:"humiditySettings"`
	LedSettings struct {
		IsEnabled bool `json:"isEnabled"`
	} `json:"ledSettings"`
	Bridge           string        `json:"bridge"`
	Camera           interface{}   `json:"camera"`
	BridgeCandidates []interface{} `json:"bridgeCandidates"`
	ID               string        `json:"id"`
	NvrMac           string        `json:"nvrMac"`
	IsConnected      bool          `json:"isConnected"`
	MarketName       string        `json:"marketName"`
	ModelKey         string        `json:"modelKey"`
}
